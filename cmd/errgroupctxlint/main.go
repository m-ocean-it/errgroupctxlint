package main

import (
	"github.com/m-ocean-it/errgroupctxlint/analyzer"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	singlechecker.Main(analyzer.NewAnalyzer())
}
