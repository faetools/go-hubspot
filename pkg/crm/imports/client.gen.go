// Package imports provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/faetools/devtool version (devel) DO NOT EDIT.
package imports

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

func (c *Client) doList(ctx context.Context, params *ListParams, reqEditors ...client.RequestEditorFn) (*http.Response, error) {
	req, err := newListRequest(c.baseURL, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.client.Do(req)
}

func (c *Client) doCreateWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...client.RequestEditorFn) (*http.Response, error) {
	req, err := newCreateRequestWithBody(c.baseURL, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.client.Do(req)
}

func (c *Client) doGetImport(ctx context.Context, importId int64, reqEditors ...client.RequestEditorFn) (*http.Response, error) {
	req, err := newGetImportRequest(c.baseURL, importId)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.client.Do(req)
}

func (c *Client) doCancelImport(ctx context.Context, importId int64, reqEditors ...client.RequestEditorFn) (*http.Response, error) {
	req, err := newCancelImportRequest(c.baseURL, importId)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.client.Do(req)
}

func (c *Client) doGetErrors(ctx context.Context, importId int64, params *GetErrorsParams, reqEditors ...client.RequestEditorFn) (*http.Response, error) {
	req, err := newGetErrorsRequest(c.baseURL, importId, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.client.Do(req)
}

// newListRequest generates requests for List
func newListRequest(baseURL *url.URL, params *ListParams) (*http.Request, error) {
	var err error

	operationPath := fmt.Sprintf("/crm/v3/imports/")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := baseURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	queryValues := queryURL.Query()

	if params.After != nil {
		if queryFrag, err := runtime.StyleParamWithLocation("form", true, "after", runtime.ParamLocationQuery, *params.After); err != nil {
			return nil, err
		} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
			return nil, err
		} else {
			for k, v := range parsed {
				for _, v2 := range v {
					queryValues.Add(k, v2)
				}
			}
		}
	}

	if params.Before != nil {
		if queryFrag, err := runtime.StyleParamWithLocation("form", true, "before", runtime.ParamLocationQuery, *params.Before); err != nil {
			return nil, err
		} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
			return nil, err
		} else {
			for k, v := range parsed {
				for _, v2 := range v {
					queryValues.Add(k, v2)
				}
			}
		}
	}

	if params.Limit != nil {
		if queryFrag, err := runtime.StyleParamWithLocation("form", true, "limit", runtime.ParamLocationQuery, *params.Limit); err != nil {
			return nil, err
		} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
			return nil, err
		} else {
			for k, v := range parsed {
				for _, v2 := range v {
					queryValues.Add(k, v2)
				}
			}
		}
	}

	queryURL.RawQuery = queryValues.Encode()

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// newCreateRequestWithBody generates requests for Create with any type of body
func newCreateRequestWithBody(baseURL *url.URL, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	operationPath := fmt.Sprintf("/crm/v3/imports/")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := baseURL.Parse(operationPath)
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

// newGetImportRequest generates requests for GetImport
func newGetImportRequest(baseURL *url.URL, importId int64) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "importId", runtime.ParamLocationPath, importId)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/crm/v3/imports/%s", pathParam0)
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

// newCancelImportRequest generates requests for CancelImport
func newCancelImportRequest(baseURL *url.URL, importId int64) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "importId", runtime.ParamLocationPath, importId)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/crm/v3/imports/%s/cancel", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := baseURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// newGetErrorsRequest generates requests for GetErrors
func newGetErrorsRequest(baseURL *url.URL, importId int64, params *GetErrorsParams) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "importId", runtime.ParamLocationPath, importId)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/crm/v3/imports/%s/errors", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := baseURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	queryValues := queryURL.Query()

	if params.After != nil {
		if queryFrag, err := runtime.StyleParamWithLocation("form", true, "after", runtime.ParamLocationQuery, *params.After); err != nil {
			return nil, err
		} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
			return nil, err
		} else {
			for k, v := range parsed {
				for _, v2 := range v {
					queryValues.Add(k, v2)
				}
			}
		}
	}

	if params.Limit != nil {
		if queryFrag, err := runtime.StyleParamWithLocation("form", true, "limit", runtime.ParamLocationQuery, *params.Limit); err != nil {
			return nil, err
		} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
			return nil, err
		} else {
			for k, v := range parsed {
				for _, v2 := range v {
					queryValues.Add(k, v2)
				}
			}
		}
	}

	queryURL.RawQuery = queryValues.Encode()

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

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
	// List request
	List(ctx context.Context, params *ListParams, reqEditors ...client.RequestEditorFn) (*ListResponse, error)

	// Create request with any body
	CreateWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...client.RequestEditorFn) (*CreateResponse, error)

	// GetImport request
	GetImport(ctx context.Context, importId int64, reqEditors ...client.RequestEditorFn) (*GetImportResponse, error)

	// CancelImport request
	CancelImport(ctx context.Context, importId int64, reqEditors ...client.RequestEditorFn) (*CancelImportResponse, error)

	// GetErrors request
	GetErrors(ctx context.Context, importId int64, params *GetErrorsParams, reqEditors ...client.RequestEditorFn) (*GetErrorsResponse, error)
}

type ListResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *CollectionResponsePublicImportResponse
}

// Status returns HTTPResponse.Status
func (r ListResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r ListResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type CreateResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *PublicImportResponse
}

// Status returns HTTPResponse.Status
func (r CreateResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r CreateResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetImportResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *PublicImportResponse
}

// Status returns HTTPResponse.Status
func (r GetImportResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetImportResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type CancelImportResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *ActionResponse
}

// Status returns HTTPResponse.Status
func (r CancelImportResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r CancelImportResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetErrorsResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *CollectionResponsePublicImportErrorForwardPaging
}

// Status returns HTTPResponse.Status
func (r GetErrorsResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetErrorsResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// List request returning *ListResponse
func (c *Client) List(ctx context.Context, params *ListParams, reqEditors ...client.RequestEditorFn) (*ListResponse, error) {
	rsp, err := c.doList(ctx, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return parseListResponse(rsp)
}

// CreateWithBody request with arbitrary body returning *CreateResponse
func (c *Client) CreateWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...client.RequestEditorFn) (*CreateResponse, error) {
	rsp, err := c.doCreateWithBody(ctx, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return parseCreateResponse(rsp)
}

// GetImport request returning *GetImportResponse
func (c *Client) GetImport(ctx context.Context, importId int64, reqEditors ...client.RequestEditorFn) (*GetImportResponse, error) {
	rsp, err := c.doGetImport(ctx, importId, reqEditors...)
	if err != nil {
		return nil, err
	}
	return parseGetImportResponse(rsp)
}

// CancelImport request returning *CancelImportResponse
func (c *Client) CancelImport(ctx context.Context, importId int64, reqEditors ...client.RequestEditorFn) (*CancelImportResponse, error) {
	rsp, err := c.doCancelImport(ctx, importId, reqEditors...)
	if err != nil {
		return nil, err
	}
	return parseCancelImportResponse(rsp)
}

// GetErrors request returning *GetErrorsResponse
func (c *Client) GetErrors(ctx context.Context, importId int64, params *GetErrorsParams, reqEditors ...client.RequestEditorFn) (*GetErrorsResponse, error) {
	rsp, err := c.doGetErrors(ctx, importId, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return parseGetErrorsResponse(rsp)
}

// parseListResponse parses an HTTP response from a List call.
func parseListResponse(rsp *http.Response) (*ListResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &ListResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest CollectionResponsePublicImportResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest
	}

	return response, nil
}

// parseCreateResponse parses an HTTP response from a Create call.
func parseCreateResponse(rsp *http.Response) (*CreateResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &CreateResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest PublicImportResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest
	}

	return response, nil
}

// parseGetImportResponse parses an HTTP response from a GetImport call.
func parseGetImportResponse(rsp *http.Response) (*GetImportResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetImportResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest PublicImportResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest
	}

	return response, nil
}

// parseCancelImportResponse parses an HTTP response from a CancelImport call.
func parseCancelImportResponse(rsp *http.Response) (*CancelImportResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &CancelImportResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest ActionResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest
	}

	return response, nil
}

// parseGetErrorsResponse parses an HTTP response from a GetErrors call.
func parseGetErrorsResponse(rsp *http.Response) (*GetErrorsResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetErrorsResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest CollectionResponsePublicImportErrorForwardPaging
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest
	}

	return response, nil
}
