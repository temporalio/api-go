package draftcompat

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

// TestOmitDraftFieldsMatchesStable checks every stable message with empty and populated payloads.
// With no feature enabled, OmitDraftFields must leave exactly what the stable tree can hold: the stable
// descriptors decode it without unknown fields at any depth, and decoding back gives the
// same message. Decoding the original while discarding unknown fields must recover the
// same stable fields. With every feature enabled, OmitDraftFields must change nothing.
func TestOmitDraftFieldsMatchesStable(t *testing.T) {
	stable := stableFiles(t)
	unmarshal := proto.UnmarshalOptions{Resolver: dynamicpb.NewTypes(stable)}
	discardUnknown := unmarshal
	discardUnknown.DiscardUnknown = true

	check := func(t *testing.T, descriptor protoreflect.MessageDescriptor, original protoreflect.Message) {
		t.Helper()

		// All gates on should preserve the payload.
		all := proto.Clone(original.Interface())
		if omitter, ok := all.(draftOmitter); ok {
			omitter.OmitDraftFields(func(string) bool { return true })
		}
		if !proto.Equal(original.Interface(), all) {
			t.Fatalf("OmitDraftFields with every feature enabled changed %s", descriptor.FullName())
		}

		// All gates off should leave no unknown fields for a stable reader.
		none := proto.Clone(original.Interface())
		if omitter, ok := none.(draftOmitter); ok {
			omitter.OmitDraftFields(func(string) bool { return false })
		}
		data, err := proto.MarshalOptions{Deterministic: true}.Marshal(none)
		if err != nil {
			t.Fatal(err)
		}
		decoded := dynamicpb.NewMessage(descriptor)
		if err := unmarshal.Unmarshal(data, decoded); err != nil {
			t.Fatal(err)
		}
		if path := unknownFields(decoded, string(descriptor.FullName())); path != "" {
			t.Fatalf("stable %s has unknown fields at %s after omitting every draft feature", descriptor.FullName(), path)
		}

		// Check that the stable fields survived omission.
		originalData, err := proto.Marshal(original.Interface())
		if err != nil {
			t.Fatal(err)
		}
		expected := dynamicpb.NewMessage(descriptor)
		if err := discardUnknown.Unmarshal(originalData, expected); err != nil {
			t.Fatal(err)
		}
		if !proto.Equal(expected, decoded) {
			t.Fatalf("OmitDraftFields with no features enabled changed stable fields of %s: got %v, want %v", descriptor.FullName(), decoded, expected)
		}

		// Round-trip back into the draft type.
		stableData, err := proto.MarshalOptions{Deterministic: true}.Marshal(decoded)
		if err != nil {
			t.Fatal(err)
		}
		back := original.New().Interface()
		if err := proto.Unmarshal(stableData, back); err != nil {
			t.Fatal(err)
		}
		if !proto.Equal(none, back) {
			t.Fatalf("%s changed on its way through the stable descriptors", descriptor.FullName())
		}
	}

	for _, descriptor := range apiMessages(stable) {
		t.Run(string(descriptor.FullName()), func(t *testing.T) {
			typeInDraft, err := protoregistry.GlobalTypes.FindMessageByName(descriptor.FullName())
			require.NoError(t, err)
			t.Parallel()
			t.Run("empty", func(t *testing.T) {
				check(t, descriptor, typeInDraft.New())
			})
			t.Run("populated", func(t *testing.T) {
				message := typeInDraft.New()
				fillFixture(message, 3)
				check(t, descriptor, message)
			})
		})
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
	return apiMessages(protoregistry.GlobalFiles)
}

// apiMessages returns every temporal.api message in the given registry.
func apiMessages(files *protoregistry.Files) []protoreflect.MessageDescriptor {
	var messages []protoreflect.MessageDescriptor
	files.RangeFiles(func(file protoreflect.FileDescriptor) bool {
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
