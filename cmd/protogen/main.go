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
	"bufio"
	"bytes"
	"context"
	"errors"
	"flag"
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"io/fs"
	"os"
	"os/exec"
	"os/signal"
	"path/filepath"
	"regexp"
	"strings"
)

var enumRgx = regexp.MustCompile(`^enum\s+(\w+)`)

type stringArr []string

func (s *stringArr) String() string {
	return "[]string"
}

func (s *stringArr) Set(value string) error {
	*s = append(*s, value)
	return nil
}

type postProcessor interface {
	Process(f *ast.File)
}

type genConfig struct {
	rootDirs    []string
	outputDir   string
	excludeDirs []string
	includes    []string
	descriptors []string
	plugins     []string
	enums       map[string]string
	// Post-processors
	rewriteString bool
	rewriteEnums  bool
	stripVersions bool
}

// walkExtension walks the directory starting from root and calls the provided callback on all files that
// have the specified extension. `exclusions` specifies a list of paths we won't search
func walkExtension(ctx context.Context, root string, extension string, exclusions []string, fn func(string) error) error {
	return filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		select {
		case <-ctx.Done():
			// Exit early. This error will be ignored because the errgroup will cancel the context only after
			// setting the true error
			return ctx.Err()
		default:
		}
		if err != nil {
			return err
		} else if !strings.HasSuffix(path, extension) {
			return nil
		}
		for i := 0; i < len(exclusions); i++ {
			if strings.HasPrefix(path, exclusions[i]) {
				return fs.SkipDir
			}
		}
		return fn(path)
	})
}

// Walk the file tree and collect all enum definitions we find
// so we can post-process the go code we're generating
func findEnums(ctx context.Context, cfg genConfig) ([]string, []string, error) {
	seen := map[string]struct{}{}
	var enums []string
	var dirs []string
	walkFunc := func(path string) error {
		// Emit unique directories containing proto files
		dir := filepath.Dir(path)
		if _, ok := seen[dir]; !ok {
			seen[dir] = struct{}{}
			dirs = append(dirs, dir)
		}

		// There isn't a good way to iteratively find matches using an io.Reader
		// so we need to pull the whole thing into memory. We could do it line by line
		// but I doubt this will take enough memory to matter
		bs, err := os.ReadFile(path)
		if err != nil {
			return err
		}
		scanner := bufio.NewScanner(bytes.NewReader(bs))
		for scanner.Scan() {
			matches := enumRgx.FindAllStringSubmatch(scanner.Text(), -1)
			for i := 0; i < len(matches); i++ {
				enums = append(enums, matches[i][1])
			}
		}
		return nil
	}
	for _, rootDir := range cfg.rootDirs {
		if err := walkExtension(ctx, rootDir, ".proto", cfg.excludeDirs, walkFunc); err != nil {
			return nil, nil, err
		}
	}
	return enums, dirs, nil
}

// Run protoc in parallel on all proto dirs we discover under the root directory
// It returns the list of unique directories contai
func runProtoc(ctx context.Context, cfg genConfig, protoDirs []string) error {
	for i := 0; i < len(protoDirs); i++ {
		dir := protoDirs[i]
		// Run protoc on each directory individually
		args := []string{
			"--fatal_warnings",
			fmt.Sprintf("--go_out=paths=source_relative:%s", cfg.outputDir),
		}
		for _, include := range cfg.includes {
			args = append(args, fmt.Sprintf("-I=%s", include))
		}
		for _, desc := range cfg.descriptors {
			args = append(args, fmt.Sprintf("--descriptor_set_in=%s", desc))
		}

		// If we need more complex plugin handling, such as per-plugin options, we can add that later.
		// For now we use the same args everywhere
		for _, plugin := range cfg.plugins {
			args = append(args, fmt.Sprintf("--%s", plugin))
		}
		files, err := filepath.Glob(filepath.Join(dir, "*.proto"))
		if err != nil {
			return err
		}
		args = append(args, files...)

		var stderr bytes.Buffer
		protoc := exec.CommandContext(ctx, "protoc", args...)
		protoc.Stderr = &stderr
		protoc.Stdout = os.Stdout
		if err := protoc.Run(); err != nil {
			if errors.Is(err, context.Canceled) {
				return err
			}
			stderrstr := strings.TrimSpace(stderr.String())
			return fmt.Errorf("failed to run `protoc %s`: %w\n%s", strings.Join(args, " "), err, stderrstr)
		}
	}
	return nil
}

func rewriteFile(fileName string, postProcessors ...postProcessor) error {
	fset := token.NewFileSet()
	// we're going to overwrite files in-place so we read it ahead of time
	srcf, err := os.Open(fileName)
	if err != nil {
		return err
	}
	ff, err := parser.ParseFile(fset, fileName, srcf, parser.ParseComments)
	if err != nil {
		return err
	}
	// Close the file so we can overwrite it
	srcf.Close()

	for _, p := range postProcessors {
		p.Process(ff)
	}

	of, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer of.Close()

	return format.Node(of, fset, ff)
}

