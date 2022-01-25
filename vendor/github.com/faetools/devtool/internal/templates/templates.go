package templates

import (
	"embed"
	"fmt"
	"text/template"

	"github.com/deepmap/oapi-codegen/pkg/codegen"
	"github.com/spf13/cobra"
)

//go:embed *.tpl
var all embed.FS

var (
	CircleCI       = parse("CircleCI")
	DestroyService = parse("DestroyService")
	Dockerfile     = parse("Dockerfile")
	ReadMe         = parse("ReadMe")
	Badges         = parse("Badges")
	PreCommitHook  = getString("PreCommitHook")
	GolangciLint   = parse("GolangciLint")
	Versions       = parse("Versions")
	Main           = parse("Main")
	ServiceMain    = parse("ServiceMain")

	// GitHub settings.
	Gitignore        = parse("Gitignore")
	GithubSettings   = parse("GithubSettings")
	Gitattributes    = parse("Gitattributes")
	DependabotConfig = parse("DependabotConfig")

	// For HTTP service:.

	OpenAPISample       = parse("OpenAPISample")
	EndpointsSample     = parse("EndpointsSample")
	Imports             = getString("Imports")
	Client              = getString("Client")
	ClientWithResponses = getString("ClientWithResponses")
	Severs              = parse("Severs")

	EndpointsTest                       = getString("EndpointsTest")
	EndpointsTestServerInterfaceWrapper = getString("EndpointsTestServerInterfaceWrapper")
	EndpointsTestHelpers                = getString("EndpointsTestHelpers")
	EndpointsTestSetup                  = getString("EndpointsTestSetup")
)

func getBytes(name string) []byte {
	b, err := all.ReadFile(fmt.Sprintf("%s.tpl", name))
	cobra.CheckErr(err)
	return b
}

func getString(name string) string { return string(getBytes(name)) }

func parse(name string) *template.Template {
	return template.Must(template.New(name).
		Funcs(codegen.TemplateFunctions).Parse(getString(name)))
}
