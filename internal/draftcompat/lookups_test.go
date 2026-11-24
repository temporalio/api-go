package draftcompat

import (
	"testing"

	"github.com/stretchr/testify/require"
	"go.temporal.io/api/operatorservice/v1"
	protometapb "go.temporal.io/api/protometa/v1"
	"go.temporal.io/api/workflowservice/v1"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"
)

// methodLookups holds the generated draft method lookup of every api_next service.
var methodLookups = map[protoreflect.FullName]func(string) (string, bool, bool){
	"temporal.api.operatorservice.v1.OperatorService": operatorservice.OperatorServiceDraftMethod,
	"temporal.api.workflowservice.v1.WorkflowService": workflowservice.WorkflowServiceDraftMethod,
}

// TestDraftMethodLookups checks every generated method lookup against the descriptors:
// the feature comes from draft_method or draft_service, and request and response report
// whether the endpoint types reach draft contents.
func TestDraftMethodLookups(t *testing.T) {
	affected := affectedMessages(nextMessages())
	services := map[protoreflect.FullName]bool{}
	rangeAPIFiles(protoregistry.GlobalFiles, func(file protoreflect.FileDescriptor) {
		for i := range file.Services().Len() {
			service := file.Services().Get(i)
			services[service.FullName()] = true
			lookup, ok := methodLookups[service.FullName()]
			if !ok {
				t.Errorf("%s has no generated draft method lookup in methodLookups", service.FullName())
				continue
			}
			serviceFeature := draftOption(service, protometapb.E_DraftService)
			for j := range service.Methods().Len() {
				method := service.Methods().Get(j)
				wantFeature := serviceFeature
				if wantFeature == "" {
					wantFeature = draftOption(method, protometapb.E_DraftMethod)
				}
				fullMethod := "/" + string(service.FullName()) + "/" + string(method.Name())
				feature, request, response := lookup(fullMethod)
				require.Equal(t, wantFeature, feature, "%s feature", fullMethod)
				require.Equal(t, affected[method.Input().FullName()], request, "%s request", fullMethod)
				require.Equal(t, affected[method.Output().FullName()], response, "%s response", fullMethod)
			}
			feature, request, response := lookup("/" + string(service.FullName()) + "/NoSuchMethod")
			require.Equal(t, "", feature)
			require.False(t, request || response)
		}
	})
	for name := range methodLookups {
		require.True(t, services[name], "methodLookups lists %s, which api_next does not declare", name)
	}
}

// TestHasOmitDraftFields checks the generated type lookup against descriptor reachability.
func TestHasOmitDraftFields(t *testing.T) {
	messages := nextMessages()
	affected := affectedMessages(messages)
	for _, message := range messages {
		require.Equal(t, affected[message.FullName()], protometapb.HasOmitDraftFields(string(message.FullName())), message.FullName())
	}
	require.False(t, protometapb.HasOmitDraftFields("google.protobuf.Any"))
}
