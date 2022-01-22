package repo

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/AlecAivazis/survey/v2"
	"github.com/faetools/devtool/devtool"
	"github.com/faetools/devtool/internal/repo/library"
	"github.com/faetools/devtool/internal/repo/service"
	"github.com/spf13/cobra"
)

// DefinitionFileMissing returns whether or not the definition is missing.
func DefinitionFileMissing() bool {
	_, err := os.Stat(infoName())
	return os.IsNotExist(err)
}

// AskUser asks the user to fill the repo definition.
func AskUser() *Definition {
	wd, err := os.Getwd()
	cobra.CheckErr(err)

	cache = &Definition{
		Name:           filepath.Base(wd),
		Version:        "0.0.1",
		DevToolVersion: devtool.Version,
	}

	repoType := ""
	cobra.CheckErr(survey.AskOne(&survey.Select{
		Message: "What type of repository are you creating?",
		Options: allTypes,
	}, &repoType))
	cache.RepoType = _TypeNameToValue[repoType]

	switch _TypeNameToValue[repoType] {
	case Microservice:
		// ask about service definition
		cache.Service = service.AskForDefinition()
	case GoLibrary:
		cache.Library = library.AskForDefinition()
	default:
		fmt.Println(repoType, "not implemented yet")
	}

	cobra.CheckErr(survey.AskOne(&survey.Select{
		Message: "Which squad will be owner of this service?",
		Options: allTeams,
	}, &cache.Team))

	return cache
}
