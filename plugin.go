package errgroupctxlint

import (
	"github.com/golangci/plugin-module-register/register"
	"github.com/m-ocean-it/errgroup-ctx-lint/analyzer"
	"golang.org/x/tools/go/analysis"
)

func init() {
	register.Plugin("errgroupctx", New)
}

type Plugin struct {
	settings analyzer.FuncVisitorConfig
}

func New(settings any) (register.LinterPlugin, error) { //nolint:ireturn
	s, err := register.DecodeSettings[analyzer.FuncVisitorConfig](settings)
	if err != nil {
		return nil, err
	}

	return &Plugin{settings: s}, nil
}

func (f *Plugin) BuildAnalyzers() ([]*analysis.Analyzer, error) {
	return []*analysis.Analyzer{
		analyzer.NewAnalyzerWithConfigProvider(func() analyzer.FuncVisitorConfig { return f.settings }),
	}, nil
}

func (f *Plugin) GetLoadMode() string {
	return register.LoadModeTypesInfo
}
