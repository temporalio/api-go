// The MIT License
//
// Copyright (c) 2022 Temporal Technologies Inc.  All rights reserved.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package main

import (
	"bytes"
	"fmt"
	"go/format"
	"go/types"
	"html/template"
	"os"
	"strings"

	"golang.org/x/tools/go/packages"
	"golang.org/x/tools/go/types/typeutil"
	"golang.org/x/tools/imports"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protodesc"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"
	"google.golang.org/protobuf/types/descriptorpb"
)

const interceptorFile = "../../proxy/interceptor.go"

const InterceptorTemplateText = `
// Code generated by proxygenerator; DO NOT EDIT.

package proxy

import (
	"context"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
    "google.golang.org/protobuf/proto"
)

// VisitPayloadsContext provides Payload context for visitor functions.
type VisitPayloadsContext struct {
	context.Context
	// The parent message for this payload.
	Parent proto.Message
	// If true, a single payload is given and a single payload must be returned.
	SinglePayloadRequired bool
}

// VisitPayloadsOptions configure visitor behaviour.
type VisitPayloadsOptions struct {
	// Context is the same for every call of a visit, callers should not store it. This must never
	// return an empty set of payloads.
	Visitor func(*VisitPayloadsContext, []*common.Payload) ([]*common.Payload, error)
	// Don't visit search attribute payloads.
	SkipSearchAttributes bool
	// Will be called for each Any encountered. If not set, the default is to recurse into the Any
	// object, unmarshal it, visit, and re-marshal it always (even if there are no changes).
	WellKnownAnyVisitor func(*VisitPayloadsContext, *anypb.Any) error
}

// VisitPayloads calls the options.Visitor function for every Payload proto within msg.
func VisitPayloads(ctx context.Context, msg proto.Message, options VisitPayloadsOptions) error {
	visitCtx := VisitPayloadsContext{Context: ctx, Parent: msg}

	return visitPayloads(&visitCtx, &options, nil, msg)
}

// PayloadVisitorInterceptorOptions configures outbound/inbound interception of Payloads within msgs.
type PayloadVisitorInterceptorOptions struct {
	// Visit options for outbound messages
	Outbound *VisitPayloadsOptions
	// Visit options for inbound messages
	Inbound *VisitPayloadsOptions
}

// NewPayloadVisitorInterceptor creates a new gRPC interceptor for workflowservice messages.
func NewPayloadVisitorInterceptor(options PayloadVisitorInterceptorOptions) (grpc.UnaryClientInterceptor, error) {
	return func(ctx context.Context, method string, req, response interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		if reqMsg, ok := req.(proto.Message); ok && options.Outbound != nil {
			err := VisitPayloads(ctx, reqMsg, *options.Outbound)
			if err != nil {
				return err
			}	
		}

		err := invoker(ctx, method, req, response, cc, opts...)
		if err != nil {
			if options.Inbound != nil {
				stat, ok := status.FromError(err)
				if !ok {
					return err
				}
				// user provided payloads can sometimes end up in the status details of
				// gRPC errors, make sure to visit those as well
				for _, detail := range stat.Details() {
					detailAny, ok := detail.(*anypb.Any)
					if {{ range $i, $name := .GrpcPayload }}{{ if $i }} || {{ end }}detailAny.MessageName() == "{{$name}}"{{ end }} {
						if !ok {
							return fmt.Errorf("Error returned from rpc invoker should be anypb.Any")
						}
						VisitPayloads(ctx, detailAny, *options.Inbound)
					}
				}
				return err
			}
		}

		if resMsg, ok := response.(proto.Message); ok && options.Inbound != nil {
			return VisitPayloads(ctx, resMsg, *options.Inbound)	
		}
		
		return nil
	}, nil
}

// VisitFailuresContext provides Failure context for visitor functions.
type VisitFailuresContext struct {
	context.Context
	// The parent message for this failure.
	Parent proto.Message
}

// VisitFailuresOptions configure visitor behaviour.
type VisitFailuresOptions struct {
	// Context is the same for every call of a visit, callers should not store it.
	// Visitor is free to mutate the passed failure struct.
	Visitor func(*VisitFailuresContext, *failure.Failure) (error)
	// Will be called for each Any encountered. If not set, the default is to recurse into the Any
	// object, unmarshal it, visit, and re-marshal it always (even if there are no changes).
	WellKnownAnyVisitor func(*VisitFailuresContext, *anypb.Any) error
}

// VisitFailures calls the options.Visitor function for every Failure proto within msg.
func VisitFailures(ctx context.Context, msg proto.Message, options VisitFailuresOptions) error {
	visitCtx := VisitFailuresContext{Context: ctx, Parent: msg}

	return visitFailures(&visitCtx, &options, msg)
}

// FailureVisitorInterceptorOptions configures outbound/inbound interception of Failures within msgs.
type FailureVisitorInterceptorOptions struct {
	// Visit options for outbound messages
	Outbound *VisitFailuresOptions
	// Visit options for inbound messages
	Inbound *VisitFailuresOptions
}

// NewFailureVisitorInterceptor creates a new gRPC interceptor for workflowservice messages.
func NewFailureVisitorInterceptor(options FailureVisitorInterceptorOptions) (grpc.UnaryClientInterceptor, error) {
	return func(ctx context.Context, method string, req, response interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		if reqMsg, ok := req.(proto.Message); ok && options.Outbound != nil {
			err := VisitFailures(ctx, reqMsg, *options.Outbound)
			if err != nil {
				return err
			}	
		}

		err := invoker(ctx, method, req, response, cc, opts...)
		if err != nil {
			if options.Inbound != nil {
				stat, ok := status.FromError(err)
				if !ok {
					return err
				}
				// user provided payloads can sometimes end up in the status details of
				// gRPC errors, make sure to visit those as well
				for _, detail := range stat.Details() {
					detailAny, ok := detail.(*anypb.Any)
					if {{ range $i, $name := .GrpcFailure }}{{ if $i }} || {{ end }}detailAny.MessageName() == "{{$name}}"{{ end }} {
						if !ok {
							return fmt.Errorf("Error returned from rpc invoker should be anypb.Any")
						}
						VisitFailures(ctx, detailAny, *options.Inbound)
					}
				}
				return err
			}
		}

		if resMsg, ok := response.(proto.Message); ok && options.Inbound != nil {
			return VisitFailures(ctx, resMsg, *options.Inbound)	
		}
		
		return nil
	}, nil
}

func (o *VisitFailuresOptions) defaultWellKnownAnyVisitor(ctx *VisitFailuresContext, p *anypb.Any) error {
	child, err := p.UnmarshalNew()
	if err != nil {
		return fmt.Errorf("failed to unmarshal any: %w", err)
	}
	// We choose to visit and re-marshal always instead of cloning, visiting,
	// and checking if anything changed before re-marshaling. It is assumed the
	// clone + equality check is not much cheaper than re-marshal.
	if err := visitFailures(ctx, o, child); err != nil {
		return err
	}
	// Confirmed this replaces both Any fields on non-error, there is nothing
	// left over
	if err := p.MarshalFrom(child); err != nil {
		return fmt.Errorf("failed to marshal any: %w", err)
	}
	return nil
}

func (o *VisitPayloadsOptions) defaultWellKnownAnyVisitor(ctx *VisitPayloadsContext, p *anypb.Any) error {
	child, err := p.UnmarshalNew()
	if err != nil {
		return fmt.Errorf("failed to unmarshal any: %w", err)
	}
	// We choose to visit and re-marshal always instead of cloning, visiting,
	// and checking if anything changed before re-marshaling. It is assumed the
	// clone + equality check is not much cheaper than re-marshal.
	if err := visitPayloads(ctx, o, p, child); err != nil {
		return err
	}
	// Confirmed this replaces both Any fields on non-error, there is nothing
	// left over
	if err := p.MarshalFrom(child); err != nil {
		return fmt.Errorf("failed to marshal any: %w", err)
	}
	return nil
}

func visitPayload(
	ctx *VisitPayloadsContext,
	options *VisitPayloadsOptions,
	parent proto.Message,
	msg *common.Payload,
) (*common.Payload, error) {
	ctx.SinglePayloadRequired, ctx.Parent = true, parent
	newPayloads, err := options.Visitor(ctx, []*common.Payload{msg})
	ctx.SinglePayloadRequired, ctx.Parent = false, nil
	if err != nil {
		return nil, err
	}

	if len(newPayloads) != 1 {
		return nil, fmt.Errorf("visitor func must return 1 payload when SinglePayloadRequired = true")
	}

	return newPayloads[0], nil
}

func visitPayloads(
	ctx *VisitPayloadsContext,
	options *VisitPayloadsOptions,
	parent proto.Message,
	objs ...interface{},
) error {
	for _, obj := range objs {
		ctx.SinglePayloadRequired = false

		switch o := obj.(type) {
			case map[string]*common.Payload:
				for ix, x := range o {
					if nx, err := visitPayload(ctx, options, parent, x); err != nil {
						return err
					} else {
						o[ix] = nx
					}
				}
			case *common.Payloads:
				if o == nil { continue }
				ctx.Parent = parent
				newPayloads, err := options.Visitor(ctx, o.Payloads)
				ctx.Parent = nil
				if err != nil { return err }
				o.Payloads = newPayloads
			case map[string]*common.Payloads:
				for _, x := range o {
					if err := visitPayloads(ctx, options, parent, x); err != nil {
						return err
					}
				}
		case *anypb.Any:
			if o == nil {
				continue
			}
			visitor := options.WellKnownAnyVisitor
			if visitor == nil {
				visitor = options.defaultWellKnownAnyVisitor
			}
			ctx.Parent = o
			err := visitor(ctx, o)
			ctx.Parent = nil
			if err != nil {
				return err
			}
{{range $type, $record := .PayloadTypes}}
		{{if $record.Slice}}
			case []{{$type}}:
				for _, x := range o {
					if err := visitPayloads(ctx, options, parent, x); err != nil {
						return err
					}
				}
		{{end}}
		{{if $record.Map}}
			case map[string]{{$type}}:
				for _, x := range o {
					if err := visitPayloads(ctx, options, parent, x); err != nil {
						return err
					}
				}
		{{end}}
			case {{$type}}:
				{{if eq $type "*common.SearchAttributes"}}
				if options.SkipSearchAttributes { continue }
				{{end}}
				if o == nil { continue }
				{{range $record.Payloads -}}
				if o.{{.}} != nil {
					no, err := visitPayload(ctx, options, o, o.{{.}})
					if err != nil { return err }
					o.{{.}} = no
				}
				{{end}}
				{{if $record.Methods}}
				if err := visitPayloads(
					ctx,
					options,
					o,
					{{range $record.Methods -}}
						o.{{.}}(),
					{{end}}
				); err != nil { return err }
				{{end}}
{{end}}
		}
	}

	return nil
}

func visitFailures(ctx *VisitFailuresContext, options *VisitFailuresOptions, objs ...interface{}) error {
	for _, obj := range objs {
		switch o := obj.(type) {
			case *failure.Failure:
				if o == nil { continue }
				if err := options.Visitor(ctx, o); err != nil { return err }
				if err := visitFailures(ctx, options, o.GetCause()); err != nil { return err }
			case *anypb.Any:
				if o == nil {
					continue
				}
				visitor := options.WellKnownAnyVisitor
				if visitor == nil {
					visitor = options.defaultWellKnownAnyVisitor
				}
				ctx.Parent = o
				err := visitor(ctx, o)
				ctx.Parent = nil
				if err != nil {
					return err
				}
{{range $type, $record := .FailureTypes}}
		{{if $record.Slice}}
			case []{{$type}}:
				for _, x := range o { if err := visitFailures(ctx, options, x); err != nil { return err } }
		{{end}}
		{{if $record.Map}}
			case map[string]{{$type}}:
				for _, x := range o { if err := visitFailures(ctx, options, x); err != nil { return err } }
		{{end}}
			case {{$type}}:
				if o == nil { continue }
				ctx.Parent = o
				if err := visitFailures(
					ctx,
					options,
					{{range $record.Methods -}}
						o.{{.}}(),
					{{end}}
				); err != nil { return err }
{{end}}
		}
	}

	return nil
}
`

