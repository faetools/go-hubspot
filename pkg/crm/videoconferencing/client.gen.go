// Package videoconferencing provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/faetools/devtool version (devel) DO NOT EDIT.
package videoconferencing

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
	"github.com/faetools/client"
)

func (c *Client) doArchiveApp(ctx context.Context, appId int32, reqEditors ...client.RequestEditorFn) (*http.Response, error) {
	req, err := newArchiveAppRequest(c.baseURL, appId)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.client.Do(req)
}

func (c *Client) doGetApp(ctx context.Context, appId int32, reqEditors ...client.RequestEditorFn) (*http.Response, error) {
	req, err := newGetAppRequest(c.baseURL, appId)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.client.Do(req)
}

func (c *Client) doReplaceAppWithBody(ctx context.Context, appId int32, contentType string, body io.Reader, reqEditors ...client.RequestEditorFn) (*http.Response, error) {
	req, err := newReplaceAppRequestWithBody(c.baseURL, appId, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.client.Do(req)
}

func (c *Client) doReplaceApp(ctx context.Context, appId int32, body ReplaceAppJSONRequestBody, reqEditors ...client.RequestEditorFn) (*http.Response, error) {
	req, err := newReplaceAppRequest(c.baseURL, appId, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.client.Do(req)
}

// newArchiveAppRequest generates requests for ArchiveApp
func newArchiveAppRequest(baseURL *url.URL, appId int32) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "appId", runtime.ParamLocationPath, appId)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/crm/v3/extensions/videoconferencing/settings/%s", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := baseURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("DELETE", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// newGetAppRequest generates requests for GetApp
func newGetAppRequest(baseURL *url.URL, appId int32) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "appId", runtime.ParamLocationPath, appId)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/crm/v3/extensions/videoconferencing/settings/%s", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := baseURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// newReplaceAppRequest calls the generic ReplaceApp builder with application/json body.
func newReplaceAppRequest(baseURL *url.URL, appId int32, body ReplaceAppJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return newReplaceAppRequestWithBody(baseURL, appId, "application/json", bodyReader)
}

// newReplaceAppRequestWithBody generates requests for ReplaceApp with any type of body
func newReplaceAppRequestWithBody(baseURL *url.URL, appId int32, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "appId", runtime.ParamLocationPath, appId)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/crm/v3/extensions/videoconferencing/settings/%s", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := baseURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PUT", queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

	return req, nil
}

func (c *Client) applyEditors(ctx context.Context, req *http.Request, additionalEditors []client.RequestEditorFn) error {
	for _, r := range c.requestEditors {
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

// compile time assert that it fulfils the interface
var _ ClientInterface = (*Client)(nil)

// Client conforms to the OpenAPI3 specification for this service.
type Client struct {
	// The endpoint of the server conforming to this interface, with scheme,
	// https://api.deepmap.com for example. This can contain a path relative
	// to the server, such as https://api.deepmap.com/dev-test, and all the
	// paths in the swagger spec will be appended to the server.
	baseURL *url.URL

	// Doer for performing requests, typically a *http.Client with any
	// customized settings, such as certificate chains.
	client client.HTTPRequestDoer

	// A list of callbacks for modifying requests which are generated before sending over
	// the network.
	requestEditors []client.RequestEditorFn
}

// SetClient sets the underlying client.
func (c *Client) SetClient(doer client.HTTPRequestDoer) {
	c.client = doer
}

// AddRequestEditor adds a request editor to the client.
func (c *Client) AddRequestEditor(fn client.RequestEditorFn) {
	c.requestEditors = append(c.requestEditors, fn)
}

// SetBaseURL overrides the baseURL.
func (c *Client) SetBaseURL(baseURL *url.URL) {
	c.baseURL = baseURL
}

// NewClient creates a new Client, with reasonable defaults.
func NewClient(opts ...client.Option) (*Client, error) {
	// create a client
	c := Client{}

	// mutate client and add all optional params
	for _, o := range opts {
		if err := o(&c); err != nil {
			return nil, err
		}
	}

	// add default server
	if c.baseURL == nil {
		if err := client.WithBaseURL(DefaultServer)(&c); err != nil {
			return nil, err
		}
	}

	// create httpClient, if not already present
	if c.client == nil {
		c.client = &http.Client{}
	}

	return &c, nil
}

// ClientInterface interface specification for the client.
type ClientInterface interface {
	client.Client
	// ArchiveApp request
	ArchiveApp(ctx context.Context, appId int32, reqEditors ...client.RequestEditorFn) (*ArchiveAppResponse, error)

	// GetApp request
	GetApp(ctx context.Context, appId int32, reqEditors ...client.RequestEditorFn) (*GetAppResponse, error)

	// ReplaceApp request with any body
	ReplaceAppWithBody(ctx context.Context, appId int32, contentType string, body io.Reader, reqEditors ...client.RequestEditorFn) (*ReplaceAppResponse, error)
	ReplaceApp(ctx context.Context, appId int32, body ReplaceAppJSONRequestBody, reqEditors ...client.RequestEditorFn) (*ReplaceAppResponse, error)
}

type ArchiveAppResponse struct {
	Body         []byte
	HTTPResponse *http.Response
}

// Status returns HTTPResponse.Status
func (r ArchiveAppResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r ArchiveAppResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetAppResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *ExternalSettings
}

// Status returns HTTPResponse.Status
func (r GetAppResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetAppResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type ReplaceAppResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *ExternalSettings
}

// Status returns HTTPResponse.Status
func (r ReplaceAppResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r ReplaceAppResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// ArchiveApp request returning *ArchiveAppResponse
func (c *Client) ArchiveApp(ctx context.Context, appId int32, reqEditors ...client.RequestEditorFn) (*ArchiveAppResponse, error) {
	rsp, err := c.doArchiveApp(ctx, appId, reqEditors...)
	if err != nil {
		return nil, err
	}
	return parseArchiveAppResponse(rsp)
}

// GetApp request returning *GetAppResponse
func (c *Client) GetApp(ctx context.Context, appId int32, reqEditors ...client.RequestEditorFn) (*GetAppResponse, error) {
	rsp, err := c.doGetApp(ctx, appId, reqEditors...)
	if err != nil {
		return nil, err
	}
	return parseGetAppResponse(rsp)
}

// ReplaceAppWithBody request with arbitrary body returning *ReplaceAppResponse
func (c *Client) ReplaceAppWithBody(ctx context.Context, appId int32, contentType string, body io.Reader, reqEditors ...client.RequestEditorFn) (*ReplaceAppResponse, error) {
	rsp, err := c.doReplaceAppWithBody(ctx, appId, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return parseReplaceAppResponse(rsp)
}

func (c *Client) ReplaceApp(ctx context.Context, appId int32, body ReplaceAppJSONRequestBody, reqEditors ...client.RequestEditorFn) (*ReplaceAppResponse, error) {
	rsp, err := c.doReplaceApp(ctx, appId, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return parseReplaceAppResponse(rsp)
}

// parseArchiveAppResponse parses an HTTP response from a ArchiveApp call.
func parseArchiveAppResponse(rsp *http.Response) (*ArchiveAppResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &ArchiveAppResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	return response, nil
}

// parseGetAppResponse parses an HTTP response from a GetApp call.
func parseGetAppResponse(rsp *http.Response) (*GetAppResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetAppResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest ExternalSettings
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest
	}

	return response, nil
}

// parseReplaceAppResponse parses an HTTP response from a ReplaceApp call.
func parseReplaceAppResponse(rsp *http.Response) (*ReplaceAppResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &ReplaceAppResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest ExternalSettings
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest
	}

	return response, nil
}
