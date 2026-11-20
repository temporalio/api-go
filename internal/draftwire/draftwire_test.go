package draftwire

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"slices"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	protometapb "go.temporal.io/api/protometa/v1"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protodesc"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"
	"google.golang.org/protobuf/types/descriptorpb"
	"google.golang.org/protobuf/types/dynamicpb"
	"pgregory.net/rapid"
)

// draftOmitter is implemented by api_next messages that can reach a draft field.
type draftOmitter interface {
	OmitDraftFields(enabled func(string) bool)
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

// TestMessageSets requires the stable tree to hold every api_next message except draft
// ones, and nothing else.
func TestMessageSets(t *testing.T) {
	stable := stableFiles(t)
	next := map[protoreflect.FullName]bool{}
	for _, message := range nextMessages() {
		next[message.FullName()] = true
		_, err := stable.FindDescriptorByName(message.FullName())
		switch draft := draftMessage(message); {
		case draft && err == nil:
			t.Errorf("draft message %s is in the stable tree", message.FullName())
		case !draft && err != nil:
			t.Errorf("message %s is missing from the stable tree", message.FullName())
		}
	}
	stable.RangeFiles(func(file protoreflect.FileDescriptor) bool {
		if strings.HasPrefix(string(file.Package()), "temporal.api.") {
			eachMessage(file.Messages(), func(message protoreflect.MessageDescriptor) {
				if !next[message.FullName()] {
					t.Errorf("stable message %s is not in api_next", message.FullName())
				}
			})
		}
		return true
	})
}

// TestOmitDraftFieldsMatchesStable fills random api_next messages. With no feature
// enabled, OmitDraftFields must leave exactly what the stable tree can hold: the stable
// descriptors decode it without unknown fields at any depth, and decoding back gives the
// same message. Decoding the original while discarding unknown fields must recover the
// same stable fields. With every feature enabled, OmitDraftFields must change nothing.
func TestOmitDraftFieldsMatchesStable(t *testing.T) {
	stable := stableFiles(t)
	unmarshal := proto.UnmarshalOptions{Resolver: dynamicpb.NewTypes(stable)}
	discardUnknown := unmarshal
	discardUnknown.DiscardUnknown = true
	// Few messages reach a draft field, so draw from those half of the time.
	var affected, unaffected []protoreflect.MessageDescriptor
	for _, message := range nextMessages() {
		if draftMessage(message) {
			continue
		}
		messageType, err := protoregistry.GlobalTypes.FindMessageByName(message.FullName())
		require.NoError(t, err)
		if _, ok := messageType.New().Interface().(draftOmitter); ok {
			affected = append(affected, message)
		} else {
			unaffected = append(unaffected, message)
		}
	}
	require.NotEmpty(t, affected)
	require.NotEmpty(t, unaffected)
	rapid.Check(t, func(rt *rapid.T) {
		candidates := unaffected
		if rapid.Bool().Draw(rt, "affected") {
			candidates = affected
		}
		descriptor := rapid.SampledFrom(candidates).Draw(rt, "message")
		messageType, err := protoregistry.GlobalTypes.FindMessageByName(descriptor.FullName())
		if err != nil {
			rt.Fatal(err)
		}
		original := messageType.New()
		fill(rt, original, 0)

		all := proto.Clone(original.Interface())
		if omitter, ok := all.(draftOmitter); ok {
			omitter.OmitDraftFields(func(string) bool { return true })
		}
		if !proto.Equal(original.Interface(), all) {
			rt.Fatalf("OmitDraftFields with every feature enabled changed %s", descriptor.FullName())
		}

		none := proto.Clone(original.Interface())
		if omitter, ok := none.(draftOmitter); ok {
			omitter.OmitDraftFields(func(string) bool { return false })
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
		originalData, err := proto.Marshal(original.Interface())
		if err != nil {
			rt.Fatal(err)
		}
		expected := dynamicpb.NewMessage(stableMessage)
		if err := discardUnknown.Unmarshal(originalData, expected); err != nil {
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
	})
}

// fill sets a random selection of m's fields, nesting at most three messages deep.
func fill(rt *rapid.T, m protoreflect.Message, depth int) {
	fields := m.Descriptor().Fields()
	for i := 0; i < fields.Len(); i++ {
		field := fields.Get(i)
		if oneof := field.ContainingOneof(); oneof != nil && m.WhichOneof(oneof) != nil {
			continue
		}
		if field.Message() != nil && depth >= 3 {
			continue
		}
		if !rapid.Bool().Draw(rt, string(field.FullName())) {
			continue
		}
		switch {
		case field.IsList():
			list := m.Mutable(field).List()
			for range rapid.IntRange(1, 2).Draw(rt, "length") {
				list.Append(value(rt, list.NewElement, field, depth))
			}
		case field.IsMap():
			entries := m.Mutable(field).Map()
			for range rapid.IntRange(1, 2).Draw(rt, "entries") {
				key := value(rt, nil, field.MapKey(), depth).MapKey()
				entries.Set(key, value(rt, entries.NewValue, field.MapValue(), depth))
			}
		default:
			m.Set(field, value(rt, func() protoreflect.Value { return m.NewField(field) }, field, depth))
		}
	}
}

func value(rt *rapid.T, newMessage func() protoreflect.Value, field protoreflect.FieldDescriptor, depth int) protoreflect.Value {
	switch field.Kind() {
	case protoreflect.BoolKind:
		return protoreflect.ValueOfBool(rapid.Bool().Draw(rt, "bool"))
	case protoreflect.EnumKind:
		values := field.Enum().Values()
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
		fill(rt, message.Message(), depth+1)
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
		if m, ok := descriptor.(protoreflect.MessageDescriptor); ok {
			if feature, _ := proto.GetExtension(m.Options(), protometapb.E_DraftMessage).(string); feature != "" {
				return true
			}
		}
	}
	return false
}

// nextMessages returns every temporal.api message linked from the api_next Go types.
func nextMessages() []protoreflect.MessageDescriptor {
	var messages []protoreflect.MessageDescriptor
	protoregistry.GlobalFiles.RangeFiles(func(file protoreflect.FileDescriptor) bool {
		if strings.HasPrefix(string(file.Package()), "temporal.api.") {
			eachMessage(file.Messages(), func(message protoreflect.MessageDescriptor) {
				messages = append(messages, message)
			})
		}
		return true
	})
	slices.SortFunc(messages, func(a, b protoreflect.MessageDescriptor) int {
		return strings.Compare(string(a.FullName()), string(b.FullName()))
	})
	return messages
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
	root := protoRoot
	inputs := protoFiles(t, "temporal/api")
	out := filepath.Join(t.TempDir(), "stable.binpb")
	output, err := exec.Command("protoc", append([]string{"-I" + root, "--include_imports", "--descriptor_set_out=" + out}, inputs...)...).CombinedOutput()
	require.NoError(t, err, "%s", output)
	data, err := os.ReadFile(out)
	require.NoError(t, err)
	var set descriptorpb.FileDescriptorSet
	require.NoError(t, proto.Unmarshal(data, &set))
	files, err := protodesc.NewFiles(&set)
	require.NoError(t, err)
	return files
}
