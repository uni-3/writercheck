package main

import (
	"fmt"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/analysis"
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


	return nil, nil
}
