package json

import (
	"fmt"
	"reflect"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func TestTemporalNexusProtoRoundTripAnnotatedTypes(t *testing.T) {
	t.Parallel()

	annotatedTypes := sortedTemporalNexusAnnotatedTypes()
	require.NotEmpty(t, annotatedTypes)

	for _, generatedType := range annotatedTypes {
		generatedType := generatedType
		t.Run(generatedType.Name(), func(t *testing.T) {
			t.Parallel()

			protoValue := buildProtoSample(temporalNexusProtoTypes[generatedType]())
			generatedPtr := reflect.New(generatedType).Interface()

			require.NoError(t, FromTemporalNexusProto(protoValue, generatedPtr))
			require.NotNil(t, GetTemporalNexusProtoMessage(generatedPtr))

			roundTrippedProto, err := ToTemporalNexusProto(reflect.ValueOf(generatedPtr).Elem().Interface())
			require.NoError(t, err)
			require.True(t, proto.Equal(protoValue, roundTrippedProto))
		})
	}
}

func sortedTemporalNexusAnnotatedTypes() []reflect.Type {
	types := make([]reflect.Type, 0, len(temporalNexusProtoTypes))
	for generatedType := range temporalNexusProtoTypes {
		types = append(types, generatedType)
	}
	sort.Slice(types, func(i, j int) bool {
		return types[i].String() < types[j].String()
	})
	return types
}

func buildProtoSample(message proto.Message) proto.Message {
	populateProtoSample(message.ProtoReflect(), "value")
	return message
}

func populateProtoSample(message protoreflect.Message, path string) {
	seenOneofs := map[protoreflect.Name]struct{}{}
	fields := message.Descriptor().Fields()
	for i := 0; i < fields.Len(); i++ {
		field := fields.Get(i)
		if oneof := field.ContainingOneof(); oneof != nil {
			if _, seen := seenOneofs[oneof.Name()]; seen {
				continue
			}
			seenOneofs[oneof.Name()] = struct{}{}
		}
		switch {
		case field.IsList():
			list := message.Mutable(field).List()
			if field.Kind() == protoreflect.MessageKind || field.Kind() == protoreflect.GroupKind {
				item := list.NewElement().Message()
				populateProtoSample(item, fmt.Sprintf("%s.%s[0]", path, field.Name()))
				list.Append(protoreflect.ValueOfMessage(item))
			} else {
				list.Append(sampleScalarValue(field, fmt.Sprintf("%s.%s[0]", path, field.Name())))
			}
		case field.IsMap():
			populateProtoMapEntry(message.Mutable(field).Map(), field, path)
		case field.Kind() == protoreflect.MessageKind || field.Kind() == protoreflect.GroupKind:
			populateProtoSample(message.Mutable(field).Message(), fmt.Sprintf("%s.%s", path, field.Name()))
		default:
			message.Set(field, sampleScalarValue(field, fmt.Sprintf("%s.%s", path, field.Name())))
		}
	}
}

func populateProtoMapEntry(m protoreflect.Map, field protoreflect.FieldDescriptor, path string) {
	keyField := field.MapKey()
	valueField := field.MapValue()
	key := sampleMapKey(keyField, fmt.Sprintf("%s.%s.key", path, field.Name()))
	if valueField.Kind() == protoreflect.MessageKind || valueField.Kind() == protoreflect.GroupKind {
		value := m.NewValue()
		populateProtoSample(value.Message(), fmt.Sprintf("%s.%s[%v]", path, field.Name(), key.Interface()))
		m.Set(key, value)
		return
	}
	m.Set(key, sampleScalarValue(valueField, fmt.Sprintf("%s.%s[%v]", path, field.Name(), key.Interface())))
}

func sampleMapKey(field protoreflect.FieldDescriptor, path string) protoreflect.MapKey {
	value := sampleScalarValue(field, path)
	return value.MapKey()
}

func sampleScalarValue(field protoreflect.FieldDescriptor, path string) protoreflect.Value {
	switch field.Kind() {
	case protoreflect.BoolKind:
		return protoreflect.ValueOfBool(true)
	case protoreflect.EnumKind:
		values := field.Enum().Values()
		for i := 0; i < values.Len(); i++ {
			if number := values.Get(i).Number(); number != 0 {
				return protoreflect.ValueOfEnum(number)
			}
		}
		return protoreflect.ValueOfEnum(values.Get(0).Number())
	case protoreflect.Int32Kind, protoreflect.Sint32Kind, protoreflect.Sfixed32Kind:
		return protoreflect.ValueOfInt32(1)
	case protoreflect.Int64Kind, protoreflect.Sint64Kind, protoreflect.Sfixed64Kind:
		return protoreflect.ValueOfInt64(1)
	case protoreflect.Uint32Kind, protoreflect.Fixed32Kind:
		return protoreflect.ValueOfUint32(1)
	case protoreflect.Uint64Kind, protoreflect.Fixed64Kind:
		return protoreflect.ValueOfUint64(1)
	case protoreflect.FloatKind:
		return protoreflect.ValueOfFloat32(1.5)
	case protoreflect.DoubleKind:
		return protoreflect.ValueOfFloat64(1.5)
	case protoreflect.StringKind:
		return protoreflect.ValueOfString(path + "-value")
	case protoreflect.BytesKind:
		return protoreflect.ValueOfBytes([]byte("test"))
	default:
		panic(fmt.Sprintf("unhandled proto scalar sample at %s: %s", path, field.Kind()))
	}
}
