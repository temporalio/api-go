package main

import (
	"flag"
	"fmt"
	"runtime"

	"go.temporal.io/api/cmd/protoc-gen-go-nexus/plugin"
	"google.golang.org/protobuf/compiler/protogen"
)

var (
	version = "dev"
	commit  = "latest"
)

func main() {
	showVersion := flag.Bool("version", false, "print the version and exit")
	flag.Parse()
	if *showVersion {
		fmt.Printf("protoc-gen-go-nexus: %s (%s)\n", version, commit)
		fmt.Printf("go: %s\n", runtime.Version())
		return
	}

	p := plugin.New(version, commit)

	opts := protogen.Options{
		ParamFunc: p.Param,
	}

	opts.Run(p.Run)
}
