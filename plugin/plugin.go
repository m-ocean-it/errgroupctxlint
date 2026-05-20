package plugin

import (
	"fmt"
	"strings"

	"github.com/golangci/plugin-module-register/register"
	"github.com/m-ocean-it/errgroupctxlint/analyzer"
	"golang.org/x/tools/go/analysis"
)

//nolint:gochecknoinits
func init() {
	register.Plugin("errgroupctxlint", New)
}

type Settings struct {
	Packages []string `json:"pkgs"`
}

type Plugin struct {
	settings Settings
}

//nolint:ireturn
func New(settings any) (register.LinterPlugin, error) {
	s, err := register.DecodeSettings[Settings](settings)
	if err != nil {
		return nil, err
	}

	return &Plugin{settings: s}, nil
}

func (f *Plugin) BuildAnalyzers() ([]*analysis.Analyzer, error) {
	a := analyzer.NewAnalyzer()

	if len(f.settings.Packages) > 0 {
		err := a.Flags.Set(analyzer.PackagePathsFlg, strings.Join(f.settings.Packages, ","))
		if err != nil {
			return nil, fmt.Errorf("configuration error: %w", err)
		}
	}

	return []*analysis.Analyzer{analyzer.NewAnalyzer()}, nil
}

func (f *Plugin) GetLoadMode() string {
	return register.LoadModeTypesInfo
}
