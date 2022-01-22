package golang

import (
	"bytes"
	"go/ast"
	"go/parser"
	"go/token"
	"regexp"

	goformat "go/format"

	"github.com/faetools/devtool/internal/version"
	"github.com/pkg/errors"
	"golang.org/x/tools/imports"
	"mvdan.cc/gofumpt/format"
)

var (
	FormatOptions = format.Options{LangVersion: version.Go}

	rxCodeGenerated = regexp.MustCompile(`^// Code generated .* DO NOT EDIT\.$`)
)

// Format formats golang code.
func Format(filepath string, src []byte) ([]byte, error) {
	res, err := imports.Process(filepath, src, nil)
	if err != nil {
		return nil, errors.Wrapf(err, "running 'imports'")
	}

	// code below taken from format.Source in order to
	// inject a line allowing formatting of generated files

	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, "", res, parser.ParseComments)
	if err != nil {
		return nil, err
	}

	// hack to allow formatting of generated files
	line, text := maskAutoGenerated(file.Comments, file.Package)

	format.File(fset, file, FormatOptions)

	if line != nil {
		line.Text = text // add comment back in
	}

	var buf bytes.Buffer
	if err := goformat.Node(&buf, fset, file); err != nil {
		return nil, errors.Wrapf(err, "running 'gofumpt'")
	}

	return buf.Bytes(), nil
}

func maskAutoGenerated(comments []*ast.CommentGroup, until token.Pos) (*ast.Comment, string) {
	for _, cg := range comments {
		if cg.Pos() > until {
			break
		}

		for _, line := range cg.List {
			if rxCodeGenerated.MatchString(line.Text) {
				text := line.Text
				line.Text = "masked comment"
				return line, text
			}
		}
	}

	return nil, ""
}
