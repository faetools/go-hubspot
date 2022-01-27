// Package transactional provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/faetools/devtool version (devel) DO NOT EDIT.
package transactional

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/faetools/client"
)

// operation paths

const (
	opPathArchiveTokenFormat               = "./marketing/v3/transactional/smtp-tokens/%s"
	opPathGetTokenByIdFormat               = "./marketing/v3/transactional/smtp-tokens/%s"
	opPathResetPasswordPasswordResetFormat = "./marketing/v3/transactional/smtp-tokens/%s/password-reset"
)

var (
	opPathSendEmail               = client.MustParseURL("./marketing/v3/transactional/single-email/send")
	opPathGetTokensPageSmtpTokens = client.MustParseURL("./marketing/v3/transactional/smtp-tokens")
	opPathCreateTokenSmtpTokens   = client.MustParseURL("./marketing/v3/transactional/smtp-tokens")
)

// ClientInterface interface specification for the client.
type ClientInterface interface {
	// SendEmail request with any body
	SendEmailWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...client.RequestEditorFn) (*SendEmailResponse, error)
	SendEmail(ctx context.Context, body SendEmailJSONRequestBody, reqEditors ...client.RequestEditorFn) (*SendEmailResponse, error)

	// GetTokensPageSmtpTokens request
	GetTokensPageSmtpTokens(ctx context.Context, params *GetTokensPageSmtpTokensParams, reqEditors ...client.RequestEditorFn) (*GetTokensPageSmtpTokensResponse, error)

	// CreateTokenSmtpTokens request with any body
	CreateTokenSmtpTokensWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...client.RequestEditorFn) (*CreateTokenSmtpTokensResponse, error)
	CreateTokenSmtpTokens(ctx context.Context, body CreateTokenSmtpTokensJSONRequestBody, reqEditors ...client.RequestEditorFn) (*CreateTokenSmtpTokensResponse, error)

	// ArchiveToken request
	ArchiveToken(ctx context.Context, tokenId string, reqEditors ...client.RequestEditorFn) (*ArchiveTokenResponse, error)

	// GetTokenById request
	GetTokenById(ctx context.Context, tokenId string, reqEditors ...client.RequestEditorFn) (*GetTokenByIdResponse, error)

	// ResetPasswordPasswordReset request
	ResetPasswordPasswordReset(ctx context.Context, tokenId string, reqEditors ...client.RequestEditorFn) (*ResetPasswordPasswordResetResponse, error)
}

// Client definition

// compile time assert that it fulfils the interface
var _ ClientInterface = (*Client)(nil)

// Client conforms to the OpenAPI3 specification for this service.
type Client client.Client

// NewClient creates a new Client with reasonable defaults.
func NewClient(opts ...client.Option) (*Client, error) {
	c, err := client.NewClient(opts...)
	if err != nil {
		return nil, err
	}

	if c.BaseURL == nil {
		if err := client.WithBaseURL(DefaultServer)(c); err != nil {
			return nil, err
		}
	}

	return (*Client)(c), nil
}

func (c *Client) applyEditors(ctx context.Context, req *http.Request, additionalEditors []client.RequestEditorFn) error {
	for _, r := range c.RequestEditors {
		if err := r(ctx, req); err != nil {
			return err
		}
	}

	for _, r := range additionalEditors {
		if err := r(ctx, req); err != nil {
			return err
		}
	}

	return nil
}

// SendEmail: POST /marketing/v3/transactional/single-email/send

type SendEmailResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *EmailSendStatusView
}

// Status returns HTTPResponse.Status
func (r SendEmailResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r SendEmailResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// newSendEmailRequestWithBody generates requests for SendEmail with any type of body
func newSendEmailRequestWithBody(baseURL *url.URL, contentType string, body io.Reader) (*http.Request, error) {
	queryURL := baseURL.ResolveReference(opPathSendEmail)

	req, err := http.NewRequest(http.MethodPost, queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add(client.ContentType, contentType)

	return req, nil
}

// SendEmailWithBody request with arbitrary body returning *SendEmailResponse
func (c *Client) SendEmailWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...client.RequestEditorFn) (*SendEmailResponse, error) {
	rsp, err := c.doSendEmailWithBody(ctx, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}

	return parseSendEmailResponse(rsp)
}

func (c *Client) doSendEmailWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...client.RequestEditorFn) (*http.Response, error) {
	req, err := newSendEmailRequestWithBody(c.BaseURL, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}

	return c.Client.Do(req)
}

func (c *Client) SendEmail(ctx context.Context, body SendEmailJSONRequestBody, reqEditors ...client.RequestEditorFn) (*SendEmailResponse, error) {
	rsp, err := c.doSendEmail(ctx, body, reqEditors...)
	if err != nil {
		return nil, err
	}

	return parseSendEmailResponse(rsp)
}

// newSendEmailRequest calls the generic SendEmail builder with application/json body.
func newSendEmailRequest(baseURL *url.URL, body SendEmailJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return newSendEmailRequestWithBody(baseURL, client.MIMEApplicationJSON, bodyReader)
}

