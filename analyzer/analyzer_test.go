package analyzer

import (
	"path/filepath"
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"
)

func Test_Analyzer(t *testing.T) {
	t.Parallel()

	cfg := FuncVisitorConfig{
		ErrgroupPackagePaths: []string{
			"github.com/m-ocean-it/errgroup-ctx-lint/testdata/base/errgroup",
		},
	}

	analysistest.Run(
		t,
		filepath.Join(analysistest.TestData(), "base"),
		NewAnalyzerWithConfigProvider(func() FuncVisitorConfig { return cfg }),
	)
}
