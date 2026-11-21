package main

import (
	"fmt"
	"strconv"
	"strings"

	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/encoding/protowire"
	"google.golang.org/protobuf/proto"
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

func fieldValue(f *protogen.Field) *protogen.Field {
	if f.Desc.IsMap() {
		return f.Message.Fields[1]
	}
	return f
}

func (p *omissionPlan) feature(d protoreflect.Descriptor) string {
	for current := d; current != nil; current = current.Parent() {
		if feature := p.features[current.FullName()]; feature != "" {
			return feature
		}
		if field, ok := current.(protoreflect.FieldDescriptor); ok {
			if oneof := field.ContainingOneof(); oneof != nil && !oneof.IsSynthetic() {
				if feature := p.features[oneof.FullName()]; feature != "" {
					return feature
				}
			}
		}
	}
	return ""
}

func (p *omissionPlan) draftEnumValues(f *protogen.Field) []draftEnumValue {
	if enum := fieldValue(f).Enum; enum != nil {
		return p.enumValues[enum.Desc.FullName()]
	}
	return nil
}

func draftPlan(gen *protogen.Plugin) (*omissionPlan, error) {
	expected := map[string]string{
		"draft":            "google.protobuf.FieldOptions",
		"draft_message":    "google.protobuf.MessageOptions",
		"draft_enum":       "google.protobuf.EnumOptions",
		"draft_enum_value": "google.protobuf.EnumValueOptions",
		"draft_service":    "google.protobuf.ServiceOptions",
		"draft_method":     "google.protobuf.MethodOptions",
		"draft_oneof":      "google.protobuf.OneofOptions",
	}
	schema := map[protoreflect.FullName]map[protowire.Number]bool{}
	found := false
	for _, f := range gen.Files {
		if err := visit(f.Desc, func(d protoreflect.Descriptor) error {
			e, ok := d.(protoreflect.ExtensionDescriptor)
			if !ok || !e.IsExtension() {
				return nil
			}
			name := strings.TrimPrefix(string(e.FullName()), draftPrefix)
			extendee, known := expected[name]
			if !known || string(e.FullName()) != draftPrefix+name {
				return nil
			}
			if e.Kind() != protoreflect.StringKind || e.Cardinality() != protoreflect.Optional || string(e.ContainingMessage().FullName()) != extendee {
				return fmt.Errorf("%s: malformed draft option schema", e.FullName())
			}
			if name == "draft" {
				found = true
			}
			key := e.ContainingMessage().FullName()
			if schema[key] == nil {
				schema[key] = map[protowire.Number]bool{}
			}
			schema[key][protowire.Number(e.Number())] = true
			return nil
		}); err != nil {
			return nil, err
		}
	}
	if !found {
		return nil, fmt.Errorf("draft_omission requires the temporal.api.protometa.v1.draft field option schema")
	}
	p := &omissionPlan{
		features:   map[protoreflect.FullName]string{},
		enumValues: map[protoreflect.FullName][]draftEnumValue{},
		affected:   map[protoreflect.FullName]bool{},
	}
	classify := func(d protoreflect.Descriptor) error {
		options := d.Options()
		if options == nil {
			return nil
		}
		raw, err := proto.Marshal(options)
		if err != nil {
			return err
		}
		for len(raw) > 0 {
			n, t, k := protowire.ConsumeTag(raw)
			if k < 0 {
				return protowire.ParseError(k)
			}
			raw = raw[k:]
			if schema[options.ProtoReflect().Descriptor().FullName()][n] {
				if t != protowire.BytesType {
					return fmt.Errorf("%s: malformed draft option value", d.FullName())
				}
				value, k := protowire.ConsumeString(raw)
				if k < 0 {
					return protowire.ParseError(k)
				}
				p.features[d.FullName()] = value
			}
			k = protowire.ConsumeFieldValue(n, t, raw)
			if k < 0 {
				return protowire.ParseError(k)
			}
			raw = raw[k:]
		}
		return nil
	}
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
			if err := classify(d); err != nil {
				return err
			}
			value, ok := d.(protoreflect.EnumValueDescriptor)
			if feature := p.features[d.FullName()]; ok && feature != "" {
				enum := value.Parent().(protoreflect.EnumDescriptor)
				if value.Number() == 0 {
					return fmt.Errorf("%s: draft enum value cannot have number zero", value.FullName())
				}
				for i := range enum.Values().Len() {
					other := enum.Values().Get(i)
					if other.Number() == value.Number() && other.FullName() != value.FullName() {
						return fmt.Errorf("%s: aliases involving draft enum values are not allowed (also used by %s)", value.FullName(), other.FullName())
					}
				}
				p.enumValues[enum.FullName()] = append(p.enumValues[enum.FullName()], draftEnumValue{value.Number(), feature})
			}
			return nil
		}); err != nil {
			return nil, err
		}
		collect(f.Messages)
		generated[f.Desc.Path()] = f.Generate
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
				child := fieldValue(f).Message
				if p.feature(f.Desc) != "" || len(p.draftEnumValues(f)) != 0 || child != nil && p.affected[child.Desc.FullName()] {
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

func generateDraft(gen *protogen.Plugin) error {
	p, err := draftPlan(gen)
	if err != nil {
		return err
	}
	for _, m := range p.messages {
		for _, field := range m.Fields {
			if (field.Desc.IsList() || field.Desc.IsMap()) && len(p.draftEnumValues(field)) != 0 {
				return fmt.Errorf("%s: draft enum values in collections are unsupported", field.Desc.FullName())
			}
		}
	}
	for _, file := range gen.Files {
		if !file.Generate {
			continue
		}
		var g *protogen.GeneratedFile
		for _, m := range p.messages {
			if m.Desc.ParentFile().Path() != file.Desc.Path() || !p.affected[m.Desc.FullName()] {
				continue
			}
			if g == nil {
				g = gen.NewGeneratedFile(file.GeneratedFilenamePrefix+".draft.pb.go", file.GoImportPath)
				g.P("// Code generated by protoc-gen-go-helpers. DO NOT EDIT.")
				g.P()
				g.P("package ", file.GoPackageName)
			}
			p.generateOmitter(g, m)
		}
	}
	return nil
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
	child := fieldValue(f).Message
	traverse := child != nil && p.affected[child.Desc.FullName()]
	values := p.draftEnumValues(f)
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
		g.P("if !enabled(", strconv.Quote(feature), ") {")
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
			generateChildOmission(g, f, field)
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

func generateChildOmission(g *protogen.GeneratedFile, f *protogen.Field, field string) {
	if f.Desc.IsList() || f.Desc.IsMap() {
		g.P("for _, child := range ", field, " {")
		field = "child"
	}
	g.P("if ", field, ".OmitDraftFields(enabled) {")
	g.P("removedEnumValue = true")
	g.P("}")
	if f.Desc.IsList() || f.Desc.IsMap() {
		g.P("}")
	}
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
