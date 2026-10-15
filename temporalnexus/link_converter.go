// Package temporalnexus provides helpers for converting between Temporal
// common.v1 links and Nexus links.
package temporalnexus

import (
	"fmt"
	"net/url"
	"regexp"
	"strconv"
	"strings"

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
	// linkComponentPathKey carries Link_Callback.component_path, repeated once per segment, in
	// order.
	linkComponentPathKey = "componentPath"
)

// linkType describes the Nexus URL form of one common.v1.Link variant. Every variant except
// Callback has the path /namespaces/{namespace}/{collection}/{id}/{runID}, optionally followed by
// a fixed tail segment, so encoding and decoding differ between those variants only by this table.
type linkType struct {
	// name is the proto full name, carried as nexus.Link.Type.
	name string
	// label names the variant in error messages, e.g. "Link_Activity".
	label string
	// template is the path with a %s for each variable segment, e.g. namespace, id and runID.
	template string
	// pathRE matches an escaped path against template, capturing each variable segment.
	pathRE *regexp.Regexp
}

func newLinkType(m proto.Message, template string) linkType {
	segments := make([]any, strings.Count(template, "%s"))
	for i := range segments {
		segments[i] = `([^/]+)`
	}
	desc := m.ProtoReflect().Descriptor()
	return linkType{
		name:     string(desc.FullName()),
		label:    "Link_" + string(desc.Name()),
		template: template,
		pathRE:   regexp.MustCompile("^" + fmt.Sprintf(template, segments...) + "$"),
	}
}

var (
	workflowEventLinkType  = newLinkType(&commonpb.Link_WorkflowEvent{}, "/namespaces/%s/workflows/%s/%s/history")
	workflowLinkType       = newLinkType(&commonpb.Link_Workflow{}, "/namespaces/%s/workflows/%s/%s")
	nexusOperationLinkType = newLinkType(&commonpb.Link_NexusOperation{}, "/namespaces/%s/nexus-operations/%s/%s/details")
	activityLinkType       = newLinkType(&commonpb.Link_Activity{}, "/namespaces/%s/activities/%s/%s/details")
	callbackLinkType       = newLinkType(&commonpb.Link_Callback{}, "/namespaces/%s/%s/%s/%s/callbacks/%s")

	executionTypeSegments = map[enumspb.ExecutionType]string{
		enumspb.EXECUTION_TYPE_WORKFLOW:        "workflows",
		enumspb.EXECUTION_TYPE_ACTIVITY:        "activities",
		enumspb.EXECUTION_TYPE_NEXUS_OPERATION: "nexus-operations",
	}
	// segmentExecutionTypes is the inverse of executionTypeSegments.
	segmentExecutionTypes = func() map[string]enumspb.ExecutionType {
		m := make(map[string]enumspb.ExecutionType, len(executionTypeSegments))
		for t, segment := range executionTypeSegments {
			m[segment] = t
		}
		return m
	}()

	eventReferenceType     = string((&commonpb.Link_WorkflowEvent_EventReference{}).ProtoReflect().Descriptor().Name())
	requestIDReferenceType = string((&commonpb.Link_WorkflowEvent_RequestIdReference{}).ProtoReflect().Descriptor().Name())
)

