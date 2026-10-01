// Package temporalnexus provides helpers for converting between Temporal
// common.v1 links and Nexus links.
package temporalnexus

import (
	"fmt"
	"net/url"
	"regexp"
	"strconv"

	"github.com/nexus-rpc/sdk-go/nexus"
	commonpb "go.temporal.io/api/common/v1"
	enumspb "go.temporal.io/api/enums/v1"
	nexuspb "go.temporal.io/api/nexus/v1"
	"google.golang.org/protobuf/proto"
)

const (
	urlSchemeTemporalKey = "temporal"

	linkWorkflowEventReferenceTypeKey = "referenceType"
	linkEventIDKey                    = "eventID"
	linkEventTypeKey                  = "eventType"
	linkRequestIDKey                  = "requestID"
	// linkReasonKey carries Link_Workflow.reason: an optional, human-readable
	// explanation of why the link exists (e.g. the reason a Query or Update
	// created the link). Encoded as a URL query param on the Nexus link.
	linkReasonKey = "reason"
)

// linkType describes the Nexus URL form of one common.v1.Link variant. Every variant's path is
// /namespaces/{namespace}/{collection}/{id}/{runID}, optionally followed by a fixed tail segment,
// so encoding and decoding differ between variants only by this table.
type linkType struct {
	// name is the proto full name, carried as nexus.Link.Type.
	name string
	// label names the variant in error messages, e.g. "Link_Activity".
	label string
	// template is the path with a %s for each of namespace, id and runID.
	template string
	// pathRE matches an escaped path against template, capturing namespace, id and runID.
	pathRE *regexp.Regexp
}

func newLinkType(m proto.Message, template string) linkType {
	const segment = `([^/]+)`
	desc := m.ProtoReflect().Descriptor()
	return linkType{
		name:     string(desc.FullName()),
		label:    "Link_" + string(desc.Name()),
		template: template,
		pathRE:   regexp.MustCompile("^" + fmt.Sprintf(template, segment, segment, segment) + "$"),
	}
}

var (
	workflowEventLinkType  = newLinkType(&commonpb.Link_WorkflowEvent{}, "/namespaces/%s/workflows/%s/%s/history")
	workflowLinkType       = newLinkType(&commonpb.Link_Workflow{}, "/namespaces/%s/workflows/%s/%s")
	nexusOperationLinkType = newLinkType(&commonpb.Link_NexusOperation{}, "/namespaces/%s/nexus-operations/%s/%s/details")
	activityLinkType       = newLinkType(&commonpb.Link_Activity{}, "/namespaces/%s/activities/%s/%s/details")

	eventReferenceType     = string((&commonpb.Link_WorkflowEvent_EventReference{}).ProtoReflect().Descriptor().Name())
	requestIDReferenceType = string((&commonpb.Link_WorkflowEvent_RequestIdReference{}).ProtoReflect().Descriptor().Name())
)

// encode builds the Nexus link for this variant. Path holds the raw IDs and RawPath their escaped
// form, so an ID containing a slash stays a single segment.
func (lt linkType) encode(namespace, id, runID, rawQuery string) nexus.Link {
	return nexus.Link{
		URL: &url.URL{
			Scheme:   urlSchemeTemporalKey,
			Path:     fmt.Sprintf(lt.template, namespace, id, runID),
			RawPath:  fmt.Sprintf(lt.template, url.PathEscape(namespace), url.PathEscape(id), url.PathEscape(runID)),
			RawQuery: rawQuery,
		},
		Type: lt.name,
	}
}

// decodedLink holds the path IDs and query of a Nexus link that matched its linkType.
type decodedLink struct {
	namespace, id, runID string
	query                url.Values
}

