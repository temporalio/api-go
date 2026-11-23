package main

import (
	"encoding/json"
	"fmt"
	"maps"
	"os"
	"slices"
	"strconv"
	"strings"
	"unicode"

	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/gofeaturespb"
)

const draftPrefix = "temporal.api.protometa.v1."

// omissionPlan drives the typed traversal and direct clears adapted from the SaaS
// protoc-gen-go-codegen sanitizer emitter independently of its version policy,
// Merge, and private dependencies.
type omissionPlan struct {
	features   map[protoreflect.FullName]string
	enumValues map[protoreflect.FullName][]draftEnumValue
	affected   map[protoreflect.FullName]bool
	messages   []*protogen.Message
}

type draftEnumValue struct {
	number  protoreflect.EnumNumber
	feature string
}

func (p *omissionPlan) feature(d protoreflect.Descriptor) string {
	return p.features[d.FullName()]
}

func (p *omissionPlan) draftEnumValues(f *protogen.Field) []draftEnumValue {
	if enum := f.Enum; enum != nil {
		return p.enumValues[enum.Desc.FullName()]
	}
	return nil
}

func (p *omissionPlan) retainedFieldOmission(f *protogen.Field) (bool, []draftEnumValue) {
	if f.Desc.IsList() || f.Desc.IsMap() {
		return false, nil
	}
	return f.Message != nil && p.affected[f.Message.Desc.FullName()], p.draftEnumValues(f)
}

func loadDraftManifest(path string) (map[protoreflect.FullName]string, error) {
	if path == "" {
		return nil, fmt.Errorf("draft_omission requires draft_manifest")
	}
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("read draft manifest: %w", err)
	}
	var manifest struct {
		Version int                              `json:"version"`
		Gates   map[protoreflect.FullName]string `json:"gates"`
	}
	if err := json.Unmarshal(data, &manifest); err != nil {
		return nil, fmt.Errorf("parse draft manifest: %w", err)
	}
	if manifest.Version != 2 || manifest.Gates == nil {
		return nil, fmt.Errorf("draft manifest requires version 2 and a gates object")
	}
	for _, name := range slices.Sorted(maps.Keys(manifest.Gates)) {
		if gate := manifest.Gates[name]; !name.IsValid() || gate == "" || strings.ContainsFunc(gate, unicode.IsSpace) {
			return nil, fmt.Errorf("invalid draft manifest entry for %q", name)
		}
	}
	return manifest.Gates, nil
}