// Post-process generated protobuf go files for useability
func postProcess(ctx context.Context, cfg genConfig) error {
	if !cfg.rewriteEnums && !cfg.rewriteString && !cfg.stripVersions {
		// nothing to do
		return nil
	}
	return walkExtension(ctx, cfg.outputDir, ".pb.go", cfg.excludeDirs, func(path string) error {
		var postProcessors []postProcessor
		if cfg.rewriteEnums {
			postProcessors = append(postProcessors, NewConstRewriter(cfg.enums))
		}
		if cfg.rewriteString {
			postProcessors = append(postProcessors, NewStringRewriter())
		}
		if cfg.stripVersions {
			postProcessors = append(postProcessors, NewVersionRemover())
		}
		return rewriteFile(path, postProcessors...)
	})
}

func compileProtos(ctx context.Context, cfg genConfig) error {
	enums, dirs, err := findEnums(ctx, cfg)
	if err != nil {
		return err
	}

	if err := runProtoc(ctx, cfg, dirs); err != nil {
		return err
	}

	// Add Enum:Prefix pairs for all identified pairs that weren't specifically overridden
	for i := 0; i < len(enums); i++ {
		e := enums[i]
		if _, ok := cfg.enums[e]; !ok {
			cfg.enums[e] = e
		}
	}
	return postProcess(ctx, cfg)
}

func fail(msg string, args ...any) {
	fmt.Fprintf(os.Stderr, msg, args...)
	os.Exit(1)
}

// Contains reports whether v is present in s.
func sliceContains[S ~[]E, E comparable](haystack S, needle E) bool {
	for i := 0; i < len(haystack); i++ {
		if needle == haystack[i] {
			return true
		}
	}
	return false
}

func main() {
	var outputDir, enumPrefixPairs string
	var protoRootDirs, protoPlugins, protoIncludes, descriptorSetIn, excludeDirs stringArr
	var noRewriteString, noRewriteEnum, noStripVersion bool
	var concurrency int
	flag.Var(&protoRootDirs, "root", "Root directories containing the protos to generate code for")
	flag.StringVar(&outputDir, "output", "api", "Base directory in which to output generated proto files")
	flag.Var(&protoIncludes, "I", "Directory to include when compiling protos")
	flag.Var(&descriptorSetIn, "descriptor_set_in", "Files containing binary FileDescriptorSet messages as inputs")
	flag.Var(&protoPlugins, "p", "Plugin=Options pairs of protobuf plugins, like grpc-gateway_out=allow_patch_feature=false")
	flag.Var(&excludeDirs, "exclude", "Directory to exclude when compiling post-processing protos")
	flag.StringVar(&enumPrefixPairs, "rewrite-enum", "",
		"Comma-separated list of additional EnumType:CodePrefix pairs to remove when rewriting golang enums. Example: BuildId_State:BuildId")
	flag.IntVar(&concurrency, "concurrency", 0, "Maximum number of goroutines to run. Defaults to 0, meaning unlimited")
	flag.BoolVar(&noRewriteEnum, "no-rewrite-enum-const", false, "Don't rewrite enum constants")
	flag.BoolVar(&noRewriteString, "no-rewrite-enum-string", false, "Don't rewrite enum String methods")
	flag.BoolVar(&noStripVersion, "no-strip-version", false, "Don't remove protoc plugin versions from generated files")

	flag.Parse()

	if len(protoRootDirs) == 0 {
		fail("must specify at least one root directory")
	}
	if outputDir == "" {
		fail("must specify the output dir for proto files")
	}

	// Always include the root dir
	for _, protoRootDir := range protoRootDirs {
		if !sliceContains(protoIncludes, protoRootDir) {
			protoIncludes = append(protoIncludes, protoRootDir)
		}
	}
	// Cancel everything if we receive SIGINT or SIGSTOP
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
	defer cancel()

	enums := make(map[string]string)
	if enumPrefixPairs != "" {
		pairs := strings.Split(enumPrefixPairs, ",")
		for _, pair := range pairs {
			parts := strings.Split(pair, ":")
			if len(parts) != 2 {
				fail("invalid enum:prefix pair %q", pair)
			}
			enums[parts[0]] = parts[1]
		}
	}
	err := compileProtos(ctx, genConfig{
		rootDirs:      protoRootDirs,
		outputDir:     outputDir,
		excludeDirs:   excludeDirs,
		includes:      protoIncludes,
		descriptors:   descriptorSetIn,
		plugins:       protoPlugins,
		enums:         enums,
		rewriteString: !noRewriteString,
		rewriteEnums:  !noRewriteEnum,
		stripVersions: !noStripVersion,
	})
	if err != nil {
		fail(err.Error())
	}
}
