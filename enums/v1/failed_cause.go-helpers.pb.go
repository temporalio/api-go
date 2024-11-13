// Code generated by protoc-gen-go-helpers. DO NOT EDIT.
package enums

import (
	"fmt"
)

var (
	WorkflowTaskFailedCause_shorthandValue = map[string]int32{
		"Unspecified":                                         0,
		"UnhandledCommand":                                    1,
		"BadScheduleActivityAttributes":                       2,
		"BadRequestCancelActivityAttributes":                  3,
		"BadStartTimerAttributes":                             4,
		"BadCancelTimerAttributes":                            5,
		"BadRecordMarkerAttributes":                           6,
		"BadCompleteWorkflowExecutionAttributes":              7,
		"BadFailWorkflowExecutionAttributes":                  8,
		"BadCancelWorkflowExecutionAttributes":                9,
		"BadRequestCancelExternalWorkflowExecutionAttributes": 10,
		"BadContinueAsNewAttributes":                          11,
		"StartTimerDuplicateId":                               12,
		"ResetStickyTaskQueue":                                13,
		"WorkflowWorkerUnhandledFailure":                      14,
		"BadSignalWorkflowExecutionAttributes":                15,
		"BadStartChildExecutionAttributes":                    16,
		"ForceCloseCommand":                                   17,
		"FailoverCloseCommand":                                18,
		"BadSignalInputSize":                                  19,
		"ResetWorkflow":                                       20,
		"BadBinary":                                           21,
		"ScheduleActivityDuplicateId":                         22,
		"BadSearchAttributes":                                 23,
		"NonDeterministicError":                               24,
		"BadModifyWorkflowPropertiesAttributes":               25,
		"PendingChildWorkflowsLimitExceeded":                  26,
		"PendingActivitiesLimitExceeded":                      27,
		"PendingSignalsLimitExceeded":                         28,
		"PendingRequestCancelLimitExceeded":                   29,
		"BadUpdateWorkflowExecutionMessage":                   30,
		"UnhandledUpdate":                                     31,
		"BadScheduleNexusOperationAttributes":                 32,
		"PendingNexusOperationsLimitExceeded":                 33,
		"BadRequestCancelNexusOperationAttributes":            34,
		"FeatureDisabled":                                     35,
	}
)

// WorkflowTaskFailedCauseFromString parses a WorkflowTaskFailedCause value from  either the protojson
// canonical SCREAMING_CASE enum or the traditional temporal PascalCase enum to WorkflowTaskFailedCause
func WorkflowTaskFailedCauseFromString(s string) (WorkflowTaskFailedCause, error) {
	if v, ok := WorkflowTaskFailedCause_value[s]; ok {
		return WorkflowTaskFailedCause(v), nil
	} else if v, ok := WorkflowTaskFailedCause_shorthandValue[s]; ok {
		return WorkflowTaskFailedCause(v), nil
	}
	return WorkflowTaskFailedCause(0), fmt.Errorf("%s is not a valid WorkflowTaskFailedCause", s)
}

var (
	StartChildWorkflowExecutionFailedCause_shorthandValue = map[string]int32{
		"Unspecified":           0,
		"WorkflowAlreadyExists": 1,
		"NamespaceNotFound":     2,
	}
)

// StartChildWorkflowExecutionFailedCauseFromString parses a StartChildWorkflowExecutionFailedCause value from  either the protojson
// canonical SCREAMING_CASE enum or the traditional temporal PascalCase enum to StartChildWorkflowExecutionFailedCause
func StartChildWorkflowExecutionFailedCauseFromString(s string) (StartChildWorkflowExecutionFailedCause, error) {
	if v, ok := StartChildWorkflowExecutionFailedCause_value[s]; ok {
		return StartChildWorkflowExecutionFailedCause(v), nil
	} else if v, ok := StartChildWorkflowExecutionFailedCause_shorthandValue[s]; ok {
		return StartChildWorkflowExecutionFailedCause(v), nil
	}
	return StartChildWorkflowExecutionFailedCause(0), fmt.Errorf("%s is not a valid StartChildWorkflowExecutionFailedCause", s)
}

var (
	CancelExternalWorkflowExecutionFailedCause_shorthandValue = map[string]int32{
		"Unspecified":                       0,
		"ExternalWorkflowExecutionNotFound": 1,
		"NamespaceNotFound":                 2,
	}
)

