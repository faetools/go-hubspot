// Package timeline provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/faetools/devtool version (devel) DO NOT EDIT.
package timeline

import (
	"encoding/json"
	"fmt"
	"time"
)

const (
	Developer_hapikeyScopes = "developer_hapikey.Scopes"
	Oauth2_legacyScopes     = "oauth2_legacy.Scopes"
)

// Defines values for BatchResponseTimelineEventResponseStatus.
const (
	BatchResponseTimelineEventResponseStatusCANCELED BatchResponseTimelineEventResponseStatus = "CANCELED"

	BatchResponseTimelineEventResponseStatusCOMPLETE BatchResponseTimelineEventResponseStatus = "COMPLETE"

	BatchResponseTimelineEventResponseStatusPENDING BatchResponseTimelineEventResponseStatus = "PENDING"

	BatchResponseTimelineEventResponseStatusPROCESSING BatchResponseTimelineEventResponseStatus = "PROCESSING"
)

// Defines values for BatchResponseTimelineEventResponseWithErrorsStatus.
const (
	BatchResponseTimelineEventResponseWithErrorsStatusCANCELED BatchResponseTimelineEventResponseWithErrorsStatus = "CANCELED"

	BatchResponseTimelineEventResponseWithErrorsStatusCOMPLETE BatchResponseTimelineEventResponseWithErrorsStatus = "COMPLETE"

	BatchResponseTimelineEventResponseWithErrorsStatusPENDING BatchResponseTimelineEventResponseWithErrorsStatus = "PENDING"

	BatchResponseTimelineEventResponseWithErrorsStatusPROCESSING BatchResponseTimelineEventResponseWithErrorsStatus = "PROCESSING"
)

