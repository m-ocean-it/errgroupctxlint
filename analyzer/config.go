package analyzer

import "errors"

const DefaultPkgPath = "golang.org/x/sync/errgroup"

type FuncVisitorConfig struct {
	ErrgroupPackagePaths []string `json:"errgroup_package_paths"`
}

func (c *FuncVisitorConfig) prepare() error {
	if c == nil {
		return errors.New("config is nil")
	}

	if len(c.ErrgroupPackagePaths) == 0 {
		c.ErrgroupPackagePaths = []string{
			DefaultPkgPath,
		}
	}

	return nil
}
