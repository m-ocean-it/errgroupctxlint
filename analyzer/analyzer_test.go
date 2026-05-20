package analyzer

import (
	"path/filepath"
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"
)

func Test_Analyzer(t *testing.T) {
	t.Parallel()

	analyzer := NewAnalyzer()

	err := analyzer.Flags.Set("pkgs", "github.com/m-ocean-it/errgroupctxlint/testdata/base/errgroup")
	if err != nil {
		t.Fatal(err)
	}

	analysistest.Run(
		t,
		filepath.Join(analysistest.TestData(), "base"),
		analyzer,
	)
}