// Defines values for ErrorCategoryHttpStatus.
const (
	ErrorCategoryHttpStatusACCEPTED ErrorCategoryHttpStatus = "ACCEPTED"

	ErrorCategoryHttpStatusALREADYREPORTED ErrorCategoryHttpStatus = "ALREADY_REPORTED"

	ErrorCategoryHttpStatusBADGATEWAY ErrorCategoryHttpStatus = "BAD_GATEWAY"

	ErrorCategoryHttpStatusBADREQUEST ErrorCategoryHttpStatus = "BAD_REQUEST"

	ErrorCategoryHttpStatusCONFLICT ErrorCategoryHttpStatus = "CONFLICT"

	ErrorCategoryHttpStatusCONTINUE ErrorCategoryHttpStatus = "CONTINUE"

	ErrorCategoryHttpStatusCREATED ErrorCategoryHttpStatus = "CREATED"

	ErrorCategoryHttpStatusEXPECTATIONFAILED ErrorCategoryHttpStatus = "EXPECTATION_FAILED"

	ErrorCategoryHttpStatusFAILEDDEPENDENCY ErrorCategoryHttpStatus = "FAILED_DEPENDENCY"

	ErrorCategoryHttpStatusFORBIDDEN ErrorCategoryHttpStatus = "FORBIDDEN"

	ErrorCategoryHttpStatusFOUND ErrorCategoryHttpStatus = "FOUND"

	ErrorCategoryHttpStatusGATEWAYTIMEOUT ErrorCategoryHttpStatus = "GATEWAY_TIMEOUT"

	ErrorCategoryHttpStatusGONE ErrorCategoryHttpStatus = "GONE"

	ErrorCategoryHttpStatusHTTPVERSIONNOTSUPPORTED ErrorCategoryHttpStatus = "HTTP_VERSION_NOT_SUPPORTED"

	ErrorCategoryHttpStatusIMATEAPOT ErrorCategoryHttpStatus = "IM_A_TEAPOT"

	ErrorCategoryHttpStatusIMUSED ErrorCategoryHttpStatus = "IM_USED"

	ErrorCategoryHttpStatusINSUFFICIENTSTORAGE ErrorCategoryHttpStatus = "INSUFFICIENT_STORAGE"

	ErrorCategoryHttpStatusINTERNALSERVERERROR ErrorCategoryHttpStatus = "INTERNAL_SERVER_ERROR"

	ErrorCategoryHttpStatusINTERNALSTALESERVICEDISCOVERY ErrorCategoryHttpStatus = "INTERNAL_STALE_SERVICE_DISCOVERY"

	ErrorCategoryHttpStatusLENGTHREQUIRED ErrorCategoryHttpStatus = "LENGTH_REQUIRED"

	ErrorCategoryHttpStatusLOCKED ErrorCategoryHttpStatus = "LOCKED"

	ErrorCategoryHttpStatusLOOPDETECTED ErrorCategoryHttpStatus = "LOOP_DETECTED"

	ErrorCategoryHttpStatusMETHODNOTALLOWED ErrorCategoryHttpStatus = "METHOD_NOT_ALLOWED"

	ErrorCategoryHttpStatusMISDIRECTEDREQUEST ErrorCategoryHttpStatus = "MISDIRECTED_REQUEST"

	ErrorCategoryHttpStatusMOVEDPERMANENTLY ErrorCategoryHttpStatus = "MOVED_PERMANENTLY"

	ErrorCategoryHttpStatusMULTIPLECHOICES ErrorCategoryHttpStatus = "MULTIPLE_CHOICES"

	ErrorCategoryHttpStatusMULTISTATUS ErrorCategoryHttpStatus = "MULTI_STATUS"

	ErrorCategoryHttpStatusNETWORKAUTHENTICATIONREQUIRED ErrorCategoryHttpStatus = "NETWORK_AUTHENTICATION_REQUIRED"

	ErrorCategoryHttpStatusNOCONTENT ErrorCategoryHttpStatus = "NO_CONTENT"

	ErrorCategoryHttpStatusNONAUTHORITATIVEINFORMATION ErrorCategoryHttpStatus = "NON_AUTHORITATIVE_INFORMATION"

	ErrorCategoryHttpStatusNOTACCEPTABLE ErrorCategoryHttpStatus = "NOT_ACCEPTABLE"

	ErrorCategoryHttpStatusNOTEXTENDED ErrorCategoryHttpStatus = "NOT_EXTENDED"

	ErrorCategoryHttpStatusNOTFOUND ErrorCategoryHttpStatus = "NOT_FOUND"

	ErrorCategoryHttpStatusNOTIMPLEMENTED ErrorCategoryHttpStatus = "NOT_IMPLEMENTED"

	ErrorCategoryHttpStatusNOTMODIFIED ErrorCategoryHttpStatus = "NOT_MODIFIED"

	ErrorCategoryHttpStatusOK ErrorCategoryHttpStatus = "OK"

	ErrorCategoryHttpStatusPARTIALCONTENT ErrorCategoryHttpStatus = "PARTIAL_CONTENT"

	ErrorCategoryHttpStatusPAYMENTREQUIRED ErrorCategoryHttpStatus = "PAYMENT_REQUIRED"

	ErrorCategoryHttpStatusPERMANENTREDIRECT ErrorCategoryHttpStatus = "PERMANENT_REDIRECT"

	ErrorCategoryHttpStatusPRECONDITIONFAILED ErrorCategoryHttpStatus = "PRECONDITION_FAILED"

	ErrorCategoryHttpStatusPRECONDITIONREQUIRED ErrorCategoryHttpStatus = "PRECONDITION_REQUIRED"

	ErrorCategoryHttpStatusPROCESSING ErrorCategoryHttpStatus = "PROCESSING"

	ErrorCategoryHttpStatusPROXYAUTHENTICATIONREQUIRED ErrorCategoryHttpStatus = "PROXY_AUTHENTICATION_REQUIRED"

	ErrorCategoryHttpStatusREQUESTEDRANGENOTSATISFIABLE ErrorCategoryHttpStatus = "REQUESTED_RANGE_NOT_SATISFIABLE"

	ErrorCategoryHttpStatusREQUESTENTITYTOOLARGE ErrorCategoryHttpStatus = "REQUEST_ENTITY_TOO_LARGE"

	ErrorCategoryHttpStatusREQUESTHEADERSFIELDSTOOLARGE ErrorCategoryHttpStatus = "REQUEST_HEADERS_FIELDS_TOO_LARGE"

	ErrorCategoryHttpStatusREQUESTTIMEOUT ErrorCategoryHttpStatus = "REQUEST_TIMEOUT"

	ErrorCategoryHttpStatusREQUESTURITOOLONG ErrorCategoryHttpStatus = "REQUEST_URI_TOO_LONG"

	ErrorCategoryHttpStatusRESETCONTENT ErrorCategoryHttpStatus = "RESET_CONTENT"

	ErrorCategoryHttpStatusSEEOTHER ErrorCategoryHttpStatus = "SEE_OTHER"

	ErrorCategoryHttpStatusSERVICEUNAVAILABLE ErrorCategoryHttpStatus = "SERVICE_UNAVAILABLE"

	ErrorCategoryHttpStatusSWITCHINGPROTOCOLS ErrorCategoryHttpStatus = "SWITCHING_PROTOCOLS"

	ErrorCategoryHttpStatusTEMPORARYREDIRECT ErrorCategoryHttpStatus = "TEMPORARY_REDIRECT"

	ErrorCategoryHttpStatusTOOMANYREQUESTS ErrorCategoryHttpStatus = "TOO_MANY_REQUESTS"

	ErrorCategoryHttpStatusUNAUTHORIZED ErrorCategoryHttpStatus = "UNAUTHORIZED"

	ErrorCategoryHttpStatusUNAVAILABLEFORLEGALREASONS ErrorCategoryHttpStatus = "UNAVAILABLE_FOR_LEGAL_REASONS"

	ErrorCategoryHttpStatusUNPROCESSABLEENTITY ErrorCategoryHttpStatus = "UNPROCESSABLE_ENTITY"

	ErrorCategoryHttpStatusUNSUPPORTEDMEDIATYPE ErrorCategoryHttpStatus = "UNSUPPORTED_MEDIA_TYPE"

	ErrorCategoryHttpStatusUPGRADEREQUIRED ErrorCategoryHttpStatus = "UPGRADE_REQUIRED"

	ErrorCategoryHttpStatusUSEPROXY ErrorCategoryHttpStatus = "USE_PROXY"

	ErrorCategoryHttpStatusVARIANTALSONEGOTIATES ErrorCategoryHttpStatus = "VARIANT_ALSO_NEGOTIATES"
)

