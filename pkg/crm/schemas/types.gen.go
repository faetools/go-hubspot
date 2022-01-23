// Package schemas provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/faetools/devtool version (devel) DO NOT EDIT.
package schemas

import (
	"encoding/json"
	"fmt"
	"time"
)

const (
	HapikeyScopes       = "hapikey.Scopes"
	Oauth2_legacyScopes = "oauth2_legacy.Scopes"
)

// Defines values for ObjectTypePropertyCreateType.
const (
	ObjectTypePropertyCreateTypeBool ObjectTypePropertyCreateType = "bool"

	ObjectTypePropertyCreateTypeDate ObjectTypePropertyCreateType = "date"

	ObjectTypePropertyCreateTypeDatetime ObjectTypePropertyCreateType = "datetime"

	ObjectTypePropertyCreateTypeEnumeration ObjectTypePropertyCreateType = "enumeration"

	ObjectTypePropertyCreateTypeNumber ObjectTypePropertyCreateType = "number"

	ObjectTypePropertyCreateTypeString ObjectTypePropertyCreateType = "string"
)

// Defines an association between two object types.
type AssociationDefinition struct {
	// When the association was defined.
	CreatedAt *time.Time `json:"createdAt,omitempty"`

	// ID of the primary object type to link from.
	FromObjectTypeId string `json:"fromObjectTypeId"`

	// A unique ID for this association.
	Id string `json:"id"`

	// A unique name for this association.
	Name *string `json:"name,omitempty"`

	// ID of the target object type ID to link to.
	ToObjectTypeId string `json:"toObjectTypeId"`

	// When the association was last updated.
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
}

// Defines an association between two object types.
type AssociationDefinitionEgg struct {
	// ID of the primary object type to link from.
	FromObjectTypeId string `json:"fromObjectTypeId"`

	// A unique name for this association.
	Name *string `json:"name,omitempty"`

	// ID of the target object type ID to link to.
	ToObjectTypeId string `json:"toObjectTypeId"`
}

// CollectionResponseObjectSchemaNoPaging defines model for CollectionResponseObjectSchemaNoPaging.
type CollectionResponseObjectSchemaNoPaging struct {
	Results []ObjectSchema `json:"results"`
}

// Error defines model for Error.
type Error struct {
	// The error category
	Category string `json:"category"`

	// Context about the error condition
	Context *Error_Context `json:"context,omitempty"`

	// A unique identifier for the request. Include this value with any error reports or support tickets
	CorrelationId string `json:"correlationId"`

	// further information about the error
	Errors *[]ErrorDetail `json:"errors,omitempty"`

	// A map of link names to associated URIs containing documentation about the error or recommended remediation steps
	Links *Error_Links `json:"links,omitempty"`

	// A human readable message describing the error along with remediation steps where appropriate
	Message string `json:"message"`

	// A specific category that contains more specific detail about the error
	SubCategory *string `json:"subCategory,omitempty"`
}

// Context about the error condition
type Error_Context struct {
	AdditionalProperties map[string][]string `json:"-"`
}

// A map of link names to associated URIs containing documentation about the error or recommended remediation steps
type Error_Links struct {
	AdditionalProperties map[string]string `json:"-"`
}

// ErrorDetail defines model for ErrorDetail.
type ErrorDetail struct {
	// The status code associated with the error detail
	Code *string `json:"code,omitempty"`

	// Context about the error condition
	Context *ErrorDetail_Context `json:"context,omitempty"`

	// The name of the field or parameter in which the error was found.
	In *string `json:"in,omitempty"`

	// A human readable message describing the error along with remediation steps where appropriate
	Message string `json:"message"`

	// A specific category that contains more specific detail about the error
	SubCategory *string `json:"subCategory,omitempty"`
}

// Context about the error condition
type ErrorDetail_Context struct {
	AdditionalProperties map[string][]string `json:"-"`
}

