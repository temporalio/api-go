package draftwire

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"reflect"
	"slices"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	protometapb "go.temporal.io/api/protometa/v1"
	"google.golang.org/protobuf/encoding/prototext"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protodesc"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"
	"google.golang.org/protobuf/types/descriptorpb"
	"google.golang.org/protobuf/types/dynamicpb"
	"pgregory.net/rapid"
)

// draftOmitter is implemented by api_next messages that can reach a draft field or enum value.
type draftOmitter interface {
	OmitDraftFields(enabled func(string) bool) (removedEnumValue bool)
}

// TestAllFilesLinked requires every proto file of the api_next tree to be linked, so the
// tests below see all of it.
func TestAllFilesLinked(t *testing.T) {
	for _, path := range protoFiles(t, "temporal/api_next") {
		if _, err := protoregistry.GlobalFiles.FindFileByPath(path); err != nil {
			t.Errorf("%s is not linked; add its Go package to links_test.go", path)
		}
	}
}

// TestMessageSets requires the stable tree to contain exactly the non-draft
// messages, fields, enums, enum values, and real oneofs in api_next.
func TestMessageSets(t *testing.T) {
	want := collectMessageSets(protoregistry.GlobalFiles, true)
	got := collectMessageSets(stableFiles(t), false)
	require.Len(t, got.messages, len(want.messages), "messages")
	for name, wantFields := range want.messages {
		gotFields, ok := got.messages[name]
		require.True(t, ok, "missing message %s", name)
		require.Equal(t, wantFields, gotFields, "fields of %s", name)
	}
	require.Equal(t, want.enums, got.enums, "enums and enum values")
	require.Equal(t, want.oneofs, got.oneofs, "oneofs and their members")
}

type messageSets struct {
	messages map[protoreflect.FullName]map[protoreflect.Name]protoreflect.FieldNumber
	enums    map[protoreflect.FullName]map[protoreflect.Name]protoreflect.EnumNumber
	oneofs   map[protoreflect.FullName]map[protoreflect.Name]protoreflect.FieldNumber
}

func collectMessageSets(files *protoregistry.Files, project bool) messageSets {
	result := messageSets{
		messages: map[protoreflect.FullName]map[protoreflect.Name]protoreflect.FieldNumber{},
		enums:    map[protoreflect.FullName]map[protoreflect.Name]protoreflect.EnumNumber{},
		oneofs:   map[protoreflect.FullName]map[protoreflect.Name]protoreflect.FieldNumber{},
	}
	addEnums := func(enums protoreflect.EnumDescriptors) {
		for i := range enums.Len() {
			enum := enums.Get(i)
			if project && hasDraftOption(enum, protometapb.E_DraftEnum) {
				continue
			}
			values := map[protoreflect.Name]protoreflect.EnumNumber{}
			for j := range enum.Values().Len() {
				value := enum.Values().Get(j)
				if !project || !hasDraftOption(value, protometapb.E_DraftEnumValue) {
					values[value.Name()] = value.Number()
				}
			}
			result.enums[enum.FullName()] = values
		}
	}

	rangeAPIFiles(files, func(file protoreflect.FileDescriptor) {
		addEnums(file.Enums())
		eachMessage(file.Messages(), func(message protoreflect.MessageDescriptor) {
			if project && draftMessage(message) {
				return
			}
			result.messages[message.FullName()] = collectFields(message.Fields(), project)
			addEnums(message.Enums())
			for i := range message.Oneofs().Len() {
				oneof := message.Oneofs().Get(i)
				if oneof.IsSynthetic() || project && hasDraftOption(oneof, protometapb.E_DraftOneof) {
					continue
				}
				result.oneofs[oneof.FullName()] = collectFields(oneof.Fields(), project)
			}
		})
	})
	return result
}

func collectFields(fields protoreflect.FieldDescriptors, project bool) map[protoreflect.Name]protoreflect.FieldNumber {
	result := map[protoreflect.Name]protoreflect.FieldNumber{}
	for i := range fields.Len() {
		field := fields.Get(i)
		if !project || !draftField(field) {
			result[field.Name()] = field.Number()
		}
	}
	return result
}

