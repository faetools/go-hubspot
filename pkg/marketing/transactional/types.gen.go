// Package transactional provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/faetools/devtool version (devel) DO NOT EDIT.
package transactional

import (
	"encoding/json"
	"fmt"
	"time"
)

const (
	HapikeyScopes       = "hapikey.Scopes"
	Oauth2_legacyScopes = "oauth2_legacy.Scopes"
)

// Defines values for EmailSendStatusViewSendResult.
const (
	EmailSendStatusViewSendResultADDRESSLISTBOMBED EmailSendStatusViewSendResult = "ADDRESS_LIST_BOMBED"

	EmailSendStatusViewSendResultADDRESSONLYACCEPTEDONPROD EmailSendStatusViewSendResult = "ADDRESS_ONLY_ACCEPTED_ON_PROD"

	EmailSendStatusViewSendResultADDRESSOPTEDOUT EmailSendStatusViewSendResult = "ADDRESS_OPTED_OUT"

	EmailSendStatusViewSendResultBLOCKEDADDRESS EmailSendStatusViewSendResult = "BLOCKED_ADDRESS"

	EmailSendStatusViewSendResultBLOCKEDDOMAIN EmailSendStatusViewSendResult = "BLOCKED_DOMAIN"

	EmailSendStatusViewSendResultCAMPAIGNCANCELLED EmailSendStatusViewSendResult = "CAMPAIGN_CANCELLED"

	EmailSendStatusViewSendResultCANCELLEDABUSE EmailSendStatusViewSendResult = "CANCELLED_ABUSE"

	EmailSendStatusViewSendResultCORRUPTINPUT EmailSendStatusViewSendResult = "CORRUPT_INPUT"

	EmailSendStatusViewSendResultEMAILDISABLED EmailSendStatusViewSendResult = "EMAIL_DISABLED"

	EmailSendStatusViewSendResultEMAILUNCONFIRMED EmailSendStatusViewSendResult = "EMAIL_UNCONFIRMED"

	EmailSendStatusViewSendResultGRAYMAILSUPPRESSED EmailSendStatusViewSendResult = "GRAYMAIL_SUPPRESSED"

	EmailSendStatusViewSendResultIDEMPOTENTFAIL EmailSendStatusViewSendResult = "IDEMPOTENT_FAIL"

	EmailSendStatusViewSendResultIDEMPOTENTIGNORE EmailSendStatusViewSendResult = "IDEMPOTENT_IGNORE"

	EmailSendStatusViewSendResultINVALIDFROMADDRESS EmailSendStatusViewSendResult = "INVALID_FROM_ADDRESS"

	EmailSendStatusViewSendResultINVALIDTOADDRESS EmailSendStatusViewSendResult = "INVALID_TO_ADDRESS"

	EmailSendStatusViewSendResultMISSINGCONTENT EmailSendStatusViewSendResult = "MISSING_CONTENT"

	EmailSendStatusViewSendResultMISSINGREQUIREDPARAMETER EmailSendStatusViewSendResult = "MISSING_REQUIRED_PARAMETER"

	EmailSendStatusViewSendResultMISSINGTEMPLATEPROPERTIES EmailSendStatusViewSendResult = "MISSING_TEMPLATE_PROPERTIES"

	EmailSendStatusViewSendResultMTAIGNORE EmailSendStatusViewSendResult = "MTA_IGNORE"

	EmailSendStatusViewSendResultNONMARKETABLECONTACT EmailSendStatusViewSendResult = "NON_MARKETABLE_CONTACT"

	EmailSendStatusViewSendResultPORTALAUTHENTICATIONFAILURE EmailSendStatusViewSendResult = "PORTAL_AUTHENTICATION_FAILURE"

	EmailSendStatusViewSendResultPORTALEXPIRED EmailSendStatusViewSendResult = "PORTAL_EXPIRED"

	EmailSendStatusViewSendResultPORTALMISSINGMARKETINGSCOPE EmailSendStatusViewSendResult = "PORTAL_MISSING_MARKETING_SCOPE"

	EmailSendStatusViewSendResultPORTALNOTAUTHORIZEDFORAPPLICATION EmailSendStatusViewSendResult = "PORTAL_NOT_AUTHORIZED_FOR_APPLICATION"

	EmailSendStatusViewSendResultPORTALOVERLIMIT EmailSendStatusViewSendResult = "PORTAL_OVER_LIMIT"

	EmailSendStatusViewSendResultPORTALSUSPENDED EmailSendStatusViewSendResult = "PORTAL_SUSPENDED"

	EmailSendStatusViewSendResultPREVIOUSLYBOUNCED EmailSendStatusViewSendResult = "PREVIOUSLY_BOUNCED"

	EmailSendStatusViewSendResultPREVIOUSLYUNSUBSCRIBEDBRAND EmailSendStatusViewSendResult = "PREVIOUSLY_UNSUBSCRIBED_BRAND"

	EmailSendStatusViewSendResultPREVIOUSLYUNSUBSCRIBEDMESSAGE EmailSendStatusViewSendResult = "PREVIOUSLY_UNSUBSCRIBED_MESSAGE"

	EmailSendStatusViewSendResultPREVIOUSLYUNSUBSCRIBEDPORTAL EmailSendStatusViewSendResult = "PREVIOUSLY_UNSUBSCRIBED_PORTAL"

	EmailSendStatusViewSendResultPREVIOUSSPAM EmailSendStatusViewSendResult = "PREVIOUS_SPAM"

	EmailSendStatusViewSendResultQUARANTINEDADDRESS EmailSendStatusViewSendResult = "QUARANTINED_ADDRESS"

	EmailSendStatusViewSendResultQUEUED EmailSendStatusViewSendResult = "QUEUED"

	EmailSendStatusViewSendResultRECIPIENTFATIGUESUPPRESSED EmailSendStatusViewSendResult = "RECIPIENT_FATIGUE_SUPPRESSED"

	EmailSendStatusViewSendResultSENT EmailSendStatusViewSendResult = "SENT"

	EmailSendStatusViewSendResultTEMPLATERENDEREXCEPTION EmailSendStatusViewSendResult = "TEMPLATE_RENDER_EXCEPTION"

	EmailSendStatusViewSendResultTHROTTLED EmailSendStatusViewSendResult = "THROTTLED"

	EmailSendStatusViewSendResultTOOMANYRECIPIENTS EmailSendStatusViewSendResult = "TOO_MANY_RECIPIENTS"

	EmailSendStatusViewSendResultUNCONFIGUREDSENDINGDOMAIN EmailSendStatusViewSendResult = "UNCONFIGURED_SENDING_DOMAIN"

	EmailSendStatusViewSendResultUNDELIVERABLE EmailSendStatusViewSendResult = "UNDELIVERABLE"

	EmailSendStatusViewSendResultVALIDATIONFAILED EmailSendStatusViewSendResult = "VALIDATION_FAILED"
)