// Defines values for TimelineEventTemplateTokenType.
const (
	TimelineEventTemplateTokenTypeDate TimelineEventTemplateTokenType = "date"

	TimelineEventTemplateTokenTypeEnumeration TimelineEventTemplateTokenType = "enumeration"

	TimelineEventTemplateTokenTypeNumber TimelineEventTemplateTokenType = "number"

	TimelineEventTemplateTokenTypeString TimelineEventTemplateTokenType = "string"
)

// Used to create timeline events in batches.
type BatchInputTimelineEvent struct {
	// A collection of timeline events we want to create.
	Inputs []TimelineEvent `json:"inputs"`
}

// The state of the batch event request.
type BatchResponseTimelineEventResponse struct {
	// The time the request was completed.
	CompletedAt time.Time                                 `json:"completedAt"`
	Links       *BatchResponseTimelineEventResponse_Links `json:"links,omitempty"`

	// The time the request occurred.
	RequestedAt *time.Time `json:"requestedAt,omitempty"`

	// Successfully created events.
	Results []TimelineEventResponse `json:"results"`

	// The time the request began processing.
	StartedAt time.Time `json:"startedAt"`

	// The status of the batch response. Should always be COMPLETED if processed.
	Status BatchResponseTimelineEventResponseStatus `json:"status"`
}

