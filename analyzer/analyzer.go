package analyzer

import (
	"errors"
	"go/ast"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

func DefaultConfig() FuncVisitorConfig {
	return FuncVisitorConfig{
		ErrgroupPackagePaths: []string{
			"golang.org/x/sync/errgroup",
		},
	}
}

func NewAnalyzer() *analysis.Analyzer {
	return NewAnalyzerWithConfigProvider(DefaultConfig)
}

func NewAnalyzerWithConfigProvider(cfg func() FuncVisitorConfig) *analysis.Analyzer {
	return newAnalyzer(getRunFuncWithConfigProvider(cfg))
}

func newAnalyzer(runFunc func(*analysis.Pass) (any, error)) *analysis.Analyzer {
	return &analysis.Analyzer{ //nolint:exhaustruct
		Name:     "errgroupctx",
		Doc:      "Checks that errgroup closures use the context derived from a corresponding errgroup",
		Run:      runFunc,
		Requires: []*analysis.Analyzer{inspect.Analyzer},
	}
}

func getRunFuncWithConfigProvider(cfg func() FuncVisitorConfig) func(*analysis.Pass) (any, error) {
	return func(pass *analysis.Pass) (any, error) {
		inspector, ok := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)
		if !ok {
			return nil, errors.New("unexpectedly type is not *inspector.Inspector")
		}

		nodeFilter := []ast.Node{
			(*ast.FuncDecl)(nil),
			(*ast.AssignStmt)(nil),
			(*ast.DeclStmt)(nil),
			(*ast.CallExpr)(nil),
		}

		thisFuncVisitor := newFuncVisitor(pass, cfg())

		inspector.WithStack(nodeFilter, thisFuncVisitor.Visit)

		return nil, nil
	}
}