func TestMessageSetsDetectsFieldDrift(t *testing.T) {
	field := func(name string, number int32) *descriptorpb.FieldDescriptorProto {
		return &descriptorpb.FieldDescriptorProto{
			Name: proto.String(name), Number: proto.Int32(number),
			Label: descriptorpb.FieldDescriptorProto_LABEL_OPTIONAL.Enum(),
			Type:  descriptorpb.FieldDescriptorProto_TYPE_STRING.Enum(),
		}
	}
	stable := &descriptorpb.FileDescriptorProto{
		Name: proto.String("field_check.proto"), Package: proto.String("temporal.api.fieldcheck.v1"), Syntax: proto.String("proto3"),
		MessageType: []*descriptorpb.DescriptorProto{{Name: proto.String("Message"), Field: []*descriptorpb.FieldDescriptorProto{field("stable", 1)}}},
	}
	next := proto.CloneOf(stable)
	next.Dependency = []string{protometapb.File_temporal_api_next_protometa_v1_draft_proto.Path()}
	msg := next.MessageType[0]
	draft := field("draft", 2)
	draft.Options = &descriptorpb.FieldOptions{}
	proto.SetExtension(draft.Options, protometapb.E_Draft, "preview")
	optional := proto.CloneOf(draft)
	optional.Name, optional.Number = proto.String("preview_optional"), proto.Int32(3)
	optional.Proto3Optional, optional.OneofIndex = proto.Bool(true), proto.Int32(1)
	member := field("preview_member", 4)
	member.OneofIndex = proto.Int32(0)
	oneof := &descriptorpb.OneofDescriptorProto{Name: proto.String("preview"), Options: &descriptorpb.OneofOptions{}}
	proto.SetExtension(oneof.Options, protometapb.E_DraftOneof, "preview")
	msg.OneofDecl = []*descriptorpb.OneofDescriptorProto{oneof, {Name: proto.String("_preview_optional")}}
	msg.Field = append(msg.Field, draft, optional, member)

	collect := func(t *testing.T, file *descriptorpb.FileDescriptorProto, project bool) messageSets {
		t.Helper()
		descriptor, err := protodesc.NewFile(file, protoregistry.GlobalFiles)
		require.NoError(t, err)
		var files protoregistry.Files
		require.NoError(t, files.RegisterFile(descriptor))
		return collectMessageSets(&files, project)
	}
	want := collect(t, next, true)
	require.Equal(t, want, collect(t, stable, false))
	for _, tt := range []struct {
		name   string
		mutate func(*descriptorpb.DescriptorProto)
	}{
		{"ordinary draft field leaked", func(m *descriptorpb.DescriptorProto) { m.Field = append(m.Field, field("draft", 2)) }},
		{"optional draft field leaked", func(m *descriptorpb.DescriptorProto) {
			m.OneofDecl = []*descriptorpb.OneofDescriptorProto{{Name: proto.String("_preview_optional")}}
			leak := field("preview_optional", 3)
			leak.Proto3Optional, leak.OneofIndex = proto.Bool(true), proto.Int32(0)
			m.Field = append(m.Field, leak)
		}},
		{"draft oneof member leaked as ordinary field", func(m *descriptorpb.DescriptorProto) { m.Field = append(m.Field, field("preview_member", 4)) }},
		{"stable field missing", func(m *descriptorpb.DescriptorProto) { m.Field = nil }},
		{"stable field renumbered", func(m *descriptorpb.DescriptorProto) { m.Field[0].Number = proto.Int32(5) }},
	} {
		t.Run(tt.name, func(t *testing.T) {
			broken := proto.CloneOf(stable)
			tt.mutate(broken.MessageType[0])
			require.NotEqual(t, want, collect(t, broken, false))
		})
	}
}

// TestServiceMethodSets requires the stable tree to contain exactly the non-draft
// services and methods in api_next.
func TestServiceMethodSets(t *testing.T) {
	collect := func(files *protoregistry.Files, project bool) map[protoreflect.FullName]map[protoreflect.Name]bool {
		result := map[protoreflect.FullName]map[protoreflect.Name]bool{}
		rangeAPIFiles(files, func(file protoreflect.FileDescriptor) {
			for i := range file.Services().Len() {
				service := file.Services().Get(i)
				if project && hasDraftOption(service, protometapb.E_DraftService) {
					continue
				}
				methods := map[protoreflect.Name]bool{}
				for j := range service.Methods().Len() {
					method := service.Methods().Get(j)
					if !project || !hasDraftOption(method, protometapb.E_DraftMethod) {
						methods[method.Name()] = true
					}
				}
				result[service.FullName()] = methods
			}
		})
		return result
	}
	require.Equal(t, collect(protoregistry.GlobalFiles, true), collect(stableFiles(t), false))
}

