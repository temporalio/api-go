// The MIT License
//
// Copyright (c) 2022 Temporal Technologies Inc.  All rights reserved.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

// Code generated by protoc-gen-go-helpers. DO NOT EDIT.
package enums

import (
	"fmt"
)

var (
	DeploymentReachability_shorthandValue = map[string]int32{
		"Unspecified":         0,
		"Reachable":           1,
		"ClosedWorkflowsOnly": 2,
		"Unreachable":         3,
	}
)

// DeploymentReachabilityFromString parses a DeploymentReachability value from  either the protojson
// canonical SCREAMING_CASE enum or the traditional temporal PascalCase enum to DeploymentReachability
func DeploymentReachabilityFromString(s string) (DeploymentReachability, error) {
	if v, ok := DeploymentReachability_value[s]; ok {
		return DeploymentReachability(v), nil
	} else if v, ok := DeploymentReachability_shorthandValue[s]; ok {
		return DeploymentReachability(v), nil
	}
	return DeploymentReachability(0), fmt.Errorf("%s is not a valid DeploymentReachability", s)
}

var (
	VersionDrainageStatus_shorthandValue = map[string]int32{
		"Unspecified": 0,
		"Draining":    1,
		"Drained":     2,
	}
)

// VersionDrainageStatusFromString parses a VersionDrainageStatus value from  either the protojson
// canonical SCREAMING_CASE enum or the traditional temporal PascalCase enum to VersionDrainageStatus
func VersionDrainageStatusFromString(s string) (VersionDrainageStatus, error) {
	if v, ok := VersionDrainageStatus_value[s]; ok {
		return VersionDrainageStatus(v), nil
	} else if v, ok := VersionDrainageStatus_shorthandValue[s]; ok {
		return VersionDrainageStatus(v), nil
	}
	return VersionDrainageStatus(0), fmt.Errorf("%s is not a valid VersionDrainageStatus", s)
}

var (
	VersionRoutingState_shorthandValue = map[string]int32{
		"Unspecified": 0,
		"None":        1,
		"Current":     2,
		"Ramping":     3,
	}
)

// VersionRoutingStateFromString parses a VersionRoutingState value from  either the protojson
// canonical SCREAMING_CASE enum or the traditional temporal PascalCase enum to VersionRoutingState
func VersionRoutingStateFromString(s string) (VersionRoutingState, error) {
	if v, ok := VersionRoutingState_value[s]; ok {
		return VersionRoutingState(v), nil
	} else if v, ok := VersionRoutingState_shorthandValue[s]; ok {
		return VersionRoutingState(v), nil
	}
	return VersionRoutingState(0), fmt.Errorf("%s is not a valid VersionRoutingState", s)
}

var (
	WorkflowVersioningMode_shorthandValue = map[string]int32{
		"Unspecified":         0,
		"Unversioned":         1,
		"VersioningBehaviors": 2,
	}
)

// WorkflowVersioningModeFromString parses a WorkflowVersioningMode value from  either the protojson
// canonical SCREAMING_CASE enum or the traditional temporal PascalCase enum to WorkflowVersioningMode
func WorkflowVersioningModeFromString(s string) (WorkflowVersioningMode, error) {
	if v, ok := WorkflowVersioningMode_value[s]; ok {
		return WorkflowVersioningMode(v), nil
	} else if v, ok := WorkflowVersioningMode_shorthandValue[s]; ok {
		return WorkflowVersioningMode(v), nil
	}
	return WorkflowVersioningMode(0), fmt.Errorf("%s is not a valid WorkflowVersioningMode", s)
}
