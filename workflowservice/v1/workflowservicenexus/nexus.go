// Code generated from proto service definition. DO NOT EDIT.

// Package workflowservicenexus provides a Nexus service handler for the WorkflowService.
package workflowservicenexus

import (
	"github.com/nexus-rpc/sdk-go/nexus"
	workflowservice "go.temporal.io/api/workflowservice/v1"
)

const (
	// WorkflowServiceServiceName is the Nexus service name for the WorkflowService proto service.
	WorkflowServiceServiceName = "temporal.api.workflowservice.v1.WorkflowService"

	// WorkflowServiceSignalWithStartWorkflowExecutionOperationName is the Nexus operation name
	// for the SignalWithStartWorkflowExecution RPC.
	WorkflowServiceSignalWithStartWorkflowExecutionOperationName = "SignalWithStartWorkflowExecution"
)

// WorkflowServiceNexusHandler is the interface that must be implemented to handle Nexus operations
// for the WorkflowService.
type WorkflowServiceNexusHandler interface {
	SignalWithStartWorkflowExecution(name string) nexus.Operation[*workflowservice.SignalWithStartWorkflowExecutionRequest, *workflowservice.SignalWithStartWorkflowExecutionResponse]
}

// NewWorkflowServiceNexusService creates a [nexus.Service] with all exposed operations registered
// from the provided handler.
func NewWorkflowServiceNexusService(handler WorkflowServiceNexusHandler) (*nexus.Service, error) {
	svc := nexus.NewService(WorkflowServiceServiceName)
	if err := svc.Register(
		handler.SignalWithStartWorkflowExecution(WorkflowServiceSignalWithStartWorkflowExecutionOperationName),
	); err != nil {
		return nil, err
	}
	return svc, nil
}
