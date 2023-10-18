package temporalproto_test

import (
	"testing"

	commonpb "go.temporal.io/api/common/v1"
	"go.temporal.io/api/temporalproto"
	workflowpb "go.temporal.io/api/workflow/v1"
)

const myUUID = "deb7b204-b384-4fde-85c6-e5a56c42336a"

type noProtoNoCry struct {
	A float64
	B string
	C []bool
}

type canHazProto struct {
	A float64
	B *commonpb.WorkflowExecution
}

func TestDeepEqual_Equal(t *testing.T) {
	for _, tc := range []struct {
		Name string
		A    any
		B    any
	}{{
		Name: "Shallow proto object",
		A: &commonpb.WorkflowExecution{
			WorkflowId: "some random workflow ID",
			RunId:      myUUID,
		},
		B: &commonpb.WorkflowExecution{
			WorkflowId: "some random workflow ID",
			RunId:      myUUID,
		},
	}, {
		Name: "Struct containing proto",
		A: canHazProto{
			A: 12,
			B: &commonpb.WorkflowExecution{
				WorkflowId: "some random workflow ID",
				RunId:      myUUID,
			},
		},
		B: canHazProto{
			A: 12,
			B: &commonpb.WorkflowExecution{
				WorkflowId: "some random workflow ID",
				RunId:      myUUID,
			},
		},
	}, {
		Name: "Simple struct",
		A: noProtoNoCry{
			A: 12.34,
			B: "foo",
			C: []bool{true, false},
		},
		B: noProtoNoCry{
			A: 12.34,
			B: "foo",
			C: []bool{true, false},
		},
	}, {
		Name: "Nested proto struct",
		A: &workflowpb.WorkflowExecutionInfo{
			Execution: &commonpb.WorkflowExecution{
				WorkflowId: "some random workflow ID",
				RunId:      myUUID,
			},
		},
		B: &workflowpb.WorkflowExecutionInfo{
			Execution: &commonpb.WorkflowExecution{
				WorkflowId: "some random workflow ID",
				RunId:      myUUID,
			},
		},
	}, {
		Name: "Primitive bool",
		A:    true,
		B:    true,
	}} {
		if !temporalproto.DeepEqual(tc.A, tc.B) {
			t.Errorf("%s: expected equality", tc.Name)
		}
	}
}

func TestDeepProtoEqual_NotEqual(t *testing.T) {
	objects := []any{
		&commonpb.WorkflowExecution{
			WorkflowId: "some random workflow ID",
			RunId:      myUUID,
		}, canHazProto{
			A: 12,
			B: &commonpb.WorkflowExecution{
				WorkflowId: "some random workflow ID",
				RunId:      myUUID,
			},
		}, noProtoNoCry{
			A: 12.34,
			B: "foo",
			C: []bool{true, false},
		}, &workflowpb.WorkflowExecutionInfo{
			Execution: &commonpb.WorkflowExecution{
				WorkflowId: "some random workflow ID",
				RunId:      myUUID,
			},
		},
		true,
		12.34,
	}
	for i := range objects {
		for j := range objects {
			if i == j {
				continue
			}

			if temporalproto.DeepEqual(objects[i], objects[j]) {
				t.Errorf("Values should not be equal: %v, %v", objects[i], objects[j])
			}
		}
	}

}
