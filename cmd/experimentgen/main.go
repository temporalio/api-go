package main

import (
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	os.Exit(runMain(os.Args[1:], os.Stdin, os.Stdout, os.Stderr))
}

func runMain(args []string, stdin io.Reader, stdout io.Writer, stderr io.Writer) int {
	var variant string
	var outDir string
	flagSet := flag.NewFlagSet("experimentgen", flag.ContinueOnError)
	flagSet.SetOutput(stderr)
	flagSet.StringVar(&variant, "variant", "", "experimental variant name")
	flagSet.StringVar(&outDir, "out", ".", "module root to write generated files into")
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
	if err := gen.generate(data, variant, outDir); err != nil {
		fmt.Fprintln(stderr, err)
		return 1
	}
	fmt.Fprintf(stdout, "generated %s\n", outDir)
	return 0
}