// decode validates a Nexus link against this variant: the declared type, the scheme, and a path
// matching the template exactly.
func (lt linkType) decode(link nexus.Link) (decodedLink, error) {
	if link.Type != lt.name {
		return decodedLink{}, fmt.Errorf("cannot parse link type %q to %q", link.Type, lt.name)
	}
	if link.URL == nil {
		return decodedLink{}, lt.errorf("empty URL")
	}
	if link.URL.Scheme != urlSchemeTemporalKey {
		return decodedLink{}, lt.errorf("invalid scheme: %s", link.URL.Scheme)
	}
	matches := lt.pathRE.FindStringSubmatch(link.URL.EscapedPath())
	if len(matches) != 4 {
		return decodedLink{}, lt.errorf("malformed URL path")
	}
	var ids [3]string
	for i, match := range matches[1:] {
		id, err := url.PathUnescape(match)
		if err != nil {
			return decodedLink{}, lt.errorf("%w", err)
		}
		ids[i] = id
	}
	return decodedLink{namespace: ids[0], id: ids[1], runID: ids[2], query: link.URL.Query()}, nil
}

func (lt linkType) errorf(format string, args ...any) error {
	return fmt.Errorf("failed to parse link to "+lt.label+": "+format, args...)
}

// Outbound: common.v1.Link -> Nexus link.

// ConvertLinkWorkflowEventToNexusLink converts a Link_WorkflowEvent type to Nexus Link.
//
// NOTE: Experimental
func ConvertLinkWorkflowEventToNexusLink(we *commonpb.Link_WorkflowEvent) nexus.Link {
	var rawQuery string
	switch ref := we.GetReference().(type) {
	case *commonpb.Link_WorkflowEvent_EventRef:
		rawQuery = convertLinkWorkflowEventEventReferenceToURLQuery(ref.EventRef)
	case *commonpb.Link_WorkflowEvent_RequestIdRef:
		rawQuery = convertLinkWorkflowEventRequestIdReferenceToURLQuery(ref.RequestIdRef)
	}
	return workflowEventLinkType.encode(we.GetNamespace(), we.GetWorkflowId(), we.GetRunId(), rawQuery)
}

// ConvertLinkWorkflowToNexusLink converts a Link_Workflow type to Nexus Link.
//
// NOTE: Experimental
func ConvertLinkWorkflowToNexusLink(w *commonpb.Link_Workflow) nexus.Link {
	var rawQuery string
	if w.GetReason() != "" {
		rawQuery = url.Values{linkReasonKey: {w.GetReason()}}.Encode()
	}
	return workflowLinkType.encode(w.GetNamespace(), w.GetWorkflowId(), w.GetRunId(), rawQuery)
}

// ConvertLinkNexusOperationToNexusLink converts a Link_NexusOperation type to Nexus Link.
//
// NOTE: Experimental
func ConvertLinkNexusOperationToNexusLink(no *commonpb.Link_NexusOperation) nexus.Link {
	return nexusOperationLinkType.encode(no.GetNamespace(), no.GetOperationId(), no.GetRunId(), "")
}

// ConvertLinkActivityToNexusLink converts a Link_Activity type to Nexus Link.
//
// NOTE: Experimental
func ConvertLinkActivityToNexusLink(a *commonpb.Link_Activity) nexus.Link {
	return activityLinkType.encode(a.GetNamespace(), a.GetActivityId(), a.GetRunId(), "")
}

// CommonLinkToNexusLink converts a common.v1.Link into a nexus.v1.Link, dispatching on the link's
// variant. Returns (nil, false) for any variant not handled here.
//
// NOTE: Experimental
func CommonLinkToNexusLink(link *commonpb.Link) (*nexuspb.Link, bool) {
	var nexusLink nexus.Link
	switch v := link.GetVariant().(type) {
	case *commonpb.Link_WorkflowEvent_:
		if v.WorkflowEvent == nil {
			return nil, false
		}
		nexusLink = ConvertLinkWorkflowEventToNexusLink(v.WorkflowEvent)
	case *commonpb.Link_NexusOperation_:
		if v.NexusOperation == nil {
			return nil, false
		}
		nexusLink = ConvertLinkNexusOperationToNexusLink(v.NexusOperation)
	case *commonpb.Link_Activity_:
		if v.Activity == nil {
			return nil, false
		}
		nexusLink = ConvertLinkActivityToNexusLink(v.Activity)
	case *commonpb.Link_Workflow_:
		if v.Workflow == nil {
			return nil, false
		}
		nexusLink = ConvertLinkWorkflowToNexusLink(v.Workflow)
	default:
		return nil, false
	}
	return &nexuspb.Link{Url: nexusLink.URL.String(), Type: nexusLink.Type}, true
}

