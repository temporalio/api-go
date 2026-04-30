package main

import (
	"bytes"
	_ "embed"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

type serviceInfo struct {
	PackageName      string
	StableImportPath string
}

type renderData struct {
	Service       serviceInfo
	Overlays      []messageOverlay
	OverlayGroups []messageOverlayGroup
	Enums         []enumInfo
	// Service stubs
	VariantTitle string     // CamelCase variant name, e.g. "Example"
	ProtoPackage string     // proto package, e.g. "temporal.api.workflowservice.v1"
	Methods      []methodInfo
}

var workflowService = serviceInfo{
	PackageName:      "workflowservice",
	StableImportPath: "go.temporal.io/api/workflowservice/v1",
}

var enumPackage = serviceInfo{
	PackageName:      "enums",
	StableImportPath: "go.temporal.io/api/enums/v1",
}

//go:embed templates/overlay.go.tmpl
var overlayTemplate string

//go:embed templates/enum.go.tmpl
var enumTemplate string

//go:embed templates/service.go.tmpl
var serviceTemplate string

func writeTemplate(path string, tmpl string, data any) error {
	funcs := template.FuncMap{
		"lc": func(s string) string {
			if s == "" {
				return ""
			}
			return strings.ToLower(s[:1]) + s[1:]
		},
	}
	t, err := template.New(filepath.Base(path)).Funcs(funcs).Parse(tmpl)
	if err != nil {
		return err
	}
	var buf bytes.Buffer
	if err := t.Execute(&buf, data); err != nil {
		return err
	}
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return err
	}
	return os.WriteFile(path, buf.Bytes(), 0o644)
}
