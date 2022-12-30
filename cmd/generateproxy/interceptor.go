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
)

const interceptorFile = "../../proxy/interceptor.go"

const InterceptorTemplateText = Header + `
package proxy

import (
	"context"
	"fmt"

	"google.golang.org/grpc"
	proto "github.com/gogo/protobuf/proto"
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
}

// VisitPayloads calls the options.Visitor function for every Payload proto within msg.
func VisitPayloads(ctx context.Context, msg proto.Message, options VisitPayloadsOptions) error {
	visitCtx := VisitPayloadsContext{Context: ctx, Parent: msg}

	return visitPayloads(&visitCtx, &options, msg)
}

// PayloadVisitorInterceptorOptions configures outbound/inbound interception of Payloads within msgs.
type PayloadVisitorInterceptorOptions struct {
	// Visit options for outbound messages
	Outbound *VisitPayloadsOptions
	// Visit options for inbound messages
	Inbound *VisitPayloadsOptions
}

// NewPayloadVisitorInterceptor creates a new GRPC interceptor for workflowservice messages.
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
			return err
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

// NewFailureVisitorInterceptor creates a new GRPC interceptor for workflowservice messages.
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
			return err
		}

		if resMsg, ok := response.(proto.Message); ok && options.Inbound != nil {
			return VisitFailures(ctx, resMsg, *options.Inbound)	
		}
		
		return nil
	}, nil
}

func visitPayload(ctx *VisitPayloadsContext, options *VisitPayloadsOptions, msg *common.Payload) error {
	ctx.SinglePayloadRequired = true

	newPayloads, err := options.Visitor(ctx, []*common.Payload{msg})
	if err != nil {
		return err
	}

	if len(newPayloads) != 1 {
		return fmt.Errorf("visitor func must return 1 payload when SinglePayloadRequired = true")
	}

	*msg = *newPayloads[0]

	return nil
}

func visitPayloads(ctx *VisitPayloadsContext, options *VisitPayloadsOptions, objs ...interface{}) error {
	for _, obj := range objs {
		ctx.SinglePayloadRequired = false

		switch o := obj.(type) {
			case *common.Payload:
				if o == nil { continue }
				err := visitPayload(ctx, options, o)
				if err != nil { return err }
			case map[string]*common.Payload:
				for _, x := range o { if err := visitPayload(ctx, options, x); err != nil { return err } }
			case *common.Payloads:
				if o == nil { continue }
				newPayloads, err := options.Visitor(ctx, o.Payloads)
				if err != nil { return err }
				o.Payloads = newPayloads
			case map[string]*common.Payloads:
				for _, x := range o { if err := visitPayloads(ctx, options, x); err != nil { return err } }
{{range $type, $record := .PayloadTypes}}
		{{if $record.Slice}}
			case []{{$type}}:
				for _, x := range o { if err := visitPayloads(ctx, options, x); err != nil { return err } }
		{{end}}
		{{if $record.Map}}
			case map[string]{{$type}}:
				for _, x := range o { if err := visitPayloads(ctx, options, x); err != nil { return err } }
		{{end}}
			case {{$type}}:
				{{if eq $type "*common.SearchAttributes"}}
				if options.SkipSearchAttributes { continue }
				{{end}}
				if o == nil { continue }
				ctx.Parent = o
				if err := visitPayloads(
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

func visitFailures(ctx *VisitFailuresContext, options *VisitFailuresOptions, objs ...interface{}) error {
	for _, obj := range objs {
		switch o := obj.(type) {
			case *failure.Failure:
				if o == nil { continue }
				if err := options.Visitor(ctx, o); err != nil { return err }
				if err := visitFailures(ctx, options, o.GetCause()); err != nil { return err }
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

// TypeRecord holds the state for a type referred to by the workflow service
type TypeRecord struct {
	Methods []string // List of methods on this type that can eventually lead to Payload(s)
	Slice   bool     // The API refers to slices of this type
	Map     bool     // The API refers to maps with this type as the value
	Matches bool     // We found methods on this type that can eventually lead to Payload(s)
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

// walk iterates the methods on a type and returns whether any of them can eventually lead to Payload(s)
// The return type for each method on this type is walked recursively to decide which methods can lead to Payload(s)
func walk(desired []types.Type, typ types.Type, records *map[string]*TypeRecord) bool {
	typeName := typeName(typ)

	// If this type is a slice then walk the underlying type and then make a note we need to encode slices of this type
	if isSlice(typ) {
		result := walk(desired, elemType(typ), records)
		if result {
			record := (*records)[typeName]
			record.Slice = true
		}
		return result
	}

	// If this type is a map then walk the underlying type and then make a note we need to encode maps with values of this type
	if isMap(typ) {
		result := walk(desired, elemType(typ), records)
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

		// Check if this method returns a Payload(s) or if it leads (eventually) to a Type which refers to a Payload(s)
		if typeMatches(resultType, desired...) || walk(desired, resultType, records) {
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
		return result, err
	}
	scope := pkgs[0].Types.Scope()

	for _, t := range typeNames {
		result = append(result, scope.Lookup(t).Type())
	}

	return result, nil
}

func generateInterceptor(cfg config) error {
	payloadTypes, err := lookupTypes("go.temporal.io/api/common/v1", []string{"Payloads", "Payload"})
	if err != nil {
		return err
	}

	failureTypes, err := lookupTypes("go.temporal.io/api/failure/v1", []string{"Failure"})
	if err != nil {
		return err
	}

	// UnimplementedWorkflowServiceServer is auto-generated via our API package
	// The methods on this type refer to all possible Request/Response types so we can use this to walk through all of our protobuf types
	serviceTypes, err := lookupTypes("go.temporal.io/api/workflowservice/v1", []string{"UnimplementedWorkflowServiceServer"})
	if err != nil {
		return err
	}

	service := serviceTypes[0]

	payloadRecords := map[string]*TypeRecord{}
	failureRecords := map[string]*TypeRecord{}

	for _, meth := range typeutil.IntuitiveMethodSet(service, nil) {
		if !meth.Obj().Exported() {
			continue
		}

		sig := meth.Obj().Type().(*types.Signature)
		walk(payloadTypes, sig.Params().At(1).Type(), &payloadRecords)
		walk(failureTypes, sig.Params().At(1).Type(), &failureRecords)
		walk(payloadTypes, sig.Results().At(0).Type(), &payloadRecords)
		walk(failureTypes, sig.Results().At(0).Type(), &failureRecords)
	}

	payloadRecords = pruneRecords(payloadRecords)
	failureRecords = pruneRecords(failureRecords)

	buf := &bytes.Buffer{}

	err = interceptorTemplate.Execute(buf, map[string]map[string]*TypeRecord{"PayloadTypes": payloadRecords, "FailureTypes": failureRecords})
	if err != nil {
		return err
	}

	src, err := imports.Process(interceptorFile, buf.Bytes(), nil)
	if err != nil {
		return err
	}

	src, err = format.Source(src)
	if err != nil {
		return err
	}

	if cfg.verifyOnly {
		currentSrc, err := os.ReadFile(interceptorFile)
		if err != nil {
			return err
		}

		if !bytes.Equal(src, currentSrc) {
			return fmt.Errorf("generated file does not match existing file: %s", interceptorFile)
		}

		return nil
	}

	return os.WriteFile(interceptorFile, src, 0666)
}
