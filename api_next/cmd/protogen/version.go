package main

import (
	"go/ast"
	"regexp"
)

type versionRemover struct{}

var rgx = regexp.MustCompile(`(//\s+(- )?protoc[-\w]*\s+)([\w.]+)`)

// NewRemover returns an ast visitor that rewrites the
// comments added by protoc that contain plugin versions.
// Our CI servers may run a different version of protoc than our local boxes
// (alpine is woefully out of date) and we don't want that to fail our buids.
func NewVersionRemover() versionRemover {
	return versionRemover{}
}

// Process the ast.File, identifying protoc plugin version comments
// and rewriting them without the version
func (v versionRemover) Process(f *ast.File) {
	for i := 0; i < len(f.Comments); i++ {
		cg := f.Comments[i]
		for j := 0; j < len(cg.List); j++ {
			c := cg.List[j]
			if c.Text == "// versions:" {
				c.Text = "// plugins:"
				continue
			}
			matches := rgx.FindStringSubmatch(c.Text)
			if matches == nil {
				continue
			}
			c.Text = matches[1]
		}
	}
}
