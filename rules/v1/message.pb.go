// The MIT License
//
// Copyright (c) 2025 Temporal Technologies Inc.  All rights reserved.
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

// Code generated by protoc-gen-go. DO NOT EDIT.
// plugins:
// 	protoc-gen-go
// 	protoc
// source: temporal/api/rules/v1/message.proto

package rules

import (
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type WorkflowRuleAction struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Supported actions.
	//
	// Types that are valid to be assigned to Variant:
	//
	//	*WorkflowRuleAction_ActivityPause
	Variant       isWorkflowRuleAction_Variant `protobuf_oneof:"variant"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *WorkflowRuleAction) Reset() {
	*x = WorkflowRuleAction{}
	mi := &file_temporal_api_rules_v1_message_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *WorkflowRuleAction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WorkflowRuleAction) ProtoMessage() {}

func (x *WorkflowRuleAction) ProtoReflect() protoreflect.Message {
	mi := &file_temporal_api_rules_v1_message_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WorkflowRuleAction.ProtoReflect.Descriptor instead.
func (*WorkflowRuleAction) Descriptor() ([]byte, []int) {
	return file_temporal_api_rules_v1_message_proto_rawDescGZIP(), []int{0}
}

func (x *WorkflowRuleAction) GetVariant() isWorkflowRuleAction_Variant {
	if x != nil {
		return x.Variant
	}
	return nil
}

func (x *WorkflowRuleAction) GetActivityPause() *WorkflowRuleAction_ActionActivityPause {
	if x != nil {
		if x, ok := x.Variant.(*WorkflowRuleAction_ActivityPause); ok {
			return x.ActivityPause
		}
	}
	return nil
}

type isWorkflowRuleAction_Variant interface {
	isWorkflowRuleAction_Variant()
}

type WorkflowRuleAction_ActivityPause struct {
	ActivityPause *WorkflowRuleAction_ActionActivityPause `protobuf:"bytes,1,opt,name=activity_pause,json=activityPause,proto3,oneof"`
}

func (*WorkflowRuleAction_ActivityPause) isWorkflowRuleAction_Variant() {}

type WorkflowRuleSpec struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The id of the new workflow rule. Must be unique within the namespace.
	// Can be set by the user, and can have business meaning.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Specifies how the rule should be triggered and evaluated.
	// Currently, only "activity start" type is supported.
	//
	// Types that are valid to be assigned to Trigger:
	//
	//	*WorkflowRuleSpec_ActivityStart
	Trigger isWorkflowRuleSpec_Trigger `protobuf_oneof:"trigger"`
	// Restricted Visibility query.
	// This query is used to filter workflows in this namespace to which this rule should apply.
	// It is applied to any running workflow each time a triggering event occurs, before the trigger predicate is evaluated.
	// The following workflow attributes are supported:
	// - WorkflowType
	// - WorkflowId
	// - StartTime
	// - ExecutionStatus
	VisibilityQuery string `protobuf:"bytes,3,opt,name=visibility_query,json=visibilityQuery,proto3" json:"visibility_query,omitempty"`
	// WorkflowRuleAction to be taken when the rule is triggered and predicate is matched.
	Actions []*WorkflowRuleAction `protobuf:"bytes,4,rep,name=actions,proto3" json:"actions,omitempty"`
	// Expiration time of the rule. After this time, the rule will be deleted.
	// Can be empty if the rule should never expire.
	ExpirationTime *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=expiration_time,json=expirationTime,proto3" json:"expiration_time,omitempty"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *WorkflowRuleSpec) Reset() {
	*x = WorkflowRuleSpec{}
	mi := &file_temporal_api_rules_v1_message_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *WorkflowRuleSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WorkflowRuleSpec) ProtoMessage() {}

func (x *WorkflowRuleSpec) ProtoReflect() protoreflect.Message {
	mi := &file_temporal_api_rules_v1_message_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WorkflowRuleSpec.ProtoReflect.Descriptor instead.
func (*WorkflowRuleSpec) Descriptor() ([]byte, []int) {
	return file_temporal_api_rules_v1_message_proto_rawDescGZIP(), []int{1}
}

func (x *WorkflowRuleSpec) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *WorkflowRuleSpec) GetTrigger() isWorkflowRuleSpec_Trigger {
	if x != nil {
		return x.Trigger
	}
	return nil
}

func (x *WorkflowRuleSpec) GetActivityStart() *WorkflowRuleSpec_ActivityStartingTrigger {
	if x != nil {
		if x, ok := x.Trigger.(*WorkflowRuleSpec_ActivityStart); ok {
			return x.ActivityStart
		}
	}
	return nil
}

func (x *WorkflowRuleSpec) GetVisibilityQuery() string {
	if x != nil {
		return x.VisibilityQuery
	}
	return ""
}

func (x *WorkflowRuleSpec) GetActions() []*WorkflowRuleAction {
	if x != nil {
		return x.Actions
	}
	return nil
}

func (x *WorkflowRuleSpec) GetExpirationTime() *timestamppb.Timestamp {
	if x != nil {
		return x.ExpirationTime
	}
	return nil
}

type isWorkflowRuleSpec_Trigger interface {
	isWorkflowRuleSpec_Trigger()
}

type WorkflowRuleSpec_ActivityStart struct {
	ActivityStart *WorkflowRuleSpec_ActivityStartingTrigger `protobuf:"bytes,2,opt,name=activity_start,json=activityStart,proto3,oneof"`
}

func (*WorkflowRuleSpec_ActivityStart) isWorkflowRuleSpec_Trigger() {}

// WorkflowRule describes a rule that can be applied to any workflow in this namespace.
type WorkflowRule struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Rule creation time.
	CreateTime *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	// Rule specification
	Spec          *WorkflowRuleSpec `protobuf:"bytes,2,opt,name=spec,proto3" json:"spec,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *WorkflowRule) Reset() {
	*x = WorkflowRule{}
	mi := &file_temporal_api_rules_v1_message_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *WorkflowRule) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WorkflowRule) ProtoMessage() {}

