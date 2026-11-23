package main

import (
	"bytes"
	"compress/gzip"
	"flag"
	"fmt"
	"go/format"
	"html/template"
	"io"
	"os"
	"path/filepath"
	"strings"
)

type templateInput struct {
	Version int
	Format  string
	Spec    string
}

const tmpl = `package openapi

// OpenAPIV{{.Version}}{{.Format}}Spec contains a gzip-compressed {{.Format}} file specifying the Temporal HTTP API
var OpenAPIV{{.Version}}{{.Format}}Spec = {{.Spec}}`

func die(msg string, args ...any) {
	fmt.Fprintf(os.Stderr, msg+"\n", args...)
	os.Exit(1)
}

func prepareSpec(version int, input, output string) {
	extension := strings.TrimPrefix(filepath.Ext(input), ".")
	f, err := os.Open(input)
	if err != nil {
		die("Failed to open spec file %q: %v", input, err)
	}
	defer f.Close()

	var b bytes.Buffer
	w := gzip.NewWriter(&b)
	if _, err := io.Copy(w, f); err != nil {
		die("Failed to compress v%d spec: %s", version, err)
	}
	if err := w.Close(); err != nil {
		die("Failed to compress v%d spec: %s", version, err)
	}

	var src bytes.Buffer
	t := template.Must(template.New("spec").Parse(tmpl))
	t.Execute(&src, templateInput{
		Version: version,
		Format:  strings.ToTitle(extension),
		Spec:    fmt.Sprintf("%#v", b.Bytes()),
	})

	fmtd, err := format.Source(src.Bytes())
	if err != nil {
		die("Failed to format generated v%d code: %s", version, err)
	}

	out, err := os.Create(output)
	if err != nil {
		die("Failed to open %q: %s", output, err)
	}
	defer out.Close()
	if _, err := out.Write(fmtd); err != nil {
		die("Failed to write v%d code: %s", version, err)
	}

}

func main() {
	var v2Path, v3Path, v2Out, v3Out string
	flag.StringVar(&v2Path, "v2", "", "The path to the OpenAPI v2 spec file. Required.")
	flag.StringVar(&v3Path, "v3", "", "The path to the OpenAPI v3 spec file. Required.")
	flag.StringVar(&v2Out, "v2-out", "", "The path to the v2 output file. Required.")
	flag.StringVar(&v3Out, "v3-out", "", "The path to the v3 output file. Required.")
	flag.Parse()
	if v2Path == "" || v3Path == "" || v2Out == "" || v3Out == "" {
		flag.Usage()
		os.Exit(127)
	}

	prepareSpec(2, v2Path, v2Out)
	prepareSpec(3, v3Path, v3Out)
}