// CancelExternalWorkflowExecutionFailedCauseFromString parses a CancelExternalWorkflowExecutionFailedCause value from  either the protojson
// canonical SCREAMING_CASE enum or the traditional temporal PascalCase enum to CancelExternalWorkflowExecutionFailedCause
func CancelExternalWorkflowExecutionFailedCauseFromString(s string) (CancelExternalWorkflowExecutionFailedCause, error) {
	if v, ok := CancelExternalWorkflowExecutionFailedCause_value[s]; ok {
		return CancelExternalWorkflowExecutionFailedCause(v), nil
	} else if v, ok := CancelExternalWorkflowExecutionFailedCause_shorthandValue[s]; ok {
		return CancelExternalWorkflowExecutionFailedCause(v), nil
	}
	return CancelExternalWorkflowExecutionFailedCause(0), fmt.Errorf("%s is not a valid CancelExternalWorkflowExecutionFailedCause", s)
}

var (
	SignalExternalWorkflowExecutionFailedCause_shorthandValue = map[string]int32{
		"Unspecified":                       0,
		"ExternalWorkflowExecutionNotFound": 1,
		"NamespaceNotFound":                 2,
		"SignalCountLimitExceeded":          3,
	}
)

// SignalExternalWorkflowExecutionFailedCauseFromString parses a SignalExternalWorkflowExecutionFailedCause value from  either the protojson
// canonical SCREAMING_CASE enum or the traditional temporal PascalCase enum to SignalExternalWorkflowExecutionFailedCause
func SignalExternalWorkflowExecutionFailedCauseFromString(s string) (SignalExternalWorkflowExecutionFailedCause, error) {
	if v, ok := SignalExternalWorkflowExecutionFailedCause_value[s]; ok {
		return SignalExternalWorkflowExecutionFailedCause(v), nil
	} else if v, ok := SignalExternalWorkflowExecutionFailedCause_shorthandValue[s]; ok {
		return SignalExternalWorkflowExecutionFailedCause(v), nil
	}
	return SignalExternalWorkflowExecutionFailedCause(0), fmt.Errorf("%s is not a valid SignalExternalWorkflowExecutionFailedCause", s)
}

var (
	ResourceExhaustedCause_shorthandValue = map[string]int32{
		"Unspecified":             0,
		"RpsLimit":                1,
		"ConcurrentLimit":         2,
		"SystemOverloaded":        3,
		"PersistenceLimit":        4,
		"BusyWorkflow":            5,
		"ApsLimit":                6,
		"PersistenceStorageLimit": 7,
		"CircuitBreakerOpen":      8,
	}
)

// ResourceExhaustedCauseFromString parses a ResourceExhaustedCause value from  either the protojson
// canonical SCREAMING_CASE enum or the traditional temporal PascalCase enum to ResourceExhaustedCause
func ResourceExhaustedCauseFromString(s string) (ResourceExhaustedCause, error) {
	if v, ok := ResourceExhaustedCause_value[s]; ok {
		return ResourceExhaustedCause(v), nil
	} else if v, ok := ResourceExhaustedCause_shorthandValue[s]; ok {
		return ResourceExhaustedCause(v), nil
	}
	return ResourceExhaustedCause(0), fmt.Errorf("%s is not a valid ResourceExhaustedCause", s)
}

var (
	ResourceExhaustedScope_shorthandValue = map[string]int32{
		"Unspecified": 0,
		"Namespace":   1,
		"System":      2,
	}
)

// ResourceExhaustedScopeFromString parses a ResourceExhaustedScope value from  either the protojson
// canonical SCREAMING_CASE enum or the traditional temporal PascalCase enum to ResourceExhaustedScope
func ResourceExhaustedScopeFromString(s string) (ResourceExhaustedScope, error) {
	if v, ok := ResourceExhaustedScope_value[s]; ok {
		return ResourceExhaustedScope(v), nil
	} else if v, ok := ResourceExhaustedScope_shorthandValue[s]; ok {
		return ResourceExhaustedScope(v), nil
	}
	return ResourceExhaustedScope(0), fmt.Errorf("%s is not a valid ResourceExhaustedScope", s)
}