// encode builds the Nexus link for this variant from its path segments, in template order. Path
// holds the raw segments and RawPath their escaped form, so a segment containing a slash stays a
// single segment.
func (lt linkType) encode(rawQuery string, segments ...string) nexus.Link {
	raw := make([]any, len(segments))
	escaped := make([]any, len(segments))
	for i, segment := range segments {
		raw[i] = segment
		escaped[i] = url.PathEscape(segment)
	}
	return nexus.Link{
		URL: &url.URL{
			Scheme:   urlSchemeTemporalKey,
			Path:     fmt.Sprintf(lt.template, raw...),
			RawPath:  fmt.Sprintf(lt.template, escaped...),
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

// decode decodes a link of any variant except Callback into its namespace, id and runID path
// segments and its URL query. Validation is done by decodeSegments.
func (lt linkType) decode(link nexus.Link) (decodedLink, error) {
	segments, err := lt.decodeSegments(link, 3)
	if err != nil {
		return decodedLink{}, err
	}
	return decodedLink{namespace: segments[0], id: segments[1], runID: segments[2], query: link.URL.Query()}, nil
}

// decodeSegments validates a Nexus link against this variant (the declared type, the scheme, and
// a path matching the template exactly) and returns its n unescaped variable path segments.
func (lt linkType) decodeSegments(link nexus.Link, n int) ([]string, error) {
	if link.Type != lt.name {
		return nil, fmt.Errorf("cannot parse link type %q to %q", link.Type, lt.name)
	}
	if link.URL == nil {
		return nil, lt.errorf("empty URL")
	}
	if link.URL.Scheme != urlSchemeTemporalKey {
		return nil, lt.errorf("invalid scheme: %s", link.URL.Scheme)
	}
	matches := lt.pathRE.FindStringSubmatch(link.URL.EscapedPath())
	if len(matches) != n+1 {
		return nil, lt.errorf("malformed URL path")
	}
	segments := matches[1:]
	for i, match := range segments {
		segment, err := url.PathUnescape(match)
		if err != nil {
			return nil, lt.errorf("%w", err)
		}
		segments[i] = segment
	}
	return segments, nil
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
	return workflowEventLinkType.encode(rawQuery, we.GetNamespace(), we.GetWorkflowId(), we.GetRunId())
}

// ConvertLinkWorkflowToNexusLink converts a Link_Workflow type to Nexus Link.
//
// NOTE: Experimental
func ConvertLinkWorkflowToNexusLink(w *commonpb.Link_Workflow) nexus.Link {
	var rawQuery string
	if w.GetReason() != "" {
		rawQuery = url.Values{linkReasonKey: {w.GetReason()}}.Encode()
	}
	return workflowLinkType.encode(rawQuery, w.GetNamespace(), w.GetWorkflowId(), w.GetRunId())
}

// ConvertLinkNexusOperationToNexusLink converts a Link_NexusOperation type to Nexus Link.
//
// NOTE: Experimental
func ConvertLinkNexusOperationToNexusLink(no *commonpb.Link_NexusOperation) nexus.Link {
	return nexusOperationLinkType.encode("", no.GetNamespace(), no.GetOperationId(), no.GetRunId())
}

// ConvertLinkActivityToNexusLink converts a Link_Activity type to Nexus Link.
//
// NOTE: Experimental
func ConvertLinkActivityToNexusLink(a *commonpb.Link_Activity) nexus.Link {
	return activityLinkType.encode("", a.GetNamespace(), a.GetActivityId(), a.GetRunId())
}

// ConvertLinkCallbackToNexusLink converts a Link_Callback type to Nexus Link.
//
// NOTE: Experimental
func ConvertLinkCallbackToNexusLink(cb *commonpb.Link_Callback) (nexus.Link, error) {
	execution := cb.GetExecution()
	executionType, ok := executionTypeSegments[execution.GetType()]
	if !ok {
		return nexus.Link{}, fmt.Errorf(
			"failed to convert Link_Callback to link: unsupported execution type: %s",
			execution.GetType(),
		)
	}

	return callbackLinkType.encode(
		url.Values{linkComponentPathKey: cb.GetComponentPath()}.Encode(),
		cb.GetNamespace(), executionType, execution.GetBusinessId(), execution.GetRunId(), cb.GetRequestId(),
	), nil
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
	case *commonpb.Link_Callback_:
		if v.Callback == nil {
			return nil, false
		}
		var err error
		if nexusLink, err = ConvertLinkCallbackToNexusLink(v.Callback); err != nil {
			return nil, false
		}
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

// ConvertNexusLinkToLinkCallback converts a Nexus Link to Link_Callback.
//
// NOTE: Experimental
func ConvertNexusLinkToLinkCallback(link nexus.Link) (*commonpb.Link_Callback, error) {
	segments, err := callbackLinkType.decodeSegments(link, 5)
	if err != nil {
		return nil, err
	}
	executionType, ok := segmentExecutionTypes[segments[1]]
	if !ok {
		return nil, callbackLinkType.errorf("unsupported execution type: %q", segments[1])
	}
	return &commonpb.Link_Callback{
		Namespace: segments[0],
		Execution: &commonpb.Execution{
			Type:       executionType,
			BusinessId: segments[2],
			RunId:      segments[3],
		},
		RequestId:     segments[4],
		ComponentPath: link.URL.Query()[linkComponentPathKey],
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
	case callbackLinkType.name:
		var cb *commonpb.Link_Callback
		cb, err = ConvertNexusLinkToLinkCallback(nexusLink)
		commonLink = &commonpb.Link{Variant: &commonpb.Link_Callback_{Callback: cb}}
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
