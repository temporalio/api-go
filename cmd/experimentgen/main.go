package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"runtime"
)

var newGenerator = func() generator {
	return generator{resolveStableVersion: defaultStableVersionResolver}
}

func main() {
	os.Exit(run(os.Args[1:], os.Stdout, os.Stderr))
}

func run(args []string, stdout io.Writer, stderr io.Writer) int {
	var apiRepo string
	var sourceSHA string
	var variant string
	flagSet := flag.NewFlagSet("experimentgen", flag.ContinueOnError)
	flagSet.SetOutput(stderr)
	flagSet.StringVar(&apiRepo, "api-repo", defaultAPIRepo(), "path to api repo checkout")
	flagSet.StringVar(&sourceSHA, "api-sha", "", "api repo commit SHA containing experimental proto source")
	flagSet.StringVar(&variant, "variant", "", "experimental variant name")
	if err := flagSet.Parse(args); err != nil {
		return 2
	}

	if sourceSHA == "" || variant == "" {
		fmt.Fprintln(stderr, "usage: go run ./cmd/experimentgen -variant <name> -api-sha <api-commit> [-api-repo <path-to-api-repo>]")
		return 2
	}

	outDir := filepath.Join("experimental", variant)
	gen := newGenerator()
	if err := gen.generate(filepath.Clean(apiRepo), sourceSHA, variant, filepath.Clean(outDir)); err != nil {
		fmt.Fprintln(stderr, err)
		return 1
	}
	fmt.Fprintf(stdout, "generated %s\n", filepath.Clean(outDir))
	return 0
}

func defaultAPIRepo() string {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		return filepath.Join(".", "proto", "api")
	}
	repoRoot := filepath.Clean(filepath.Join(filepath.Dir(filename), "..", ".."))
	candidate := filepath.Join(repoRoot, "proto", "api")
	if info, err := os.Stat(candidate); err == nil && info.IsDir() {
		return candidate
	}
	return filepath.Join(".", "proto", "api")
}