var interceptorTemplate = template.Must(template.New("interceptor").Parse(InterceptorTemplateText))

type TemplateInput struct {
	PayloadTypes map[string]*TypeRecord
	FailureTypes map[string]*TypeRecord
	GrpcPayload  []string
	GrpcFailure  []string
}

// TypeRecord holds the state for a type referred to by the workflow service
type TypeRecord struct {
	Methods  []string // List of methods on this type that can eventually lead to Payload(s)
	Payloads []string // List of attributes on this type that are of type Payload
	Slice    bool     // The API refers to slices of this type
	Map      bool     // The API refers to maps with this type as the value
	Matches  bool     // We found methods on this type that can eventually lead to Payload(s)
}

// isSlice returns true if a type is slice, false otherwise
func isSlice(t types.Type) bool {
	switch t.(type) {
	case *types.Slice:
		return true
	}
	return false
}

// isMap returns true if a type is map, false otherwise
func isMap(t types.Type) bool {
	switch t.(type) {
	case *types.Map:
		return true
	}
	return false
}

// elemType returns the elem (value) type for a slice or map
func elemType(t types.Type) types.Type {
	switch typ := t.(type) {
	case *types.Slice:
		return typ.Elem()
	case *types.Map:
		return typ.Elem()
	}
	return t
}

