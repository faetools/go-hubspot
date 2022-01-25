// Package oauth provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/faetools/devtool version (devel) DO NOT EDIT.
package oauth

import (
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

func (c *Client) doGetAccessToken(ctx context.Context, token string, reqEditors ...client.RequestEditorFn) (*http.Response, error) {
	req, err := newGetAccessTokenRequest(c.baseURL, token)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.client.Do(req)
}

func (c *Client) doArchiveRefreshToken(ctx context.Context, token string, reqEditors ...client.RequestEditorFn) (*http.Response, error) {
	req, err := newArchiveRefreshTokenRequest(c.baseURL, token)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.client.Do(req)
}

func (c *Client) doGetRefreshToken(ctx context.Context, token string, reqEditors ...client.RequestEditorFn) (*http.Response, error) {
	req, err := newGetRefreshTokenRequest(c.baseURL, token)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.client.Do(req)
}

func (c *Client) doCreateTokenWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...client.RequestEditorFn) (*http.Response, error) {
	req, err := newCreateTokenRequestWithBody(c.baseURL, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.client.Do(req)
}

// newGetAccessTokenRequest generates requests for GetAccessToken
func newGetAccessTokenRequest(baseURL *url.URL, token string) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "token", runtime.ParamLocationPath, token)
	if err != nil {
		return nil, err
	}

	opPathGetAccessToken := fmt.Sprintf("/oauth/v1/access-tokens/%s", pathParam0)
	if opPathGetAccessToken[0] == '/' {
		opPathGetAccessToken = "." + opPathGetAccessToken
	}

	queryURL, err := baseURL.Parse(opPathGetAccessToken)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// newArchiveRefreshTokenRequest generates requests for ArchiveRefreshToken
func newArchiveRefreshTokenRequest(baseURL *url.URL, token string) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "token", runtime.ParamLocationPath, token)
	if err != nil {
		return nil, err
	}

	opPathArchiveRefreshToken := fmt.Sprintf("/oauth/v1/refresh-tokens/%s", pathParam0)
	if opPathArchiveRefreshToken[0] == '/' {
		opPathArchiveRefreshToken = "." + opPathArchiveRefreshToken
	}

	queryURL, err := baseURL.Parse(opPathArchiveRefreshToken)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("DELETE", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// newGetRefreshTokenRequest generates requests for GetRefreshToken
func newGetRefreshTokenRequest(baseURL *url.URL, token string) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "token", runtime.ParamLocationPath, token)
	if err != nil {
		return nil, err
	}

	opPathGetRefreshToken := fmt.Sprintf("/oauth/v1/refresh-tokens/%s", pathParam0)
	if opPathGetRefreshToken[0] == '/' {
		opPathGetRefreshToken = "." + opPathGetRefreshToken
	}

	queryURL, err := baseURL.Parse(opPathGetRefreshToken)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

const opPathCreateToken = "./oauth/v1/token"

// newCreateTokenRequestWithBody generates requests for CreateToken with any type of body
func newCreateTokenRequestWithBody(baseURL *url.URL, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	queryURL, err := baseURL.Parse(opPathCreateToken)
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
	// GetAccessToken request
	GetAccessToken(ctx context.Context, token string, reqEditors ...client.RequestEditorFn) (*GetAccessTokenResponse, error)

	// ArchiveRefreshToken request
	ArchiveRefreshToken(ctx context.Context, token string, reqEditors ...client.RequestEditorFn) (*ArchiveRefreshTokenResponse, error)

	// GetRefreshToken request
	GetRefreshToken(ctx context.Context, token string, reqEditors ...client.RequestEditorFn) (*GetRefreshTokenResponse, error)

	// CreateToken request with any body
	CreateTokenWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...client.RequestEditorFn) (*CreateTokenResponse, error)
}

type GetAccessTokenResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *AccessTokenInfoResponse
}

// Status returns HTTPResponse.Status
func (r GetAccessTokenResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetAccessTokenResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type ArchiveRefreshTokenResponse struct {
	Body         []byte
	HTTPResponse *http.Response
}

// Status returns HTTPResponse.Status
func (r ArchiveRefreshTokenResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r ArchiveRefreshTokenResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetRefreshTokenResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *RefreshTokenInfoResponse
}

// Status returns HTTPResponse.Status
func (r GetRefreshTokenResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetRefreshTokenResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type CreateTokenResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *TokenResponseIF
}

// Status returns HTTPResponse.Status
func (r CreateTokenResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r CreateTokenResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// GetAccessToken request returning *GetAccessTokenResponse
func (c *Client) GetAccessToken(ctx context.Context, token string, reqEditors ...client.RequestEditorFn) (*GetAccessTokenResponse, error) {
	rsp, err := c.doGetAccessToken(ctx, token, reqEditors...)
	if err != nil {
		return nil, err
	}
	return parseGetAccessTokenResponse(rsp)
}

// ArchiveRefreshToken request returning *ArchiveRefreshTokenResponse
func (c *Client) ArchiveRefreshToken(ctx context.Context, token string, reqEditors ...client.RequestEditorFn) (*ArchiveRefreshTokenResponse, error) {
	rsp, err := c.doArchiveRefreshToken(ctx, token, reqEditors...)
	if err != nil {
		return nil, err
	}
	return parseArchiveRefreshTokenResponse(rsp)
}

// GetRefreshToken request returning *GetRefreshTokenResponse
func (c *Client) GetRefreshToken(ctx context.Context, token string, reqEditors ...client.RequestEditorFn) (*GetRefreshTokenResponse, error) {
	rsp, err := c.doGetRefreshToken(ctx, token, reqEditors...)
	if err != nil {
		return nil, err
	}
	return parseGetRefreshTokenResponse(rsp)
}

// CreateTokenWithBody request with arbitrary body returning *CreateTokenResponse
func (c *Client) CreateTokenWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...client.RequestEditorFn) (*CreateTokenResponse, error) {
	rsp, err := c.doCreateTokenWithBody(ctx, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return parseCreateTokenResponse(rsp)
}

// parseGetAccessTokenResponse parses an HTTP response from a GetAccessToken call.
func parseGetAccessTokenResponse(rsp *http.Response) (*GetAccessTokenResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetAccessTokenResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest AccessTokenInfoResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest
	}

	return response, nil
}

// parseArchiveRefreshTokenResponse parses an HTTP response from a ArchiveRefreshToken call.
func parseArchiveRefreshTokenResponse(rsp *http.Response) (*ArchiveRefreshTokenResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &ArchiveRefreshTokenResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	return response, nil
}

// parseGetRefreshTokenResponse parses an HTTP response from a GetRefreshToken call.
func parseGetRefreshTokenResponse(rsp *http.Response) (*GetRefreshTokenResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetRefreshTokenResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest RefreshTokenInfoResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest
	}

	return response, nil
}

// parseCreateTokenResponse parses an HTTP response from a CreateToken call.
func parseCreateTokenResponse(rsp *http.Response) (*CreateTokenResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &CreateTokenResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest TokenResponseIF
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest
	}

	return response, nil
}