// Defines an object schema, including its properties and associations.
type ObjectSchema struct {
	Archived bool `json:"archived"`

	// Associations defined for a given object type.
	Associations []AssociationDefinition `json:"associations"`

	// When the object schema was created.
	CreatedAt *time.Time `json:"createdAt,omitempty"`

	// An assigned unique ID for the object, including portal ID and object name.
	FullyQualifiedName string `json:"fullyQualifiedName"`

	// A unique ID for this schema's object type. Will be defined as {meta-type}-{unique ID}.
	Id string `json:"id"`

	// Singular and plural labels for the object. Used in CRM display.
	Labels ObjectTypeDefinitionLabels `json:"labels"`

	// A unique name for the schema's object type.
	Name         string `json:"name"`
	ObjectTypeId string `json:"objectTypeId"`

	// The name of the primary property for this object. This will be displayed as primary on the HubSpot record page for this object type.
	PrimaryDisplayProperty *string `json:"primaryDisplayProperty,omitempty"`

	// Properties defined for this object type.
	Properties []Property `json:"properties"`

	// The names of properties that should be **required** when creating an object of this type.
	RequiredProperties []string `json:"requiredProperties"`

	// Names of properties that will be indexed for this object type in by HubSpot's product search.
	SearchableProperties []string `json:"searchableProperties"`

	// The names of secondary properties for this object. These will be displayed as secondary on the HubSpot record page for this object type.
	SecondaryDisplayProperties []string `json:"secondaryDisplayProperties"`

	// When the object schema was last updated.
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
}

// Defines a new object type, its properties, and associations.
type ObjectSchemaEgg struct {
	// Associations defined for this object type.
	AssociatedObjects []string `json:"associatedObjects"`

	// Singular and plural labels for the object. Used in CRM display.
	Labels ObjectTypeDefinitionLabels `json:"labels"`

	// A unique name for this object. For internal use only.
	Name string `json:"name"`

	// The name of the primary property for this object. This will be displayed as primary on the HubSpot record page for this object type.
	PrimaryDisplayProperty *string `json:"primaryDisplayProperty,omitempty"`

	// Properties defined for this object type.
	Properties []ObjectTypePropertyCreate `json:"properties"`

	// The names of properties that should be **required** when creating an object of this type.
	RequiredProperties []string `json:"requiredProperties"`

	// Names of properties that will be indexed for this object type in by HubSpot's product search.
	SearchableProperties []string `json:"searchableProperties"`

	// The names of secondary properties for this object. These will be displayed as secondary on the HubSpot record page for this object type.
	SecondaryDisplayProperties []string `json:"secondaryDisplayProperties"`
}

// Defines an object type.
type ObjectTypeDefinition struct {
	Archived bool `json:"archived"`

	// When the object type was created.
	CreatedAt          *time.Time `json:"createdAt,omitempty"`
	FullyQualifiedName string     `json:"fullyQualifiedName"`

	// A unique ID for this object type. Will be defined as {meta-type}-{unique ID}.
	Id string `json:"id"`

	// Singular and plural labels for the object. Used in CRM display.
	Labels ObjectTypeDefinitionLabels `json:"labels"`

	// A unique name for this object. For internal use only.
	Name         string `json:"name"`
	ObjectTypeId string `json:"objectTypeId"`

	// The ID of the account that this object type is specific to.
	PortalId *int32 `json:"portalId,omitempty"`

	// The name of the primary property for this object. This will be displayed as primary on the HubSpot record page for this object type.
	PrimaryDisplayProperty *string `json:"primaryDisplayProperty,omitempty"`

	// The names of properties that should be **required** when creating an object of this type.
	RequiredProperties []string `json:"requiredProperties"`

	// Names of properties that will be indexed for this object type in by HubSpot's product search.
	SearchableProperties []string `json:"searchableProperties"`

	// The names of secondary properties for this object. These will be displayed as secondary on the HubSpot record page for this object type.
	SecondaryDisplayProperties []string `json:"secondaryDisplayProperties"`

	// When the object type was last updated.
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
}

// Singular and plural labels for the object. Used in CRM display.
type ObjectTypeDefinitionLabels struct {
	// The word for multiple objects. (There’s no way to change this later.)
	Plural *string `json:"plural,omitempty"`

	// The word for one object. (There’s no way to change this later.)
	Singular *string `json:"singular,omitempty"`
}