func draftPlan(gen *protogen.Plugin, gates map[protoreflect.FullName]string) (*omissionPlan, error) {
	p := &omissionPlan{
		features:   gates,
		enumValues: map[protoreflect.FullName][]draftEnumValue{},
		affected:   map[protoreflect.FullName]bool{},
	}
	unmatched := maps.Clone(gates)
	var collect func([]*protogen.Message)
	collect = func(ms []*protogen.Message) {
		for _, m := range ms {
			if !m.Desc.IsMapEntry() {
				p.messages = append(p.messages, m)
			}
			collect(m.Messages)
		}
	}
	generated := map[string]bool{}
	for _, f := range gen.Files {
		if err := visit(f.Desc, func(d protoreflect.Descriptor) error {
			if _, isFile := d.(protoreflect.FileDescriptor); !isFile {
				delete(unmatched, d.FullName())
			}
			value, ok := d.(protoreflect.EnumValueDescriptor)
			if feature := p.feature(d); ok && feature != "" && p.feature(value.Parent()) == "" {
				enum := value.Parent().FullName()
				p.enumValues[enum] = append(p.enumValues[enum], draftEnumValue{value.Number(), feature})
			}
			return nil
		}); err != nil {
			return nil, err
		}
		collect(f.Messages)
		generated[f.Desc.Path()] = f.Generate
	}
	if len(unmatched) > 0 {
		name := slices.Sorted(maps.Keys(unmatched))[0]
		return nil, fmt.Errorf("draft manifest references unknown declaration %s", name)
	}
	for changed := true; changed; {
		changed = false
		for _, m := range p.messages {
			if p.affected[m.Desc.FullName()] {
				continue
			}
			if p.feature(m.Desc) != "" {
				p.affected[m.Desc.FullName()] = true
				changed = true
				continue
			}
			for _, f := range m.Fields {
				traverse, enumValues := p.retainedFieldOmission(f)
				if p.feature(f.Desc) != "" || traverse || len(enumValues) != 0 {
					p.affected[m.Desc.FullName()] = true
					changed = true
					break
				}
			}
		}
	}
	for _, m := range p.messages {
		if !p.affected[m.Desc.FullName()] {
			continue
		}
		if !generated[m.Desc.ParentFile().Path()] {
			return nil, fmt.Errorf("%s: affected dependency cannot receive omission methods; include its file in generation", m.Desc.FullName())
		}
		if m.APILevel != gofeaturespb.GoFeatures_API_OPEN {
			return nil, fmt.Errorf("%s: draft omission requires open Go structs", m.Desc.FullName())
		}
		for _, f := range m.Fields {
			if f.GoName == "OmitDraftFields" || f.Oneof != nil && f.Oneof.GoName == "OmitDraftFields" {
				return nil, fmt.Errorf("%s: OmitDraftFields name collision", m.Desc.FullName())
			}
		}
	}
	return p, nil
}

// visit calls fn on d and then on every declaration lexically nested in it.
func visit(d protoreflect.Descriptor, fn func(protoreflect.Descriptor) error) error {
	if err := fn(d); err != nil {
		return err
	}
	var children []protoreflect.Descriptor
	switch d := d.(type) {
	case protoreflect.FileDescriptor:
		children = add(add(add(add(children, d.Messages()), d.Enums()), d.Services()), d.Extensions())
	case protoreflect.MessageDescriptor:
		children = add(add(add(add(add(children, d.Fields()), d.Oneofs()), d.Messages()), d.Enums()), d.Extensions())
	case protoreflect.EnumDescriptor:
		children = add(children, d.Values())
	case protoreflect.ServiceDescriptor:
		children = add(children, d.Methods())
	}
	for _, child := range children {
		if err := visit(child, fn); err != nil {
			return err
		}
	}
	return nil
}

func add[T protoreflect.Descriptor](out []protoreflect.Descriptor, list interface {
	Len() int
	Get(int) T
}) []protoreflect.Descriptor {
	for i := range list.Len() {
		out = append(out, list.Get(i))
	}
	return out
}

func generateDraft(gen *protogen.Plugin, gates map[protoreflect.FullName]string) error {
	p, err := draftPlan(gen, gates)
	if err != nil {
		return err
	}
	for _, file := range gen.Files {
		if !file.Generate {
			continue
		}
		var g *protogen.GeneratedFile
		output := func() *protogen.GeneratedFile {
			if g == nil {
				g = gen.NewGeneratedFile(file.GeneratedFilenamePrefix+".draft.pb.go", file.GoImportPath)
				g.P("// Code generated by protoc-gen-go-helpers. DO NOT EDIT.")
				g.P()
				g.P("package ", file.GoPackageName)
			}
			return g
		}
		for _, m := range p.messages {
			if m.Desc.ParentFile().Path() != file.Desc.Path() || !p.affected[m.Desc.FullName()] {
				continue
			}
			p.generateOmitter(output(), m)
		}
		for _, service := range file.Services {
			p.generateMethodLookup(output(), service)
		}
		if file.Desc.Package()+"." == draftPrefix && file.Desc.Extensions().ByName("draft") != nil {
			p.generateTypeLookup(output())
		}
	}
	return nil
}

