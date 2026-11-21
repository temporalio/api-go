package main

import (
	"fmt"
	"go/format"
	"log"
	"os"
	"strings"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	// This code fixes the generated mock service server code to properly embed
	// the necessary unimplemented struct and add the interface assertion

	if len(os.Args) != 3 {
		return fmt.Errorf("must provide service name and code filename as arguments")
	}
	// Read file
	b, err := os.ReadFile(os.Args[2])
	if err != nil {
		return fmt.Errorf("unable to read file: %w", err)
	}
	source := string(b)
	serviceName := os.Args[1]

	// Find the first "Server struct {" location
	toFind := fmt.Sprintf("Mock%vServer struct {\n", serviceName)
	structIndex := strings.Index(source, toFind)
	if structIndex < 0 || strings.LastIndex(source, toFind) != structIndex {
		return fmt.Errorf("expected single server struct in file")
	}
	structIndex += len(toFind)

	// At the first newline we need to embed the unimplemented server
	source = source[:structIndex] +
		fmt.Sprintf("\t%v.Unimplemented%vServer\n", strings.ToLower(serviceName), serviceName) +
		source[structIndex:]

	// After the closing brace, we need to add the type assertion to ensure
	// interface conformance
	endBrace := structIndex + strings.Index(source[structIndex:], "}\n") + 2
	source = source[:endBrace] +
		fmt.Sprintf(
			"\nvar _ %v.%vServer = (*Mock%vServer)(nil)\n\n",
			strings.ToLower(serviceName),
			serviceName,
			serviceName,
		) +
		source[endBrace:]

	// Format and write
	if b, err := format.Source([]byte(source)); err != nil {
		return fmt.Errorf("failed formatting: %w", err)
	} else if err := os.WriteFile(os.Args[2], b, 0644); err != nil {
		return fmt.Errorf("failed writing: %w", err)
	}
	return nil
}
