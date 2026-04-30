package main

import (
	"bytes"
	_ "embed"
	"os"
	"path/filepath"
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

func writeTemplate(path string, tmpl string, data any) error {
	t, err := template.New(filepath.Base(path)).Parse(tmpl)
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
