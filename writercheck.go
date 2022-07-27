package writercheck

import (
	"go/ast"

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
		(*ast.FuncDecl)(nil),
	}
	inspect.Preorder(nodeFilter, func(n ast.Node) {
		fn := n.(*ast.FuncDecl)
		fname := fn.Name.Name
		if fname == "Write" {
			// 引数
			if fn.Type.Params.NumFields() != 1 {
				pass.Reportf(fn.Pos(), "%s's argument length is '%d' must be 1", fname, fn.Type.Params.NumFields())
				return
			}
			for _, fi := range fn.Type.Params.List {
				if fi.Names[0].Name != "p" {
					pass.Reportf(fn.Pos(), "%s's argument name is '%s' must be 'p'", fname, fi.Names[0].Name)
					return
				}
				switch ft := fi.Type.(type) {
				case *ast.ArrayType:
					switch et := ft.Elt.(type) {
					case *ast.Ident:
						if et.Name != "byte" {
							pass.Reportf(fn.Pos(), "%s argument is invalid type '%s' must be 'byte'", fi.Names[0].Name, et.Name)
							return
						}
					default:
						pass.Reportf(fn.Pos(), "%s argument is invalid", fi.Names[0].Name)
						return
					}
				default:
					pass.Reportf(fn.Pos(), "%s argument is invalid", fi.Names[0].Name)
					return
				}
			}

			// 返り値
			results := fn.Type.Results
			if results.NumFields() != 2 {
				pass.Reportf(fn.Pos(), "%s returns length '%d' must be 2", fname, results.NumFields())
				return
			}

			resInt := results.List[0]
			// 1つめ
			if resInt.Names[0].Name != "n" {
				pass.Reportf(fn.Pos(), "%s first return name is '%s' must be 'n'", fname, resInt.Names[0].Name)
			}
			switch typ := resInt.Type.(type) {
			case *ast.Ident:
				if typ.Name != "int" {
					pass.Reportf(fn.Pos(), "%s first return type is '%s' must be 'int'", fname, typ.Name)
				}
			}

			resErr := results.List[1]
			if resErr.Names[0].Name != "err" {
				pass.Reportf(fn.Pos(), "%s second return name is '%s' must be 'p'", fname, resErr.Names[0].Name)
			}
			switch typ := resErr.Type.(type) {
			case *ast.Ident:
				if typ.Name != "error" {
					pass.Reportf(fn.Pos(), "%s second return type is '%s' must be 'byte'", resErr.Names[0].Name, typ.Name)
				}
			}
		}
	})

	return nil, nil
}