// Defines attributes to update on an object type.
type ObjectTypeDefinitionPatch struct {
	// Singular and plural labels for the object. Used in CRM display.
	Labels *ObjectTypeDefinitionLabels `json:"labels,omitempty"`

	// The name of the primary property for this object. This will be displayed as primary on the HubSpot record page for this object type.
	PrimaryDisplayProperty *string `json:"primaryDisplayProperty,omitempty"`

	// The names of properties that should be **required** when creating an object of this type.
	RequiredProperties *[]string `json:"requiredProperties,omitempty"`
	Restorable         *bool     `json:"restorable,omitempty"`

	// Names of properties that will be indexed for this object type in by HubSpot's product search.
	SearchableProperties *[]string `json:"searchableProperties,omitempty"`

	// The names of secondary properties for this object. These will be displayed as secondary on the HubSpot record page for this object type.
	SecondaryDisplayProperties *[]string `json:"secondaryDisplayProperties,omitempty"`
}

// Defines a property to create.
type ObjectTypePropertyCreate struct {
	// A description of the property that will be shown as help text in HubSpot.
	Description *string `json:"description,omitempty"`

	// The order that this property should be displayed in the HubSpot UI relative to other properties for this object type. Properties are displayed in order starting with the lowest positive integer value. A value of -1 will cause the property to be displayed **after** any positive values.
	DisplayOrder *int32 `json:"displayOrder,omitempty"`

	// Controls how the property appears in HubSpot.
	FieldType string `json:"fieldType"`

	// The name of the group this property belongs to.
	GroupName *string `json:"groupName,omitempty"`

	// Whether or not the property's value must be unique. Once set, this can't be changed.
	HasUniqueValue *bool `json:"hasUniqueValue,omitempty"`
	Hidden         *bool `json:"hidden,omitempty"`

	// A human-readable property label that will be shown in HubSpot.
	Label string `json:"label"`

	// The internal property name, which must be used when referencing the property from the API.
	Name string `json:"name"`

	// A list of available options for the property. This field is only required for enumerated properties.
	Options *[]OptionInput `json:"options,omitempty"`

	// The data type of the property.
	Type ObjectTypePropertyCreateType `json:"type"`
}

// The data type of the property.
type ObjectTypePropertyCreateType string

// The options available when a property is an enumeration
type Option struct {
	// A description of the option.
	Description *string `json:"description,omitempty"`

	// Options are displayed in order starting with the lowest positive integer value. Values of -1 will cause the option to be displayed after any positive values.
	DisplayOrder *int32 `json:"displayOrder,omitempty"`

	// Hidden options will not be displayed in HubSpot.
	Hidden bool `json:"hidden"`

	// A human-readable option label that will be shown in HubSpot.
	Label string `json:"label"`

	// The internal value of the option, which must be used when setting the property value through the API.
	Value string `json:"value"`
}

// Defines a enumeration property option
type OptionInput struct {
	// A description of the option.
	Description *string `json:"description,omitempty"`

	// Options are shown in order starting with the lowest positive integer value. Values of -1 will cause the option to be displayed after any positive values.
	DisplayOrder int32 `json:"displayOrder"`

	// Hidden options won't be shown in HubSpot.
	Hidden bool `json:"hidden"`

	// A human-readable option label that will be shown in HubSpot.
	Label string `json:"label"`

	// The internal value of the option, which must be used when setting the property value through the API.
	Value string `json:"value"`
}

