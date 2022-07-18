package writercheck

import (
	"fmt"
	"go/ast"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

var Analyzer = &Analyzer{
	Name: "writercheck",
	Doc:  "check for implementation writer interface",
	Run:  run,
	Requiers: []*analysis.Analyzer{
		inspect.Analyzer,
	},
}

func run(pass *analysis.Pass) (interface{}, error) {
	fmt.Println(pass.Analyzer)
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
	return nil, nil
}
