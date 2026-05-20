package analyzer

import (
	"path/filepath"
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"
)

func Test_Analyzer_flags(t *testing.T) {
	t.Parallel()

	analyzer := NewAnalyzer(nil)

	err := analyzer.Flags.Set(PackagePathsFlg, "github.com/m-ocean-it/errgroupctxlint/testdata/base/errgroup")
	if err != nil {
		t.Fatal(err)
	}

	analysistest.Run(
		t,
		filepath.Join(analysistest.TestData(), "base"),
		analyzer,
	)
}

func Test_Analyzer_config(t *testing.T) {
	t.Parallel()

	analyzer := NewAnalyzer([]string{"github.com/m-ocean-it/errgroupctxlint/testdata/base/errgroup"})

	analysistest.Run(
		t,
		filepath.Join(analysistest.TestData(), "base"),
		analyzer,
	)
}