// Defines a property
type Property struct {
	// Whether or not the property is archived.
	Archived *bool `json:"archived,omitempty"`

	// When the property was archived.
	ArchivedAt *time.Time `json:"archivedAt,omitempty"`

	// For default properties, true indicates that the property is calculated by a HubSpot process. It has no effect for custom properties.
	Calculated *bool `json:"calculated,omitempty"`

	// When the property was created
	CreatedAt *time.Time `json:"createdAt,omitempty"`

	// The internal ID of the user who created the property in HubSpot. This field may not exist if the property was created outside of HubSpot.
	CreatedUserId *string `json:"createdUserId,omitempty"`

	// A description of the property that will be shown as help text in HubSpot.
	Description string `json:"description"`

	// The order that this property should be displayed in the HubSpot UI relative to other properties for this object type. Properties are displayed in order starting with the lowest positive integer value. A value of -1 will cause the property to be displayed **after** any positive values.
	DisplayOrder *int32 `json:"displayOrder,omitempty"`

	// For default properties, true indicates that the options are stored externally to the property settings.
	ExternalOptions *bool `json:"externalOptions,omitempty"`

	// Controls how the property appears in HubSpot.
	FieldType string `json:"fieldType"`

	// Whether or not the property can be used in a HubSpot form.
	FormField *bool `json:"formField,omitempty"`

	// The name of the property group the property belongs to.
	GroupName string `json:"groupName"`

	// Whether or not the property's value must be unique. Once set, this can't be changed.
	HasUniqueValue *bool `json:"hasUniqueValue,omitempty"`
	Hidden         *bool `json:"hidden,omitempty"`

	// This will be true for default object properties built into HubSpot.
	HubspotDefined *bool `json:"hubspotDefined,omitempty"`

	// A human-readable property label that will be shown in HubSpot.
	Label                string                        `json:"label"`
	ModificationMetadata *PropertyModificationMetadata `json:"modificationMetadata,omitempty"`

	// The internal property name, which must be used when referencing the property via the API.
	Name string `json:"name"`

	// A list of valid options for the property. This field is required for enumerated properties, but will be empty for other property types.
	Options []Option `json:"options"`

	// If this property is related to other object(s), they'll be listed here.
	ReferencedObjectType *string `json:"referencedObjectType,omitempty"`

	// Whether the property will display the currency symbol set in the account settings.
	ShowCurrencySymbol *bool `json:"showCurrencySymbol,omitempty"`

	// The property data type.
	Type      string     `json:"type"`
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`

	// The internal user ID of the user who updated the property in HubSpot. This field may not exist if the property was updated outside of HubSpot.
	UpdatedUserId *string `json:"updatedUserId,omitempty"`
}

// PropertyModificationMetadata defines model for PropertyModificationMetadata.
type PropertyModificationMetadata struct {
	Archivable         bool  `json:"archivable"`
	ReadOnlyDefinition bool  `json:"readOnlyDefinition"`
	ReadOnlyOptions    *bool `json:"readOnlyOptions,omitempty"`
	ReadOnlyValue      bool  `json:"readOnlyValue"`
}

// GetAllSchemasParams defines parameters for GetAllSchemas.
type GetAllSchemasParams struct {
	// Whether to return only results that have been archived.
	Archived *bool `json:"archived,omitempty"`
}

// CreateSchemasJSONBody defines parameters for CreateSchemas.
type CreateSchemasJSONBody ObjectSchemaEgg

// ArchiveObjectTypeParams defines parameters for ArchiveObjectType.
type ArchiveObjectTypeParams struct {
	// Whether to return only results that have been archived.
	Archived *bool `json:"archived,omitempty"`
}

// UpdateObjectTypeJSONBody defines parameters for UpdateObjectType.
type UpdateObjectTypeJSONBody ObjectTypeDefinitionPatch

// CreateAssociationAssociationsJSONBody defines parameters for CreateAssociationAssociations.
type CreateAssociationAssociationsJSONBody AssociationDefinitionEgg

// CreateSchemasJSONRequestBody defines body for CreateSchemas for application/json ContentType.
type CreateSchemasJSONRequestBody CreateSchemasJSONBody

// UpdateObjectTypeJSONRequestBody defines body for UpdateObjectType for application/json ContentType.
type UpdateObjectTypeJSONRequestBody UpdateObjectTypeJSONBody

// CreateAssociationAssociationsJSONRequestBody defines body for CreateAssociationAssociations for application/json ContentType.
type CreateAssociationAssociationsJSONRequestBody CreateAssociationAssociationsJSONBody

// Getter for additional properties for Error_Context. Returns the specified
// element and whether it was found
func (a Error_Context) Get(fieldName string) (value []string, found bool) {
	if a.AdditionalProperties != nil {
		value, found = a.AdditionalProperties[fieldName]
	}
	return
}

// Setter for additional properties for Error_Context
func (a *Error_Context) Set(fieldName string, value []string) {
	if a.AdditionalProperties == nil {
		a.AdditionalProperties = make(map[string][]string)
	}
	a.AdditionalProperties[fieldName] = value
}

// Override default JSON handling for Error_Context to handle AdditionalProperties
func (a *Error_Context) UnmarshalJSON(b []byte) error {
	object := make(map[string]json.RawMessage)
	err := json.Unmarshal(b, &object)
	if err != nil {
		return err
	}

	if len(object) != 0 {
		a.AdditionalProperties = make(map[string][]string)
		for fieldName, fieldBuf := range object {
			var fieldVal []string
			err := json.Unmarshal(fieldBuf, &fieldVal)
			if err != nil {
				return fmt.Errorf("error unmarshaling field %s: %w", fieldName, err)
			}
			a.AdditionalProperties[fieldName] = fieldVal
		}
	}
	return nil
}

// Override default JSON handling for Error_Context to handle AdditionalProperties
func (a Error_Context) MarshalJSON() ([]byte, error) {
	var err error
	object := make(map[string]json.RawMessage)

	for fieldName, field := range a.AdditionalProperties {
		object[fieldName], err = json.Marshal(field)
		if err != nil {
			return nil, fmt.Errorf("error marshaling '%s': %w", fieldName, err)
		}
	}
	return json.Marshal(object)
}

// Getter for additional properties for Error_Links. Returns the specified
// element and whether it was found
func (a Error_Links) Get(fieldName string) (value string, found bool) {
	if a.AdditionalProperties != nil {
		value, found = a.AdditionalProperties[fieldName]
	}
	return
}

// Setter for additional properties for Error_Links
func (a *Error_Links) Set(fieldName string, value string) {
	if a.AdditionalProperties == nil {
		a.AdditionalProperties = make(map[string]string)
	}
	a.AdditionalProperties[fieldName] = value
}

// Override default JSON handling for Error_Links to handle AdditionalProperties
func (a *Error_Links) UnmarshalJSON(b []byte) error {
	object := make(map[string]json.RawMessage)
	err := json.Unmarshal(b, &object)
	if err != nil {
		return err
	}

	if len(object) != 0 {
		a.AdditionalProperties = make(map[string]string)
		for fieldName, fieldBuf := range object {
			var fieldVal string
			err := json.Unmarshal(fieldBuf, &fieldVal)
			if err != nil {
				return fmt.Errorf("error unmarshaling field %s: %w", fieldName, err)
			}
			a.AdditionalProperties[fieldName] = fieldVal
		}
	}
	return nil
}

// Override default JSON handling for Error_Links to handle AdditionalProperties
func (a Error_Links) MarshalJSON() ([]byte, error) {
	var err error
	object := make(map[string]json.RawMessage)

	for fieldName, field := range a.AdditionalProperties {
		object[fieldName], err = json.Marshal(field)
		if err != nil {
			return nil, fmt.Errorf("error marshaling '%s': %w", fieldName, err)
		}
	}
	return json.Marshal(object)
}

// Getter for additional properties for ErrorDetail_Context. Returns the specified
// element and whether it was found
func (a ErrorDetail_Context) Get(fieldName string) (value []string, found bool) {
	if a.AdditionalProperties != nil {
		value, found = a.AdditionalProperties[fieldName]
	}
	return
}

// Setter for additional properties for ErrorDetail_Context
func (a *ErrorDetail_Context) Set(fieldName string, value []string) {
	if a.AdditionalProperties == nil {
		a.AdditionalProperties = make(map[string][]string)
	}
	a.AdditionalProperties[fieldName] = value
}

// Override default JSON handling for ErrorDetail_Context to handle AdditionalProperties
func (a *ErrorDetail_Context) UnmarshalJSON(b []byte) error {
	object := make(map[string]json.RawMessage)
	err := json.Unmarshal(b, &object)
	if err != nil {
		return err
	}

	if len(object) != 0 {
		a.AdditionalProperties = make(map[string][]string)
		for fieldName, fieldBuf := range object {
			var fieldVal []string
			err := json.Unmarshal(fieldBuf, &fieldVal)
			if err != nil {
				return fmt.Errorf("error unmarshaling field %s: %w", fieldName, err)
			}
			a.AdditionalProperties[fieldName] = fieldVal
		}
	}
	return nil
}

// Override default JSON handling for ErrorDetail_Context to handle AdditionalProperties
func (a ErrorDetail_Context) MarshalJSON() ([]byte, error) {
	var err error
	object := make(map[string]json.RawMessage)

	for fieldName, field := range a.AdditionalProperties {
		object[fieldName], err = json.Marshal(field)
		if err != nil {
			return nil, fmt.Errorf("error marshaling '%s': %w", fieldName, err)
		}
	}
	return json.Marshal(object)
}