// TestOmitDraftFieldsMatchesStable fills random api_next messages. Every affected root is
// checked on its own, and the unaffected roots are sampled. Filling is steered toward
// fields that reach draft contents, so those contents appear at any depth.
//
// For a random subset of the schema's features, OmitDraftFields must match the descriptor
// walk on a copy of the original, including whether it removed a draft enum value, and a
// second call must change nothing. With no feature enabled, OmitDraftFields must leave
// exactly what the stable tree can hold: the stable descriptors decode it without unknown
// fields at any depth, and decoding back gives the same message. With every feature
// enabled, OmitDraftFields must preserve the input and return false.
func TestOmitDraftFieldsMatchesStable(t *testing.T) {
	stable := stableFiles(t)
	unmarshal := proto.UnmarshalOptions{Resolver: dynamicpb.NewTypes(stable)}
	discardUnknown := unmarshal
	discardUnknown.DiscardUnknown = true
	var affected, unaffected []protoreflect.MessageDescriptor
	messages := nextMessages()
	expectedMethods := affectedMessages(messages)
	features := draftFeatures(messages)
	require.NotEmpty(t, features)
	// A feature the schema does not declare must never matter.
	features = append(features, "undeclared-feature")
	filler := &filler{affected: expectedMethods}
	for _, message := range messages {
		name := message.FullName()
		messageType, err := protoregistry.GlobalTypes.FindMessageByName(name)
		require.NoError(t, err)
		value := messageType.New().Interface()
		_, hasMethod := reflect.TypeOf(value).MethodByName("OmitDraftFields")
		wantMethod := expectedMethods[name]
		require.Equal(t, wantMethod, hasMethod, "%s OmitDraftFields method presence", name)
		if wantMethod {
			require.Implements(t, (*draftOmitter)(nil), value, "%s OmitDraftFields signature", name)
		}
		if draftMessage(message) {
			continue
		}
		if wantMethod {
			affected = append(affected, message)
		} else {
			unaffected = append(unaffected, message)
		}
	}
	require.NotEmpty(t, affected)
	require.NotEmpty(t, unaffected)
	check := func(rt *rapid.T, descriptor protoreflect.MessageDescriptor) {
		messageType, err := protoregistry.GlobalTypes.FindMessageByName(descriptor.FullName())
		if err != nil {
			rt.Fatal(err)
		}
		original := messageType.New()
		filler.fill(rt, original, 0)

		enabledFeatures := rapid.SliceOfDistinct(rapid.SampledFrom(features), rapid.ID[string]).Draw(rt, "enabled features")
		enabled := func(feature string) bool { return slices.Contains(enabledFeatures, feature) }
		some := proto.Clone(original.Interface())
		someOmitter, ok := some.(draftOmitter)
		removedSome := ok && someOmitter.OmitDraftFields(enabled)
		wantSome := proto.Clone(original.Interface())
		wantRemovedSome := projectDraft(wantSome.ProtoReflect(), enabled)
		if removedSome != wantRemovedSome {
			rt.Fatalf("OmitDraftFields for %s with %v enabled reported enum removal %t, want %t", descriptor.FullName(), enabledFeatures, removedSome, wantRemovedSome)
		}
		if !proto.Equal(wantSome, some) {
			rt.Fatalf("OmitDraftFields for %s with %v enabled: got %v, want %v", descriptor.FullName(), enabledFeatures, some, wantSome)
		}
		if ok {
			again := proto.Clone(some)
			if again.(draftOmitter).OmitDraftFields(enabled) || !proto.Equal(some, again) {
				rt.Fatalf("OmitDraftFields for %s with %v enabled is not idempotent", descriptor.FullName(), enabledFeatures)
			}
		}

		all := proto.Clone(original.Interface())
		if omitter, ok := all.(draftOmitter); ok && omitter.OmitDraftFields(func(string) bool { return true }) {
			rt.Fatalf("OmitDraftFields with every feature enabled reported removing a draft enum value from %s", descriptor.FullName())
		}
		if !proto.Equal(original.Interface(), all) {
			rt.Fatalf("OmitDraftFields with every feature enabled changed %s", descriptor.FullName())
		}

		none := proto.Clone(original.Interface())
		omitter, ok := none.(draftOmitter)
		removedEnumValue := ok && omitter.OmitDraftFields(func(string) bool { return false })
		projected := proto.Clone(original.Interface())
		wantRemovedEnumValue := projectStableFields(projected.ProtoReflect())
		if removedEnumValue != wantRemovedEnumValue {
			rt.Fatalf("OmitDraftFields for %s reported enum removal %t, want %t", descriptor.FullName(), removedEnumValue, wantRemovedEnumValue)
		}
		data, err := proto.MarshalOptions{Deterministic: true}.Marshal(none)
		if err != nil {
			rt.Fatal(err)
		}
		stableDescriptor, err := stable.FindDescriptorByName(descriptor.FullName())
		if err != nil {
			rt.Fatal(err)
		}
		stableMessage := stableDescriptor.(protoreflect.MessageDescriptor)
		decoded := dynamicpb.NewMessage(stableMessage)
		if err := unmarshal.Unmarshal(data, decoded); err != nil {
			rt.Fatal(err)
		}
		if path := unknownFields(decoded, string(descriptor.FullName())); path != "" {
			rt.Fatalf("stable %s has unknown fields at %s after omitting every draft feature", descriptor.FullName(), path)
		}
		projectedData, err := proto.Marshal(projected)
		if err != nil {
			rt.Fatal(err)
		}
		expected := dynamicpb.NewMessage(stableMessage)
		if err := discardUnknown.Unmarshal(projectedData, expected); err != nil {
			rt.Fatal(err)
		}
		if !proto.Equal(expected, decoded) {
			rt.Fatalf("OmitDraftFields with no features enabled changed stable fields of %s: got %v, want %v", descriptor.FullName(), decoded, expected)
		}
		stableData, err := proto.MarshalOptions{Deterministic: true}.Marshal(decoded)
		if err != nil {
			rt.Fatal(err)
		}
		back := messageType.New().Interface()
		if err := proto.Unmarshal(stableData, back); err != nil {
			rt.Fatal(err)
		}
		if !proto.Equal(none, back) {
			rt.Fatalf("%s changed on its way through the stable descriptors", descriptor.FullName())
		}
	}
	for _, descriptor := range affected {
		t.Run(string(descriptor.FullName()), func(t *testing.T) {
			t.Parallel()
			rapid.Check(t, func(rt *rapid.T) { check(rt, descriptor) })
		})
	}
	t.Run("unaffected", func(t *testing.T) {
		t.Parallel()
		rapid.Check(t, func(rt *rapid.T) { check(rt, rapid.SampledFrom(unaffected).Draw(rt, "message")) })
	})
}

