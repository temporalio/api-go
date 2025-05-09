package main

import (
	"bytes"
	"fmt"
	"go/format"
	"go/types"
	"log"
	"os"
	"strings"

	"golang.org/x/tools/go/packages"
	"golang.org/x/tools/go/types/typeutil"

	_ "go.temporal.io/api/workflowservice/v1"
)

const serviceFile = "../../proxy/service.go"

const ServiceHeader = `
// Code generated by proxygenerator; DO NOT EDIT.

package proxy

import (
	"context"

	"go.temporal.io/api/workflowservice/v1"
)

// WorkflowServiceProxyOptions provides options for configuring a WorkflowServiceProxyServer.
// Client is a WorkflowServiceClient used to forward requests received by the server to the 
// Temporal Frontend.
type WorkflowServiceProxyOptions struct {
	Client workflowservice.WorkflowServiceClient
	DisableHeaderForwarding bool
}

type workflowServiceProxyServer struct {
	workflowservice.UnimplementedWorkflowServiceServer
	client workflowservice.WorkflowServiceClient
	disableHeaderForwarding bool
}

// NewWorkflowServiceProxyServer creates a WorkflowServiceServer suitable for registering with a gRPC Server. Requests will
// be forwarded to the passed in WorkflowService Client. gRPC interceptors can be added on the Server or Client to adjust
// requests and responses.
func NewWorkflowServiceProxyServer(options WorkflowServiceProxyOptions) (workflowservice.WorkflowServiceServer, error) {
	return &workflowServiceProxyServer{
		client: options.Client,
		disableHeaderForwarding: options.DisableHeaderForwarding,
	}, nil
}
`

func generateService(cfg config) error {
	buf := &bytes.Buffer{}
	fmt.Fprint(buf, ServiceHeader)

	conf := &packages.Config{Mode: packages.NeedImports | packages.NeedTypes | packages.NeedTypesInfo}
	pkgs, err := packages.Load(conf, "go.temporal.io/api/workflowservice/v1")
	if err != nil {
		return fmt.Errorf("unable to load workflowservice: %w", err)
	}

	pkg := pkgs[0]
	if len(pkg.Errors) > 0 {
		return fmt.Errorf("unable to load workflowservice: %v", pkg.Errors)
	}

	qual := func(other *types.Package) string {
		if other == pkg.Types {
			return "workflowservice"
		}
		return other.Path()
	}
	scope := pkg.Types.Scope()
	service := scope.Lookup("UnimplementedWorkflowServiceServer")
	if _, ok := service.(*types.TypeName); ok {
		for _, meth := range typeutil.IntuitiveMethodSet(service.Type(), nil) {
			if !meth.Obj().Exported() {
				continue
			}

			name := meth.Obj().Name()
			sig := meth.Obj().Type().(*types.Signature)
			params := make([]string, sig.Params().Len())

			// at some point the parsed type signature stopped carrying the variable's names, so we need to manually name them
			counter := 0
			paramDecl := make([]string, sig.Params().Len())
			for i := 0; i < sig.Params().Len(); i++ {
				typ := sig.Params().At(i).Type()
				typeName := typ.String()
				if typeName == "context.Context" {
					params[i] = "ctx"
				} else {
					params[i] = fmt.Sprintf("in%d", counter)
					counter += 1
				}
				paramDecl[i] = fmt.Sprintf("%s %s", params[i], types.TypeString(typ, qual))
				// Wrap ctx parameter in reqCtx
				if params[i] == "ctx" {
					params[i] = "s.reqCtx(ctx)"
				}
			}
			fmt.Fprintf(buf, "\nfunc (s *workflowServiceProxyServer) %s(%s) %s {\n", name, strings.Join(paramDecl, ", "), types.TypeString(sig.Results(), qual))
			fmt.Fprintf(buf, "\treturn s.client.%s(%s)\n", name, strings.Join(params, ", "))
			fmt.Fprintf(buf, "}\n")
		}
	}

	src, err := format.Source(buf.Bytes())
	if err != nil {
		log.Println(buf.String())
		return fmt.Errorf("failed to format generated workflowservice: %w", err)
	}

	if cfg.verifyOnly {
		currentSrc, err := os.ReadFile(serviceFile)
		if err != nil {
			return err
		}

		if !bytes.Equal(src, currentSrc) {
			return fmt.Errorf("generated file does not match existing file: %s", serviceFile)
		}

		return nil
	}

	if err := os.WriteFile(serviceFile, src, 0666); err != nil {
		return fmt.Errorf("failed to save generated service file: %w", err)
	}

	return nil
}
