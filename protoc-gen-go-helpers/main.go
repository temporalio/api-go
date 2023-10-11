// The MIT License
//
// Copyright (c) 2023 Temporal Technologies Inc.  All rights reserved.
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
	"bytes"
	"fmt"
	"go/format"
	"io"
	"strings"
	"text/template"

	"github.com/iancoleman/strcase"
	"google.golang.org/protobuf/compiler/protogen"
)

type (
	headerInput struct {
		Pkg     string
		Imports []string
	}

	msgInput struct {
		Type string
	}

	shortValue struct {
		Num int32
		Str string
	}

	enumInput struct {
		Type        string
		ShortValues []shortValue
	}
)

const (
	msgImport  = "google.golang.org/protobuf/proto"
	enumImport = "fmt"
	headerTmpl = `
package {{.Pkg}}

import  ({{range .Imports}}
    "{{.}}"{{end}}
)
`

	msgTmplStr = `
// Marshal an object of type {{.Type}} to the protobuf v3 wire format
func (val *{{.Type}}) Marshal() ([]byte, error) {
    return proto.Marshal(val)
}

// Unmarshal an object of type {{.Type}} from the protobuf v3 wire format
func (val *{{.Type}}) Unmarshal(buf []byte) error {
    return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *{{.Type}}) Size() int {
    return proto.Size(val)
}

// Equal returns whether two {{.Type}} values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *{{.Type}}) Equal(that interface{}) bool {
    if that == nil {
		return this == nil
	}

    var that1 *{{.Type}}
    switch t := that.(type) {
    case *{{.Type}}:
        that1 = t
    case {{.Type}}:
        that1 = &t
    default:
        return false
    }

    return proto.Equal(this, that1)
}`

	enumTmplStr = `
var (
    {{.Type}}_shorthandValue = map[string]int32{ {{range .ShortValues}}
        "{{.Str}}": {{.Num}},{{end}}
    }
    {{.Type}}_shorthandName = map[int32]string{ {{range .ShortValues}}
        {{.Num}}: "{{.Str}}",{{end}}
    }
)

// {{.Type}}FromString parses a {{.Type}} value from  either the protojson
// canonical SCREAMING_CASE enum or the traditional temporal PascalCase enum to {{.Type}}
func {{.Type}}FromString(s string) ({{.Type}}, error) {
    if v, ok := {{.Type}}_value[s]; ok {
        return {{.Type}}(v), nil
    } else if v, ok := {{.Type}}_shorthandValue[s]; ok {
        return {{.Type}}(v), nil
    }
    return {{.Type}}(0), fmt.Errorf("%s is not a valid {{.Type}}", s)
}`
)

// NOTE: If our implementation of Equal is too slow (its reflection-based) it doesn't look too
// hard to generate unrolled versions...
func main() {
	opts := protogen.Options{}
	opts.Run(func(plugin *protogen.Plugin) error {
		header, err := template.New("header").Parse(headerTmpl)
		if err != nil {
			return fmt.Errorf("failed to parse header template: %w", err)
		}
		msgTmpl, err := template.New("message").Parse(msgTmplStr)
		if err != nil {
			return fmt.Errorf("failed to parse message template: %w", err)
		}

		enumTmpl, err := template.New("enum").Parse(enumTmplStr)
		if err != nil {
			return fmt.Errorf("failed to parse enum template: %w", err)
		}

		for _, file := range plugin.Files {
			if !file.Generate || !strings.Contains(string(file.GoImportPath), "go.temporal.io") || (len(file.Proto.MessageType) == 0 && len(file.Proto.EnumType) == 0) {
				continue
			}
			hi := headerInput{
				Pkg: string(file.GoPackageName),
			}

			var body bytes.Buffer

			if len(file.Proto.MessageType) > 0 {
				hi.Imports = append(hi.Imports, msgImport)
			}
			for _, msg := range file.Proto.MessageType {
				if err := msgTmpl.Execute(&body, msgInput{Type: *msg.Name}); err != nil {
					return fmt.Errorf("failed to execute message template on type %s: %w", *msg.Name, err)
				}
			}

			if len(file.Proto.EnumType) > 0 {
				hi.Imports = append(hi.Imports, enumImport)
			}
			for _, enum := range file.Proto.EnumType {
				// Preprocess enum to create the mapping
				input := enumInput{
					Type: *enum.Name,
				}
				pfx := strcase.ToScreamingSnake(*enum.Name)
				for _, val := range enum.Value {
					input.ShortValues = append(input.ShortValues, shortValue{
						Num: *val.Number,
						Str: strcase.ToCamel(strings.TrimPrefix(val.GetName(), pfx)),
					})
				}
				if err := enumTmpl.Execute(&body, input); err != nil {
					return fmt.Errorf("failed to execute enum template on type %s: %w", *enum.Name, err)
				}
			}

			// Prepend header
			var src bytes.Buffer
			if err := header.Execute(&src, hi); err != nil {
				return fmt.Errorf("failed to execute header template: %w", err)
			}

			io.Copy(&src, &body)
			fmtd, err := format.Source(src.Bytes())
			if err != nil {
				return fmt.Errorf("failed to format generated source: \n%s\n%w", src.String(), err)
			}

			gf := plugin.NewGeneratedFile(fmt.Sprintf("%s.go-helpers.go", file.GeneratedFilenamePrefix), ".")
			gf.Write(fmtd)
		}

		return nil
	})
}
