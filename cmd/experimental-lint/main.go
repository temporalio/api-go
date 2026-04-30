// experimental-lint validates usage of [(temporal.api.options.v1.experimental)] across
// a compiled descriptor set.
//
// It reads a FileDescriptorSet from stdin (produced by `buf build -o -`) and checks
// every annotated element against the registered feature names in experimental-features.txt.
//
// Exit codes:
//   0 – all annotations are valid (or --warn-only is set)
//   1 – one or more annotations reference an unregistered feature name
//
// Usage:
//
//	buf build -o - | go run ./cmd/experimental-lint --features experimental-features.txt
//
// This tool is currently run in warn-only mode (--warn-only flag, default true) while
// the option is being bootstrapped. Set --warn-only=false to enforce in CI.
package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/encoding/protowire"
	"google.golang.org/protobuf/types/descriptorpb"
)

const experimentalFieldNumber = 77001

func main() {
	featuresFile := flag.String("features", "experimental-features.txt", "path to experimental-features.txt")
	warnOnly := flag.Bool("warn-only", true, "log warnings instead of failing on unknown feature names")
	flag.Parse()

	registered, err := loadFeatures(*featuresFile)
	if err != nil {
		log.Fatalf("loading features: %v", err)
	}

	data, err := io.ReadAll(os.Stdin)
	if err != nil {
		log.Fatalf("reading descriptor set: %v", err)
	}
	var fds descriptorpb.FileDescriptorSet
	if err := proto.Unmarshal(data, &fds); err != nil {
		log.Fatalf("unmarshaling descriptor set: %v", err)
	}

	violations := check(&fds, registered)
	for _, v := range violations {
		fmt.Fprintf(os.Stderr, "WARN: experimental option has unregistered feature %q at %s\n", v.feature, v.location)
	}
	if len(violations) > 0 && !*warnOnly {
		os.Exit(1)
	}
}

type violation struct {
	feature  string
	location string
}

func check(fds *descriptorpb.FileDescriptorSet, registered map[string]bool) []violation {
	var out []violation
	for _, f := range fds.GetFile() {
		out = append(out, checkFile(f, registered)...)
	}
	return out
}

func checkFile(f *descriptorpb.FileDescriptorProto, registered map[string]bool) []violation {
	var out []violation
	file := f.GetName()
	for _, m := range f.GetMessageType() {
		out = append(out, checkMessage(m, file, registered)...)
	}
	for _, e := range f.GetEnumType() {
		out = append(out, checkEnum(e, file, registered)...)
	}
	for _, s := range f.GetService() {
		out = append(out, checkService(s, file, registered)...)
	}
	return out
}

func checkMessage(m *descriptorpb.DescriptorProto, file string, registered map[string]bool) []violation {
	var out []violation
	if feat := experimentalValue(m.GetOptions()); feat != "" && !registered[feat] {
		out = append(out, violation{feat, file + ": message " + m.GetName()})
	}
	for _, f := range m.GetField() {
		if feat := experimentalValue(f.GetOptions()); feat != "" && !registered[feat] {
			out = append(out, violation{feat, file + ": " + m.GetName() + "." + f.GetName()})
		}
	}
	for _, nested := range m.GetNestedType() {
		out = append(out, checkMessage(nested, file, registered)...)
	}
	for _, e := range m.GetEnumType() {
		out = append(out, checkEnum(e, file, registered)...)
	}
	return out
}

func checkEnum(e *descriptorpb.EnumDescriptorProto, file string, registered map[string]bool) []violation {
	var out []violation
	if feat := experimentalValue(e.GetOptions()); feat != "" && !registered[feat] {
		out = append(out, violation{feat, file + ": enum " + e.GetName()})
	}
	for _, v := range e.GetValue() {
		if feat := experimentalValue(v.GetOptions()); feat != "" && !registered[feat] {
			out = append(out, violation{feat, file + ": " + e.GetName() + "." + v.GetName()})
		}
	}
	return out
}

func checkService(s *descriptorpb.ServiceDescriptorProto, file string, registered map[string]bool) []violation {
	var out []violation
	if feat := experimentalValue(s.GetOptions()); feat != "" && !registered[feat] {
		out = append(out, violation{feat, file + ": service " + s.GetName()})
	}
	for _, m := range s.GetMethod() {
		if feat := experimentalValue(m.GetOptions()); feat != "" && !registered[feat] {
			out = append(out, violation{feat, file + ": " + s.GetName() + "." + m.GetName()})
		}
	}
	return out
}

// experimentalValue extracts the experimental option value (field number 77001, wire type 2)
// from the unknown-fields bytes of an Options message. The option is unknown to this binary
// because the experimental.proto descriptor is not linked in, but proto preserves unknown
// fields, so we can read it directly from the wire bytes.
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

func loadFeatures(path string) (map[string]bool, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	registered := map[string]bool{}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		registered[line] = true
	}
	return registered, scanner.Err()
}
