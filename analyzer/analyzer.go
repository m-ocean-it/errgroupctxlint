package analyzer

import (
	"errors"
	"flag"
	"go/ast"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

type linter struct {
	errgroupPackagePaths PackagePaths
}

func (l *linter) run(pass *analysis.Pass) (any, error) {
	insp, ok := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)
	if !ok {
		return nil, errors.New("unexpectedly type is not *inspector.Inspector")
	}

	nodeFilter := []ast.Node{
		(*ast.FuncDecl)(nil),
		(*ast.AssignStmt)(nil),
		(*ast.DeclStmt)(nil),
		(*ast.CallExpr)(nil),
	}

	if len(l.errgroupPackagePaths) == 0 {
		l.errgroupPackagePaths = append(l.errgroupPackagePaths, DefaultPkgPath)
	}

	visitor := newFuncVisitor(pass, l.errgroupPackagePaths)

	insp.WithStack(nodeFilter, visitor.Visit)

	return nil, nil //nolint:nilnil
}

func NewAnalyzer() *analysis.Analyzer {
	//nolint:exhaustruct
	l := &linter{}

	//nolint:exhaustruct
	a := &analysis.Analyzer{
		Name:     "errgroupctx",
		Doc:      "Checks that errgroup closures use the context derived from a corresponding errgroup",
		Run:      l.run,
		Requires: []*analysis.Analyzer{inspect.Analyzer},
	}

	a.Flags.Init("errgroupctx", flag.ExitOnError)

	a.Flags.Var(
		&l.errgroupPackagePaths,
		"pkgs",
		"Comma-separated list of packages that provide an errgroup. Use in case you're dealing with a non-standard errgroup library.",
	)

	return a
}