// BatchResponseTimelineEventResponse_Links defines model for BatchResponseTimelineEventResponse.Links.
type BatchResponseTimelineEventResponse_Links struct {
	AdditionalProperties map[string]string `json:"-"`
}

// The status of the batch response. Should always be COMPLETED if processed.
type BatchResponseTimelineEventResponseStatus string

// BatchResponseTimelineEventResponseWithErrors defines model for BatchResponseTimelineEventResponseWithErrors.
type BatchResponseTimelineEventResponseWithErrors struct {
	CompletedAt time.Time                                           `json:"completedAt"`
	Errors      *[]StandardError                                    `json:"errors,omitempty"`
	Links       *BatchResponseTimelineEventResponseWithErrors_Links `json:"links,omitempty"`
	NumErrors   *int32                                              `json:"numErrors,omitempty"`
	RequestedAt *time.Time                                          `json:"requestedAt,omitempty"`
	Results     []TimelineEventResponse                             `json:"results"`
	StartedAt   time.Time                                           `json:"startedAt"`
	Status      BatchResponseTimelineEventResponseWithErrorsStatus  `json:"status"`
}

// BatchResponseTimelineEventResponseWithErrors_Links defines model for BatchResponseTimelineEventResponseWithErrors.Links.
type BatchResponseTimelineEventResponseWithErrors_Links struct {
	AdditionalProperties map[string]string `json:"-"`
}

// BatchResponseTimelineEventResponseWithErrorsStatus defines model for BatchResponseTimelineEventResponseWithErrors.Status.
type BatchResponseTimelineEventResponseWithErrorsStatus string