// Defines values for EmailSendStatusViewStatus.
const (
	EmailSendStatusViewStatusCANCELED EmailSendStatusViewStatus = "CANCELED"

	EmailSendStatusViewStatusCOMPLETE EmailSendStatusViewStatus = "COMPLETE"

	EmailSendStatusViewStatusPENDING EmailSendStatusViewStatus = "PENDING"

	EmailSendStatusViewStatusPROCESSING EmailSendStatusViewStatus = "PROCESSING"
)

// A collection of SMTP API tokens.
type CollectionResponseSmtpApiTokenView struct {
	Paging *Paging `json:"paging,omitempty"`

	// The actual collection of tokens.
	Results []SmtpApiTokenView `json:"results"`
}

// Describes the status of an email send request.
type EmailSendStatusView struct {
	// Time when the send was completed.
	CompletedAt *time.Time `json:"completedAt,omitempty"`

	// The ID of a send event.
	EventId *EventIdView `json:"eventId,omitempty"`

	// Time when the send was requested.
	RequestedAt *time.Time `json:"requestedAt,omitempty"`

	// Result of the send.
	SendResult *EmailSendStatusViewSendResult `json:"sendResult,omitempty"`

	// Time when the send began processing.
	StartedAt *time.Time `json:"startedAt,omitempty"`

	// Status of the send request.
	Status EmailSendStatusViewStatus `json:"status"`

	// Identifier used to query the status of the send.
	StatusId string `json:"statusId"`
}

// Result of the send.
type EmailSendStatusViewSendResult string

// Status of the send request.
type EmailSendStatusViewStatus string

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

// The ID of a send event.
type EventIdView struct {
	// Time of event creation.
	Created time.Time `json:"created"`

	// Identifier of event.
	Id string `json:"id"`
}

// NextPage defines model for NextPage.
type NextPage struct {
	After string  `json:"after"`
	Link  *string `json:"link,omitempty"`
}

// Paging defines model for Paging.
type Paging struct {
	Next *NextPage `json:"next,omitempty"`
}

// A JSON object containing anything you want to override.
type PublicSingleSendEmail struct {
	// List of email addresses to send as Bcc.
	Bcc []string `json:"bcc"`

	// List of email addresses to send as Cc.
	Cc []string `json:"cc"`

	// The From header for the email.
	From *string `json:"from,omitempty"`

	// List of Reply-To header values for the email.
	ReplyTo []string `json:"replyTo"`

	// ID for a particular send. No more than one email will be sent per sendId.
	SendId *string `json:"sendId,omitempty"`

	// The recipient of the email.
	To *string `json:"to,omitempty"`
}

