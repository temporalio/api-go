package main

import (
	"encoding/json"
	"os"
	"os/exec"
	"path/filepath"
	"slices"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/encoding/protowire"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protodesc"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/descriptorpb"
	"google.golang.org/protobuf/types/gofeaturespb"
	"google.golang.org/protobuf/types/pluginpb"

	"go.temporal.io/api/cmd/protoc-gen-go-helpers/testprotos/draft/child"
	"go.temporal.io/api/cmd/protoc-gen-go-helpers/testprotos/draft/parent"
)

type omitter interface{ OmitDraftFields(func(string) bool) }

func TestOmitFixtures(t *testing.T) {
	leaf := func() *child.Leaf { return &child.Leaf{Stable: "stable", A: "a", B: "b"} }
	for _, features := range [][]string{nil, {"feature-a"}, {"feature-b"}, {"feature-a", "feature-b"}} {
		t.Run(strings.Join(features, "+"), func(t *testing.T) {
			enabled := func(s string) bool {
				for _, f := range features {
					if s == f {
						return true
					}
				}
				return false
			}
			m := &parent.Root{Stable: "root", Gated: leaf(), List: []*child.Leaf{nil, leaf()}, Map: map[string]*child.Leaf{"nil": nil, "leaf": leaf()}, Optional: proto.Int32(0), Choice: &parent.Root_A{A: 0}, Nested: &parent.Root_Nested{Leaf: leaf()}, Names: []string{"name"}, Counts: map[string]int32{"count": 1}}
			expectedLeaf := &child.Leaf{Stable: "stable"}
			if enabled("feature-a") {
				expectedLeaf.A = "a"
			}
			if enabled("feature-b") {
				expectedLeaf.B = "b"
			}
			expected := &parent.Root{Stable: "root", List: []*child.Leaf{nil, proto.Clone(expectedLeaf).(*child.Leaf)}, Map: map[string]*child.Leaf{"nil": nil, "leaf": proto.Clone(expectedLeaf).(*child.Leaf)}, Nested: &parent.Root_Nested{Leaf: proto.Clone(expectedLeaf).(*child.Leaf)}}
			if enabled("feature-a") {
				expected.Gated = proto.Clone(expectedLeaf).(*child.Leaf)
				expected.Optional = proto.Int32(0)
				expected.Choice = &parent.Root_A{A: 0}
			}
			if enabled("feature-b") {
				expected.Names = []string{"name"}
				expected.Counts = map[string]int32{"count": 1}
			}
			originalLeaf := m.List[1]
			m.OmitDraftFields(enabled)
			require.True(t, proto.Equal(expected, m), "expected %v, got %v", expected, m)
			require.Same(t, originalLeaf, m.List[1])
			m.OmitDraftFields(enabled)
			require.True(t, proto.Equal(expected, m))
			m.Choice = &parent.Root_Leaf{Leaf: leaf()}
			m.OmitDraftFields(enabled)
			require.True(t, proto.Equal(expectedLeaf, m.GetLeaf()))
			m.Choice = &parent.Root_StableChoice{StableChoice: "stable"}
			m.OmitDraftFields(enabled)
			require.Equal(t, "stable", m.GetStableChoice())
			m.Choice = (*parent.Root_Leaf)(nil)
			require.NotPanics(t, func() { m.OmitDraftFields(enabled) })
			m.Choice = (*parent.Root_A)(nil)
			m.OmitDraftFields(enabled)
			if !enabled("feature-a") {
				require.Nil(t, m.Choice)
			}
		})
	}
	var nilRoot *parent.Root
	require.NotPanics(t, func() { nilRoot.OmitDraftFields(func(string) bool { return false }) })
	_, ok := any(&child.Stable{}).(omitter)
	require.False(t, ok)
	recursive := &parent.Ancestor{Child: &parent.CycleA{B: &parent.CycleB{A: &parent.CycleA{B: &parent.CycleB{Leaf: leaf()}}, Leaf: leaf()}}}
	recursive.OmitDraftFields(func(string) bool { return false })
	require.Equal(t, &child.Leaf{Stable: "stable"}, recursive.Child.B.Leaf)
	require.Equal(t, &child.Leaf{Stable: "stable"}, recursive.Child.B.A.B.Leaf)
}