func (c *Client) doSendEmail(ctx context.Context, body SendEmailJSONRequestBody, reqEditors ...client.RequestEditorFn) (*http.Response, error) {
	req, err := newSendEmailRequest(c.BaseURL, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}

	return c.Client.Do(req)
}

// parseSendEmailResponse parses an HTTP response from a SendEmail call.
func parseSendEmailResponse(rsp *http.Response) (*SendEmailResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	if err != nil {
		return nil, err
	}
	defer rsp.Body.Close()

	response := &SendEmailResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest EmailSendStatusView
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest
	}

	return response, nil
}

// GetTokensPageSmtpTokens: GET /marketing/v3/transactional/smtp-tokens

type GetTokensPageSmtpTokensResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *CollectionResponseSmtpApiTokenView
}

// Status returns HTTPResponse.Status
func (r GetTokensPageSmtpTokensResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetTokensPageSmtpTokensResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// newGetTokensPageSmtpTokensRequest generates requests for GetTokensPageSmtpTokens
func newGetTokensPageSmtpTokensRequest(baseURL *url.URL, params *GetTokensPageSmtpTokensParams) (*http.Request, error) {
	queryURL := baseURL.ResolveReference(opPathGetTokensPageSmtpTokens)

	q := queryURL.Query()

	if params.CampaignName != nil {
		if err := client.AddQueryParam(q, "campaignName", *params.CampaignName); err != nil {
			return nil, err
		}
	}

	if params.EmailCampaignId != nil {
		if err := client.AddQueryParam(q, "emailCampaignId", *params.EmailCampaignId); err != nil {
			return nil, err
		}
	}

	if params.After != nil {
		if err := client.AddQueryParam(q, "after", *params.After); err != nil {
			return nil, err
		}
	}

	if params.Limit != nil {
		if err := client.AddQueryParam(q, "limit", *params.Limit); err != nil {
			return nil, err
		}
	}

	queryURL.RawQuery = q.Encode()

	req, err := http.NewRequest(http.MethodGet, queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// GetTokensPageSmtpTokens request returning *GetTokensPageSmtpTokensResponse
func (c *Client) GetTokensPageSmtpTokens(ctx context.Context, params *GetTokensPageSmtpTokensParams, reqEditors ...client.RequestEditorFn) (*GetTokensPageSmtpTokensResponse, error) {
	req, err := newGetTokensPageSmtpTokensRequest(c.BaseURL, params)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}

	rsp, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}

	bodyBytes, err := io.ReadAll(rsp.Body)
	if err != nil {
		return nil, err
	}
	defer rsp.Body.Close()

	response := &GetTokensPageSmtpTokensResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest CollectionResponseSmtpApiTokenView
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest
	}

	return response, nil
}

// CreateTokenSmtpTokens: POST /marketing/v3/transactional/smtp-tokens

type CreateTokenSmtpTokensResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON201      *SmtpApiTokenView
}

// Status returns HTTPResponse.Status
func (r CreateTokenSmtpTokensResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r CreateTokenSmtpTokensResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// newCreateTokenSmtpTokensRequestWithBody generates requests for CreateTokenSmtpTokens with any type of body
func newCreateTokenSmtpTokensRequestWithBody(baseURL *url.URL, contentType string, body io.Reader) (*http.Request, error) {
	queryURL := baseURL.ResolveReference(opPathCreateTokenSmtpTokens)

	req, err := http.NewRequest(http.MethodPost, queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add(client.ContentType, contentType)

	return req, nil
}

// CreateTokenSmtpTokensWithBody request with arbitrary body returning *CreateTokenSmtpTokensResponse
func (c *Client) CreateTokenSmtpTokensWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...client.RequestEditorFn) (*CreateTokenSmtpTokensResponse, error) {
	rsp, err := c.doCreateTokenSmtpTokensWithBody(ctx, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}

	return parseCreateTokenSmtpTokensResponse(rsp)
}

func (c *Client) doCreateTokenSmtpTokensWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...client.RequestEditorFn) (*http.Response, error) {
	req, err := newCreateTokenSmtpTokensRequestWithBody(c.BaseURL, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}

	return c.Client.Do(req)
}

func (c *Client) CreateTokenSmtpTokens(ctx context.Context, body CreateTokenSmtpTokensJSONRequestBody, reqEditors ...client.RequestEditorFn) (*CreateTokenSmtpTokensResponse, error) {
	rsp, err := c.doCreateTokenSmtpTokens(ctx, body, reqEditors...)
	if err != nil {
		return nil, err
	}

	return parseCreateTokenSmtpTokensResponse(rsp)
}

// newCreateTokenSmtpTokensRequest calls the generic CreateTokenSmtpTokens builder with application/json body.
func newCreateTokenSmtpTokensRequest(baseURL *url.URL, body CreateTokenSmtpTokensJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return newCreateTokenSmtpTokensRequestWithBody(baseURL, client.MIMEApplicationJSON, bodyReader)
}

