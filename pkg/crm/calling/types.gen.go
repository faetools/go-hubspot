// Package calling provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/faetools/devtool version (devel) DO NOT EDIT.
package calling

import (
	"encoding/json"
	"fmt"
	"time"
)

const (
	Developer_hapikeyScopes = "developer_hapikey.Scopes"
)

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

// Settings update request
type SettingsPatchRequest struct {
	// The target height of the iframe that will contain your phone/calling UI.
	Height *int32 `json:"height,omitempty"`

	// When true, your service will appear as an option under the *Call* action in contact records of connected accounts.
	IsReady *bool `json:"isReady,omitempty"`

	// The name of your calling service to display to users.
	Name *string `json:"name,omitempty"`

	// When true, you are indicating that your service is compatible with engagement v2 service and can be used with custom objects.
	SupportsCustomObjects *bool `json:"supportsCustomObjects,omitempty"`

	// The URL to your phone/calling UI, built with the [Calling SDK](#).
	Url *string `json:"url,omitempty"`

	// The target width of the iframe that will contain your phone/calling UI.
	Width *int32 `json:"width,omitempty"`
}

// Settings create request
type SettingsRequest struct {
	// The target height of the iframe that will contain your phone/calling UI.
	Height *int32 `json:"height,omitempty"`

	// When true, your service will appear as an option under the *Call* action in contact records of connected accounts.
	IsReady *bool `json:"isReady,omitempty"`

	// The name of your calling service to display to users.
	Name string `json:"name"`

	// When true, you are indicating that your service is compatible with engagement v2 service and can be used with custom objects.
	SupportsCustomObjects *bool `json:"supportsCustomObjects,omitempty"`

	// The URL to your phone/calling UI, built with the [Calling SDK](#).
	Url string `json:"url"`

	// The target width of the iframe that will contain your phone/calling UI.
	Width *int32 `json:"width,omitempty"`
}

// Current settings state
type SettingsResponse struct {
	// When this calling extension was created.
	CreatedAt time.Time `json:"createdAt"`

	// The target height of the iframe that will contain your phone/calling UI.
	Height int32 `json:"height"`

	// When true, your service will appear as an option under the *Call* action in contact records of connected accounts.
	IsReady bool `json:"isReady"`

	// The name of your calling service to display to users.
	Name string `json:"name"`

	// When true, you are indicating that your service is compatible with engagement v2 service and can be used with custom objects.
	SupportsCustomObjects bool `json:"supportsCustomObjects"`

	// The last time the settings for this calling extension were modified.
	UpdatedAt time.Time `json:"updatedAt"`

	// The URL to your phone/calling UI, built with the [Calling SDK](#).
	Url string `json:"url"`

	// The target width of the iframe that will contain your phone/calling UI.
	Width int32 `json:"width"`
}

// UpdateSettingsJSONBody defines parameters for UpdateSettings.
type UpdateSettingsJSONBody SettingsPatchRequest

// CreateSettingsJSONBody defines parameters for CreateSettings.
type CreateSettingsJSONBody SettingsRequest

// UpdateSettingsJSONRequestBody defines body for UpdateSettings for application/json ContentType.
type UpdateSettingsJSONRequestBody UpdateSettingsJSONBody

// CreateSettingsJSONRequestBody defines body for CreateSettings for application/json ContentType.
type CreateSettingsJSONRequestBody CreateSettingsJSONBody

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