// generateMethodLookup emits a switch from a full gRPC method name to how draft annotations
// affect that RPC, so a server can gate and filter calls without reading descriptors.
func (p *omissionPlan) generateMethodLookup(g *protogen.GeneratedFile, service *protogen.Service) {
	name := service.GoName + "DraftMethod"
	g.P()
	g.P("// ", name, " describes how draft annotations affect the ", service.GoName, " RPC with the")
	g.P("// given full gRPC method name. feature is the RPC's draft feature, or \"\" for a stable")
	g.P("// RPC. request and response report whether the request and response types have")
	g.P("// OmitDraftFields methods. RPCs that draft annotations do not affect return zero values.")
	g.P("func ", name, "(fullMethod string) (feature string, request, response bool) {")
	var cases []string
	for _, method := range service.Methods {
		feature := p.feature(method.Desc)
		request := p.affected[method.Input.Desc.FullName()]
		response := p.affected[method.Output.Desc.FullName()]
		if feature == "" && !request && !response {
			continue
		}
		cases = append(cases, fmt.Sprintf("case %q:\nreturn %q, %t, %t",
			"/"+string(service.Desc.FullName())+"/"+string(method.Desc.Name()), feature, request, response))
	}
	if len(cases) > 0 {
		g.P("switch fullMethod {")
		for _, c := range cases {
			g.P(c)
		}
		g.P("}")
	}
	g.P(`return "", false, false`)
	g.P("}")
}

// generateTypeLookup emits, next to the draft options, a switch over every message with an
// OmitDraftFields method, so a server can skip decoding google.protobuf.Any payloads, such as
// gRPC status details, that cannot hold draft data.
func (p *omissionPlan) generateTypeLookup(g *protogen.GeneratedFile) {
	var names []string
	for _, m := range p.messages {
		if p.affected[m.Desc.FullName()] {
			names = append(names, string(m.Desc.FullName()))
		}
	}
	slices.Sort(names)
	g.P()
	g.P("// HasOmitDraftFields reports whether the message with the given full name has an")
	g.P("// OmitDraftFields method.")
	g.P("func HasOmitDraftFields(messageFullName string) bool {")
	if len(names) > 0 {
		g.P("switch messageFullName {")
		g.P("case ", strings.Join(quoted(names), ",\n"), ":")
		g.P("return true")
		g.P("}")
	}
	g.P("return false")
	g.P("}")
}

func quoted(names []string) []string {
	out := make([]string, len(names))
	for i, name := range names {
		out[i] = strconv.Quote(name)
	}
	return out
}

func (p *omissionPlan) generateOmitter(g *protogen.GeneratedFile, m *protogen.Message) {
	g.P("// OmitDraftFields clears disabled draft fields and enum values.")
	g.P("// It reports whether it removed a draft enum value from a retained field.")
	g.P("func (m *", m.GoIdent, ") OmitDraftFields(enabled func(string) bool) (removedEnumValue bool) {")
	g.P("if m == nil { return false }")
	rootFeature := p.feature(m.Desc)
	if rootFeature != "" {
		g.P("if !enabled(", strconv.Quote(rootFeature), ") {")
		g.P(g.QualifiedGoIdent(protogen.GoIdent{GoName: "Reset", GoImportPath: "google.golang.org/protobuf/proto"}), "(m)")
		g.P("return false")
		g.P("}")
	}
	for _, f := range m.Fields {
		p.generateFieldOmission(g, f, rootFeature)
	}
	g.P("return removedEnumValue")
	g.P("}")
}