func (c *Client) doCreateTokenSmtpTokens(ctx context.Context, body CreateTokenSmtpTokensJSONRequestBody, reqEditors ...client.RequestEditorFn) (*http.Response, error) {
	req, err := newCreateTokenSmtpTokensRequest(c.BaseURL, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}

	return c.Client.Do(req)
}

// parseCreateTokenSmtpTokensResponse parses an HTTP response from a CreateTokenSmtpTokens call.
func parseCreateTokenSmtpTokensResponse(rsp *http.Response) (*CreateTokenSmtpTokensResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	if err != nil {
		return nil, err
	}
	defer rsp.Body.Close()

	response := &CreateTokenSmtpTokensResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 201:
		var dest SmtpApiTokenView
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON201 = &dest
	}

	return response, nil
}

// ArchiveToken: DELETE /marketing/v3/transactional/smtp-tokens/{tokenId}

type ArchiveTokenResponse struct {
	Body         []byte
	HTTPResponse *http.Response
}

// Status returns HTTPResponse.Status
func (r ArchiveTokenResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r ArchiveTokenResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// newArchiveTokenRequest generates requests for ArchiveToken
func newArchiveTokenRequest(baseURL *url.URL, tokenId string) (*http.Request, error) {
	pathParam0, err := client.GetPathParam("tokenId", tokenId)
	if err != nil {
		return nil, err
	}

	opPath := fmt.Sprintf(opPathArchiveTokenFormat, pathParam0)

	queryURL, err := baseURL.Parse(opPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodDelete, queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// ArchiveToken request returning *ArchiveTokenResponse
func (c *Client) ArchiveToken(ctx context.Context, tokenId string, reqEditors ...client.RequestEditorFn) (*ArchiveTokenResponse, error) {
	req, err := newArchiveTokenRequest(c.BaseURL, tokenId)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}

	rsp, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}

	bodyBytes, err := io.ReadAll(rsp.Body)
	if err != nil {
		return nil, err
	}
	defer rsp.Body.Close()

	response := &ArchiveTokenResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	return response, nil
}

// GetTokenById: GET /marketing/v3/transactional/smtp-tokens/{tokenId}

type GetTokenByIdResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *SmtpApiTokenView
}

// Status returns HTTPResponse.Status
func (r GetTokenByIdResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetTokenByIdResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// newGetTokenByIdRequest generates requests for GetTokenById
func newGetTokenByIdRequest(baseURL *url.URL, tokenId string) (*http.Request, error) {
	pathParam0, err := client.GetPathParam("tokenId", tokenId)
	if err != nil {
		return nil, err
	}

	opPath := fmt.Sprintf(opPathGetTokenByIdFormat, pathParam0)

	queryURL, err := baseURL.Parse(opPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodGet, queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// GetTokenById request returning *GetTokenByIdResponse
func (c *Client) GetTokenById(ctx context.Context, tokenId string, reqEditors ...client.RequestEditorFn) (*GetTokenByIdResponse, error) {
	req, err := newGetTokenByIdRequest(c.BaseURL, tokenId)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}

	rsp, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}

	bodyBytes, err := io.ReadAll(rsp.Body)
	if err != nil {
		return nil, err
	}
	defer rsp.Body.Close()

	response := &GetTokenByIdResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest SmtpApiTokenView
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest
	}

	return response, nil
}

// ResetPasswordPasswordReset: POST /marketing/v3/transactional/smtp-tokens/{tokenId}/password-reset

type ResetPasswordPasswordResetResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *SmtpApiTokenView
}

// Status returns HTTPResponse.Status
func (r ResetPasswordPasswordResetResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r ResetPasswordPasswordResetResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// newResetPasswordPasswordResetRequest generates requests for ResetPasswordPasswordReset
func newResetPasswordPasswordResetRequest(baseURL *url.URL, tokenId string) (*http.Request, error) {
	pathParam0, err := client.GetPathParam("tokenId", tokenId)
	if err != nil {
		return nil, err
	}

	opPath := fmt.Sprintf(opPathResetPasswordPasswordResetFormat, pathParam0)

	queryURL, err := baseURL.Parse(opPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// ResetPasswordPasswordReset request returning *ResetPasswordPasswordResetResponse
func (c *Client) ResetPasswordPasswordReset(ctx context.Context, tokenId string, reqEditors ...client.RequestEditorFn) (*ResetPasswordPasswordResetResponse, error) {
	req, err := newResetPasswordPasswordResetRequest(c.BaseURL, tokenId)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}

	rsp, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}

	bodyBytes, err := io.ReadAll(rsp.Body)
	if err != nil {
		return nil, err
	}
	defer rsp.Body.Close()

	response := &ResetPasswordPasswordResetResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest SmtpApiTokenView
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest
	}

	return response, nil
}
