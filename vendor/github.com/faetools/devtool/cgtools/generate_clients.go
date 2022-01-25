package cgtools

import (
	"path/filepath"

	"github.com/deepmap/oapi-codegen/pkg/codegen"
	"github.com/faetools/devtool/internal/templates"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/pkg/errors"
)

const (
	filenameServers = "servers.gen.go"
	filenameClient  = "client.gen.go"
	filenameTypes   = "types.gen.go"
)

// GenerateClientsPkg generates the clients package in the specified folder.
func (g Generator) GenerateClientsPkg(dir string, swagger *openapi3.T) error {
	pkgName := filepath.Base(dir)

	if err := g.WriteTemplate(
		filepath.Join(dir, filenameServers),
		templates.Severs, map[string]interface{}{
			"PackageName": pkgName,
			"Servers":     swagger.Servers,
		}); err != nil {
		return errors.Wrap(err, "generating servers")
	}

	// Generate Client.
	if err := g.WriteSwagger(
		filepath.Join(dir, filenameClient),
		swagger, pkgName, codegen.Options{
			GenerateClient: true,
			UserTemplates: map[string]string{
				"client.tmpl":                templates.Client,
				"client-with-responses.tmpl": templates.ClientWithResponses,
			},
		}); err != nil {
		return errors.Wrap(err, "generating client")
	}

	// Generate types.
	return errors.Wrap(g.WriteSwagger(
		filepath.Join(dir, filenameTypes),
		swagger, pkgName, codegen.Options{
			GenerateTypes: true,
		}), "generating types")
}
