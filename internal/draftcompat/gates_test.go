package draftcompat

import (
	"testing"

	"github.com/stretchr/testify/require"
	protometapb "go.temporal.io/api/protometa/v1"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"
)

func TestDraftGates(t *testing.T) {
	for _, message := range nextMessages() {
		for i := range message.Fields().Len() {
			field := message.Fields().Get(i)
			feature := draftFieldFeature(field)
			if feature != "" {
				t.Run(string(field.FullName()), func(t *testing.T) {
					checkFieldGate(t, field, feature, nil)
				})
			}

			elem := field
			if field.IsMap() {
				elem = field.MapValue()
			}
			if enum := elem.Enum(); enum != nil {
				for j := range enum.Values().Len() {
					value := enum.Values().Get(j)
					gate := draftOption(value, protometapb.E_DraftEnumValue)
					if gate != "" {
						t.Run(string(field.FullName())+"/"+string(value.Name()), func(t *testing.T) {
							checkFieldGate(t, field, gate, value)
						})
					}
				}
			}
		}
	}
}

func checkFieldGate(t *testing.T, field protoreflect.FieldDescriptor, gate string, enumValue protoreflect.EnumValueDescriptor) {
	t.Helper()
	typ, err := protoregistry.GlobalTypes.FindMessageByName(field.ContainingMessage().FullName())
	require.NoError(t, err)

	message := typ.New()
	setFixtureField(message, field, enumValue, 0)
	require.True(t, message.Has(field), "test input must populate %s", field.FullName())

	omitter, ok := message.Interface().(draftOmitter)
	require.True(t, ok, "%s has no OmitDraftFields", message.Descriptor().FullName())
	want := proto.Clone(message.Interface())
	omitter.OmitDraftFields(func(string) bool { return true })
	require.True(t, proto.Equal(want, message.Interface()), "enabled gate %q changed %s", gate, field.FullName())

	omitter.OmitDraftFields(func(feature string) bool { return feature != gate })
	require.False(t, message.Has(field), "disabled gate %q left %s set", gate, field.FullName())
}

// fillFixture sets one value per field, including one entry per collection.
func fillFixture(message protoreflect.Message, depth int) {
	for i := range message.Descriptor().Fields().Len() {
		field := message.Descriptor().Fields().Get(i)
		if oneof := field.ContainingOneof(); oneof != nil && message.WhichOneof(oneof) != nil {
			continue
		}
		setFixtureField(message, field, nil, depth)
	}
}

func setFixtureField(message protoreflect.Message, field protoreflect.FieldDescriptor, enumValue protoreflect.EnumValueDescriptor, depth int) {
	switch {
	case field.IsList():
		list := message.Mutable(field).List()
		list.Append(fixtureValue(field, list.NewElement(), enumValue, depth))
	case field.IsMap():
		entries := message.Mutable(field).Map()
		key := fixtureValue(field.MapKey(), protoreflect.Value{}, nil, 0).MapKey()
		entries.Set(key, fixtureValue(field.MapValue(), entries.NewValue(), enumValue, depth))
	default:
		message.Set(field, fixtureValue(field, message.NewField(field), enumValue, depth))
	}
}

// fixtureValue supplies a present value even for fields with implicit presence.
func fixtureValue(field protoreflect.FieldDescriptor, empty protoreflect.Value, enumValue protoreflect.EnumValueDescriptor, depth int) protoreflect.Value {
	switch field.Kind() {
	case protoreflect.BoolKind:
		return protoreflect.ValueOfBool(true)
	case protoreflect.EnumKind:
		if enumValue != nil {
			return protoreflect.ValueOfEnum(enumValue.Number())
		}
		values := field.Enum().Values()
		for i := range values.Len() {
			if number := values.Get(i).Number(); number != 0 {
				return protoreflect.ValueOfEnum(number)
			}
		}
		return protoreflect.ValueOfEnum(1)
	case protoreflect.Int32Kind, protoreflect.Sint32Kind, protoreflect.Sfixed32Kind:
		return protoreflect.ValueOfInt32(1)
	case protoreflect.Int64Kind, protoreflect.Sint64Kind, protoreflect.Sfixed64Kind:
		return protoreflect.ValueOfInt64(1)
	case protoreflect.Uint32Kind, protoreflect.Fixed32Kind:
		return protoreflect.ValueOfUint32(1)
	case protoreflect.Uint64Kind, protoreflect.Fixed64Kind:
		return protoreflect.ValueOfUint64(1)
	case protoreflect.FloatKind:
		return protoreflect.ValueOfFloat32(1)
	case protoreflect.DoubleKind:
		return protoreflect.ValueOfFloat64(1)
	case protoreflect.StringKind:
		return protoreflect.ValueOfString("example")
	case protoreflect.BytesKind:
		return protoreflect.ValueOfBytes([]byte{1})
	case protoreflect.MessageKind, protoreflect.GroupKind:
		if depth > 0 {
			fillFixture(empty.Message(), depth-1)
		}
		return empty
	default:
		panic("unsupported protobuf field kind")
	}
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

func draftOption(descriptor protoreflect.Descriptor, option protoreflect.ExtensionType) string {
	feature, _ := proto.GetExtension(descriptor.Options(), option).(string)
	return feature
}