func TestOmitEveryScalar(t *testing.T) {
	m := &parent.Scalars{Double: 1, Float: 1, Int32: 1, Int64: 1, Uint32: 1, Uint64: 1, Sint32: 1, Sint64: 1, Fixed32: 1, Fixed64: 1, Sfixed32: 1, Sfixed64: 1, Bool: true, String_: "x", Bytes: []byte("x"), Enum: parent.Scalars_ONE}
	expected := proto.Clone(m)
	m.OmitDraftFields(func(string) bool { return true })
	require.True(t, proto.Equal(expected, m))
	m.OmitDraftFields(func(string) bool { return false })
	require.True(t, proto.Equal(&parent.Scalars{}, m))
}

func TestOmitDraftMessageReferences(t *testing.T) {
	for _, features := range [][]string{nil, {"feature-a"}, {"feature-b"}, {"feature-a", "feature-b"}} {
		t.Run(strings.Join(features, "+"), func(t *testing.T) {
			enabled := func(feature string) bool { return slices.Contains(features, feature) }
			value := func() *child.Draft { return &child.Draft{Value: "value", Extra: "extra"} }
			m := &parent.DraftReferences{
				Stable: "stable", Single: value(), List: []*child.Draft{nil, value()},
				ByName: map[string]*child.Draft{"nil": nil, "value": value()},
				Choice: &parent.DraftReferences_Draft{Draft: value()}, Gated: value(),
				Inner: &child.Draft_Inner{Value: "inner"}, State: child.Draft_ONE,
				States: []child.Draft_State{child.Draft_ONE}, StateByName: map[string]child.Draft_State{"value": child.Draft_ONE},
				Empty: &child.Draft_Empty{},
			}
			expected := &parent.DraftReferences{Stable: "stable"}
			if enabled("feature-a") {
				kept := &child.Draft{Value: "value"}
				if enabled("feature-b") {
					kept.Extra = "extra"
					expected.Gated = proto.Clone(kept).(*child.Draft)
					expected.Inner = &child.Draft_Inner{Value: "inner"}
				}
				expected.Single = proto.Clone(kept).(*child.Draft)
				expected.List = []*child.Draft{nil, proto.Clone(kept).(*child.Draft)}
				expected.ByName = map[string]*child.Draft{"nil": nil, "value": proto.Clone(kept).(*child.Draft)}
				expected.Choice = &parent.DraftReferences_Draft{Draft: proto.Clone(kept).(*child.Draft)}
				expected.State, expected.States = child.Draft_ONE, []child.Draft_State{child.Draft_ONE}
				expected.StateByName, expected.Empty = map[string]child.Draft_State{"value": child.Draft_ONE}, &child.Draft_Empty{}
			}
			m.OmitDraftFields(enabled)
			require.True(t, proto.Equal(expected, m), "expected %v, got %v", expected, m)
			m.Choice = &parent.DraftReferences_StableChoice{StableChoice: "stable"}
			m.OmitDraftFields(enabled)
			require.Equal(t, "stable", m.GetStableChoice())
			m.Choice = (*parent.DraftReferences_Draft)(nil)
			require.NotPanics(t, func() { m.OmitDraftFields(enabled) })
		})
	}
}

func TestOmitDirectDraftMessage(t *testing.T) {
	m := &child.Draft{Value: "value", Extra: "extra"}
	m.ProtoReflect().SetUnknown([]byte{0x98, 0x06, 0x01})
	m.OmitDraftFields(func(feature string) bool { return feature == "feature-a" })
	require.Equal(t, "value", m.Value)
	require.Empty(t, m.Extra)
	require.NotEmpty(t, m.ProtoReflect().GetUnknown())
	m.OmitDraftFields(func(string) bool { return false })
	require.True(t, proto.Equal(&child.Draft{}, m))
	inner := &child.Draft_Inner{Value: "inner"}
	inner.OmitDraftFields(func(feature string) bool { return feature == "feature-b" })
	require.Empty(t, inner.Value)
	var nilDraft *child.Draft
	require.NotPanics(t, func() { nilDraft.OmitDraftFields(func(string) bool { return false }) })
	empty := &child.Draft_Empty{}
	empty.ProtoReflect().SetUnknown([]byte{0x98, 0x06, 0x01})
	empty.OmitDraftFields(func(feature string) bool { return feature == "feature-a" })
	require.NotEmpty(t, empty.ProtoReflect().GetUnknown())
	empty.OmitDraftFields(func(string) bool { return false })
	require.Empty(t, empty.ProtoReflect().GetUnknown())
}