// CollectionResponseTimelineEventTemplateNoPaging defines model for CollectionResponseTimelineEventTemplateNoPaging.
type CollectionResponseTimelineEventTemplateNoPaging struct {
	Results []TimelineEventTemplate `json:"results"`
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

// ErrorCategory defines model for ErrorCategory.
type ErrorCategory struct {
	HttpStatus ErrorCategoryHttpStatus `json:"httpStatus"`
	Name       string                  `json:"name"`
}

// ErrorCategoryHttpStatus defines model for ErrorCategory.HttpStatus.
type ErrorCategoryHttpStatus string

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

// The details Markdown rendered as HTML.
type EventDetail struct {
	// The details Markdown rendered as HTML.
	Details string `json:"details"`
}

// StandardError defines model for StandardError.
type StandardError struct {
	Category    ErrorCategory           `json:"category"`
	Context     StandardError_Context   `json:"context"`
	Errors      []ErrorDetail           `json:"errors"`
	Id          *string                 `json:"id,omitempty"`
	Links       StandardError_Links     `json:"links"`
	Message     string                  `json:"message"`
	Status      string                  `json:"status"`
	SubCategory *map[string]interface{} `json:"subCategory,omitempty"`
}

// StandardError_Context defines model for StandardError.Context.
type StandardError_Context struct {
	AdditionalProperties map[string][]string `json:"-"`
}

// StandardError_Links defines model for StandardError.Links.
type StandardError_Links struct {
	AdditionalProperties map[string]string `json:"-"`
}

// The state of the timeline event.
type TimelineEvent struct {
	// The event domain (often paired with utk).
	Domain *string `json:"domain,omitempty"`

	// The email address used for contact-specific events. This can be used to identify existing contacts, create new ones, or change the email for an existing contact (if paired with the `objectId`).
	Email *string `json:"email,omitempty"`

	// The event template ID.
	EventTemplateId string `json:"eventTemplateId"`

	// Additional event-specific data that can be interpreted by the template's markdown.
	ExtraData *map[string]interface{} `json:"extraData,omitempty"`

	// Identifier for the event. This is optional, and we recommend you do not pass this in. We will create one for you if you omit this. You can also use `{{uuid}}` anywhere in the ID to generate a unique string, guaranteeing uniqueness.
	Id string `json:"id"`

	// The CRM object identifier. This is required for every event other than contacts (where utk or email can be used).
	ObjectId       *string              `json:"objectId,omitempty"`
	TimelineIFrame *TimelineEventIFrame `json:"timelineIFrame,omitempty"`

	// The time the event occurred. If not passed in, the curren time will be assumed. This is used to determine where an event is shown on a CRM object's timeline.
	Timestamp *time.Time `json:"timestamp,omitempty"`

	// A collection of token keys and values associated with the template tokens.
	Tokens TimelineEvent_Tokens `json:"tokens"`

	// Use the `utk` parameter to associate an event with a contact by `usertoken`. This is recommended if you don't know a user's email, but have an identifying user token in your cookie.
	Utk *string `json:"utk,omitempty"`
}

// A collection of token keys and values associated with the template tokens.
type TimelineEvent_Tokens struct {
	AdditionalProperties map[string]string `json:"-"`
}

// TimelineEventIFrame defines model for TimelineEventIFrame.
type TimelineEventIFrame struct {
	// The label of the modal window that displays the iframe contents.
	HeaderLabel string `json:"headerLabel"`

	// The height of the modal window in pixels.
	Height int32 `json:"height"`

	// The text displaying the link that will display the iframe.
	LinkLabel string `json:"linkLabel"`

	// The URI of the iframe contents.
	Url string `json:"url"`

	// The width of the modal window in pixels.
	Width int32 `json:"width"`
}

// The current state of the timeline event.
type TimelineEventResponse struct {
	CreatedAt *time.Time `json:"createdAt,omitempty"`

	// The event domain (often paired with utk).
	Domain *string `json:"domain,omitempty"`

	// The email address used for contact-specific events. This can be used to identify existing contacts, create new ones, or change the email for an existing contact (if paired with the `objectId`).
	Email *string `json:"email,omitempty"`

	// The event template ID.
	EventTemplateId string `json:"eventTemplateId"`

	// Additional event-specific data that can be interpreted by the template's markdown.
	ExtraData *map[string]interface{} `json:"extraData,omitempty"`

	// Identifier for the event. This should be unique to the app and event template. If you use the same ID for different CRM objects, the last to be processed will win and the first will not have a record. You can also use `{{uuid}}` anywhere in the ID to generate a unique string, guaranteeing uniqueness.
	Id string `json:"id"`

	// The CRM object identifier. This is required for every event other than contacts (where utk or email can be used).
	ObjectId *string `json:"objectId,omitempty"`

	// The ObjectType associated with the EventTemplate.
	ObjectType     string               `json:"objectType"`
	TimelineIFrame *TimelineEventIFrame `json:"timelineIFrame,omitempty"`

	// The time the event occurred. If not passed in, the curren time will be assumed. This is used to determine where an event is shown on a CRM object's timeline.
	Timestamp *time.Time `json:"timestamp,omitempty"`

	// A collection of token keys and values associated with the template tokens.
	Tokens TimelineEventResponse_Tokens `json:"tokens"`

	// Use the `utk` parameter to associate an event with a contact by `usertoken`. This is recommended if you don't know a user's email, but have an identifying user token in your cookie.
	Utk *string `json:"utk,omitempty"`
}

// A collection of token keys and values associated with the template tokens.
type TimelineEventResponse_Tokens struct {
	AdditionalProperties map[string]string `json:"-"`
}

// The current state of the template definition.
type TimelineEventTemplate struct {
	// The date and time that the Event Template was created, as an ISO 8601 timestamp. Will be null if the template was created before Feb 18th, 2020.
	CreatedAt *time.Time `json:"createdAt,omitempty"`

	// This uses Markdown syntax with Handlebars and event-specific data to render HTML on a timeline when you expand the details.
	DetailTemplate *string `json:"detailTemplate,omitempty"`

	// This uses Markdown syntax with Handlebars and event-specific data to render HTML on a timeline as a header.
	HeaderTemplate *string `json:"headerTemplate,omitempty"`

	// The template ID.
	Id string `json:"id"`

	// The template name.
	Name string `json:"name"`

	// The type of CRM object this template is for. [Contacts, companies, tickets, and deals] are supported.
	ObjectType string `json:"objectType"`

	// A collection of tokens that can be used as custom properties on the event and to create fully fledged CRM objects.
	Tokens []TimelineEventTemplateToken `json:"tokens"`

	// The date and time that the Event Template was last updated, as an ISO 8601 timestamp. Will be null if the template was created before Feb 18th, 2020.
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
}

// State of the template definition being created.
type TimelineEventTemplateCreateRequest struct {
	// This uses Markdown syntax with Handlebars and event-specific data to render HTML on a timeline when you expand the details.
	DetailTemplate *string `json:"detailTemplate,omitempty"`

	// This uses Markdown syntax with Handlebars and event-specific data to render HTML on a timeline as a header.
	HeaderTemplate *string `json:"headerTemplate,omitempty"`

	// The template name.
	Name string `json:"name"`

	// The type of CRM object this template is for. [Contacts, companies, tickets, and deals] are supported.
	ObjectType string `json:"objectType"`

	// A collection of tokens that can be used as custom properties on the event and to create fully fledged CRM objects.
	Tokens []TimelineEventTemplateToken `json:"tokens"`
}

// State of the token definition.
type TimelineEventTemplateToken struct {
	// The date and time that the Event Template Token was created, as an ISO 8601 timestamp. Will be null if the template was created before Feb 18th, 2020.
	CreatedAt *time.Time `json:"createdAt,omitempty"`

	// Used for list segmentation and reporting.
	Label string `json:"label"`

	// The name of the token referenced in the templates. This must be unique for the specific template. It may only contain alphanumeric characters, periods, dashes, or underscores (. - _).
	Name string `json:"name"`

	// The name of the CRM object property. This will populate the CRM object property associated with the event. With enough of these, you can fully build CRM objects via the Timeline API.
	ObjectPropertyName *string `json:"objectPropertyName,omitempty"`

	// If type is `enumeration`, we should have a list of options to choose from.
	Options []TimelineEventTemplateTokenOption `json:"options"`

	// The data type of the token. You can currently choose from [string, number, date, enumeration].
	Type TimelineEventTemplateTokenType `json:"type"`

	// The date and time that the Event Template Token was last updated, as an ISO 8601 timestamp. Will be null if the template was created before Feb 18th, 2020.
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
}

// The data type of the token. You can currently choose from [string, number, date, enumeration].
type TimelineEventTemplateTokenType string

// TimelineEventTemplateTokenOption defines model for TimelineEventTemplateTokenOption.
type TimelineEventTemplateTokenOption struct {
	Label string `json:"label"`
	Value string `json:"value"`
}

// State of the token definition for update requests.
type TimelineEventTemplateTokenUpdateRequest struct {
	// Used for list segmentation and reporting.
	Label string `json:"label"`

	// The name of the CRM object property. This will populate the CRM object property associated with the event. With enough of these, you can fully build CRM objects via the Timeline API.
	ObjectPropertyName *string `json:"objectPropertyName,omitempty"`

	// If type is `enumeration`, we should have a list of options to choose from.
	Options []TimelineEventTemplateTokenOption `json:"options"`
}

// State of the template definition being updated.
type TimelineEventTemplateUpdateRequest struct {
	// This uses Markdown syntax with Handlebars and event-specific data to render HTML on a timeline when you expand the details.
	DetailTemplate *string `json:"detailTemplate,omitempty"`

	// This uses Markdown syntax with Handlebars and event-specific data to render HTML on a timeline as a header.
	HeaderTemplate *string `json:"headerTemplate,omitempty"`

	// The template ID.
	Id string `json:"id"`

	// The template name.
	Name string `json:"name"`

	// A collection of tokens that can be used as custom properties on the event and to create fully fledged CRM objects.
	Tokens []TimelineEventTemplateToken `json:"tokens"`
}

// CreateEventsJSONBody defines parameters for CreateEvents.
type CreateEventsJSONBody TimelineEvent

// CreateBatchJSONBody defines parameters for CreateBatch.
type CreateBatchJSONBody BatchInputTimelineEvent

// GetRenderByIdParams defines parameters for GetRenderById.
type GetRenderByIdParams struct {
	// Set to 'true', we want to render the `detailTemplate` instead of the `headerTemplate`.
	Detail *bool `json:"detail,omitempty"`
}

// CreateEventTemplatesJSONBody defines parameters for CreateEventTemplates.
type CreateEventTemplatesJSONBody TimelineEventTemplateCreateRequest

// UpdateEventTemplateJSONBody defines parameters for UpdateEventTemplate.
type UpdateEventTemplateJSONBody TimelineEventTemplateUpdateRequest

// CreateTokensJSONBody defines parameters for CreateTokens.
type CreateTokensJSONBody TimelineEventTemplateToken

// UpdateTokenNameJSONBody defines parameters for UpdateTokenName.
type UpdateTokenNameJSONBody TimelineEventTemplateTokenUpdateRequest

// CreateEventsJSONRequestBody defines body for CreateEvents for application/json ContentType.
type CreateEventsJSONRequestBody CreateEventsJSONBody

// CreateBatchJSONRequestBody defines body for CreateBatch for application/json ContentType.
type CreateBatchJSONRequestBody CreateBatchJSONBody

// CreateEventTemplatesJSONRequestBody defines body for CreateEventTemplates for application/json ContentType.
type CreateEventTemplatesJSONRequestBody CreateEventTemplatesJSONBody

// UpdateEventTemplateJSONRequestBody defines body for UpdateEventTemplate for application/json ContentType.
type UpdateEventTemplateJSONRequestBody UpdateEventTemplateJSONBody

// CreateTokensJSONRequestBody defines body for CreateTokens for application/json ContentType.
type CreateTokensJSONRequestBody CreateTokensJSONBody

// UpdateTokenNameJSONRequestBody defines body for UpdateTokenName for application/json ContentType.
type UpdateTokenNameJSONRequestBody UpdateTokenNameJSONBody

// Getter for additional properties for BatchResponseTimelineEventResponse_Links. Returns the specified
// element and whether it was found
func (a BatchResponseTimelineEventResponse_Links) Get(fieldName string) (value string, found bool) {
	if a.AdditionalProperties != nil {
		value, found = a.AdditionalProperties[fieldName]
	}
	return
}

// Setter for additional properties for BatchResponseTimelineEventResponse_Links
func (a *BatchResponseTimelineEventResponse_Links) Set(fieldName string, value string) {
	if a.AdditionalProperties == nil {
		a.AdditionalProperties = make(map[string]string)
	}
	a.AdditionalProperties[fieldName] = value
}

// Override default JSON handling for BatchResponseTimelineEventResponse_Links to handle AdditionalProperties
func (a *BatchResponseTimelineEventResponse_Links) UnmarshalJSON(b []byte) error {
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

// Override default JSON handling for BatchResponseTimelineEventResponse_Links to handle AdditionalProperties
func (a BatchResponseTimelineEventResponse_Links) MarshalJSON() ([]byte, error) {
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

// Getter for additional properties for BatchResponseTimelineEventResponseWithErrors_Links. Returns the specified
// element and whether it was found
func (a BatchResponseTimelineEventResponseWithErrors_Links) Get(fieldName string) (value string, found bool) {
	if a.AdditionalProperties != nil {
		value, found = a.AdditionalProperties[fieldName]
	}
	return
}

// Setter for additional properties for BatchResponseTimelineEventResponseWithErrors_Links
func (a *BatchResponseTimelineEventResponseWithErrors_Links) Set(fieldName string, value string) {
	if a.AdditionalProperties == nil {
		a.AdditionalProperties = make(map[string]string)
	}
	a.AdditionalProperties[fieldName] = value
}

// Override default JSON handling for BatchResponseTimelineEventResponseWithErrors_Links to handle AdditionalProperties
func (a *BatchResponseTimelineEventResponseWithErrors_Links) UnmarshalJSON(b []byte) error {
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

// Override default JSON handling for BatchResponseTimelineEventResponseWithErrors_Links to handle AdditionalProperties
func (a BatchResponseTimelineEventResponseWithErrors_Links) MarshalJSON() ([]byte, error) {
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

// Getter for additional properties for StandardError_Context. Returns the specified
// element and whether it was found
func (a StandardError_Context) Get(fieldName string) (value []string, found bool) {
	if a.AdditionalProperties != nil {
		value, found = a.AdditionalProperties[fieldName]
	}
	return
}

// Setter for additional properties for StandardError_Context
func (a *StandardError_Context) Set(fieldName string, value []string) {
	if a.AdditionalProperties == nil {
		a.AdditionalProperties = make(map[string][]string)
	}
	a.AdditionalProperties[fieldName] = value
}

// Override default JSON handling for StandardError_Context to handle AdditionalProperties
func (a *StandardError_Context) UnmarshalJSON(b []byte) error {
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

// Override default JSON handling for StandardError_Context to handle AdditionalProperties
func (a StandardError_Context) MarshalJSON() ([]byte, error) {
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

// Getter for additional properties for StandardError_Links. Returns the specified
// element and whether it was found
func (a StandardError_Links) Get(fieldName string) (value string, found bool) {
	if a.AdditionalProperties != nil {
		value, found = a.AdditionalProperties[fieldName]
	}
	return
}

// Setter for additional properties for StandardError_Links
func (a *StandardError_Links) Set(fieldName string, value string) {
	if a.AdditionalProperties == nil {
		a.AdditionalProperties = make(map[string]string)
	}
	a.AdditionalProperties[fieldName] = value
}

// Override default JSON handling for StandardError_Links to handle AdditionalProperties
func (a *StandardError_Links) UnmarshalJSON(b []byte) error {
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

// Override default JSON handling for StandardError_Links to handle AdditionalProperties
func (a StandardError_Links) MarshalJSON() ([]byte, error) {
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

// Getter for additional properties for TimelineEvent_Tokens. Returns the specified
// element and whether it was found
func (a TimelineEvent_Tokens) Get(fieldName string) (value string, found bool) {
	if a.AdditionalProperties != nil {
		value, found = a.AdditionalProperties[fieldName]
	}
	return
}

// Setter for additional properties for TimelineEvent_Tokens
func (a *TimelineEvent_Tokens) Set(fieldName string, value string) {
	if a.AdditionalProperties == nil {
		a.AdditionalProperties = make(map[string]string)
	}
	a.AdditionalProperties[fieldName] = value
}

// Override default JSON handling for TimelineEvent_Tokens to handle AdditionalProperties
func (a *TimelineEvent_Tokens) UnmarshalJSON(b []byte) error {
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

// Override default JSON handling for TimelineEvent_Tokens to handle AdditionalProperties
func (a TimelineEvent_Tokens) MarshalJSON() ([]byte, error) {
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

// Getter for additional properties for TimelineEventResponse_Tokens. Returns the specified
// element and whether it was found
func (a TimelineEventResponse_Tokens) Get(fieldName string) (value string, found bool) {
	if a.AdditionalProperties != nil {
		value, found = a.AdditionalProperties[fieldName]
	}
	return
}

// Setter for additional properties for TimelineEventResponse_Tokens
func (a *TimelineEventResponse_Tokens) Set(fieldName string, value string) {
	if a.AdditionalProperties == nil {
		a.AdditionalProperties = make(map[string]string)
	}
	a.AdditionalProperties[fieldName] = value
}

// Override default JSON handling for TimelineEventResponse_Tokens to handle AdditionalProperties
func (a *TimelineEventResponse_Tokens) UnmarshalJSON(b []byte) error {
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

// Override default JSON handling for TimelineEventResponse_Tokens to handle AdditionalProperties
func (a TimelineEventResponse_Tokens) MarshalJSON() ([]byte, error) {
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
