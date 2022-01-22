package repo

import (
	"bytes"

	"github.com/faetools/devtool/internal/templates"
	"github.com/spf13/cobra"
)

// Badges returns the badges in the ReadMe.
func (d Definition) Badges() string {
	b := &bytes.Buffer{}
	cobra.CheckErr(templates.Badges.Execute(b, d))
	return b.String()
}

// CoverageColor returns the color the badge for coverage is displayed in.
func (d Definition) CoverageColor() string {
	switch {
	case d.Coverage < 40:
		return "red"
	case d.Coverage < 60:
		return "orange"
	case d.Coverage < 80:
		return "green"
	default:
		return "brightgreen"
	}
}
