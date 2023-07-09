package testcasesmap

import (
	"go/types"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
)

const doc = "testcasesmap finds testcases that are not defined with map"

var Analyzer = &analysis.Analyzer{
	Name: "testcasesmap",
	Doc:  doc,
	Run:  run,
	Requires: []*analysis.Analyzer{
		inspect.Analyzer,
	},
}

func run(pass *analysis.Pass) (interface{}, error) {
	for ident, obj := range pass.TypesInfo.Defs {
		switch ident.Name {
		case "testcases":
			_, isMap := obj.Type().Underlying().(*types.Map)
			if !isMap {
				pass.Reportf(ident.Pos(), "use map for %s", ident.Name)
			}
		}
	}
	return nil, nil
}