func TestOmitConjunctiveMessageGates(t *testing.T) {
	for _, missing := range []string{"", "feature-a", "feature-b", "feature-c"} {
		t.Run("missing="+missing, func(t *testing.T) {
			m := &parent.DraftReferences{GatedInner: &child.Draft_Inner{Value: "inner"}}
			m.OmitDraftFields(func(feature string) bool { return feature != missing })
			if missing == "" {
				require.Equal(t, "inner", m.GatedInner.Value)
			} else {
				require.Nil(t, m.GatedInner)
			}
		})
	}
}

func fixtureRequest() *pluginpb.CodeGeneratorRequest {
	var files []*descriptorpb.FileDescriptorProto
	seen := map[string]bool{}
	var visit func(protoreflect.FileDescriptor)
	visit = func(f protoreflect.FileDescriptor) {
		if seen[f.Path()] {
			return
		}
		seen[f.Path()] = true
		for i := 0; i < f.Imports().Len(); i++ {
			visit(f.Imports().Get(i))
		}
		files = append(files, protodesc.ToFileDescriptorProto(f))
	}
	visit(parent.File_parent_message_proto)
	return &pluginpb.CodeGeneratorRequest{ProtoFile: files, FileToGenerate: []string{"annotations/draft.proto", "child/message.proto", "parent/message.proto"}, Parameter: proto.String("paths=source_relative")}
}

func TestDraftGenerationModes(t *testing.T) {
	for _, enabled := range []bool{false, true} {
		gen, err := (protogen.Options{}).New(fixtureRequest())
		require.NoError(t, err)
		require.NoError(t, generate(gen, enabled))
		response := gen.Response()
		require.Empty(t, response.GetError())
		count := 0
		var helpers, omissions strings.Builder
		for _, f := range response.File {
			if strings.HasSuffix(f.GetName(), ".draft.pb.go") {
				count++
				omissions.WriteString(f.GetContent())
			} else if strings.HasSuffix(f.GetName(), ".go-helpers.pb.go") {
				helpers.WriteString(f.GetContent())
			}
		}
		require.Contains(t, helpers.String(), "func (val *Root) Marshal()")
		require.NotContains(t, helpers.String(), "func (val *Root_Nested) Marshal()")
		require.NotContains(t, helpers.String(), "func Draft_StateFromString(")
		require.Contains(t, helpers.String(), "func LevelFromString(")
		if enabled {
			require.Equal(t, 2, count)
			require.Contains(t, omissions.String(), "func (m *Root_Nested) OmitDraftFields(")
		} else {
			require.Zero(t, count)
		}
		require.NotZero(t, response.GetSupportedFeatures()&uint64(pluginpb.CodeGeneratorResponse_FEATURE_PROTO3_OPTIONAL))
	}
	req := fixtureRequest()
	req.ProtoFile = req.ProtoFile[:1]
	req.FileToGenerate = []string{req.ProtoFile[0].GetName()}
	gen, err := (protogen.Options{}).New(req)
	require.NoError(t, err)
	require.NoError(t, generate(gen, false))
	require.ErrorContains(t, generateDraft(gen), "requires")
}

