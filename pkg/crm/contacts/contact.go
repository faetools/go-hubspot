package contacts

import (
	"strings"
)

const (
	keyFirstName        = "firstname"
	keyLastName         = "lastname"
	keyEmail            = "email"
	keyAdditionalEmails = "hs_additional_emails"

	// separator for multiple values
	sep = ";"
)

// PropertyKeys contains all the general property keys.
// Ask for at least these to be able to use the convenience functions.
var PropertyKeys = []string{
	keyFirstName,
	keyLastName,
	keyEmail,
	keyAdditionalEmails,
}

// FirstName returns the first name of the contact.
func (c SimplePublicObjectWithAssociations) FirstName() string {
	return c.Properties.AdditionalProperties[keyFirstName]
}

// LastName returns the last name of the contact.
func (c SimplePublicObjectWithAssociations) LastName() string {
	return c.Properties.AdditionalProperties[keyLastName]
}

// Email returns the email addresses of the contact.
func (c SimplePublicObjectWithAssociations) Emails() []string {
	primary := c.Properties.AdditionalProperties[keyEmail]
	if primary == "" {
		return nil
	}

	out := []string{primary}

	if add := c.Properties.AdditionalProperties[keyAdditionalEmails]; add != "" {
		out = append(out, strings.Split(add, sep)...)
	}

	return out
}