// Inbound: Nexus link -> common.v1.Link.

// ConvertNexusLinkToLinkWorkflowEvent converts a Nexus Link to Link_WorkflowEvent.
//
// NOTE: Experimental
func ConvertNexusLinkToLinkWorkflowEvent(link nexus.Link) (*commonpb.Link_WorkflowEvent, error) {
	d, err := workflowEventLinkType.decode(link)
	if err != nil {
		return nil, err
	}
	we := &commonpb.Link_WorkflowEvent{
		Namespace:  d.namespace,
		WorkflowId: d.id,
		RunId:      d.runID,
	}

	switch refType := d.query.Get(linkWorkflowEventReferenceTypeKey); refType {
	case eventReferenceType:
		eventRef, err := convertURLQueryToLinkWorkflowEventEventReference(d.query)
		if err != nil {
			return nil, workflowEventLinkType.errorf("%w", err)
		}
		we.Reference = &commonpb.Link_WorkflowEvent_EventRef{EventRef: eventRef}
	case requestIDReferenceType:
		requestIDRef, err := convertURLQueryToLinkWorkflowEventRequestIdReference(d.query)
		if err != nil {
			return nil, workflowEventLinkType.errorf("%w", err)
		}
		we.Reference = &commonpb.Link_WorkflowEvent_RequestIdRef{RequestIdRef: requestIDRef}
	default:
		return nil, workflowEventLinkType.errorf("unknown reference type: %q", refType)
	}
	return we, nil
}

// ConvertNexusLinkToLinkWorkflow converts a Nexus Link to Link_Workflow.
//
// NOTE: Experimental
func ConvertNexusLinkToLinkWorkflow(link nexus.Link) (*commonpb.Link_Workflow, error) {
	d, err := workflowLinkType.decode(link)
	if err != nil {
		return nil, err
	}
	return &commonpb.Link_Workflow{
		Namespace:  d.namespace,
		WorkflowId: d.id,
		RunId:      d.runID,
		Reason:     d.query.Get(linkReasonKey),
	}, nil
}

// ConvertNexusLinkToLinkNexusOperation converts a Nexus Link to Link_NexusOperation.
//
// NOTE: Experimental
func ConvertNexusLinkToLinkNexusOperation(link nexus.Link) (*commonpb.Link_NexusOperation, error) {
	d, err := nexusOperationLinkType.decode(link)
	if err != nil {
		return nil, err
	}
	return &commonpb.Link_NexusOperation{
		Namespace:   d.namespace,
		OperationId: d.id,
		RunId:       d.runID,
	}, nil
}

// ConvertNexusLinkToLinkActivity converts a Nexus Link to Link_Activity.
//
// NOTE: Experimental
func ConvertNexusLinkToLinkActivity(link nexus.Link) (*commonpb.Link_Activity, error) {
	d, err := activityLinkType.decode(link)
	if err != nil {
		return nil, err
	}
	return &commonpb.Link_Activity{
		Namespace:  d.namespace,
		ActivityId: d.id,
		RunId:      d.runID,
	}, nil
}