func TestDraftPlannerFailures(t *testing.T) {
	cases := []struct {
		name, want string
		mutate     func(*pluginpb.CodeGeneratorRequest)
		mutateGen  func(*protogen.Plugin)
	}{
		{name: "schema type", want: "malformed", mutate: func(r *pluginpb.CodeGeneratorRequest) {
			r.ProtoFile[1].Extension[0].Type = descriptorpb.FieldDescriptorProto_TYPE_BOOL.Enum()
		}},
		{name: "schema extendee", want: "malformed", mutate: func(r *pluginpb.CodeGeneratorRequest) {
			r.ProtoFile[1].Extension[0].Extendee = proto.String(".google.protobuf.FileOptions")
		}},
		{name: "dependency", want: "affected dependency", mutate: func(r *pluginpb.CodeGeneratorRequest) { r.FileToGenerate = []string{"parent/message.proto"} }},
		{name: "collision", want: "name collision", mutate: func(r *pluginpb.CodeGeneratorRequest) {
			r.ProtoFile[len(r.ProtoFile)-1].MessageType[0].Field[0].Name = proto.String("omit_draft_fields")
		}},
		{name: "opaque", want: "open Go structs", mutateGen: func(g *protogen.Plugin) {
			g.FilesByPath["parent/message.proto"].Messages[0].APILevel = gofeaturespb.GoFeatures_API_OPAQUE
		}},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			req := fixtureRequest()
			if tc.mutate != nil {
				tc.mutate(req)
			}
			gen, err := (protogen.Options{}).New(req)
			require.NoError(t, err)
			if tc.mutateGen != nil {
				tc.mutateGen(gen)
			}
			_, err = draftPlan(gen)
			require.ErrorContains(t, err, tc.want)
		})
	}
	req := fixtureRequest()
	for _, file := range req.ProtoFile {
		for _, m := range file.MessageType {
			for _, f := range m.Field {
				if f.Options != nil {
					raw, err := proto.Marshal(f.Options)
					require.NoError(t, err)
					f.Options = &descriptorpb.FieldOptions{}
					f.Options.ProtoReflect().SetUnknown(raw)
				}
			}
		}
	}
	gen, err := (protogen.Options{}).New(req)
	require.NoError(t, err)
	plan, err := draftPlan(gen)
	require.NoError(t, err)
	require.Equal(t, "feature-a", plan.features["omission.child.Leaf.a"])
	req = fixtureRequest()
	f := req.ProtoFile[len(req.ProtoFile)-1].MessageType[0].Field[0]
	f.Options = &descriptorpb.FieldOptions{}
	f.Options.ProtoReflect().SetUnknown(protowire.AppendVarint(protowire.AppendTag(nil, 98008, protowire.VarintType), 1))
	gen, err = (protogen.Options{}).New(req)
	require.NoError(t, err)
	_, err = draftPlan(gen)
	require.ErrorContains(t, err, "malformed")
}

func TestDraftSharedClassification(t *testing.T) {
	corpus := "testdata/classification"
	entries, err := os.ReadDir(corpus)
	require.NoError(t, err)
	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}
		t.Run(entry.Name(), func(t *testing.T) {
			dir := filepath.Join(corpus, entry.Name())
			var expected struct {
				Error    string            `json:"error"`
				Fields   map[string]string `json:"fields"`
				Messages map[string]string `json:"messages"`
			}
			raw, err := os.ReadFile(filepath.Join(dir, "expected.json"))
			require.NoError(t, err)
			require.NoError(t, json.Unmarshal(raw, &expected))
			var inputs []string
			require.NoError(t, filepath.WalkDir(filepath.Join(dir, "input"), func(path string, d os.DirEntry, err error) error {
				if err != nil {
					return err
				}
				if !d.IsDir() && strings.HasSuffix(path, ".proto") {
					inputs = append(inputs, path)
				}
				return nil
			}))
			output := filepath.Join(t.TempDir(), "schema.pb")
			args := []string{"-I" + filepath.Join(dir, "input"), "--include_imports", "--descriptor_set_out=" + output}
			args = append(args, inputs...)
			out, err := exec.Command("protoc", args...).CombinedOutput()
			require.NoError(t, err, "%s", out)
			raw, err = os.ReadFile(output)
			require.NoError(t, err)
			var set descriptorpb.FileDescriptorSet
			require.NoError(t, proto.Unmarshal(raw, &set))
			req := &pluginpb.CodeGeneratorRequest{ProtoFile: set.File}
			for _, f := range set.File {
				if f.GetName() != "google/protobuf/descriptor.proto" {
					if f.Options == nil {
						f.Options = &descriptorpb.FileOptions{}
					}
					f.Options.GoPackage = proto.String("go.temporal.io/api/testclassification/" + strings.ReplaceAll(f.GetPackage(), ".", "/"))
					req.FileToGenerate = append(req.FileToGenerate, f.GetName())
				}
			}
			gen, err := (protogen.Options{}).New(req)
			require.NoError(t, err)
			plan, err := draftPlan(gen)
			if expected.Error != "" {
				require.ErrorContains(t, err, expected.Error)
				return
			}
			require.NoError(t, err)
			actual := map[string]string{}
			for name, feature := range plan.features {
				actual[string(name)] = feature
			}
			if expected.Fields == nil {
				expected.Fields = map[string]string{}
			}
			require.Equal(t, expected.Fields, actual)
			if expected.Messages == nil {
				expected.Messages = map[string]string{}
			}
			actual = map[string]string{}
			for name, feature := range plan.messageFeatures {
				actual[string(name)] = feature
			}
			require.Equal(t, expected.Messages, actual)
		})
	}
}