// typeName returns a normalized path for a type
func typeName(t types.Type) string {
	return types.TypeString(elemType(t), func(p *types.Package) string {
		return p.Name()
	})
}

// typeMatches returns true if a type:
// Is equal to or is a pointer to any of the desired types
// Is a slice or slice of pointers to any of the desired types
// Is a map where the value is any of the desired types or a pointer to any of the desired types
func typeMatches(t types.Type, desired ...types.Type) bool {
	resolved := resolveType(t).String()
	for _, f := range desired {
		if resolved == f.String() {
			return true
		}
	}

	return false
}

// resolveType returns the underlying type for pointers, slices and maps
func resolveType(t types.Type) types.Type {
	switch typ := t.(type) {
	case *types.Pointer:
		return resolveType(typ.Elem())
	case *types.Slice:
		return resolveType(typ.Elem())
	case *types.Map:
		return resolveType(typ.Elem())
	}

	return t
}

func pruneRecords(input map[string]*TypeRecord) map[string]*TypeRecord {
	result := map[string]*TypeRecord{}

	for typ, record := range input {
		if record.Matches {
			result[typ] = record
		}
	}

	return result
}

// isMatchingMessage returns true if the message descriptor is one of the target types
func isMatchingMessage(md protoreflect.MessageDescriptor, targetNames []string) bool {
	fullName := string(md.FullName())
	for _, targetName := range targetNames {
		if fullName == targetName {
			return true
		}
	}
	return false
}

