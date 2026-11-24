package draftcompat

import (
	"testing"

	upstream "github.com/nexus-rpc/nexus-proto-annotations/go/nexusannotations/v1"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/reflect/protoregistry"

	local "go.temporal.io/api/nexusannotations/v1"
	_ "go.temporal.io/api/workflowservice/v1"
)

func TestNexusAnnotationsUseUpstreamBinding(t *testing.T) {
	var _ *upstream.OperationOptions = new(local.OperationOptions)
	var _ *upstream.ServiceOptions = new(local.ServiceOptions)
	registered, err := protoregistry.GlobalFiles.FindFileByPath("nexusannotations/v1/options.proto")
	require.NoError(t, err)
	require.Same(t, upstream.File_nexusannotations_v1_options_proto, registered)
	require.Same(t, registered, local.File_nexusannotations_v1_options_proto)
	require.Same(t, upstream.E_Service, local.E_Service)
	require.Same(t, upstream.E_Operation, local.E_Operation)
}
