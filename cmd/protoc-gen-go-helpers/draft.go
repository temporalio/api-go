package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"

	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/encoding/protowire"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/gofeaturespb"
)

const draftPrefix = "temporal.api.protometa.v1."

// The typed traversal and direct clears adapt the SaaS protoc-gen-go-codegen
// sanitizer emitter, without its version policy, Merge, or private dependencies.
type omissionPlan struct {
	features        map[protoreflect.FullName]string
	messageFeatures map[protoreflect.FullName]string
	affected        map[protoreflect.FullName]bool
	messages        []*protogen.Message
}

func childMessage(f *protogen.Field) *protogen.Message {
	if f.Desc.IsMap() {
		return f.Message.Fields[1].Message
	}
	return f.Message
}

func (p *omissionPlan) typeFeatures(d protoreflect.Descriptor) []string {
	var features []string
	for current := d; current != nil; current = current.Parent() {
		if feature := p.messageFeatures[current.FullName()]; feature != "" {
			features = append(features, feature)
		}
	}
	return features
}

func (p *omissionPlan) fieldFeatures(f *protogen.Field) []string {
	var features []string
	if feature := p.features[f.Desc.FullName()]; feature != "" {
		features = append(features, feature)
	}
	if f.Desc.IsMap() {
		f = f.Message.Fields[1]
	}
	if f.Message != nil {
		features = append(features, p.typeFeatures(f.Message.Desc)...)
	} else if f.Enum != nil {
		features = append(features, p.typeFeatures(f.Enum.Desc)...)
	}
	return features
}

func disabledCondition(features []string) string {
	seen := map[string]bool{}
	var conditions []string
	for _, feature := range features {
		if !seen[feature] {
			conditions = append(conditions, "!enabled("+strconv.Quote(feature)+")")
			seen[feature] = true
		}
	}
	return strings.Join(conditions, " || ")
}

