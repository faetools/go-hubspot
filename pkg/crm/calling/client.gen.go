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
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/deepmap/oapi-codegen/pkg/runtime"
)

// RequestEditorFn  is the function signature for the RequestEditor callback function
type RequestEditorFn func(ctx context.Context, req *http.Request) error

// Doer performs HTTP requests.
//
// The standard http.Client implements this interface.
type HttpRequestDoer interface {
	Do(req *http.Request) (*http.Response, error)
}

// ClientOption allows setting custom parameters during construction.
type ClientOption func(*Client) error

// WithHTTPClient allows overriding the default Doer, which is
// automatically created using http.Client. This is useful for tests.
func WithHTTPClient(doer HttpRequestDoer) ClientOption {
	return func(c *Client) error {
		c.Client = doer
		return nil
	}
}

// WithRequestEditorFn allows setting up a callback function, which will be
// called right before sending the request. This can be used to mutate the request.
func WithRequestEditorFn(fn RequestEditorFn) ClientOption {
	return func(c *Client) error {
		c.RequestEditors = append(c.RequestEditors, fn)
		return nil
	}
}

func (c *Client) doArchiveSettings(ctx context.Context, appId int32, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := newArchiveSettingsRequest(c.Server, appId)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) doGetByIdSettings(ctx context.Context, appId int32, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := newGetByIdSettingsRequest(c.Server, appId)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) doUpdateSettingsWithBody(ctx context.Context, appId int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := newUpdateSettingsRequestWithBody(c.Server, appId, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) doUpdateSettings(ctx context.Context, appId int32, body UpdateSettingsJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := newUpdateSettingsRequest(c.Server, appId, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) doCreateSettingsWithBody(ctx context.Context, appId int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := newCreateSettingsRequestWithBody(c.Server, appId, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) doCreateSettings(ctx context.Context, appId int32, body CreateSettingsJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := newCreateSettingsRequest(c.Server, appId, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// newArchiveSettingsRequest generates requests for ArchiveSettings
func newArchiveSettingsRequest(server string, appId int32) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "appId", runtime.ParamLocationPath, appId)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/crm/v3/extensions/calling/%s/settings", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("DELETE", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// newGetByIdSettingsRequest generates requests for GetByIdSettings
func newGetByIdSettingsRequest(server string, appId int32) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "appId", runtime.ParamLocationPath, appId)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/crm/v3/extensions/calling/%s/settings", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// newUpdateSettingsRequest calls the generic UpdateSettings builder with application/json body.
func newUpdateSettingsRequest(server string, appId int32, body UpdateSettingsJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return newUpdateSettingsRequestWithBody(server, appId, "application/json", bodyReader)
}

// newUpdateSettingsRequestWithBody generates requests for UpdateSettings with any type of body
func newUpdateSettingsRequestWithBody(server string, appId int32, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "appId", runtime.ParamLocationPath, appId)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/crm/v3/extensions/calling/%s/settings", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PATCH", queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

	return req, nil
}

// newCreateSettingsRequest calls the generic CreateSettings builder with application/json body.
func newCreateSettingsRequest(server string, appId int32, body CreateSettingsJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return newCreateSettingsRequestWithBody(server, appId, "application/json", bodyReader)
}

// newCreateSettingsRequestWithBody generates requests for CreateSettings with any type of body
func newCreateSettingsRequestWithBody(server string, appId int32, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "appId", runtime.ParamLocationPath, appId)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/crm/v3/extensions/calling/%s/settings", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

	return req, nil
}

func (c *Client) applyEditors(ctx context.Context, req *http.Request, additionalEditors []RequestEditorFn) error {
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

// Client conforms to the OpenAPI3 specification for this service.
type Client struct {
	// The endpoint of the server conforming to this interface, with scheme,
	// https://api.deepmap.com for example. This can contain a path relative
	// to the server, such as https://api.deepmap.com/dev-test, and all the
	// paths in the swagger spec will be appended to the server.
	Server string

	// Doer for performing requests, typically a *http.Client with any
	// customized settings, such as certificate chains.
	Client HttpRequestDoer

	// A list of callbacks for modifying requests which are generated before sending over
	// the network.
	RequestEditors []RequestEditorFn
}

// NewClient creates a new Client, with reasonable defaults.
func NewClient(opts ...ClientOption) (*Client, error) {
	// create a client with default server
	client := Client{Server: DefaultServer}

	// mutate client and add all optional params
	for _, o := range opts {
		if err := o(&client); err != nil {
			return nil, err
		}
	}

	// ensure the server URL always has a trailing slash
	if !strings.HasSuffix(client.Server, "/") {
		client.Server += "/"
	}

	// create httpClient, if not already present
	if client.Client == nil {
		client.Client = &http.Client{}
	}

	return &client, nil
}

// WithBaseURL overrides the baseURL.
func WithBaseURL(baseURL string) ClientOption {
	return func(c *Client) error {
		newBaseURL, err := url.Parse(baseURL)
		if err != nil {
			return err
		}
		c.Server = newBaseURL.String()
		return nil
	}
}

// ClientInterface interface specification for the client.
type ClientInterface interface {
	// ArchiveSettings request
	ArchiveSettings(ctx context.Context, appId int32, reqEditors ...RequestEditorFn) (*ArchiveSettingsResponse, error)

	// GetByIdSettings request
	GetByIdSettings(ctx context.Context, appId int32, reqEditors ...RequestEditorFn) (*GetByIdSettingsResponse, error)

	// UpdateSettings request with any body
	UpdateSettingsWithBody(ctx context.Context, appId int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*UpdateSettingsResponse, error)
	UpdateSettings(ctx context.Context, appId int32, body UpdateSettingsJSONRequestBody, reqEditors ...RequestEditorFn) (*UpdateSettingsResponse, error)

	// CreateSettings request with any body
	CreateSettingsWithBody(ctx context.Context, appId int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*CreateSettingsResponse, error)
	CreateSettings(ctx context.Context, appId int32, body CreateSettingsJSONRequestBody, reqEditors ...RequestEditorFn) (*CreateSettingsResponse, error)
}

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

type GetByIdSettingsResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *SettingsResponse
}

// Status returns HTTPResponse.Status
func (r GetByIdSettingsResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetByIdSettingsResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

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

// ArchiveSettings request returning *ArchiveSettingsResponse
func (c *Client) ArchiveSettings(ctx context.Context, appId int32, reqEditors ...RequestEditorFn) (*ArchiveSettingsResponse, error) {
	rsp, err := c.doArchiveSettings(ctx, appId, reqEditors...)
	if err != nil {
		return nil, err
	}
	return parseArchiveSettingsResponse(rsp)
}

// GetByIdSettings request returning *GetByIdSettingsResponse
func (c *Client) GetByIdSettings(ctx context.Context, appId int32, reqEditors ...RequestEditorFn) (*GetByIdSettingsResponse, error) {
	rsp, err := c.doGetByIdSettings(ctx, appId, reqEditors...)
	if err != nil {
		return nil, err
	}
	return parseGetByIdSettingsResponse(rsp)
}

// UpdateSettingsWithBody request with arbitrary body returning *UpdateSettingsResponse
func (c *Client) UpdateSettingsWithBody(ctx context.Context, appId int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*UpdateSettingsResponse, error) {
	rsp, err := c.doUpdateSettingsWithBody(ctx, appId, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return parseUpdateSettingsResponse(rsp)
}

func (c *Client) UpdateSettings(ctx context.Context, appId int32, body UpdateSettingsJSONRequestBody, reqEditors ...RequestEditorFn) (*UpdateSettingsResponse, error) {
	rsp, err := c.doUpdateSettings(ctx, appId, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return parseUpdateSettingsResponse(rsp)
}

// CreateSettingsWithBody request with arbitrary body returning *CreateSettingsResponse
func (c *Client) CreateSettingsWithBody(ctx context.Context, appId int32, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*CreateSettingsResponse, error) {
	rsp, err := c.doCreateSettingsWithBody(ctx, appId, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return parseCreateSettingsResponse(rsp)
}

func (c *Client) CreateSettings(ctx context.Context, appId int32, body CreateSettingsJSONRequestBody, reqEditors ...RequestEditorFn) (*CreateSettingsResponse, error) {
	rsp, err := c.doCreateSettings(ctx, appId, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return parseCreateSettingsResponse(rsp)
}

// parseArchiveSettingsResponse parses an HTTP response from a ArchiveSettings call.
func parseArchiveSettingsResponse(rsp *http.Response) (*ArchiveSettingsResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &ArchiveSettingsResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	return response, nil
}

// parseGetByIdSettingsResponse parses an HTTP response from a GetByIdSettings call.
func parseGetByIdSettingsResponse(rsp *http.Response) (*GetByIdSettingsResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetByIdSettingsResponse{
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

// parseUpdateSettingsResponse parses an HTTP response from a UpdateSettings call.
func parseUpdateSettingsResponse(rsp *http.Response) (*UpdateSettingsResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

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

// parseCreateSettingsResponse parses an HTTP response from a CreateSettings call.
func parseCreateSettingsResponse(rsp *http.Response) (*CreateSettingsResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

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