// The MIT License
//
// Copyright (c) 2022 Temporal Technologies Inc.  All rights reserved.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

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
