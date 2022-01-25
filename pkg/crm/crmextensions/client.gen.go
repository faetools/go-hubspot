// Package crmextensions provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/faetools/devtool version (devel) DO NOT EDIT.
package crmextensions

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
func WithRequestEditorFn(fn client.RequestEditorFn) ClientOption {
	return func(c *Client) error {
		c.RequestEditors = append(c.RequestEditors, fn)
		return nil
	}
}

func (c *Client) doGetCardsSampleResponseSampleResponse(ctx context.Context, reqEditors ...client.RequestEditorFn) (*http.Response, error) {
	req, err := newGetCardsSampleResponseSampleResponseRequest(c.Server)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) doGetAllApp(ctx context.Context, appId int32, reqEditors ...client.RequestEditorFn) (*http.Response, error) {
	req, err := newGetAllAppRequest(c.Server, appId)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) doCreateAppWithBody(ctx context.Context, appId int32, contentType string, body io.Reader, reqEditors ...client.RequestEditorFn) (*http.Response, error) {
	req, err := newCreateAppRequestWithBody(c.Server, appId, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) doCreateApp(ctx context.Context, appId int32, body CreateAppJSONRequestBody, reqEditors ...client.RequestEditorFn) (*http.Response, error) {
	req, err := newCreateAppRequest(c.Server, appId, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) doArchiveCard(ctx context.Context, appId int32, cardId string, reqEditors ...client.RequestEditorFn) (*http.Response, error) {
	req, err := newArchiveCardRequest(c.Server, appId, cardId)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) doGetCard(ctx context.Context, appId int32, cardId string, reqEditors ...client.RequestEditorFn) (*http.Response, error) {
	req, err := newGetCardRequest(c.Server, appId, cardId)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) doUpdateCardWithBody(ctx context.Context, appId int32, cardId string, contentType string, body io.Reader, reqEditors ...client.RequestEditorFn) (*http.Response, error) {
	req, err := newUpdateCardRequestWithBody(c.Server, appId, cardId, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) doUpdateCard(ctx context.Context, appId int32, cardId string, body UpdateCardJSONRequestBody, reqEditors ...client.RequestEditorFn) (*http.Response, error) {
	req, err := newUpdateCardRequest(c.Server, appId, cardId, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// newGetCardsSampleResponseSampleResponseRequest generates requests for GetCardsSampleResponseSampleResponse
func newGetCardsSampleResponseSampleResponseRequest(server string) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/crm/v3/extensions/cards/sample-response")
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

// newGetAllAppRequest generates requests for GetAllApp
func newGetAllAppRequest(server string, appId int32) (*http.Request, error) {
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

	operationPath := fmt.Sprintf("/crm/v3/extensions/cards/%s", pathParam0)
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

// newCreateAppRequest calls the generic CreateApp builder with application/json body.
func newCreateAppRequest(server string, appId int32, body CreateAppJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return newCreateAppRequestWithBody(server, appId, "application/json", bodyReader)
}

// newCreateAppRequestWithBody generates requests for CreateApp with any type of body
func newCreateAppRequestWithBody(server string, appId int32, contentType string, body io.Reader) (*http.Request, error) {
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

	operationPath := fmt.Sprintf("/crm/v3/extensions/cards/%s", pathParam0)
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

// newArchiveCardRequest generates requests for ArchiveCard
func newArchiveCardRequest(server string, appId int32, cardId string) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "appId", runtime.ParamLocationPath, appId)
	if err != nil {
		return nil, err
	}

	var pathParam1 string

	pathParam1, err = runtime.StyleParamWithLocation("simple", false, "cardId", runtime.ParamLocationPath, cardId)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/crm/v3/extensions/cards/%s/%s", pathParam0, pathParam1)
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

// newGetCardRequest generates requests for GetCard
func newGetCardRequest(server string, appId int32, cardId string) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "appId", runtime.ParamLocationPath, appId)
	if err != nil {
		return nil, err
	}

	var pathParam1 string

	pathParam1, err = runtime.StyleParamWithLocation("simple", false, "cardId", runtime.ParamLocationPath, cardId)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/crm/v3/extensions/cards/%s/%s", pathParam0, pathParam1)
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

// newUpdateCardRequest calls the generic UpdateCard builder with application/json body.
func newUpdateCardRequest(server string, appId int32, cardId string, body UpdateCardJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return newUpdateCardRequestWithBody(server, appId, cardId, "application/json", bodyReader)
}

// newUpdateCardRequestWithBody generates requests for UpdateCard with any type of body
func newUpdateCardRequestWithBody(server string, appId int32, cardId string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "appId", runtime.ParamLocationPath, appId)
	if err != nil {
		return nil, err
	}

	var pathParam1 string

	pathParam1, err = runtime.StyleParamWithLocation("simple", false, "cardId", runtime.ParamLocationPath, cardId)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/crm/v3/extensions/cards/%s/%s", pathParam0, pathParam1)
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
	RequestEditors []client.RequestEditorFn
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
	// GetCardsSampleResponseSampleResponse request
	GetCardsSampleResponseSampleResponse(ctx context.Context, reqEditors ...client.RequestEditorFn) (*GetCardsSampleResponseSampleResponseResponse, error)

	// GetAllApp request
	GetAllApp(ctx context.Context, appId int32, reqEditors ...client.RequestEditorFn) (*GetAllAppResponse, error)

	// CreateApp request with any body
	CreateAppWithBody(ctx context.Context, appId int32, contentType string, body io.Reader, reqEditors ...client.RequestEditorFn) (*CreateAppResponse, error)
	CreateApp(ctx context.Context, appId int32, body CreateAppJSONRequestBody, reqEditors ...client.RequestEditorFn) (*CreateAppResponse, error)

	// ArchiveCard request
	ArchiveCard(ctx context.Context, appId int32, cardId string, reqEditors ...client.RequestEditorFn) (*ArchiveCardResponse, error)

	// GetCard request
	GetCard(ctx context.Context, appId int32, cardId string, reqEditors ...client.RequestEditorFn) (*GetCardResponse, error)

	// UpdateCard request with any body
	UpdateCardWithBody(ctx context.Context, appId int32, cardId string, contentType string, body io.Reader, reqEditors ...client.RequestEditorFn) (*UpdateCardResponse, error)
	UpdateCard(ctx context.Context, appId int32, cardId string, body UpdateCardJSONRequestBody, reqEditors ...client.RequestEditorFn) (*UpdateCardResponse, error)
}

type GetCardsSampleResponseSampleResponseResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *IntegratorCardPayloadResponse
}

// Status returns HTTPResponse.Status
func (r GetCardsSampleResponseSampleResponseResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetCardsSampleResponseSampleResponseResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetAllAppResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *CardListResponse
}

// Status returns HTTPResponse.Status
func (r GetAllAppResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetAllAppResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type CreateAppResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON201      *CardResponse
}

// Status returns HTTPResponse.Status
func (r CreateAppResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r CreateAppResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type ArchiveCardResponse struct {
	Body         []byte
	HTTPResponse *http.Response
}

// Status returns HTTPResponse.Status
func (r ArchiveCardResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r ArchiveCardResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetCardResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *CardResponse
}

// Status returns HTTPResponse.Status
func (r GetCardResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetCardResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type UpdateCardResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *CardResponse
}

// Status returns HTTPResponse.Status
func (r UpdateCardResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r UpdateCardResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// GetCardsSampleResponseSampleResponse request returning *GetCardsSampleResponseSampleResponseResponse
func (c *Client) GetCardsSampleResponseSampleResponse(ctx context.Context, reqEditors ...client.RequestEditorFn) (*GetCardsSampleResponseSampleResponseResponse, error) {
	rsp, err := c.doGetCardsSampleResponseSampleResponse(ctx, reqEditors...)
	if err != nil {
		return nil, err
	}
	return parseGetCardsSampleResponseSampleResponseResponse(rsp)
}

// GetAllApp request returning *GetAllAppResponse
func (c *Client) GetAllApp(ctx context.Context, appId int32, reqEditors ...client.RequestEditorFn) (*GetAllAppResponse, error) {
	rsp, err := c.doGetAllApp(ctx, appId, reqEditors...)
	if err != nil {
		return nil, err
	}
	return parseGetAllAppResponse(rsp)
}

// CreateAppWithBody request with arbitrary body returning *CreateAppResponse
func (c *Client) CreateAppWithBody(ctx context.Context, appId int32, contentType string, body io.Reader, reqEditors ...client.RequestEditorFn) (*CreateAppResponse, error) {
	rsp, err := c.doCreateAppWithBody(ctx, appId, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return parseCreateAppResponse(rsp)
}

func (c *Client) CreateApp(ctx context.Context, appId int32, body CreateAppJSONRequestBody, reqEditors ...client.RequestEditorFn) (*CreateAppResponse, error) {
	rsp, err := c.doCreateApp(ctx, appId, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return parseCreateAppResponse(rsp)
}

// ArchiveCard request returning *ArchiveCardResponse
func (c *Client) ArchiveCard(ctx context.Context, appId int32, cardId string, reqEditors ...client.RequestEditorFn) (*ArchiveCardResponse, error) {
	rsp, err := c.doArchiveCard(ctx, appId, cardId, reqEditors...)
	if err != nil {
		return nil, err
	}
	return parseArchiveCardResponse(rsp)
}

// GetCard request returning *GetCardResponse
func (c *Client) GetCard(ctx context.Context, appId int32, cardId string, reqEditors ...client.RequestEditorFn) (*GetCardResponse, error) {
	rsp, err := c.doGetCard(ctx, appId, cardId, reqEditors...)
	if err != nil {
		return nil, err
	}
	return parseGetCardResponse(rsp)
}

// UpdateCardWithBody request with arbitrary body returning *UpdateCardResponse
func (c *Client) UpdateCardWithBody(ctx context.Context, appId int32, cardId string, contentType string, body io.Reader, reqEditors ...client.RequestEditorFn) (*UpdateCardResponse, error) {
	rsp, err := c.doUpdateCardWithBody(ctx, appId, cardId, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return parseUpdateCardResponse(rsp)
}

func (c *Client) UpdateCard(ctx context.Context, appId int32, cardId string, body UpdateCardJSONRequestBody, reqEditors ...client.RequestEditorFn) (*UpdateCardResponse, error) {
	rsp, err := c.doUpdateCard(ctx, appId, cardId, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return parseUpdateCardResponse(rsp)
}

// parseGetCardsSampleResponseSampleResponseResponse parses an HTTP response from a GetCardsSampleResponseSampleResponse call.
func parseGetCardsSampleResponseSampleResponseResponse(rsp *http.Response) (*GetCardsSampleResponseSampleResponseResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetCardsSampleResponseSampleResponseResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest IntegratorCardPayloadResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest
	}

	return response, nil
}

// parseGetAllAppResponse parses an HTTP response from a GetAllApp call.
func parseGetAllAppResponse(rsp *http.Response) (*GetAllAppResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetAllAppResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest CardListResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest
	}

	return response, nil
}

// parseCreateAppResponse parses an HTTP response from a CreateApp call.
func parseCreateAppResponse(rsp *http.Response) (*CreateAppResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &CreateAppResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 201:
		var dest CardResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON201 = &dest
	}

	return response, nil
}

// parseArchiveCardResponse parses an HTTP response from a ArchiveCard call.
func parseArchiveCardResponse(rsp *http.Response) (*ArchiveCardResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &ArchiveCardResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	return response, nil
}

// parseGetCardResponse parses an HTTP response from a GetCard call.
func parseGetCardResponse(rsp *http.Response) (*GetCardResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetCardResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest CardResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest
	}

	return response, nil
}

// parseUpdateCardResponse parses an HTTP response from a UpdateCard call.
func parseUpdateCardResponse(rsp *http.Response) (*UpdateCardResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &UpdateCardResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest CardResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest
	}

	return response, nil
}