// containsMessage recursively checks whether the given message descriptor (or any of its fields)
// contains (transitively) a target message.
func containsMessage(
	md protoreflect.MessageDescriptor,
	targetMessages []string,
	memo map[protoreflect.FullName]bool,
) bool {
	fullName := md.FullName()
	// If we've already computed for this message, return the cached result.
	if res, ok := memo[fullName]; ok {
		return res
	}
	// Mark this message as not containing a payload to break cycles.
	memo[fullName] = false

	// Check every field of this message.
	for i := 0; i < md.Fields().Len(); i++ {
		field := md.Fields().Get(i)
		// Only care about message-type fields.
		if field.Kind() == protoreflect.MessageKind && field.Message() != nil {
			child := field.Message()
			// If the field is directly a payload (or Any) then mark and return true.
			if isMatchingMessage(child, targetMessages) {
				memo[fullName] = true
				return true
			}
			// Otherwise, recursively check if the field's message type contains a payload.
			if containsMessage(child, targetMessages, memo) {
				memo[fullName] = true
				return true
			}
		}
	}
	return false
}

// checkMessage examines the given message descriptor md and, if it (transitively) contains a
// payload, appends its result slice.
func checkMessage(md protoreflect.MessageDescriptor,
	targetMessages []string,
	memo map[protoreflect.FullName]bool,
	result *[]protoreflect.MessageDescriptor,
) {
	// Avoid reporting the target types directly
	if !isMatchingMessage(md, targetMessages) && containsMessage(md, targetMessages, memo) {
		*result = append(*result, md)
	}
	nested := md.Messages()
	for i := 0; i < nested.Len(); i++ {
		checkMessage(nested.Get(i), targetMessages, memo, result)
	}
}