// draftFeatures lists every feature named by a draft annotation reachable from messages.
func draftFeatures(messages []protoreflect.MessageDescriptor) []string {
	seen := map[string]bool{}
	add := func(descriptor protoreflect.Descriptor, option protoreflect.ExtensionType) {
		if feature := draftOption(descriptor, option); feature != "" {
			seen[feature] = true
		}
	}
	addEnum := func(enum protoreflect.EnumDescriptor) {
		if enum == nil {
			return
		}
		add(enum, protometapb.E_DraftEnum)
		for i := range enum.Values().Len() {
			add(enum.Values().Get(i), protometapb.E_DraftEnumValue)
		}
	}
	for _, message := range messages {
		add(message, protometapb.E_DraftMessage)
		for i := range message.Oneofs().Len() {
			add(message.Oneofs().Get(i), protometapb.E_DraftOneof)
		}
		for i := range message.Fields().Len() {
			field := message.Fields().Get(i)
			add(field, protometapb.E_Draft)
			addEnum(field.Enum())
			if field.IsMap() {
				addEnum(field.MapValue().Enum())
			}
		}
	}
	features := make([]string, 0, len(seen))
	for feature := range seen {
		features = append(features, feature)
	}
	slices.Sort(features)
	return features
}

func TestAffectedMessages(t *testing.T) {
	var messages []protoreflect.MessageDescriptor
	eachMessage(draftProjectionFixture(t).Messages(), func(message protoreflect.MessageDescriptor) {
		messages = append(messages, message)
	})
	require.Equal(t, map[protoreflect.FullName]bool{
		"draftwire.fixture.Values":           true,
		"draftwire.fixture.Envelope":         true,
		"draftwire.fixture.List":             true,
		"draftwire.fixture.Map":              true,
		"draftwire.fixture.Choice":           true,
		"draftwire.fixture.CycleA":           true,
		"draftwire.fixture.CycleB":           true,
		"draftwire.fixture.DraftRoot":        true,
		"draftwire.fixture.DraftRoot.Nested": true,
	}, affectedMessages(messages))
}

