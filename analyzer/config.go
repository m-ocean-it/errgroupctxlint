package analyzer

import (
	"errors"
	"strings"
)

const DefaultPkgPath = "golang.org/x/sync/errgroup"

type PackagePaths []string

func (p *PackagePaths) String() string {
	if p == nil {
		return ""
	}

	return strings.Join(*p, ",")
}

func (p *PackagePaths) Set(value string) error {
	if value == "" {
		return errors.New("empty value")
	}

	for v := range strings.SplitSeq(value, ",") {
		*p = append(*p, strings.TrimSpace(v))
	}

	return nil
}
