package writercheck

import (
	"fmt"
	"go/ast"
	"go/types"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

var Analyzer = &analysis.Analyzer{
	Name: "writercheck",
	Doc:  "check for implementation writer interface",
	Run:  run,
	Requires: []*analysis.Analyzer{
		inspect.Analyzer,
	},
}

func run(pass *analysis.Pass) (interface{}, error) {
	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	nodeFilter := []ast.Node{
		(*ast.Ident)(nil),
	}
	inspect.Preorder(nodeFilter, func(n ast.Node) {
		switch n := n.(type) {
		case *ast.Ident:
			if n.Name == "Gopher" {
				pass.Reportf(n.Pos(), "name of identifier must not be 'Gopher'")
			}
		}
	})

	for expr, typ := range pass.TypesInfo.Types {
		switch expr.(type) {
		case *ast.BinaryExpr:
			obj := types.Universe.Lookup("string").Type().(*types.Interface)
			if types.Implements(typ.Type, obj) {
				fmt.Println(typ.Value, "implements error")
			}
		}
	}

	/*
		for _, f := range pass.Files {
			for _, decl := range f.Decls {
				if decl, ok := decl.(*ast.StructType); ok {
					ast.Print(nil, decl.Fields.List)
					for _, p := range decl.Fields.List {
						obj := types.Universe.Lookup("string").Type().(*types.Interface)
						if types.Implements(p.Type, obj) {
							fmt.Println(f, "implements error")
						}
					}
				}
			}
		}
	*/

	return nil, nil
}

type foundFact struct{}

func (*foundFact) String() string { return "found" }
func (*foundFact) AFact()         {}