// affectedMessages follows descriptor references backward from messages with draft contents.
func affectedMessages(messages []protoreflect.MessageDescriptor) map[protoreflect.FullName]bool {
	seen := map[protoreflect.FullName]bool{}
	parents := map[protoreflect.FullName][]protoreflect.FullName{}
	affected := map[protoreflect.FullName]bool{}
	var pending []protoreflect.FullName
	markAffected := func(name protoreflect.FullName) {
		if !affected[name] {
			affected[name] = true
			pending = append(pending, name)
		}
	}

	var index func(protoreflect.MessageDescriptor)
	index = func(message protoreflect.MessageDescriptor) {
		name := message.FullName()
		if seen[name] {
			return
		}
		seen[name] = true
		if draftMessage(message) {
			markAffected(name)
		}
		fields := message.Fields()
		for i := range fields.Len() {
			field := fields.Get(i)
			value := field
			if field.IsMap() {
				value = field.MapValue()
			}
			if draftField(field) || hasDraftEnumValues(value.Enum()) {
				markAffected(name)
			}
			if child := value.Message(); child != nil {
				parents[child.FullName()] = append(parents[child.FullName()], name)
				index(child)
			}
		}
	}
	for _, message := range messages {
		index(message)
	}
	for i := 0; i < len(pending); i++ {
		for _, parent := range parents[pending[i]] {
			markAffected(parent)
		}
	}
	return affected
}

func hasDraftEnumValues(enum protoreflect.EnumDescriptor) bool {
	if enum == nil {
		return false
	}
	values := enum.Values()
	for i := range values.Len() {
		if hasDraftOption(values.Get(i), protometapb.E_DraftEnumValue) {
			return true
		}
	}
	return false
}

func TestProjectStableFieldsRemovesDraftEnumValues(t *testing.T) {
	file := draftProjectionFixture(t)
	for _, tt := range []struct {
		name         string
		message      protoreflect.Name
		input, want  string
		removedValue bool
	}{
		{"singular", "Values", "singular: MODE_PREVIEW", "", true},
		{"optional", "Values", "optional_mode: MODE_PREVIEW", "", true},
		{"oneof", "Values", "selected: MODE_PREVIEW", "", true},
		{"list", "Values", "list: [MODE_PREVIEW, MODE_STABLE, 99, MODE_PREVIEW]", "list: [MODE_STABLE, 99]", true},
		{"map", "Values", `by_name: [{key: "preview", value: MODE_PREVIEW}, {key: "stable", value: MODE_STABLE}, {key: "unknown", value: 99}]`,
			`by_name: [{key: "stable", value: MODE_STABLE}, {key: "unknown", value: 99}]`, true},
		{"retained message", "Envelope", "retained: {singular: MODE_PREVIEW}", "retained: {}", true},
		{"disabled field", "Envelope", "hidden: {singular: MODE_PREVIEW}", "", false},
		{"disabled oneof", "Envelope", "hidden_oneof: {singular: MODE_PREVIEW}", "", false},
		{"retained numbers", "Values", "singular: MODE_STABLE optional_mode: 99", "singular: MODE_STABLE optional_mode: 99", false},
	} {
		t.Run(tt.name, func(t *testing.T) {
			message := file.Messages().ByName(tt.message)
			got := dynamicpb.NewMessage(message)
			want := dynamicpb.NewMessage(message)
			require.NoError(t, prototext.Unmarshal([]byte(tt.input), got))
			require.NoError(t, prototext.Unmarshal([]byte(tt.want), want))
			require.Equal(t, tt.removedValue, projectStableFields(got))
			require.True(t, proto.Equal(want, got), "projected %v, want %v", got, want)
		})
	}
}

func draftProjectionFixture(t *testing.T) protoreflect.FileDescriptor {
	t.Helper()
	const name = "draft_projection.proto"
	files := compileFiles(t, []string{"-I" + protoRoot, "-Itestdata"}, []string{name})
	file, err := files.FindFileByPath(name)
	require.NoError(t, err)
	return file
}

// projectStableFields removes draft fields and enum values using api_next descriptors.
func projectStableFields(m protoreflect.Message) (removedEnumValue bool) {
	return projectDraft(m, func(string) bool { return false })
}

