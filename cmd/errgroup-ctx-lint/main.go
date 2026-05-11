package main

import (
	"flag"
	"strings"

	"github.com/m-ocean-it/errgroup-ctx-lint/analyzer"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	pkgsVar := flag.String("pkgs",
		"golang.org/x/sync/errgroup", // Default.
		"Comma-separated list of packages that provide an errgroup. Use in case you're dealing with a non-standard errgroup library.",
	)
	// flag.Parse will be called inside singlechecker.Main below.

	configProvider := func() analyzer.FuncVisitorConfig {
		cfg := analyzer.DefaultConfig()

		if strings.TrimSpace(*pkgsVar) == "" {
			return cfg
		}

		cfg.ErrgroupPackagePaths = []string{}
		for p := range strings.SplitSeq(*pkgsVar, ",") {
			pTrimmed := strings.TrimSpace(p)
			cfg.ErrgroupPackagePaths = append(cfg.ErrgroupPackagePaths, pTrimmed)
		}

		return cfg
	}

	analyzerInstance := analyzer.NewAnalyzerWithConfigProvider(configProvider)

	singlechecker.Main(analyzerInstance)
}