func (x *WorkflowRule) ProtoReflect() protoreflect.Message {
	mi := &file_temporal_api_rules_v1_message_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WorkflowRule.ProtoReflect.Descriptor instead.
func (*WorkflowRule) Descriptor() ([]byte, []int) {
	return file_temporal_api_rules_v1_message_proto_rawDescGZIP(), []int{2}
}

func (x *WorkflowRule) GetCreateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

func (x *WorkflowRule) GetSpec() *WorkflowRuleSpec {
	if x != nil {
		return x.Spec
	}
	return nil
}

type WorkflowRuleAction_ActionActivityPause struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *WorkflowRuleAction_ActionActivityPause) Reset() {
	*x = WorkflowRuleAction_ActionActivityPause{}
	mi := &file_temporal_api_rules_v1_message_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *WorkflowRuleAction_ActionActivityPause) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WorkflowRuleAction_ActionActivityPause) ProtoMessage() {}

func (x *WorkflowRuleAction_ActionActivityPause) ProtoReflect() protoreflect.Message {
	mi := &file_temporal_api_rules_v1_message_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WorkflowRuleAction_ActionActivityPause.ProtoReflect.Descriptor instead.
func (*WorkflowRuleAction_ActionActivityPause) Descriptor() ([]byte, []int) {
	return file_temporal_api_rules_v1_message_proto_rawDescGZIP(), []int{0, 0}
}

// Activity trigger will be triggered when an activity is about to start.
type WorkflowRuleSpec_ActivityStartingTrigger struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Activity predicate is a SQL-like string filter parameter.
	// It is used to match against workflow data.
	// The following activity attributes are supported as part of the predicate:
	// - ActivityType: An Activity Type is the mapping of a name to an Activity Definition..
	// - ActivityId: The ID of the activity.
	// - ActivityAttempt: The number attempts of the activity.
	// - BackoffInterval: The current amount of time between scheduled attempts of the activity.
	// - ActivityStatus: The status of the activity. Can be one of "Scheduled", "Started", "Paused".
	// - TaskQueue: The name of the task queue the workflow specified that the activity should run on.
	// Activity predicate support the following operators:
	//   - =, !=, >, >=, <, <=
	//   - AND, OR, ()
	//   - BETWEEN ... AND
	//     STARTS_WITH
	Predicate     string `protobuf:"bytes,1,opt,name=predicate,proto3" json:"predicate,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *WorkflowRuleSpec_ActivityStartingTrigger) Reset() {
	*x = WorkflowRuleSpec_ActivityStartingTrigger{}
	mi := &file_temporal_api_rules_v1_message_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *WorkflowRuleSpec_ActivityStartingTrigger) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WorkflowRuleSpec_ActivityStartingTrigger) ProtoMessage() {}

func (x *WorkflowRuleSpec_ActivityStartingTrigger) ProtoReflect() protoreflect.Message {
	mi := &file_temporal_api_rules_v1_message_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WorkflowRuleSpec_ActivityStartingTrigger.ProtoReflect.Descriptor instead.
func (*WorkflowRuleSpec_ActivityStartingTrigger) Descriptor() ([]byte, []int) {
	return file_temporal_api_rules_v1_message_proto_rawDescGZIP(), []int{1, 0}
}

func (x *WorkflowRuleSpec_ActivityStartingTrigger) GetPredicate() string {
	if x != nil {
		return x.Predicate
	}
	return ""
}

var File_temporal_api_rules_v1_message_proto protoreflect.FileDescriptor

const file_temporal_api_rules_v1_message_proto_rawDesc = "" +
	"\n" +
	"#temporal/api/rules/v1/message.proto\x12\x15temporal.api.rules.v1\x1a\x1fgoogle/protobuf/timestamp.proto\"\x9e\x01\n" +
	"\x12WorkflowRuleAction\x12f\n" +
	"\x0eactivity_pause\x18\x01 \x01(\v2=.temporal.api.rules.v1.WorkflowRuleAction.ActionActivityPauseH\x00R\ractivityPause\x1a\x15\n" +
	"\x13ActionActivityPauseB\t\n" +
	"\avariant\"\x85\x03\n" +
	"\x10WorkflowRuleSpec\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x12h\n" +
	"\x0eactivity_start\x18\x02 \x01(\v2?.temporal.api.rules.v1.WorkflowRuleSpec.ActivityStartingTriggerH\x00R\ractivityStart\x12)\n" +
	"\x10visibility_query\x18\x03 \x01(\tR\x0fvisibilityQuery\x12C\n" +
	"\aactions\x18\x04 \x03(\v2).temporal.api.rules.v1.WorkflowRuleActionR\aactions\x12C\n" +
	"\x0fexpiration_time\x18\x05 \x01(\v2\x1a.google.protobuf.TimestampR\x0eexpirationTime\x1a7\n" +
	"\x17ActivityStartingTrigger\x12\x1c\n" +
	"\tpredicate\x18\x01 \x01(\tR\tpredicateB\t\n" +
	"\atrigger\"\x88\x01\n" +
	"\fWorkflowRule\x12;\n" +
	"\vcreate_time\x18\x01 \x01(\v2\x1a.google.protobuf.TimestampR\n" +
	"createTime\x12;\n" +
	"\x04spec\x18\x02 \x01(\v2'.temporal.api.rules.v1.WorkflowRuleSpecR\x04specB\x84\x01\n" +
	"\x18io.temporal.api.rules.v1B\fMessageProtoP\x01Z!go.temporal.io/api/rules/v1;rules\xaa\x02\x17Temporalio.Api.Rules.V1\xea\x02\x1aTemporalio::Api::Rules::V1b\x06proto3"

var (
	file_temporal_api_rules_v1_message_proto_rawDescOnce sync.Once
	file_temporal_api_rules_v1_message_proto_rawDescData []byte
)

func file_temporal_api_rules_v1_message_proto_rawDescGZIP() []byte {
	file_temporal_api_rules_v1_message_proto_rawDescOnce.Do(func() {
		file_temporal_api_rules_v1_message_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_temporal_api_rules_v1_message_proto_rawDesc), len(file_temporal_api_rules_v1_message_proto_rawDesc)))
	})
	return file_temporal_api_rules_v1_message_proto_rawDescData
}

var file_temporal_api_rules_v1_message_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_temporal_api_rules_v1_message_proto_goTypes = []any{
	(*WorkflowRuleAction)(nil),                       // 0: temporal.api.rules.v1.WorkflowRuleAction
	(*WorkflowRuleSpec)(nil),                         // 1: temporal.api.rules.v1.WorkflowRuleSpec
	(*WorkflowRule)(nil),                             // 2: temporal.api.rules.v1.WorkflowRule
	(*WorkflowRuleAction_ActionActivityPause)(nil),   // 3: temporal.api.rules.v1.WorkflowRuleAction.ActionActivityPause
	(*WorkflowRuleSpec_ActivityStartingTrigger)(nil), // 4: temporal.api.rules.v1.WorkflowRuleSpec.ActivityStartingTrigger
	(*timestamppb.Timestamp)(nil),                    // 5: google.protobuf.Timestamp
}
var file_temporal_api_rules_v1_message_proto_depIdxs = []int32{
	3, // 0: temporal.api.rules.v1.WorkflowRuleAction.activity_pause:type_name -> temporal.api.rules.v1.WorkflowRuleAction.ActionActivityPause
	4, // 1: temporal.api.rules.v1.WorkflowRuleSpec.activity_start:type_name -> temporal.api.rules.v1.WorkflowRuleSpec.ActivityStartingTrigger
	0, // 2: temporal.api.rules.v1.WorkflowRuleSpec.actions:type_name -> temporal.api.rules.v1.WorkflowRuleAction
	5, // 3: temporal.api.rules.v1.WorkflowRuleSpec.expiration_time:type_name -> google.protobuf.Timestamp
	5, // 4: temporal.api.rules.v1.WorkflowRule.create_time:type_name -> google.protobuf.Timestamp
	1, // 5: temporal.api.rules.v1.WorkflowRule.spec:type_name -> temporal.api.rules.v1.WorkflowRuleSpec
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_temporal_api_rules_v1_message_proto_init() }
func file_temporal_api_rules_v1_message_proto_init() {
	if File_temporal_api_rules_v1_message_proto != nil {
		return
	}
	file_temporal_api_rules_v1_message_proto_msgTypes[0].OneofWrappers = []any{
		(*WorkflowRuleAction_ActivityPause)(nil),
	}
	file_temporal_api_rules_v1_message_proto_msgTypes[1].OneofWrappers = []any{
		(*WorkflowRuleSpec_ActivityStart)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_temporal_api_rules_v1_message_proto_rawDesc), len(file_temporal_api_rules_v1_message_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_temporal_api_rules_v1_message_proto_goTypes,
		DependencyIndexes: file_temporal_api_rules_v1_message_proto_depIdxs,
		MessageInfos:      file_temporal_api_rules_v1_message_proto_msgTypes,
	}.Build()
	File_temporal_api_rules_v1_message_proto = out.File
	file_temporal_api_rules_v1_message_proto_goTypes = nil
	file_temporal_api_rules_v1_message_proto_depIdxs = nil
}