// projectDraft removes draft fields and enum values whose features are disabled, using
// api_next descriptors. It skips the contents of removed fields.
func projectDraft(m protoreflect.Message, enabled func(string) bool) (removedEnumValue bool) {
	disabledValue := func(field protoreflect.FieldDescriptor, value protoreflect.Value) bool {
		feature := draftEnumValueFeature(field, value)
		return feature != "" && !enabled(feature)
	}
	m.Range(func(field protoreflect.FieldDescriptor, value protoreflect.Value) bool {
		if feature := draftFieldFeature(field); feature != "" && !enabled(feature) {
			m.Clear(field)
			return true
		}
		switch {
		case field.IsList():
			list := value.List()
			kept := 0
			for i := range list.Len() {
				item := list.Get(i)
				if disabledValue(field, item) {
					removedEnumValue = true
					continue
				}
				if field.Message() != nil && projectDraft(item.Message(), enabled) {
					removedEnumValue = true
				}
				if kept != i {
					list.Set(kept, item)
				}
				kept++
			}
			if kept != list.Len() {
				list.Truncate(kept)
			}
		case field.IsMap():
			entries := value.Map()
			entries.Range(func(key protoreflect.MapKey, item protoreflect.Value) bool {
				if disabledValue(field.MapValue(), item) {
					entries.Clear(key)
					removedEnumValue = true
				} else if field.MapValue().Message() != nil && projectDraft(item.Message(), enabled) {
					removedEnumValue = true
				}
				return true
			})
		case disabledValue(field, value):
			m.Clear(field)
			removedEnumValue = true
		case field.Message() != nil:
			if projectDraft(value.Message(), enabled) {
				removedEnumValue = true
			}
		}
		return true
	})
	return removedEnumValue
}

func draftField(field protoreflect.FieldDescriptor) bool {
	return draftFieldFeature(field) != ""
}

// draftFieldFeature returns the feature of a field annotated as draft or a member of a
// draft oneof, or "".
func draftFieldFeature(field protoreflect.FieldDescriptor) string {
	if feature := draftOption(field, protometapb.E_Draft); feature != "" {
		return feature
	}
	if oneof := field.ContainingOneof(); oneof != nil && !oneof.IsSynthetic() {
		return draftOption(oneof, protometapb.E_DraftOneof)
	}
	return ""
}

// draftEnumValueFeature returns the feature of value if it is a draft value of field's enum, or "".
func draftEnumValueFeature(field protoreflect.FieldDescriptor, value protoreflect.Value) string {
	if field.Enum() == nil {
		return ""
	}
	descriptor := field.Enum().Values().ByNumber(value.Enum())
	if descriptor == nil {
		return ""
	}
	return draftOption(descriptor, protometapb.E_DraftEnumValue)
}

func hasDraftOption(descriptor protoreflect.Descriptor, option protoreflect.ExtensionType) bool {
	return draftOption(descriptor, option) != ""
}

func draftOption(descriptor protoreflect.Descriptor, option protoreflect.ExtensionType) string {
	feature, _ := proto.GetExtension(descriptor.Options(), option).(string)
	return feature
}

// Unaffected message fields stop at shallowDepth. Fields that reach draft contents continue
// to maxDepth, so draft contents appear deep in the tree.
const (
	shallowDepth = 3
	maxDepth     = 10
)

// filler fills random messages, favoring fields that reach draft contents.
type filler struct {
	affected map[protoreflect.FullName]bool
}

// reachesDraft reports whether field is draft or can hold a draft enum value or draft field.
func (f *filler) reachesDraft(field protoreflect.FieldDescriptor) bool {
	value := field
	if field.IsMap() {
		value = field.MapValue()
	}
	return draftField(field) || hasDraftEnumValues(value.Enum()) ||
		value.Message() != nil && f.affected[value.Message().FullName()]
}

// allowed reports whether field may be set at depth.
func (f *filler) allowed(field protoreflect.FieldDescriptor, depth int) bool {
	value := field
	if field.IsMap() {
		value = field.MapValue()
	}
	if value.Message() == nil {
		return true
	}
	return depth < shallowDepth || depth < maxDepth && f.reachesDraft(field)
}

// chance draws true with the given percent probability.
func chance(rt *rapid.T, percent int, label string) bool {
	return rapid.IntRange(0, 99).Draw(rt, label) < percent
}

