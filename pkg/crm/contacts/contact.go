package contacts

const (
	keyFirstName = "firstname"
	keyLastName  = "lastname"
	keyEmail     = "email"
)

// FirstName returns the first name of the contact.
func (c SimplePublicObjectWithAssociations) FirstName() string {
	return c.Properties.AdditionalProperties[keyFirstName]
}

// LastName returns the last name of the contact.
func (c SimplePublicObjectWithAssociations) LastName() string {
	return c.Properties.AdditionalProperties[keyLastName]
}

// Email returns the email of the contact.
func (c SimplePublicObjectWithAssociations) Email() string {
	return c.Properties.AdditionalProperties[keyEmail]
}
