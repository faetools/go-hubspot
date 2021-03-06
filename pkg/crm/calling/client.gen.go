// Package calling provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/faetools/devtool version (devel) DO NOT EDIT.
package calling

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
	opPathArchiveSettingsFormat = "./crm/v3/extensions/calling/%s/settings"
	opPathGetSettingsFormat     = "./crm/v3/extensions/calling/%s/settings"
	opPathUpdateSettingsFormat  = "./crm/v3/extensions/calling/%s/settings"
	opPathCreateSettingsFormat  = "./crm/v3/extensions/calling/%s/settings"
)

// ClientInterface interface specification for the client.
type ClientInterface interface {
	// ArchiveSettings request
	ArchiveSettings(ctx context.Context, appId int32, reqEditors ...client.RequestEditorFn) (*ArchiveSettingsResponse, error)

	// GetSettings request
	GetSettings(ctx context.Context, appId int32, reqEditors ...client.RequestEditorFn) (*GetSettingsResponse, error)

	// UpdateSettings request with any body
	UpdateSettingsWithBody(ctx context.Context, appId int32, contentType string, body io.Reader, reqEditors ...client.RequestEditorFn) (*UpdateSettingsResponse, error)
	UpdateSettings(ctx context.Context, appId int32, body UpdateSettingsJSONRequestBody, reqEditors ...client.RequestEditorFn) (*UpdateSettingsResponse, error)

	// CreateSettings request with any body
	CreateSettingsWithBody(ctx context.Context, appId int32, contentType string, body io.Reader, reqEditors ...client.RequestEditorFn) (*CreateSettingsResponse, error)
	CreateSettings(ctx context.Context, appId int32, body CreateSettingsJSONRequestBody, reqEditors ...client.RequestEditorFn) (*CreateSettingsResponse, error)
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

// ArchiveSettings: DELETE /crm/v3/extensions/calling/{appId}/settings

type ArchiveSettingsResponse struct {
	Body         []byte
	HTTPResponse *http.Response
}

// Status returns HTTPResponse.Status
func (r ArchiveSettingsResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r ArchiveSettingsResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// newArchiveSettingsRequest generates requests for ArchiveSettings
func newArchiveSettingsRequest(baseURL *url.URL, appId int32) (*http.Request, error) {
	pathParam0, err := client.GetPathParam("appId", appId)
	if err != nil {
		return nil, err
	}

	opPath := fmt.Sprintf(opPathArchiveSettingsFormat, pathParam0)

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

// ArchiveSettings request returning *ArchiveSettingsResponse
func (c *Client) ArchiveSettings(ctx context.Context, appId int32, reqEditors ...client.RequestEditorFn) (*ArchiveSettingsResponse, error) {
	req, err := newArchiveSettingsRequest(c.BaseURL, appId)
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

	response := &ArchiveSettingsResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	return response, nil
}

// GetSettings: GET /crm/v3/extensions/calling/{appId}/settings

type GetSettingsResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *SettingsResponse
}

// Status returns HTTPResponse.Status
func (r GetSettingsResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetSettingsResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// newGetSettingsRequest generates requests for GetSettings
func newGetSettingsRequest(baseURL *url.URL, appId int32) (*http.Request, error) {
	pathParam0, err := client.GetPathParam("appId", appId)
	if err != nil {
		return nil, err
	}

	opPath := fmt.Sprintf(opPathGetSettingsFormat, pathParam0)

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

// GetSettings request returning *GetSettingsResponse
func (c *Client) GetSettings(ctx context.Context, appId int32, reqEditors ...client.RequestEditorFn) (*GetSettingsResponse, error) {
	req, err := newGetSettingsRequest(c.BaseURL, appId)
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

	response := &GetSettingsResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest SettingsResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest
	}

	return response, nil
}

// UpdateSettings: PATCH /crm/v3/extensions/calling/{appId}/settings

type UpdateSettingsResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *SettingsResponse
}

// Status returns HTTPResponse.Status
func (r UpdateSettingsResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r UpdateSettingsResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// newUpdateSettingsRequestWithBody generates requests for UpdateSettings with any type of body
func newUpdateSettingsRequestWithBody(baseURL *url.URL, appId int32, contentType string, body io.Reader) (*http.Request, error) {
	pathParam0, err := client.GetPathParam("appId", appId)
	if err != nil {
		return nil, err
	}

	opPath := fmt.Sprintf(opPathUpdateSettingsFormat, pathParam0)

	queryURL, err := baseURL.Parse(opPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPatch, queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add(client.ContentType, contentType)

	return req, nil
}

// UpdateSettingsWithBody request with arbitrary body returning *UpdateSettingsResponse
func (c *Client) UpdateSettingsWithBody(ctx context.Context, appId int32, contentType string, body io.Reader, reqEditors ...client.RequestEditorFn) (*UpdateSettingsResponse, error) {
	rsp, err := c.doUpdateSettingsWithBody(ctx, appId, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}

	return parseUpdateSettingsResponse(rsp)
}

func (c *Client) doUpdateSettingsWithBody(ctx context.Context, appId int32, contentType string, body io.Reader, reqEditors ...client.RequestEditorFn) (*http.Response, error) {
	req, err := newUpdateSettingsRequestWithBody(c.BaseURL, appId, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}

	return c.Client.Do(req)
}

func (c *Client) UpdateSettings(ctx context.Context, appId int32, body UpdateSettingsJSONRequestBody, reqEditors ...client.RequestEditorFn) (*UpdateSettingsResponse, error) {
	rsp, err := c.doUpdateSettings(ctx, appId, body, reqEditors...)
	if err != nil {
		return nil, err
	}

	return parseUpdateSettingsResponse(rsp)
}

// newUpdateSettingsRequest calls the generic UpdateSettings builder with application/json body.
func newUpdateSettingsRequest(baseURL *url.URL, appId int32, body UpdateSettingsJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return newUpdateSettingsRequestWithBody(baseURL, appId, client.MIMEApplicationJSON, bodyReader)
}

func (c *Client) doUpdateSettings(ctx context.Context, appId int32, body UpdateSettingsJSONRequestBody, reqEditors ...client.RequestEditorFn) (*http.Response, error) {
	req, err := newUpdateSettingsRequest(c.BaseURL, appId, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}

	return c.Client.Do(req)
}

// parseUpdateSettingsResponse parses an HTTP response from a UpdateSettings call.
func parseUpdateSettingsResponse(rsp *http.Response) (*UpdateSettingsResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	if err != nil {
		return nil, err
	}
	defer rsp.Body.Close()

	response := &UpdateSettingsResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest SettingsResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest
	}

	return response, nil
}

// CreateSettings: POST /crm/v3/extensions/calling/{appId}/settings

type CreateSettingsResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *SettingsResponse
}

// Status returns HTTPResponse.Status
func (r CreateSettingsResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r CreateSettingsResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// newCreateSettingsRequestWithBody generates requests for CreateSettings with any type of body
func newCreateSettingsRequestWithBody(baseURL *url.URL, appId int32, contentType string, body io.Reader) (*http.Request, error) {
	pathParam0, err := client.GetPathParam("appId", appId)
	if err != nil {
		return nil, err
	}

	opPath := fmt.Sprintf(opPathCreateSettingsFormat, pathParam0)

	queryURL, err := baseURL.Parse(opPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add(client.ContentType, contentType)

	return req, nil
}

// CreateSettingsWithBody request with arbitrary body returning *CreateSettingsResponse
func (c *Client) CreateSettingsWithBody(ctx context.Context, appId int32, contentType string, body io.Reader, reqEditors ...client.RequestEditorFn) (*CreateSettingsResponse, error) {
	rsp, err := c.doCreateSettingsWithBody(ctx, appId, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}

	return parseCreateSettingsResponse(rsp)
}

func (c *Client) doCreateSettingsWithBody(ctx context.Context, appId int32, contentType string, body io.Reader, reqEditors ...client.RequestEditorFn) (*http.Response, error) {
	req, err := newCreateSettingsRequestWithBody(c.BaseURL, appId, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}

	return c.Client.Do(req)
}

func (c *Client) CreateSettings(ctx context.Context, appId int32, body CreateSettingsJSONRequestBody, reqEditors ...client.RequestEditorFn) (*CreateSettingsResponse, error) {
	rsp, err := c.doCreateSettings(ctx, appId, body, reqEditors...)
	if err != nil {
		return nil, err
	}

	return parseCreateSettingsResponse(rsp)
}

// newCreateSettingsRequest calls the generic CreateSettings builder with application/json body.
func newCreateSettingsRequest(baseURL *url.URL, appId int32, body CreateSettingsJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return newCreateSettingsRequestWithBody(baseURL, appId, client.MIMEApplicationJSON, bodyReader)
}

func (c *Client) doCreateSettings(ctx context.Context, appId int32, body CreateSettingsJSONRequestBody, reqEditors ...client.RequestEditorFn) (*http.Response, error) {
	req, err := newCreateSettingsRequest(c.BaseURL, appId, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}

	return c.Client.Do(req)
}

// parseCreateSettingsResponse parses an HTTP response from a CreateSettings call.
func parseCreateSettingsResponse(rsp *http.Response) (*CreateSettingsResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	if err != nil {
		return nil, err
	}
	defer rsp.Body.Close()

	response := &CreateSettingsResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest SettingsResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest
	}

	return response, nil
}