// fill sets a random selection of m's fields. It sets at most one member of each real
// oneof, choosing a member that reaches draft contents most of the time.
func (f *filler) fill(rt *rapid.T, m protoreflect.Message, depth int) {
	oneofs := m.Descriptor().Oneofs()
	for i := range oneofs.Len() {
		oneof := oneofs.Get(i)
		if oneof.IsSynthetic() || !rapid.Bool().Draw(rt, string(oneof.FullName())) {
			continue
		}
		var members, drafty []protoreflect.FieldDescriptor
		for j := range oneof.Fields().Len() {
			if member := oneof.Fields().Get(j); f.allowed(member, depth) {
				members = append(members, member)
				if f.reachesDraft(member) {
					drafty = append(drafty, member)
				}
			}
		}
		if len(drafty) > 0 && chance(rt, 75, "draft member") {
			members = drafty
		}
		if len(members) > 0 {
			f.set(rt, m, rapid.SampledFrom(members).Draw(rt, "member"), depth)
		}
	}
	fields := m.Descriptor().Fields()
	for i := range fields.Len() {
		field := fields.Get(i)
		if oneof := field.ContainingOneof(); oneof != nil && !oneof.IsSynthetic() || !f.allowed(field, depth) {
			continue
		}
		percent := 50
		if f.reachesDraft(field) {
			percent = 85
		}
		if chance(rt, percent, string(field.FullName())) {
			f.set(rt, m, field, depth)
		}
	}
}

func (f *filler) set(rt *rapid.T, m protoreflect.Message, field protoreflect.FieldDescriptor, depth int) {
	// Keep deep collections small so recursive draft types stay bounded.
	maxLength := 2
	if depth >= shallowDepth {
		maxLength = 1
	}
	switch {
	case field.IsList():
		list := m.Mutable(field).List()
		for range rapid.IntRange(1, maxLength).Draw(rt, "length") {
			list.Append(f.value(rt, list.NewElement, field, depth))
		}
	case field.IsMap():
		entries := m.Mutable(field).Map()
		for range rapid.IntRange(1, maxLength).Draw(rt, "entries") {
			key := f.value(rt, nil, field.MapKey(), depth).MapKey()
			entries.Set(key, f.value(rt, entries.NewValue, field.MapValue(), depth))
		}
	default:
		m.Set(field, f.value(rt, func() protoreflect.Value { return m.NewField(field) }, field, depth))
	}
}

func (f *filler) value(rt *rapid.T, newMessage func() protoreflect.Value, field protoreflect.FieldDescriptor, depth int) protoreflect.Value {
	switch field.Kind() {
	case protoreflect.BoolKind:
		return protoreflect.ValueOfBool(rapid.Bool().Draw(rt, "bool"))
	case protoreflect.EnumKind:
		values := field.Enum().Values()
		if !field.Enum().IsClosed() && rapid.Bool().Draw(rt, "unrecognized enum number") {
			for number := protoreflect.EnumNumber(-1); ; number-- {
				if values.ByNumber(number) == nil {
					return protoreflect.ValueOfEnum(number)
				}
			}
		}
		return protoreflect.ValueOfEnum(values.Get(rapid.IntRange(0, values.Len()-1).Draw(rt, "enum")).Number())
	case protoreflect.Int32Kind, protoreflect.Sint32Kind, protoreflect.Sfixed32Kind:
		return protoreflect.ValueOfInt32(rapid.Int32().Draw(rt, "int32"))
	case protoreflect.Int64Kind, protoreflect.Sint64Kind, protoreflect.Sfixed64Kind:
		return protoreflect.ValueOfInt64(rapid.Int64().Draw(rt, "int64"))
	case protoreflect.Uint32Kind, protoreflect.Fixed32Kind:
		return protoreflect.ValueOfUint32(rapid.Uint32().Draw(rt, "uint32"))
	case protoreflect.Uint64Kind, protoreflect.Fixed64Kind:
		return protoreflect.ValueOfUint64(rapid.Uint64().Draw(rt, "uint64"))
	case protoreflect.FloatKind:
		return protoreflect.ValueOfFloat32(float32(rapid.IntRange(-1000, 1000).Draw(rt, "float")))
	case protoreflect.DoubleKind:
		return protoreflect.ValueOfFloat64(float64(rapid.IntRange(-1000, 1000).Draw(rt, "double")))
	case protoreflect.StringKind:
		return protoreflect.ValueOfString(rapid.StringN(0, 8, -1).Draw(rt, "string"))
	case protoreflect.BytesKind:
		return protoreflect.ValueOfBytes(rapid.SliceOfN(rapid.Byte(), 0, 8).Draw(rt, "bytes"))
	case protoreflect.MessageKind, protoreflect.GroupKind:
		message := newMessage()
		f.fill(rt, message.Message(), depth+1)
		return message
	default:
		panic(fmt.Sprintf("unhandled kind %v", field.Kind()))
	}
}