// gatherMessagesContainingTargets walks all proto file descriptors in the registry,
// and returns a slice of full message names that (transitively) contain the target message types.
// The excludedPathPrefixes are used to skip files that match the given prefixes.
func gatherMessagesContainingTargets(
	protoFiles *protoregistry.Files,
	targetMessages []string,
	excludedPathPrefixes []string,
) ([]protoreflect.MessageDescriptor, error) {
	messagesWithPayload := make([]protoreflect.MessageDescriptor, 0)
	memo := make(map[protoreflect.FullName]bool)
	protoFiles.RangeFiles(func(fd protoreflect.FileDescriptor) bool {
		for _, prefix := range excludedPathPrefixes {
			if strings.HasPrefix(fd.Path(), prefix) {
				return true
			}
		}
		msgs := fd.Messages()
		for i := 0; i < msgs.Len(); i++ {
			checkMessage(msgs.Get(i), targetMessages, memo, &messagesWithPayload)
		}
		return true
	})
	return messagesWithPayload, nil
}

// getProtoRegistryFromDescriptor reads a file descriptor set from the given path and returns a proto registry.
func getProtoRegistryFromDescriptor(path string) (*protoregistry.Files, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("reading descriptor set: %w", err)
	}
	var fdSet descriptorpb.FileDescriptorSet
	if err := proto.Unmarshal(data, &fdSet); err != nil {
		return nil, fmt.Errorf("unmarshalling descriptor set: %w", err)
	}
	files, err := protodesc.NewFiles(&fdSet)
	if err != nil {
		return nil, fmt.Errorf("creating file registry: %w", err)
	}
	return files, nil
}

// protoFullNameToGoPackageAndType converts a proto full name to a Go package path and type name.
// You will need to adjust this function to suit your project’s naming conventions.
func protoFullNameToGoPackageAndType(md protoreflect.MessageDescriptor) (pkgPath, typeName string, err error) {
	if md.IsMapEntry() {
		// For map entries, what we actually want to search is their parent.
		parent := md.Parent()
		if parent != nil {
			if msgParent, ok := parent.(protoreflect.MessageDescriptor); ok {
				return protoFullNameToGoPackageAndType(msgParent)
			}
		}
		return "", "", fmt.Errorf("map entry has no parent: %s", md.FullName())
	}

	fullName := string(md.FullName())
	// Ex: "temporal.api.common.v1.Payload"
	parts := strings.Split(fullName, ".")
	if len(parts) < 4 {
		return "", "", fmt.Errorf("unexpected proto full name: %s", fullName)
	}
	// Fix up the descriptor name to match the Go package path.
	if parts[0] == "temporal" {
		parts[0] = "go.temporal.io"
	}
	pkgPath = strings.Join(parts[0:len(parts)-1], "/")
	typeName = parts[len(parts)-1]
	return pkgPath, typeName, nil
}

