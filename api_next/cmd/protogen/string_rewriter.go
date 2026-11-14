package main

import (
	"fmt"
	"go/ast"
	"go/token"
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

type stringRewriter struct {
	// imports. will add to them as needed later
	imports *ast.GenDecl
	// type name -> String() method
	stringMethods map[string]method
	// enum type name -> variant name and value
	possibleEnums map[string][]enumVariant
}

// NewStringRewriter returns an ast visitor that rewrites the
// String methods on protobuf-generated enums
func NewStringRewriter() *stringRewriter {
	return &stringRewriter{
		stringMethods: map[string]method{},
		possibleEnums: map[string][]enumVariant{},
	}
}

// Visit all AST nodes, changing behavior based on the node's kind.
// This is the first pass of enum rewriting: we collect possible enum variants
// and string methods so that in [Finalize] we can rewrite them
func (v *stringRewriter) visit(n ast.Node) bool {
	switch n := n.(type) {
	case *ast.GenDecl:
		switch n.Tok {
		case token.IMPORT:
			// We'll add strconv later if needed
			v.imports = n
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
				if _, found := v.possibleEnums[tspec.Name.Name]; found {
					continue
				}
				v.possibleEnums[tspec.Name.Name] = nil
			}
		case token.CONST:
			// Enum variants are simple ValueSpecs with a single token.INT value
			for i := range n.Specs {
				vspec, ok := n.Specs[i].(*ast.ValueSpec)
				if !ok {
					continue
				}
				if len(vspec.Values) != 1 || len(vspec.Names) != 1 {
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
				v.possibleEnums[valType.Name] = append(v.possibleEnums[valType.Name], enumVariant{
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
		v.stringMethods[typeIdent.Name] = method{
			SelfVar: recv.Names[0].Name,
			Decl:    n,
		}

	}
	return true
}

// Process the AST by rewriting the String method for all identified enums.
// This produces a String method that returns the old-style CamelCase enum variant for display
// purposes
func (v *stringRewriter) Process(f *ast.File) {
	ast.Inspect(f, v.visit)

	if len(v.possibleEnums) == 0 || len(v.stringMethods) == 0 {
		return
	}
	v.imports.Specs = append([]ast.Spec{&ast.ImportSpec{
		Path: &ast.BasicLit{
			Kind:  token.STRING,
			Value: "\"strconv\"",
		},
	}}, v.imports.Specs...)

	for typeName, variants := range v.possibleEnums {
		stringMthd, found := v.stringMethods[typeName]
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
