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
	"fmt"
	"go/ast"
	"go/token"
	"strconv"
	"strings"

	"google.golang.org/protobuf/encoding/protowire"
)

type disableUtf8Validation struct{}

type xform struct {
	filters map[protowire.Number]func([]byte) []byte
	xforms  map[protowire.Number]*xform
}

var xformFileDescriptorProto = &xform{
	xforms: map[protowire.Number]*xform{
		4: &xform{
			xforms: map[protowire.Number]*xform{
				2: &xform{
					filters: map[protowire.Number]func([]byte) []byte{
						8: processFieldOptionsProto,
					},
				},
			},
		},
	},
}

func NewDisableUtf8Validation() *disableUtf8Validation {
	return &disableUtf8Validation{}
}

func (v *disableUtf8Validation) Process(f *ast.File) {
	ast.Inspect(f, v.visit)
}

func (v *disableUtf8Validation) visit(n ast.Node) bool {
	switch n := n.(type) {
	case *ast.File:
		return true
	case *ast.GenDecl:
		if n.Tok != token.VAR {
			break
		}
		for _, spec := range n.Specs {
			spec := spec.(*ast.ValueSpec)
			if len(spec.Names) != 1 || !strings.HasSuffix(spec.Names[0].Name, "_rawDesc") {
				continue
			}
			lit := spec.Values[0].(*ast.CompositeLit)
			byteArr := make([]byte, len(lit.Elts))
			for i, e := range lit.Elts {
				v := e.(*ast.BasicLit).Value
				if strings.HasPrefix(v, "0x") {
					v = v[2:]
				}
				byteVal, err := strconv.ParseUint(v, 16, 8)
				if err != nil {
					panic(err)
				}
				byteArr[i] = byte(byteVal)
			}
			newArr := transform(byteArr, xformFileDescriptorProto)
			newElts := make([]ast.Expr, len(newArr))
			for i, v := range newArr {
				newElts[i] = &ast.BasicLit{
					Kind:  token.INT,
					Value: fmt.Sprintf("%#02x", v),
				}
			}
			lit.Elts = newElts
		}
	}
	return false
}

func transform(b []byte, xf *xform) []byte {
	out := make([]byte, 0, len(b)*3/2)
	seen := make(map[protowire.Number]bool)
	for len(b) > 0 {
		num, typ, n := protowire.ConsumeTag(b)
		if n < 0 {
			panic("ConsumeTag")
		}
		b = b[n:]
		out = protowire.AppendTag(out, num, typ)
		switch typ {
		case protowire.VarintType:
			v, n := protowire.ConsumeVarint(b)
			b = b[n:]
			out = protowire.AppendVarint(out, v)
		case protowire.Fixed32Type:
			v, n := protowire.ConsumeFixed32(b)
			b = b[n:]
			out = protowire.AppendFixed32(out, v)
		case protowire.Fixed64Type:
			v, n := protowire.ConsumeFixed64(b)
			b = b[n:]
			out = protowire.AppendFixed64(out, v)
		case protowire.BytesType:
			v, n := protowire.ConsumeBytes(b)
			if filt := xf.filters[num]; filt != nil {
				v = filt(v)
				seen[num] = true
			} else if xf2 := xf.xforms[num]; xf2 != nil {
				v = transform(v, xf2)
			}
			b = b[n:]
			out = protowire.AppendBytes(out, v)
		default:
			panic("bad type")
		}
	}
	for num, filt := range xf.filters {
		if !seen[num] {
			out = protowire.AppendTag(out, num, protowire.BytesType)
			out = protowire.AppendBytes(out, filt(nil))
		}
	}
	return out
}

func processFieldOptionsProto(b []byte) []byte {
	out := make([]byte, len(b), len(b)+2)
	copy(out, b)
	out = protowire.AppendTag(out, 13, protowire.VarintType)
	out = protowire.AppendVarint(out, 0)
	return out
}