// NexusLinkToCommonLink converts a nexus.v1.Link into a common.v1.Link, dispatching on link.Type.
// Returns (nil, false) for any link type not handled here.
//
// NOTE: Experimental
func NexusLinkToCommonLink(link *nexuspb.Link) (*commonpb.Link, bool) {
	nexusLink := nexus.Link{Type: link.GetType()}
	if link.GetUrl() != "" {
		u, err := url.Parse(link.GetUrl())
		if err != nil {
			return nil, false
		}
		nexusLink.URL = u
	}

	var commonLink *commonpb.Link
	var err error
	switch nexusLink.Type {
	case workflowEventLinkType.name:
		var we *commonpb.Link_WorkflowEvent
		we, err = ConvertNexusLinkToLinkWorkflowEvent(nexusLink)
		commonLink = &commonpb.Link{Variant: &commonpb.Link_WorkflowEvent_{WorkflowEvent: we}}
	case nexusOperationLinkType.name:
		var no *commonpb.Link_NexusOperation
		no, err = ConvertNexusLinkToLinkNexusOperation(nexusLink)
		commonLink = &commonpb.Link{Variant: &commonpb.Link_NexusOperation_{NexusOperation: no}}
	case activityLinkType.name:
		var a *commonpb.Link_Activity
		a, err = ConvertNexusLinkToLinkActivity(nexusLink)
		commonLink = &commonpb.Link{Variant: &commonpb.Link_Activity_{Activity: a}}
	case workflowLinkType.name:
		var w *commonpb.Link_Workflow
		w, err = ConvertNexusLinkToLinkWorkflow(nexusLink)
		commonLink = &commonpb.Link{Variant: &commonpb.Link_Workflow_{Workflow: w}}
	default:
		return nil, false
	}
	if err != nil {
		return nil, false
	}
	return commonLink, true
}

// convertLinkWorkflowEventEventReferenceToURLQuery flattens a WorkflowEvent reference into the
// link's query. The reference names which event in the workflow's history the link points at,
// either by event ID or by the request ID that produced the event, and is encoded as referenceType
// plus eventID or requestID, and eventType.
func convertLinkWorkflowEventEventReferenceToURLQuery(eventRef *commonpb.Link_WorkflowEvent_EventReference) string {
	values := url.Values{}
	values.Set(linkWorkflowEventReferenceTypeKey, eventReferenceType)
	// Zero is not a valid event ID, so an unset ID is omitted rather than sent.
	if eventRef.GetEventId() > 0 {
		values.Set(linkEventIDKey, strconv.FormatInt(eventRef.GetEventId(), 10))
	}
	values.Set(linkEventTypeKey, eventRef.GetEventType().String())
	return values.Encode()
}

func convertURLQueryToLinkWorkflowEventEventReference(queryValues url.Values) (*commonpb.Link_WorkflowEvent_EventReference, error) {
	var err error
	eventRef := &commonpb.Link_WorkflowEvent_EventReference{}
	if eventIDValue := queryValues.Get(linkEventIDKey); eventIDValue != "" {
		eventRef.EventId, err = strconv.ParseInt(eventIDValue, 10, 64)
		if err != nil {
			return nil, err
		}
	}
	eventRef.EventType, err = enumspb.EventTypeFromString(queryValues.Get(linkEventTypeKey))
	if err != nil {
		return nil, err
	}
	return eventRef, nil
}

func convertLinkWorkflowEventRequestIdReferenceToURLQuery(requestIDRef *commonpb.Link_WorkflowEvent_RequestIdReference) string {
	values := url.Values{}
	values.Set(linkWorkflowEventReferenceTypeKey, requestIDReferenceType)
	values.Set(linkRequestIDKey, requestIDRef.GetRequestId())
	values.Set(linkEventTypeKey, requestIDRef.GetEventType().String())
	return values.Encode()
}

func convertURLQueryToLinkWorkflowEventRequestIdReference(queryValues url.Values) (*commonpb.Link_WorkflowEvent_RequestIdReference, error) {
	eventType, err := enumspb.EventTypeFromString(queryValues.Get(linkEventTypeKey))
	if err != nil {
		return nil, err
	}
	return &commonpb.Link_WorkflowEvent_RequestIdReference{
		RequestId: queryValues.Get(linkRequestIDKey),
		EventType: eventType,
	}, nil
}
