package main

import (
	"go/ast"
	"go/token"
	"strings"
)

type constRewriter struct {
	// imports. will add to them as needed later
	imports *ast.GenDecl
	// enum type name -> prefix to remove
	typePrefixes map[string]string
	// variant -> type name
	variants map[string]string
}

// NewConstRewriter returns an ast visitor that rewrites the
// Go constants that define enum variants
// Takes in a map of known enum type -> prefix to trim to handle any edge cases
func NewConstRewriter(knownEnums map[string]string) *constRewriter {
	// Copy it just in case
	enums := make(map[string]string, len(knownEnums))
	for k, v := range knownEnums {
		// Need a trailing _ or the rewrite will be incorrect
		if !strings.HasSuffix(v, "_") {
			v = v + "_"
		}
		enums[k] = v
	}
	return &constRewriter{
		typePrefixes: enums,
		variants:     map[string]string{},
	}
}

func (v *constRewriter) collectEnums(n ast.Node) bool {
	switch n := n.(type) {
	case *ast.ValueSpec:
		if len(n.Values) == 0 || len(n.Values) > 1 || len(n.Names) > 1 {
			return true
		}

		lit, ok := n.Values[0].(*ast.BasicLit)
		if !ok || lit.Kind != token.INT {
			return true
		}

		valType, ok := n.Type.(*ast.Ident)
		if !ok {
			return true
		}
		v.variants[n.Names[0].Name] = valType.Name
		newName := strings.TrimPrefix(n.Names[0].Name, v.typePrefixes[valType.Name])
		*n = *renameValspec(n, newName, valType.Name)
	case *ast.GenDecl:
		switch n.Tok {
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

				if _, ok := v.typePrefixes[tspec.Name.Name]; ok {
					continue
				}
				v.typePrefixes[tspec.Name.Name] = tspec.Name.Name + "_"
			}
		}
	}
	return true
}

// Rewrite references to known variants, removing the type prefix
func (v *constRewriter) rewriteReferences(n ast.Node) bool {
	switch n := n.(type) {
	case *ast.Ident:
		typeName, known := v.variants[n.Name]
		if !known {
			return true
		}

		prefix, known := v.typePrefixes[typeName]
		if !known {
			return true
		}

		name := strings.TrimPrefix(n.Name, prefix)
		ident := &ast.Ident{
			Name: name,
		}
		if n.Obj != nil {
			ident.Obj = &ast.Object{
				Kind: n.Obj.Kind,
				Name: name,
				Data: n.Obj.Data,
				Type: n.Obj.Type,
			}
		}
		*n = *ident
	}
	return true
}

// Visit the given AST node, changing behavior based on the node's kind.
// This is the first pass of enum rewriting: we collect possible enum variants
// and string methods so that in [Finalize] we can rewrite them
func (v *constRewriter) Process(f *ast.File) {
	ast.Inspect(f, v.collectEnums)
	ast.Inspect(f, v.rewriteReferences)
}

// Clone the first Name (and values, if necessary)
func renameValspec(v *ast.ValueSpec, name, typ string) *ast.ValueSpec {
	ni := ast.NewIdent(name)
	ni.NamePos = v.Names[0].NamePos
	return &ast.ValueSpec{
		Doc:     v.Doc,
		Names:   []*ast.Ident{ni},
		Values:  v.Values,
		Type:    v.Type,
		Comment: v.Comment,
	}
}