// A request to send a single transactional email asynchronously.
type PublicSingleSendRequestEgg struct {
	// The contactProperties field is a map of contact property values. Each contact property value contains a name and value property. Each property will get set on the contact record and will be visible in the template under {{ contact.NAME }}. Use these properties when you want to set a contact property while you’re sending the email. For example, when sending a reciept you may want to set a last_paid_date property, as the sending of the receipt will have information about the last payment.
	ContactProperties PublicSingleSendRequestEgg_ContactProperties `json:"contactProperties"`

	// The customProperties field is a map of property values. Each property value contains a name and value property. Each property will be visible in the template under {{ custom.NAME }}.
	// Note: Custom properties do not currently support arrays. To provide a listing in an email, one workaround is to build an HTML list (either with tables or ul) and specify it as a custom property.
	CustomProperties map[string]interface{} `json:"customProperties"`

	// The content ID for the transactional email, which can be found in email tool UI.
	EmailId int32 `json:"emailId"`

	// A JSON object containing anything you want to override.
	Message PublicSingleSendEmail `json:"message"`
}

// The contactProperties field is a map of contact property values. Each contact property value contains a name and value property. Each property will get set on the contact record and will be visible in the template under {{ contact.NAME }}. Use these properties when you want to set a contact property while you’re sending the email. For example, when sending a reciept you may want to set a last_paid_date property, as the sending of the receipt will have information about the last payment.
type PublicSingleSendRequestEgg_ContactProperties struct {
	AdditionalProperties map[string]string `json:"-"`
}

// A request object to create a SMTP API token
type SmtpApiTokenRequestEgg struct {
	// A name for the campaign tied to the SMTP API token.
	CampaignName string `json:"campaignName"`

	// Indicates whether a contact should be created for recipients of emails.
	CreateContact bool `json:"createContact"`
}

// A SMTP API token provides both an ID and password that can be used to send email through the HubSpot SMTP API.
type SmtpApiTokenView struct {
	// A name for the campaign tied to the token.
	CampaignName string `json:"campaignName"`

	// Indicates whether a contact should be created for recipients of emails.
	CreateContact bool `json:"createContact"`

	// Timestamp generated when a token is created.
	CreatedAt time.Time `json:"createdAt"`

	// Email address of the user that sent the token creation request.
	CreatedBy string `json:"createdBy"`

	// Identifier assigned to the campaign provided in the token creation request.
	EmailCampaignId string `json:"emailCampaignId"`

	// User name to log into the HubSpot SMTP server.
	Id string `json:"id"`

	// Password used to log into the HubSpot SMTP server.
	Password *string `json:"password,omitempty"`
}

// SendEmailJSONBody defines parameters for SendEmail.
type SendEmailJSONBody PublicSingleSendRequestEgg

// GetTokensPageSmtpTokensParams defines parameters for GetTokensPageSmtpTokens.
type GetTokensPageSmtpTokensParams struct {
	// A name for the campaign tied to the SMTP API token.
	CampaignName *string `json:"campaignName,omitempty"`

	// Identifier assigned to the campaign provided during the token creation.
	EmailCampaignId *string `json:"emailCampaignId,omitempty"`

	// Starting point to get the next set of results.
	After *string `json:"after,omitempty"`

	// Maximum number of tokens to return.
	Limit *int32 `json:"limit,omitempty"`
}

// CreateTokenSmtpTokensJSONBody defines parameters for CreateTokenSmtpTokens.
type CreateTokenSmtpTokensJSONBody SmtpApiTokenRequestEgg

// SendEmailJSONRequestBody defines body for SendEmail for application/json ContentType.
type SendEmailJSONRequestBody SendEmailJSONBody

// CreateTokenSmtpTokensJSONRequestBody defines body for CreateTokenSmtpTokens for application/json ContentType.
type CreateTokenSmtpTokensJSONRequestBody CreateTokenSmtpTokensJSONBody

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

// Getter for additional properties for PublicSingleSendRequestEgg_ContactProperties. Returns the specified
// element and whether it was found
func (a PublicSingleSendRequestEgg_ContactProperties) Get(fieldName string) (value string, found bool) {
	if a.AdditionalProperties != nil {
		value, found = a.AdditionalProperties[fieldName]
	}
	return
}

// Setter for additional properties for PublicSingleSendRequestEgg_ContactProperties
func (a *PublicSingleSendRequestEgg_ContactProperties) Set(fieldName string, value string) {
	if a.AdditionalProperties == nil {
		a.AdditionalProperties = make(map[string]string)
	}
	a.AdditionalProperties[fieldName] = value
}

// Override default JSON handling for PublicSingleSendRequestEgg_ContactProperties to handle AdditionalProperties
func (a *PublicSingleSendRequestEgg_ContactProperties) UnmarshalJSON(b []byte) error {
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

// Override default JSON handling for PublicSingleSendRequestEgg_ContactProperties to handle AdditionalProperties
func (a PublicSingleSendRequestEgg_ContactProperties) MarshalJSON() ([]byte, error) {
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
