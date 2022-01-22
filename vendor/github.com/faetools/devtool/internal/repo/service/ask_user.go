package service

import (
	"github.com/AlecAivazis/survey/v2"
	"github.com/spf13/cobra"
)

// AskForDefinition asks the user to fill the service definition.
func AskForDefinition() *Definition {
	def := &Definition{
		// HTTP enabled by default for now
		HTTPS: &HTTPS{Enabled: true},
	}

	cobra.CheckErr(survey.AskOne(&survey.Select{
		Message: "To which tier does your service belong?",
		Options: allTiers,
	}, &def.Tier))

	return def
}