func (p *omissionPlan) generateFieldOmission(g *protogen.GeneratedFile, f *protogen.Field, rootFeature string) {
	traverse, values := p.retainedFieldOmission(f)
	filter := traverse || len(values) != 0
	feature := p.feature(f.Desc)
	if feature == rootFeature {
		feature = ""
	}
	if feature == "" && !filter {
		return
	}

	field := "m." + f.GoName
	clear := field + " = " + clearValue(f)
	oneof := f.Oneof != nil && !f.Oneof.Desc.IsSynthetic()
	if oneof {
		binding := "_"
		if filter {
			binding = "v"
		}
		wrapper := "m." + f.Oneof.GoName
		g.P("if ", binding, ", ok := ", wrapper, ".(*", f.GoIdent, "); ok {")
		field = "v." + f.GoName
		clear = wrapper + " = nil"
	}
	if feature != "" {
		// Ask about the feature only when the field is set, so callers can tell from
		// the predicate's calls whether omission would change anything. A oneof
		// member is set when its wrapper matched above.
		condition := "!enabled(" + strconv.Quote(feature) + ")"
		if !oneof {
			condition = presence(g, f, field) + " && " + condition
		}
		g.P("if ", condition, " {")
		g.P(clear)
		if filter {
			g.P("} else {")
		} else {
			g.P("}")
		}
	}
	if filter {
		if oneof {
			g.P("if v != nil {")
		}
		if traverse {
			generateChildOmission(g, field)
		} else {
			generateEnumOmission(g, f, field, clear, values)
		}
		if oneof {
			g.P("}")
		}
		if feature != "" {
			g.P("}")
		}
	}
	if oneof {
		g.P("}")
	}
}

func generateChildOmission(g *protogen.GeneratedFile, field string) {
	g.P("if ", field, ".OmitDraftFields(enabled) {")
	g.P("removedEnumValue = true")
	g.P("}")
}

func generateEnumOmission(g *protogen.GeneratedFile, f *protogen.Field, field, clear string, values []draftEnumValue) {
	condition := disabledEnumValue(field, values)
	if f.Desc.HasPresence() && (f.Oneof == nil || f.Oneof.Desc.IsSynthetic()) {
		condition = field + " != nil && (" + disabledEnumValue("*"+field, values) + ")"
	}
	g.P("if ", condition, " {")
	g.P(clear)
	g.P("removedEnumValue = true")
	g.P("}")
}

func disabledEnumValue(value string, options []draftEnumValue) string {
	var conditions []string
	for _, option := range options {
		conditions = append(conditions, fmt.Sprintf("(%s == %d && !enabled(%q))", value, option.number, option.feature))
	}
	return strings.Join(conditions, " || ")
}

// presence returns a Go condition that holds when the field is set to a non-default value.
func presence(g *protogen.GeneratedFile, f *protogen.Field, field string) string {
	if f.Desc.IsList() || f.Desc.IsMap() {
		return "len(" + field + ") > 0"
	}
	if f.Desc.HasPresence() {
		return field + " != nil"
	}
	switch f.Desc.Kind() {
	case protoreflect.BoolKind:
		return field
	// Compare bits: protobuf serializes -0.0, which equals 0.
	case protoreflect.FloatKind:
		return g.QualifiedGoIdent(protogen.GoIdent{GoName: "Float32bits", GoImportPath: "math"}) + "(" + field + ") != 0"
	case protoreflect.DoubleKind:
		return g.QualifiedGoIdent(protogen.GoIdent{GoName: "Float64bits", GoImportPath: "math"}) + "(" + field + ") != 0"
	case protoreflect.StringKind:
		return field + ` != ""`
	case protoreflect.BytesKind:
		return "len(" + field + ") > 0"
	case protoreflect.MessageKind, protoreflect.GroupKind:
		return field + " != nil"
	default:
		return field + " != 0"
	}
}

func clearValue(f *protogen.Field) string {
	if f.Desc.IsList() || f.Desc.IsMap() || f.Desc.HasPresence() {
		return "nil"
	}
	switch f.Desc.Kind() {
	case protoreflect.BoolKind:
		return "false"
	case protoreflect.StringKind:
		return `""`
	case protoreflect.BytesKind, protoreflect.MessageKind, protoreflect.GroupKind:
		return "nil"
	default:
		return "0"
	}
}