func gatherMatchesToTypeRecords(
	mds []protoreflect.MessageDescriptor,
	targetTypes []types.Type,
	directMatchTypes []types.Type,
) (map[string]*TypeRecord, error) {
	matchingRecords := map[string]*TypeRecord{}
	packagesToTypes := map[string][]string{}
	for _, md := range mds {
		pkgPath, typeName, err := protoFullNameToGoPackageAndType(md)
		if pkgPath == "" {
			continue
		}
		if err != nil {
			return nil, err
		}
		if _, ok := packagesToTypes[pkgPath]; !ok {
			packagesToTypes[pkgPath] = []string{}
		}
		packagesToTypes[pkgPath] = append(packagesToTypes[pkgPath], typeName)
	}
	for pkgPath, typeNames := range packagesToTypes {
		typesList, err := lookupTypes(pkgPath, typeNames)
		if err != nil {
			return nil, fmt.Errorf("failed to lookup Go types for %q: %w", pkgPath, err)
		}
		for _, t := range typesList {
			walk(targetTypes, directMatchTypes, types.NewPointer(t), &matchingRecords)
		}
	}
	matchingRecords = pruneRecords(matchingRecords)
	return matchingRecords, nil
}

// walk iterates the methods on a type and returns whether any of them can eventually lead the
// desired type(s). The return type for each method on this type is walked recursively to decide
// which methods can lead to the desired type.
func walk(desired []types.Type, directMatchTypes []types.Type, typ types.Type, records *map[string]*TypeRecord) bool {
	typeName := typeName(typ)

	// If this type is a slice then walk the underlying type and then make a note we need to encode slices of this type
	if isSlice(typ) {
		result := walk(desired, directMatchTypes, elemType(typ), records)
		if result {
			record := (*records)[typeName]
			record.Slice = true
		}
		return result
	}

	// If this type is a map then walk the underlying type and then make a note we need to encode maps with values of this type
	if isMap(typ) {
		result := walk(desired, directMatchTypes, elemType(typ), records)
		if result {
			record := (*records)[typeName]
			record.Map = true
		}
		return result
	}

	// If we've walked this type before, return the previous result
	if record, ok := (*records)[typeName]; ok {
		return record.Matches
	}

	record := TypeRecord{}
	(*records)[typeName] = &record

	// Look for all functions with this `typ` type
	for _, meth := range typeutil.IntuitiveMethodSet(elemType(typ), nil) {
		// Ignore non-exported methods
		if !meth.Obj().Exported() {
			continue
		}

		methodName := meth.Obj().Name()
		// The protobuf types have a Get.. method for every protobuf they refer to
		// We walk only these methods to avoid cycles or other nasty issues
		if !strings.HasPrefix(methodName, "Get") {
			continue
		}

		sig := meth.Obj().Type().(*types.Signature)
		// All the Get... methods return the relevant protobuf as the first result
		resultType := sig.Results().At(0).Type()

		hasDirectMatch := false
		for _, directMatchType := range directMatchTypes {
			if resultType.String() == types.NewPointer(directMatchType).String() {
				hasDirectMatch = true
				break
			}
		}
		if hasDirectMatch {
			record.Matches = true
			prefix, ok := strings.CutPrefix(methodName, "Get")
			if !ok {
				panic(fmt.Errorf("expected method to have a Get prefix: %s", methodName))
			}
			record.Payloads = append(record.Payloads, prefix)
			continue
		}

		// Check if this method returns a desired type or if it leads (eventually) to a Type which
		// refers to a desired type
		if typeMatches(resultType, desired...) || walk(desired, directMatchTypes, resultType, records) {
			record.Matches = true
			record.Methods = append(record.Methods, methodName)
		}
	}

	// Return whether this Type can (eventually) lead to Payload(s)
	// This is used in the walk logic above so that our encoding handles intermediate Types between our Request/Response objects and Payload(s)
	return record.Matches
}

func lookupTypes(pkgName string, typeNames []string) ([]types.Type, error) {
	conf := &packages.Config{Mode: packages.NeedImports | packages.NeedTypes | packages.NeedTypesInfo}
	result := []types.Type{}

	pkgs, err := packages.Load(conf, pkgName)
	if err != nil {
		return result, fmt.Errorf("failed to load package %s: %w", pkgName, err)
	}
	scope := pkgs[0].Types.Scope()

	for _, t := range typeNames {
		lookedUpType := scope.Lookup(t)
		if lookedUpType != nil {
			result = append(result, lookedUpType.Type())
		}
	}

	return result, nil
}

