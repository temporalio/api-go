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
	"go/format"
	"go/parser"
	"go/token"
	"log"
	"os"
	"strings"

	"go.temporal.io/api/internal/strcase"
)

type method struct {
	SelfVar string
	Decl    *ast.FuncDecl
}

// All protobuf enums are 32-bit ints thankfully
type enumVariant struct {
	Name string
	// Int32 as string
	Value string
}

func defaultClause(enumTypeName, selfVar string) *ast.CaseClause {
	return &ast.CaseClause{
		Body: []ast.Stmt{
			&ast.ReturnStmt{
				Results: []ast.Expr{
					&ast.CallExpr{
						Fun: &ast.SelectorExpr{
							X: &ast.Ident{
								Name: "strconv",
							},
							Sel: &ast.Ident{
								Name: "Itoa",
							},
						},
						Args: []ast.Expr{
							&ast.CallExpr{
								Fun: &ast.Ident{Name: "int"},
								Args: []ast.Expr{
									&ast.Ident{Name: selfVar},
								},
							},
						},
					},
				},
			},
		},
	}
}

func rewriteFile(fileName string) error {
	fset := token.NewFileSet()
	// we're going to overwrite files in-place so we read it ahead of time
	srcf, err := os.Open(fileName)
	if err != nil {
		return err
	}
	ff, err := parser.ParseFile(fset, fileName, srcf, parser.ParseComments)
	if err != nil {
		return err
	}
	// Close the file so we can overwrite it
	srcf.Close()

	// Pass 1: collect details
	var imports *ast.GenDecl
	// type name -> String() method
	stringMethods := map[string]method{}
	// enum type name -> variant name and value
	possibleEnums := map[string][]enumVariant{}
	ast.Inspect(ff, func(n ast.Node) bool {
		switch n := n.(type) {
		case *ast.GenDecl:
			switch n.Tok {
			case token.IMPORT:
				// We'll add strconv later if needed
				imports = n
			case token.TYPE:
				// Enums are just 32-bit ints
				for i := range n.Specs {
					tspec, ok := n.Specs[i].(*ast.TypeSpec)
					if !ok {
						continue
					}
					specType, ok := tspec.Type.(*ast.Ident)
					if !ok {
						continue
					}
					if specType.Name != "int32" {
						continue
					}
					// We don't assume the order of declarations or of Inspect, so only add
					// the enum if we haven't already encountered a variant
					if _, found := possibleEnums[tspec.Name.Name]; found {
						continue
					}
					possibleEnums[tspec.Name.Name] = nil
				}
			case token.CONST:
				// Enum variants are simple ValueSpecs with a single token.INT value
				for i := range n.Specs {
					vspec, ok := n.Specs[i].(*ast.ValueSpec)
					if !ok {
						continue
					}
					if len(vspec.Values) > 1 || len(vspec.Names) > 1 {
						continue
					}

					lit, ok := vspec.Values[0].(*ast.BasicLit)
					if !ok {
						continue
					}
					if lit.Kind != token.INT {
						continue
					}

					valType, ok := vspec.Type.(*ast.Ident)
					if !ok {
						continue
					}

					// I don't really care if we capture non-enums here, those will lack a String
					possibleEnums[valType.Name] = append(possibleEnums[valType.Name], enumVariant{
						Name:  vspec.Names[0].Name,
						Value: lit.Value,
					})
				}
			}
		// Collect String methods defined on types. We'll union this with the enums map later
		// to filter down to only the types we care about
		case *ast.FuncDecl:
			if n.Name.Name != "String" || n.Recv == nil || len(n.Recv.List) != 1 {
				return true
			}
			recv := n.Recv.List[0]
			if len(recv.Names) > 1 {
				return true
			}
			typeIdent, ok := recv.Type.(*ast.Ident)
			if !ok {
				return true
			}
			stringMethods[typeIdent.Name] = method{
				SelfVar: recv.Names[0].Name,
				Decl:    n,
			}

		}
		return true
	})

	// Pass 2: rewrite string methods for our enums
	if len(possibleEnums) == 0 || len(stringMethods) == 0 {
		return nil
	}
	imports.Specs = append([]ast.Spec{&ast.ImportSpec{
		Path: &ast.BasicLit{
			Kind:  token.STRING,
			Value: "\"strconv\"",
		},
	}}, imports.Specs...)

	for typeName, variants := range possibleEnums {
		stringMthd, found := stringMethods[typeName]
		if !found {
			continue
		}

		switchStmt := &ast.SwitchStmt{
			Tag: &ast.Ident{
				Name: stringMthd.SelfVar,
			},
			Body: &ast.BlockStmt{},
		}
		for i := range variants {
			displayName := strings.TrimPrefix(strcase.ToCamel(variants[i].Name), typeName)
			clause := &ast.CaseClause{
				List: []ast.Expr{
					&ast.Ident{
						Name: variants[i].Name,
					},
				},
				Body: []ast.Stmt{
					&ast.ReturnStmt{
						Results: []ast.Expr{
							&ast.BasicLit{
								Kind:  token.STRING,
								Value: fmt.Sprintf("\"%s\"", displayName),
							},
						},
					},
				},
			}
			switchStmt.Body.List = append(switchStmt.Body.List, clause)
		}
		switchStmt.Body.List = append(switchStmt.Body.List, defaultClause(typeName, stringMthd.SelfVar))
		stringMthd.Decl.Body.List = []ast.Stmt{switchStmt}
	}

	of, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer of.Close()

	return format.Node(of, fset, ff)
}

func main() {
	for _, fileName := range os.Args[1:] {
		// When using `go run` you must provide this before your files
		// otherwise Go will think you're trying to execute both files...
		// ex: `go run enum-rewriter/main.go -- other.go`
		if fileName == "--" {
			continue
		}
		if err := rewriteFile(fileName); err != nil {
			log.Fatalf("Failed to rewrite %s: %s", fileName, err)
		}
	}

}
