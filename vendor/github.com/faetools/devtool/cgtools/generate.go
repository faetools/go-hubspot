package cgtools

import (
	"bytes"
	"flag"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"text/template"

	"github.com/deepmap/oapi-codegen/pkg/codegen"
	"github.com/faetools/devtool/format"
	"github.com/faetools/kit/terminal"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/logrusorgru/aurora"
	"github.com/pkg/errors"
	"github.com/spf13/afero"
	"github.com/spf13/viper"
)

// Generator is a file system that can generate files.
type Generator struct{ afero.Fs }

// WriteBytes writes the bytes to the given path, creating any directories and overwriting existing files.
func (g Generator) WriteBytes(path string, content []byte, mode fs.FileMode) (err error) {
	// Format contents according to type.
	content, err = format.Format(path, content)
	if err != nil {
		return errors.Wrapf(err, "formatting %s", path)
	}

	// Check if we need to write.
	current, readErr := afero.ReadFile(g.Fs, path)

	if bytes.Equal(content, current) {
		if viper.GetBool("verbose") {
			terminal.Printf(aurora.Green, "  • %v is unchanged\n", path)
		}
		return nil
	}

	newFile := errors.Is(readErr, os.ErrNotExist)
	if newFile {
		if err := g.MkdirAll(filepath.Dir(path), os.ModePerm); err != nil {
			return errors.Wrapf(err, "making directories %s", filepath.Dir(path))
		}
	}

	f, err := g.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, mode)
	if err != nil {
		return errors.Wrapf(err, "opening or creating file %s", path)
	}
	defer f.Close()

	_, err = f.Write(content)

	if err == nil && flag.Lookup("test.v") == nil {
		if newFile {
			terminal.Printf(aurora.Green, "  • %v generated\n", path)
		} else {
			terminal.Printf(aurora.Green, "  • %v regenerated\n", path)
		}
	}

	return errors.Wrapf(err, "writing to file %s", path)
}

// WriteTemplate writes the template to the given path, creating any directories and overwriting existing files.
func (g Generator) WriteTemplate(path string,
	tpl *template.Template, data interface{}) (err error,
) {
	b := &bytes.Buffer{}
	if err = tpl.Execute(b, data); err != nil {
		return errors.Wrapf(err, "executing template %s", tpl.Name())
	}

	return g.WriteBytes(path, b.Bytes(), 0o644)
}

// WriteSwagger generates code according to specified options and a swagger file.
func (g Generator) WriteSwagger(path string,
	swagger *openapi3.T, packageName string, opts codegen.Options,
) error {
	opts.SkipFmt = true // We do that in WriteBytes anyway.

	code, err := codegen.Generate(swagger, packageName, opts)
	if err != nil {
		return fmt.Errorf("error generating code: %w", err)
	}

	return g.WriteBytes(path, []byte(code), 0o644)
}

