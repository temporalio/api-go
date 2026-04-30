// experimental-lint validates usage of [(temporal.api.experimental_*)] across
// a compiled descriptor set.
//
// It reads a FileDescriptorSet from stdin (produced by `buf build -o -`) and
// reports any element annotated with a non-empty experimental option value.
// Currently used to enumerate all experimental items in CI; set --warn-only=false
// to fail the build when unrecognised annotations appear.
//
// Exit codes:
//   0 – no violations (or --warn-only is set)
//   1 – one or more violations found
//
// Usage:
//
//	buf build -o - | go run ./cmd/experimental-lint
package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"

	"google.golang.org/protobuf/encoding/protowire"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/descriptorpb"
)

const experimentalFieldNumber = 77001

func main() {
	warnOnly := flag.Bool("warn-only", true, "log warnings instead of exiting 1 on violations")
	flag.Parse()

	data, err := io.ReadAll(os.Stdin)
	if err != nil {
		log.Fatalf("reading descriptor set: %v", err)
	}
	var fds descriptorpb.FileDescriptorSet
	if err := proto.Unmarshal(data, &fds); err != nil {
		log.Fatalf("unmarshaling descriptor set: %v", err)
	}

	violations := check(&fds)
	for _, v := range violations {
		fmt.Fprintf(os.Stderr, "experimental: feature=%q at %s\n", v.feature, v.location)
	}
	if len(violations) > 0 && !*warnOnly {
		os.Exit(1)
	}
}

type violation struct {
	feature  string
	location string
}

func check(fds *descriptorpb.FileDescriptorSet) []violation {
	var out []violation
	for _, f := range fds.GetFile() {
		out = append(out, checkFile(f)...)
	}
	return out
}

func checkFile(f *descriptorpb.FileDescriptorProto) []violation {
	var out []violation
	file := f.GetName()
	for _, m := range f.GetMessageType() {
		out = append(out, checkMessage(m, file)...)
	}
	for _, e := range f.GetEnumType() {
		out = append(out, checkEnum(e, file)...)
	}
	for _, s := range f.GetService() {
		out = append(out, checkService(s, file)...)
	}
	return out
}

func checkMessage(m *descriptorpb.DescriptorProto, file string) []violation {
	var out []violation
	if feat := experimentalValue(m.GetOptions()); feat != "" {
		out = append(out, violation{feat, file + ": message " + m.GetName()})
	}
	for _, f := range m.GetField() {
		if feat := experimentalValue(f.GetOptions()); feat != "" {
			out = append(out, violation{feat, file + ": " + m.GetName() + "." + f.GetName()})
		}
	}
	for _, nested := range m.GetNestedType() {
		out = append(out, checkMessage(nested, file)...)
	}
	for _, e := range m.GetEnumType() {
		out = append(out, checkEnum(e, file)...)
	}
	return out
}

func checkEnum(e *descriptorpb.EnumDescriptorProto, file string) []violation {
	var out []violation
	if feat := experimentalValue(e.GetOptions()); feat != "" {
		out = append(out, violation{feat, file + ": enum " + e.GetName()})
	}
	for _, v := range e.GetValue() {
		if feat := experimentalValue(v.GetOptions()); feat != "" {
			out = append(out, violation{feat, file + ": " + e.GetName() + "." + v.GetName()})
		}
	}
	return out
}

func checkService(s *descriptorpb.ServiceDescriptorProto, file string) []violation {
	var out []violation
	if feat := experimentalValue(s.GetOptions()); feat != "" {
		out = append(out, violation{feat, file + ": service " + s.GetName()})
	}
	for _, m := range s.GetMethod() {
		if feat := experimentalValue(m.GetOptions()); feat != "" {
			out = append(out, violation{feat, file + ": " + s.GetName() + "." + m.GetName()})
		}
	}
	return out
}

// experimentalValue extracts the experimental option value (field number 77001,
// wire type 2 / length-delimited) from the unknown-field bytes of an Options
// message. The option is unknown to this binary because experimental.proto is
// not linked in, but proto preserves unknown fields on encode/decode.
func experimentalValue(opts proto.Message) string {
	if opts == nil {
		return ""
	}
	unknown := opts.ProtoReflect().GetUnknown()
	for len(unknown) > 0 {
		num, typ, n := protowire.ConsumeTag(unknown)
		if n < 0 {
			break
		}
		unknown = unknown[n:]
		if num == experimentalFieldNumber && typ == protowire.BytesType {
			val, n := protowire.ConsumeBytes(unknown)
			if n < 0 {
				break
			}
			return string(val)
		}
		n = protowire.ConsumeFieldValue(num, typ, unknown)
		if n < 0 {
			break
		}
		unknown = unknown[n:]
	}
	return ""
}