func generateInterceptor(cfg config) error {
	payloadTypes, err := lookupTypes("go.temporal.io/api/common/v1", []string{"Payloads", "Payload"})
	payloadDirectMatchType, err := lookupTypes("go.temporal.io/api/common/v1", []string{"Payload"})
	if err != nil {
		return err
	}

	failureTypes, err := lookupTypes("go.temporal.io/api/failure/v1", []string{"Failure"})
	if err != nil {
		return err
	}

	// For the purposes of payloads and failures, we also consider the Any well known type as
	// possible
	if anyTypes, err := lookupTypes("google.golang.org/protobuf/types/known/anypb", []string{"Any"}); err != nil {
		return err
	} else {
		payloadTypes = append(payloadTypes, anyTypes...)
		failureTypes = append(failureTypes, anyTypes...)
	}

	protoFiles, err := getProtoRegistryFromDescriptor(cfg.descriptorPath)
	if err != nil {
		return fmt.Errorf("loading descriptor set: %w", err)
	}

	// Cloud protos currently not included in interceptor
	excludedEntryPoints := []string{
		"temporal/api/cloud",
	}
	payloadMessageNames := []string{
		"temporal.api.common.v1.Payload",
		"temporal.api.common.v1.Payloads",
		"google.protobuf.Any",
	}
	allPayloadContainingMessages, err := gatherMessagesContainingTargets(protoFiles, payloadMessageNames, excludedEntryPoints)
	failureMessageNames := []string{
		"temporal.api.failure.v1.Failure",
		"google.protobuf.Any",
	}
	allFailureContainingMessages, err := gatherMessagesContainingTargets(protoFiles, failureMessageNames, excludedEntryPoints)

	payloadRecords, err := gatherMatchesToTypeRecords(allPayloadContainingMessages, payloadTypes, payloadDirectMatchType)
	if err != nil {
		return err
	}
	failureRecords, err := gatherMatchesToTypeRecords(allFailureContainingMessages, failureTypes, make([]types.Type, 0))
	if err != nil {
		return err
	}

	// gather a list of errordetails that can contain user payloads when included in
	// gRPC error messages
	var grpcPayload []string
	for _, msg := range allPayloadContainingMessages {
		if strings.Contains(string(msg.FullName()), "temporal.api.errordetails.v1.") && strings.Count(string(msg.FullName()), ".") == 4 {
			grpcPayload = append(grpcPayload, string(msg.FullName()))
		}
	}

	var grpcFailure []string
	for _, msg := range allFailureContainingMessages {
		if strings.Contains(string(msg.FullName()), "temporal.api.errordetails.v1.") && strings.Count(string(msg.FullName()), ".") == 4 {
			grpcFailure = append(grpcFailure, string(msg.FullName()))
		}
	}

	buf := &bytes.Buffer{}
	fmt.Fprint(buf, cfg.license)

	err = interceptorTemplate.Execute(buf, TemplateInput{
		PayloadTypes: payloadRecords,
		FailureTypes: failureRecords,
		GrpcFailure:  grpcFailure,
		GrpcPayload:  grpcPayload,
	})
	if err != nil {
		return err
	}

	src, err := imports.Process(interceptorFile, buf.Bytes(), nil)
	if err != nil {
		return fmt.Errorf("failed to process interceptor imports: %w", err)
	}

	src, err = format.Source(src)
	if err != nil {
		return fmt.Errorf("failed to format interceptor: %w", err)
	}

	if cfg.verifyOnly {
		currentSrc, err := os.ReadFile(interceptorFile)
		if err != nil {
			return fmt.Errorf("failed to read previously generated interceptor: %w", err)
		}

		if !bytes.Equal(src, currentSrc) {
			return fmt.Errorf("generated file does not match existing file: %s", interceptorFile)
		}

		return nil
	}

	if err := os.WriteFile(interceptorFile, src, 0666); err != nil {
		return fmt.Errorf("failed to write generated interceptor: %w", err)
	}

	return nil
}
