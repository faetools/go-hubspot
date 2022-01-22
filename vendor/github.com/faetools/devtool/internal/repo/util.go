package repo

import (
	"regexp"
	"strings"

	"github.com/faetools/devtool/devtool"
	"github.com/faetools/devtool/internal/version"
	"github.com/spf13/afero"
)

func (d Definition) NameLowerCleaned() string {
	return strings.ReplaceAll(strings.ToLower(d.Name), "-", "")
}

func (d Definition) GoVersion() string           { return version.Go }
func (d Definition) DockerVersion() string       { return version.Docker }
func (d Definition) DevtoolVersion() string      { return devtool.Version }
func (d Definition) GolangciLintVersion() string { return version.GolangciLint }

func (d Definition) DependabotScheduledAt() string { return randTime(d.Name) }

var reCircleToken = regexp.MustCompile(`circle-token=([^\)]*)\)`)

const apiTokenPlaceholder = "YOUR_STATUS_API_TOKEN"

func (d Definition) CicleCIToken() string {
	fs := afero.NewOsFs()

	src, err := afero.ReadFile(fs, "README.md")
	if err != nil {
		return ""
	}

	match := reCircleToken.FindStringSubmatch(string(src))
	if len(match) > 1 &&
		!strings.Contains(match[1], apiTokenPlaceholder) {

		return match[1]
	}

	return ""
}
