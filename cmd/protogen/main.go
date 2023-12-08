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

	"go.temporal.io/api/internal/errgroup"

	"go.temporal.io/api/internal/protogen/enum"
	"go.temporal.io/api/internal/protogen/version"
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
	rootDir     string
	outputDir   string
	excludeDirs []string
	includes    []string
	plugins     []string
	enums       map[string]string
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
func findEnums(ctx context.Context, eg *errgroup.Group, cfg genConfig, dirs chan<- string) []string {
	seen := map[string]struct{}{}
	var enums []string
	enumsFound := make(chan string)
	eg.Go(func() error {
		defer close(enumsFound)
		defer close(dirs)
		return walkExtension(ctx, cfg.rootDir, ".proto", cfg.excludeDirs, func(path string) error {
			// Emit unique directories containing proto files
			dir := filepath.Dir(path)
			if _, ok := seen[dir]; !ok {
				seen[dir] = struct{}{}
				select {
				case dirs <- dir:
				case <-ctx.Done():
					return ctx.Err()
				}
			}

			// process the file and find enums, placing them in the enum channel
			eg.Go(func() error {
				// There isn't a good way to iteratively find matches using an io.Reader
				// so we need to pull the whole thing into memory. We could do it line by line
				// but I doubt this will take enough memory to matter
				bs, err := os.ReadFile(path)
				if err != nil {
					return err
				}
				matches := enumRgx.FindAllStringSubmatch(string(bs), -1)
				for i := 0; i < len(matches); i++ {
					// group 0 is the full match, group 1 has the enum's name
					select {
					case <-ctx.Done():
						return ctx.Err()
					case enumsFound <- matches[i][1]:
					}
				}
				return nil
			})
			return nil
		})
	})
	for enum := range enumsFound {
		enums = append(enums, enum)
	}
	return enums
}

// Run protoc in parallel on all proto dirs we discover under the root directory
// It returns the list of unique directories contai
func runProtoc(ctx context.Context, eg *errgroup.Group, cfg genConfig, protoDirs <-chan string) {
	for dir := range protoDirs {
		dir := dir
		// Run protoc on each directory individually
		eg.Go(func() error {
			args := []string{
				"--fatal_warnings",
				fmt.Sprintf("--go_out=paths=source_relative:%s", cfg.outputDir),
			}
			for _, include := range cfg.includes {
				args = append(args, fmt.Sprintf("-I=%s", include))
			}

			// If we need more complex plugin handling, such as per-plugin options, we can add that later.
			// For now we use the same args everywhere
			for _, plugin := range cfg.plugins {
				args = append(args, fmt.Sprintf("--%s", plugin))
			}
			files, err := filepath.Glob(filepath.Join(dir, "*"))
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
			return nil
		})
	}
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
func postProcess(ctx context.Context, eg *errgroup.Group, cfg genConfig) error {
	return walkExtension(ctx, cfg.outputDir, ".pb.go", cfg.excludeDirs, func(path string) error {
		return rewriteFile(path,
			version.NewRemover(),
			enum.NewConstRewriter(cfg.enums),
			enum.NewStringRewriter())
	})
}

func compileProtos(ctx context.Context, cfg genConfig) error {
	eg, ectx := errgroup.WithContext(ctx)
	var enums []string
	dirCh := make(chan string)
	eg.Go(func() error {
		enums = findEnums(ctx, eg, cfg, dirCh)
		return nil
	})

	eg.Go(func() error {
		runProtoc(ectx, eg, cfg, dirCh)
		return nil
	})

	if err := eg.Wait(); err != nil {
		return err
	}

	// Add Enum:Prefix pairs for all identified pairs that weren't specifically overridden
	for i := 0; i < len(enums); i++ {
		e := enums[i]
		if _, ok := cfg.enums[e]; !ok {
			cfg.enums[e] = e
		}
	}
	return postProcess(ctx, eg, cfg)
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
	var protoRootDir, outputDir, enumPrefixPairs string
	var protoPlugins, protoIncludes, excludeDirs stringArr
	var concurrency int
	flag.StringVar(&protoRootDir, "root", "proto", "Root directory containing the protos to generate code for")
	flag.StringVar(&outputDir, "output", "api", "Base directory in which to output generated proto files")
	flag.Var(&protoIncludes, "I", "Directory to include when compiling protos")
	flag.Var(&protoPlugins, "p", "Plugin=Options pairs of protobuf plugins, like grpc-gateway_out=allow_patch_feature=false")
	flag.Var(&excludeDirs, "exclude", "Directory to exclude when compiling post-processing protos")
	flag.StringVar(&enumPrefixPairs, "rewrite-enum", "",
		"Comma-separated list of additional EnumType:CodePrefix pairs to remove when rewriting golang enums. Example: BuildId_State:BuildId")
	flag.IntVar(&concurrency, "concurrency", 0, "Maximum number of goroutines to run. Defaults to 0, meaning unlimited")
	flag.Parse()

	if protoRootDir == "" {
		fail("must specify the root directory")
	}
	if outputDir == "" {
		fail("must specify the output dir for proto files")
	}

	// Always include the root dir
	if !sliceContains(protoIncludes, protoRootDir) {
		protoIncludes = append(protoIncludes, protoRootDir)
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
		rootDir:     protoRootDir,
		outputDir:   outputDir,
		excludeDirs: excludeDirs,
		includes:    protoIncludes,
		plugins:     protoPlugins,
		enums:       enums,
	})
	if err != nil {
		fail(err.Error())
	}
}
