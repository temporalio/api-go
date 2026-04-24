package main

import (
	"flag"
	"fmt"
	"go/format"
	"log"
	"os"
	"strings"
)

// Print the code being modified to allow debugging any formatting-related problems.
// e.g. $ go run ./cmd/mockgen-fix --verbose Service service_grpc.pb.mock.go
var verbose = flag.Bool("verbose", false, "Print the results of the mockgen fix")

// This code fixes the generated mock service server code to properly embed
// the necessary unimplemented struct and add the interface assertion.
func main() {
	// Parse flags and args.
	flag.Parse()

	if flag.NArg() != 2 {
		fmt.Println("Usage: mockgen-fix [service] [filename]")
		fmt.Println("Error: No service or filename provided.")
	}
	service := flag.Arg(0)
	filename := flag.Arg(1)

	if err := run(service, filename); err != nil {
		log.Fatal(err)
	}
}

func run(serviceName, filename string) error {

	// Read file
	b, err := os.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("unable to read file: %w", err)
	}

	source := string(b)
	if *verbose {
		fmt.Println("Source Mockgen Code")
		fmt.Println("===================")
		printSourceFile(source)
		fmt.Println()
	}

	// Find the first "Server struct {" location.
	toFind := fmt.Sprintf("Mock%vServer struct {\n", serviceName)
	structIndex := strings.Index(source, toFind)
	if structIndex < 0 || strings.LastIndex(source, toFind) != structIndex {
		return fmt.Errorf("expected single server struct in file")
	}
	structIndex += len(toFind)

	// At the first newline we need to embed the unimplemented server.
	source = source[:structIndex] +
		fmt.Sprintf("\t%v.Unimplemented%vServer\n", strings.ToLower(serviceName), serviceName) +
		source[structIndex:]

	// After the closing brace ending the struct's definition, we need to add a type
	// assertion to ensure interface conformance.
	endBrace := structIndex + strings.Index(source[structIndex:], "\n}\n") + 2
	source = source[:endBrace] +
		fmt.Sprintf(
			"\nvar _ %v.%vServer = (*Mock%vServer)(nil)\n\n",
			strings.ToLower(serviceName),
			serviceName,
			serviceName,
		) +
		source[endBrace:]

	// Format and write
	formatted, err := format.Source([]byte(source))
	if err != nil {
		if *verbose {
			fmt.Println("Resulting Code (pre-formatting)")
			fmt.Println("===============================")
			printSourceFile(source)
			fmt.Println()
		}
		return fmt.Errorf("failed formatting: %w", err)
	}

	if err := os.WriteFile(filename, formatted, 0644); err != nil {
		return fmt.Errorf("failed writing: %w", err)
	}
	return nil
}

// Prints the source file's contexts to STDOUT with line numbers.
func printSourceFile(source string) {
	lines := strings.Split(source, "\n")
	for i, line := range lines {
		fmt.Printf("%4d: %s\n", i+1, line)
	}
}