// unknownFields returns the path to the first message with unknown fields, or "".
func unknownFields(m protoreflect.Message, path string) string {
	if len(m.GetUnknown()) > 0 {
		return path
	}
	found := ""
	m.Range(func(field protoreflect.FieldDescriptor, v protoreflect.Value) bool {
		at := path + "." + string(field.Name())
		switch {
		case field.IsList() && field.Message() != nil:
			for i := 0; i < v.List().Len() && found == ""; i++ {
				found = unknownFields(v.List().Get(i).Message(), fmt.Sprintf("%s[%d]", at, i))
			}
		case field.IsMap() && field.MapValue().Message() != nil:
			v.Map().Range(func(key protoreflect.MapKey, value protoreflect.Value) bool {
				found = unknownFields(value.Message(), fmt.Sprintf("%s[%v]", at, key))
				return found == ""
			})
		case field.Message() != nil && !field.IsList() && !field.IsMap():
			found = unknownFields(v.Message(), at)
		}
		return found == ""
	})
	return found
}

// draftMessage reports whether message or a message it is nested in is draft.
func draftMessage(message protoreflect.MessageDescriptor) bool {
	for descriptor := protoreflect.Descriptor(message); descriptor != nil; descriptor = descriptor.Parent() {
		if m, ok := descriptor.(protoreflect.MessageDescriptor); ok && hasDraftOption(m, protometapb.E_DraftMessage) {
			return true
		}
	}
	return false
}

// nextMessages returns every temporal.api message linked from the api_next Go types.
func nextMessages() []protoreflect.MessageDescriptor {
	var messages []protoreflect.MessageDescriptor
	rangeAPIFiles(protoregistry.GlobalFiles, func(file protoreflect.FileDescriptor) {
		eachMessage(file.Messages(), func(message protoreflect.MessageDescriptor) {
			messages = append(messages, message)
		})
	})
	slices.SortFunc(messages, func(a, b protoreflect.MessageDescriptor) int {
		return strings.Compare(string(a.FullName()), string(b.FullName()))
	})
	return messages
}

func rangeAPIFiles(files *protoregistry.Files, visit func(protoreflect.FileDescriptor)) {
	files.RangeFiles(func(file protoreflect.FileDescriptor) bool {
		if strings.HasPrefix(string(file.Package()), "temporal.api.") {
			visit(file)
		}
		return true
	})
}

func eachMessage(messages protoreflect.MessageDescriptors, visit func(protoreflect.MessageDescriptor)) {
	for i := 0; i < messages.Len(); i++ {
		if message := messages.Get(i); !message.IsMapEntry() {
			visit(message)
			eachMessage(message.Messages(), visit)
		}
	}
}

const protoRoot = "../../proto/api"

// protoFiles lists the proto files under dir of the proto/api submodule, relative to it.
func protoFiles(t *testing.T, dir string) []string {
	t.Helper()
	var files []string
	require.NoError(t, filepath.WalkDir(filepath.Join(protoRoot, dir), func(path string, entry os.DirEntry, err error) error {
		if err == nil && !entry.IsDir() && strings.HasSuffix(path, ".proto") {
			relative, relErr := filepath.Rel(protoRoot, path)
			files = append(files, filepath.ToSlash(relative))
			return relErr
		}
		return err
	}))
	require.NotEmpty(t, files, "%s missing; check out the proto/api submodule", dir)
	return files
}

// stableFiles compiles the stable proto tree of the proto/api submodule with protoc.
func stableFiles(t *testing.T) *protoregistry.Files {
	t.Helper()
	return compileFiles(t, []string{"-I" + protoRoot}, protoFiles(t, "temporal/api"))
}

func compileFiles(t *testing.T, includes, inputs []string) *protoregistry.Files {
	t.Helper()
	out := filepath.Join(t.TempDir(), "schema.binpb")
	args := append(includes, "--include_imports", "--descriptor_set_out="+out)
	output, err := exec.Command("protoc", append(args, inputs...)...).CombinedOutput()
	require.NoError(t, err, "%s", output)
	data, err := os.ReadFile(out)
	require.NoError(t, err)
	var set descriptorpb.FileDescriptorSet
	require.NoError(t, proto.Unmarshal(data, &set))
	files, err := protodesc.NewFiles(&set)
	require.NoError(t, err)
	return files
}
