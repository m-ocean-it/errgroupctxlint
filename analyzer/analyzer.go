package analyzer

import (
	"errors"
	"go/ast"

	"github.com/m-ocean-it/errgroup-ctx-lint/analyzer/func_visitor"
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

func DefaultConfig() func_visitor.Config {
	return func_visitor.Config{
		ErrgroupPackagePaths: []string{
			"golang.org/x/sync/errgroup",
		},
	}
}

func NewAnalyzer() *analysis.Analyzer {
	return NewAnalyzerWithConfig(DefaultConfig())
}

func NewAnalyzerWithConfig(cfg func_visitor.Config) *analysis.Analyzer {
	return newAnalyzer(cfg)
}

func newAnalyzer(cfg func_visitor.Config) *analysis.Analyzer {
	return &analysis.Analyzer{ //nolint:exhaustruct
		Name:     "errgroupctx",
		Doc:      "Checks that errgroup closures use the context derived from a corresponding errgroup",
		Run:      Run(cfg),
		Requires: []*analysis.Analyzer{inspect.Analyzer},
	}
}

func Run(cfg func_visitor.Config) func(*analysis.Pass) (any, error) {
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

		thisFuncVisitor := func_visitor.New(pass, cfg)

		inspector.WithStack(nodeFilter, thisFuncVisitor.Visit)

		return nil, nil
	}
}
