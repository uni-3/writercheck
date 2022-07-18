package writercheck

import (
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
	Requiers: []*analysis.Analyzer{
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

	for _, f := range pass.Files {
		for _, decl := range f.Decls {
			if decl, ok := decl.(*ast.FuncDecl); ok && decl.Name.Name == name {
				if obj, ok := pass.TypesInfo.Defs[decl.Name].(*types.Func); ok {
					pass.ExportObjectFact(obj, new(foundFact))
				}
			}
		}

	}
	return nil, nil
}

type foundFact struct{}

func (*foundFact) String() string { return "found" }
func (*foundFact) AFact()         {}
