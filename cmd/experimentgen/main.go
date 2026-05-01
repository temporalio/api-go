package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	os.Exit(runMain(os.Args[1:], os.Stdin, os.Stdout, os.Stderr))
}

func runMain(args []string, stdin io.Reader, stdout io.Writer, stderr io.Writer) int {
	var variant string
	var outDir string
	var stableVersion string
	flagSet := flag.NewFlagSet("experimentgen", flag.ContinueOnError)
	flagSet.SetOutput(stderr)
	flagSet.StringVar(&variant, "variant", "", "experimental variant name")
	flagSet.StringVar(&outDir, "out", ".", "module root to write generated files into")
	flagSet.StringVar(&stableVersion, "stable-version", "", "stable module version to use in go.mod (default: resolved from git tags)")
	if err := flagSet.Parse(args); err != nil {
		return 2
	}

	if variant == "" {
		fmt.Fprintln(stderr, "usage: buf build -o - | go run ./cmd/experimentgen -variant <name> [-out <module-root>]")
		return 2
	}

	data, err := io.ReadAll(stdin)
	if err != nil {
		fmt.Fprintln(stderr, err)
		return 1
	}

	gen := generator{}
	if stableVersion != "" {
		// Use the provided stable version; still resolve Go version from go.mod.
		gen.resolveStableVersion = func(apiRepo string) (moduleVersion, error) {
			goVersion := "1.21"
			if ver, err := readGoVersionFromMod(apiRepo); err == nil && ver != "" {
				goVersion = ver
			}
			return moduleVersion{Tag: strings.TrimSpace(stableVersion), GoVersion: goVersion}, nil
		}
	}

	if err := gen.generate(data, variant, outDir); err != nil {
		fmt.Fprintln(stderr, err)
		return 1
	}
	fmt.Fprintf(stdout, "generated %s\n", outDir)
	return 0
}
