package main

import (
	"cmp"
	"fmt"
	"regexp"
	"slices"
	"strings"

	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"
	"google.golang.org/protobuf/types/descriptorpb"
)

const (
	httpRuleOption      = "google.api.http"
	requestHeaderOption = draftPrefix + "request_header"
	fieldPathPattern    = `([a-zA-Z_][a-zA-Z_0-9]*(?:\.[a-zA-Z_][a-zA-Z_0-9]*)*)`
)

var (
	httpPathVariable = regexp.MustCompile(`\{` + fieldPathPattern + `(?:=[^{}]+)?\}`)
	headerVariable   = regexp.MustCompile(`\{` + fieldPathPattern + `\}`)
)

func validRPCOption(option protoreflect.ExtensionDescriptor) bool {
	if option.ContainingMessage().FullName() != "google.protobuf.MethodOptions" || option.Kind() != protoreflect.MessageKind {
		return false
	}
	fields := option.Message().Fields()
	isString := func(field protoreflect.FieldDescriptor) bool {
		return field != nil && field.Kind() == protoreflect.StringKind && field.Cardinality() == protoreflect.Optional
	}
	if option.FullName() == requestHeaderOption {
		return option.Cardinality() == protoreflect.Repeated && isString(fields.ByName("value"))
	}
	if option.Cardinality() != protoreflect.Optional {
		return false
	}
	for _, name := range []protoreflect.Name{"get", "put", "post", "delete", "patch", "body", "response_body"} {
		if !isString(fields.ByName(name)) {
			return false
		}
	}
	custom := fields.ByName("custom")
	bindings := fields.ByName("additional_bindings")
	return custom != nil && custom.Kind() == protoreflect.MessageKind && custom.Cardinality() == protoreflect.Optional &&
		isString(custom.Message().Fields().ByName("path")) &&
		bindings != nil && bindings.Kind() == protoreflect.MessageKind && bindings.Cardinality() == protoreflect.Repeated &&
		bindings.Message() == option.Message()
}

func (p *omissionPlan) validateMethodFieldReferences(method *protogen.Method, extensions *protoregistry.Types) error {
	raw, err := proto.Marshal(method.Desc.Options())
	if err != nil {
		return fmt.Errorf("%s: encoding RPC options: %w", method.Desc.FullName(), err)
	}
	options := &descriptorpb.MethodOptions{}
	if err := (proto.UnmarshalOptions{Resolver: extensions}).Unmarshal(raw, options); err != nil {
		return fmt.Errorf("%s: decoding RPC options: %w", method.Desc.FullName(), err)
	}
	message := options.ProtoReflect()
	for _, field := range populatedFieldsByNumber(message) {
		value := message.Get(field)
		switch field.FullName() {
		case httpRuleOption:
			if path := p.draftHTTPField(method, value.Message()); path != "" {
				return fmt.Errorf("%s: http rule references draft field %s", method.Desc.FullName(), path)
			}
		case requestHeaderOption:
			headers := value.List()
			for i := range headers.Len() {
				header := headers.Get(i).Message()
				template := header.Get(header.Descriptor().Fields().ByName("value")).String()
				if path := p.draftTemplateField(method.Input, template, headerVariable); path != "" {
					return fmt.Errorf("%s: request header references draft field %s", method.Desc.FullName(), path)
				}
			}
		}
	}
	return nil
}

func (p *omissionPlan) draftHTTPField(method *protogen.Method, rule protoreflect.Message) string {
	for _, field := range populatedFieldsByNumber(rule) {
		value := rule.Get(field)
		var path string
		switch field.Name() {
		case "get", "put", "post", "delete", "patch":
			path = p.draftTemplateField(method.Input, value.String(), httpPathVariable)
		case "custom":
			pattern := value.Message()
			template := pattern.Get(pattern.Descriptor().Fields().ByName("path")).String()
			path = p.draftTemplateField(method.Input, template, httpPathVariable)
		case "body":
			if p.referencesDraftField(method.Input, value.String()) {
				path = value.String()
			}
		case "response_body":
			if p.referencesDraftField(method.Output, value.String()) {
				path = value.String()
			}
		case "additional_bindings":
			bindings := value.List()
			for i := range bindings.Len() {
				if path = p.draftHTTPField(method, bindings.Get(i).Message()); path != "" {
					break
				}
			}
		}
		if path != "" {
			return path
		}
	}
	return ""
}

func populatedFieldsByNumber(message protoreflect.Message) []protoreflect.FieldDescriptor {
	var fields []protoreflect.FieldDescriptor
	message.Range(func(field protoreflect.FieldDescriptor, _ protoreflect.Value) bool {
		fields = append(fields, field)
		return true
	})
	slices.SortFunc(fields, func(a, b protoreflect.FieldDescriptor) int {
		return cmp.Compare(a.Number(), b.Number())
	})
	return fields
}

func (p *omissionPlan) draftTemplateField(message *protogen.Message, template string, variable *regexp.Regexp) string {
	for _, match := range variable.FindAllStringSubmatch(template, -1) {
		if p.referencesDraftField(message, match[1]) {
			return match[1]
		}
	}
	return ""
}

func (p *omissionPlan) referencesDraftField(message *protogen.Message, path string) bool {
	for name := range strings.SplitSeq(path, ".") {
		if message == nil {
			return false
		}
		i := slices.IndexFunc(message.Fields, func(field *protogen.Field) bool {
			return field.Desc.Name() == protoreflect.Name(name)
		})
		if i < 0 {
			return false
		}
		field := message.Fields[i]
		if len(p.fieldFeatures(field)) > 0 {
			return true
		}
		message = field.Message
	}
	return false
}
