package writercheck

import (
	"fmt"
	"go/ast"
	"go/types"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
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
	/*
		inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

		nodeFilter := []ast.Node{
			(*ast.Ident)(nil),
		}
		//ast.Print(nil, pass.Files)
		//fmt.Println("types", pass.TypesInfo)
			inspect.Preorder(nodeFilter, func(n ast.Node) {
				switch n := n.(type) {
				case *ast.Ident:
					if n.Name == "Gopher" {
						pass.Reportf(n.Pos(), "name of identifier must not be 'Gopher'")
					}
				}
			})
	*/

	fmt.Println(types.Universe.Names())
	for expr, typ := range pass.TypesInfo.Types {
		//ast.Print(nil, expr)
		//ast.Print(nil, typ.Type)
		switch e := expr.(type) {
		case *ast.BinaryExpr:
			obj := types.Universe.Lookup("string").Type()
			if types.Implements(typ.Type, obj.(*types.Interface)) {
				fmt.Println(typ.Value, "implements error")
			}
		case *ast.StructType:
			fmt.Println("str", expr.Pos(), e)
		case *ast.FuncType:
			fmt.Println("func", expr.Pos(), e)

		case *ast.FuncLit:
			fmt.Println("if", expr.Pos(), e)
		default:
			//fmt.Println(expr.Pos(), e)
			obj := types.Universe.Lookup("error").Type().Underlying()
			if !types.Implements(typ.Type, obj.(*types.Interface)) {
				//fmt.Println(typ.Type, "implements error")
				pass.Reportf(expr.Pos(), "must implement stringer")
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
