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
	Requires: []*analysis.Analyzer{
		inspect.Analyzer,
	},
}

func run(pass *analysis.Pass) (interface{}, error) {
	// 普通に愚直にinterface満たすか確認していく
	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	nodeFilter := []ast.Node{
		//(*ast.Ident)(nil),
		(*ast.FuncDecl)(nil),
	}
	//ast.Print(nil, pass.Files)
	//fmt.Println("types", pass.TypesInfo)
	inspect.Preorder(nodeFilter, func(n ast.Node) {
		switch n := n.(type) {
		case *ast.Ident:
			if n.Name == "Gopher" {
				pass.Reportf(n.Pos(), "name of identifier must not be 'Gopher'")
			}
		case *ast.FuncDecl:
			if n.Name.Name == "Write" {
				// 引数
				if n.Type.Params.NumFields() != 1 {
					pass.Reportf(n.Pos(), "%s arg length '%d' must be 1", n.Name.Name, n.Type.Params.NumFields())
				}
				for _, fi := range n.Type.Params.List {
					if fi.Names[0].Name != "p" {
						pass.Reportf(n.Pos(), "%s's an argument name is '%s' must be 'p'", n.Name.Name, fi.Names[0].Name)
					}
					switch ft := fi.Type.(type) {
					case *ast.ArrayType:
						switch et := ft.Elt.(type) {
						case *ast.Ident:
							if et.Name != "byte" {
								pass.Reportf(n.Pos(), "%s arg is invalid type '%s' must be 'byte'", fi.Names[0].Name, et.Name)
							}
						default:
							pass.Reportf(n.Pos(), "%s arg is invalid", fi.Names[0].Name)
						}

					default:
						pass.Reportf(n.Pos(), "%s arg is invalid", fi.Names[0].Name)
					}
				}

				// 返り値
				results := n.Type.Results
				if results.NumFields() != 2 {
					pass.Reportf(n.Pos(), "%s arg length '%d' must be 1", n.Name.Name, results.NumFields())
					return
				}

				resInt := results.List[0]
				// 1つめ
				if resInt.Names[0].Name != "n" {
					pass.Reportf(n.Pos(), "%s's an argument name is '%s' must be 'n'", n.Name.Name, resInt.Names[0].Name)
				}
				switch typ := resInt.Type.(type) {
				case *ast.Ident:
					if typ.Name != "int" {
						pass.Reportf(n.Pos(), "%s arg is invalid type '%s' must be 'int'", resInt.Names[0].Name, typ.Name)
					}
				default:
					pass.Reportf(n.Pos(), "%s arg is invalid", resInt.Names[1].Name)
				}

				resErr := results.List[1]
				// 2つめ
				if resErr.Names[0].Name != "err" {
					pass.Reportf(n.Pos(), "%s's an argument name is '%s' must be 'p'", n.Name.Name, resErr.Names[0].Name)
				}
				switch typ := resErr.Type.(type) {
				case *ast.Ident:
					if typ.Name != "error" {
						pass.Reportf(n.Pos(), "%s arg is invalid type '%s' must be 'byte'", resErr.Names[0].Name, typ.Name)
					}
				default:
					pass.Reportf(n.Pos(), "%s arg is invalid", resErr.Names[0].Name)
				}
			}
		}
	})

	return nil, nil
}

type foundFact struct{}

func (*foundFact) String() string { return "found" }
func (*foundFact) AFact()         {}

var errTyp = types.Universe.Lookup("error").Type().Underlying().(*types.Interface)

func isErrType(t types.Type) bool {
	return types.Implements(t, errTyp)
}