func draftPlan(gen *protogen.Plugin) (*omissionPlan, error) {
	expected := map[string]string{
		"draft":            "google.protobuf.FieldOptions",
		"draft_message":    "google.protobuf.MessageOptions",
		"draft_enum":       "google.protobuf.EnumOptions",
		"draft_enum_value": "google.protobuf.EnumValueOptions",
		"draft_service":    "google.protobuf.ServiceOptions",
		"draft_method":     "google.protobuf.MethodOptions",
		"draft_file":       "google.protobuf.FileOptions",
		"draft_oneof":      "google.protobuf.OneofOptions",
		"draft_extension":  "google.protobuf.FieldOptions",
	}
	schema := map[protoreflect.FullName]map[protowire.Number]string{}
	schemaFiles := map[string]bool{}
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
			schemaFiles[e.ParentFile().Path()] = true
			if e.Kind() != protoreflect.StringKind || e.Cardinality() != protoreflect.Optional || string(e.ContainingMessage().FullName()) != extendee {
				return fmt.Errorf("%s: malformed draft option schema", e.FullName())
			}
			if name == "draft" {
				found = true
			}
			key := e.ContainingMessage().FullName()
			if schema[key] == nil {
				schema[key] = map[protowire.Number]string{}
			}
			schema[key][protowire.Number(e.Number())] = name
			return nil
		}); err != nil {
			return nil, err
		}
	}
	for _, f := range gen.Files {
		for i := range f.Desc.Imports().Len() {
			if imp := f.Desc.Imports().Get(i); imp.IsPublic && schemaFiles[imp.Path()] {
				return nil, fmt.Errorf("%s: public import of %s is not supported; import it directly", f.Desc.Path(), imp.Path())
			}
		}
	}
	if !found {
		return nil, fmt.Errorf("draft_omission requires the temporal.api.protometa.v1.draft field option schema")
	}
	p := &omissionPlan{
		features: map[protoreflect.FullName]string{}, messageFeatures: map[protoreflect.FullName]string{},
		affected: map[protoreflect.FullName]bool{},
	}
	validate := func(d protoreflect.Descriptor) error {
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
			option, recognized := schema[options.ProtoReflect().Descriptor().FullName()][n]
			if recognized {
				if t != protowire.BytesType {
					return fmt.Errorf("%s: malformed draft option value", d.FullName())
				}
				value, k := protowire.ConsumeString(raw)
				if k < 0 {
					return protowire.ParseError(k)
				}
				f, isField := d.(protoreflect.FieldDescriptor)
				_, isMessage := d.(protoreflect.MessageDescriptor)
				if !(option == "draft" && isField && !f.IsExtension()) && !(option == "draft_message" && isMessage) {
					return fmt.Errorf("%s: unsupported draft annotation on %s", d.FullName(), annotationPlacement(d))
				}
				if value == "" || strings.ContainsFunc(value, unicode.IsSpace) {
					return fmt.Errorf("%s: draft feature must be nonempty and contain no whitespace", d.FullName())
				}
				if isMessage {
					p.messageFeatures[d.FullName()] = value
				} else {
					p.features[d.FullName()] = value
				}
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
		if err := visit(f.Desc, validate); err != nil {
			return nil, err
		}
		collect(f.Messages)
		generated[f.Desc.Path()] = f.Generate
	}
	for _, file := range gen.Files {
		if err := visit(file.Desc, func(d protoreflect.Descriptor) error {
			extension, ok := d.(protoreflect.ExtensionDescriptor)
			if ok && extension.IsExtension() && (len(p.typeFeatures(extension.Message())) > 0 || len(p.typeFeatures(extension.Enum())) > 0 ||
				len(p.typeFeatures(extension.ContainingMessage())) > 0 || len(p.typeFeatures(extension.Parent())) > 0) {
				return fmt.Errorf("%s: extension references draft type", extension.FullName())
			}
			return nil
		}); err != nil {
			return nil, err
		}
	}
	for changed := true; changed; {
		changed = false
		for _, m := range p.messages {
			if p.affected[m.Desc.FullName()] {
				continue
			}
			if len(p.typeFeatures(m.Desc)) > 0 {
				p.affected[m.Desc.FullName()] = true
				changed = true
				continue
			}
			for _, f := range m.Fields {
				child := childMessage(f)
				if len(p.fieldFeatures(f)) > 0 || child != nil && p.affected[child.Desc.FullName()] {
					p.affected[m.Desc.FullName()] = true
					changed = true
					break
				}
			}
		}
	}
	for _, file := range gen.Files {
		for _, service := range file.Services {
			for _, method := range service.Methods {
				if len(p.typeFeatures(method.Input.Desc)) > 0 || len(p.typeFeatures(method.Output.Desc)) > 0 {
					return nil, fmt.Errorf("%s: RPC references draft type", method.Desc.FullName())
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

func annotationPlacement(d protoreflect.Descriptor) string {
	switch d := d.(type) {
	case protoreflect.FileDescriptor:
		return "file"
	case protoreflect.MessageDescriptor:
		return "message"
	case protoreflect.EnumDescriptor:
		return "enum"
	case protoreflect.EnumValueDescriptor:
		return "enum value"
	case protoreflect.OneofDescriptor:
		return "oneof"
	case protoreflect.ServiceDescriptor:
		return "service"
	case protoreflect.MethodDescriptor:
		return "method"
	case protoreflect.FieldDescriptor:
		if d.IsExtension() {
			return "extension"
		}
		return "field"
	default:
		return "unknown declaration"
	}
}

func generateDraft(gen *protogen.Plugin) error {
	p, err := draftPlan(gen)
	if err != nil {
		return err
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
			g.P("func (m *", m.GoIdent, ") OmitDraftFields(enabled func(string) bool) {")
			g.P("if m == nil { return }")
			if condition := disabledCondition(p.typeFeatures(m.Desc)); condition != "" {
				g.P("if ", condition, " {")
				g.P(g.QualifiedGoIdent(protogen.GoIdent{GoName: "Reset", GoImportPath: "google.golang.org/protobuf/proto"}), "(m)")
				g.P("return")
				g.P("}")
			}
			for _, f := range m.Fields {
				child := childMessage(f)
				traverse := child != nil && p.affected[child.Desc.FullName()]
				condition := disabledCondition(p.fieldFeatures(f))
				if condition == "" && !traverse {
					continue
				}
				oneof := f.Oneof != nil && !f.Oneof.Desc.IsSynthetic()
				if oneof {
					binding := "_"
					if traverse {
						binding = "v"
					}
					g.P("if ", binding, ", ok := m.", f.Oneof.GoName, ".(*", f.GoIdent, "); ok {")
				}
				if condition != "" {
					g.P("if ", condition, " {")
					if oneof {
						g.P("m.", f.Oneof.GoName, " = nil")
					} else {
						g.P("m.", f.GoName, " = ", clearValue(f))
					}
					if traverse {
						g.P("} else {")
					} else {
						g.P("}")
					}
				}
				if traverse {
					if oneof {
						g.P("if v != nil {")
						g.P("v.", f.GoName, ".OmitDraftFields(enabled)")
						g.P("}")
					} else if f.Desc.IsMap() || f.Desc.IsList() {
						g.P("for _, v := range m.", f.GoName, " { v.OmitDraftFields(enabled) }")
					} else {
						g.P("m.", f.GoName, ".OmitDraftFields(enabled)")
					}
					if condition != "" {
						g.P("}")
					}
				}
				if oneof {
					g.P("}")
				}
			}
			g.P("}")
		}
	}
	return nil
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
