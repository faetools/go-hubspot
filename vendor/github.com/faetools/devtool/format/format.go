package format

import (
	"bytes"
	"fmt"
	"path/filepath"
	"strings"

	"github.com/faetools/devtool/format/golang"
	"github.com/faetools/devtool/format/markdown"
	"github.com/faetools/devtool/format/yaml"
	"github.com/faetools/kit/terminal"
	"github.com/logrusorgru/aurora"
	dockerfile "github.com/moby/buildkit/frontend/dockerfile/parser"
	"github.com/pkg/errors"
	json "github.com/tidwall/pretty"
)

type testingT int

func (t testingT) Errorf(format string, args ...interface{}) {
	fmt.Printf(format, args...)
}

// Format formats contents according to type.
func Format(path string, src []byte) ([]byte, error) {
	var err error
	switch filepath.Ext(path) {
	case ".go":
		src, err = golang.Format(path, src)
	case ".yml", ".yaml":
		src, err = yaml.Format(src)
	case ".md":
		src, err = markdown.Format(src)
	case ".json":
		src = json.PrettyOptions(src, &json.Options{Indent: "  "})
	case "":
		if filepath.Base(path) == "Dockerfile" {
			docker, err := dockerfile.Parse(bytes.NewReader(src))
			if err != nil {
				return nil, errors.Wrap(err, "parsing dockerfile")
			}

			if len(docker.Warnings) != 0 {
				terminal.Printf(aurora.Red, "Dockerfile warnings:\n  • %s\n",
					strings.Join(docker.Warnings, "\n  • "))
			}
		}
	}

	return src, errors.Wrapf(err, "formatting %s", path)
}
