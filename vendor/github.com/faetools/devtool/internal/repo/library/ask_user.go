package library

import (
	"github.com/AlecAivazis/survey/v2"
	"github.com/spf13/cobra"
)

// AskForDefinition asks the user to fill the service definition.
func AskForDefinition() *Definition {
	def := &Definition{}

	tp := ""
	cobra.CheckErr(survey.AskOne(&survey.Select{
		Message: "To which type does your library belong?",
		Options: allTypes,
	}, &tp))
	def.SubType = _SubTypeNameToValue[tp]

	return def
}
