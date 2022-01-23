// Package quotes provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/faetools/devtool version (devel) DO NOT EDIT.
package quotes

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

func (c *Client) doGetPageQuotes(ctx context.Context, params *GetPageQuotesParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := newGetPageQuotesRequest(c.Server, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) doReadBatchWithBody(ctx context.Context, params *ReadBatchParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := newReadBatchRequestWithBody(c.Server, params, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) doReadBatch(ctx context.Context, params *ReadBatchParams, body ReadBatchJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := newReadBatchRequest(c.Server, params, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) doDoSearchWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := newDoSearchRequestWithBody(c.Server, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) doDoSearch(ctx context.Context, body DoSearchJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := newDoSearchRequest(c.Server, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) doGetByIdQuote(ctx context.Context, quoteId string, params *GetByIdQuoteParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := newGetByIdQuoteRequest(c.Server, quoteId, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) doGetAllToObjectType(ctx context.Context, quoteId string, toObjectType string, params *GetAllToObjectTypeParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := newGetAllToObjectTypeRequest(c.Server, quoteId, toObjectType, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// newGetPageQuotesRequest generates requests for GetPageQuotes
func newGetPageQuotesRequest(server string, params *GetPageQuotesParams) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/crm/v3/objects/quotes")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	queryValues := queryURL.Query()

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

	if params.Properties != nil {
		if queryFrag, err := runtime.StyleParamWithLocation("form", true, "properties", runtime.ParamLocationQuery, *params.Properties); err != nil {
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

	if params.Associations != nil {
		if queryFrag, err := runtime.StyleParamWithLocation("form", true, "associations", runtime.ParamLocationQuery, *params.Associations); err != nil {
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

	if params.Archived != nil {
		if queryFrag, err := runtime.StyleParamWithLocation("form", true, "archived", runtime.ParamLocationQuery, *params.Archived); err != nil {
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

// newReadBatchRequest calls the generic ReadBatch builder with application/json body.
func newReadBatchRequest(server string, params *ReadBatchParams, body ReadBatchJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return newReadBatchRequestWithBody(server, params, "application/json", bodyReader)
}

// newReadBatchRequestWithBody generates requests for ReadBatch with any type of body
func newReadBatchRequestWithBody(server string, params *ReadBatchParams, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/crm/v3/objects/quotes/batch/read")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	queryValues := queryURL.Query()

	if params.Archived != nil {
		if queryFrag, err := runtime.StyleParamWithLocation("form", true, "archived", runtime.ParamLocationQuery, *params.Archived); err != nil {
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

	req, err := http.NewRequest("POST", queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

	return req, nil
}

// newDoSearchRequest calls the generic DoSearch builder with application/json body.
func newDoSearchRequest(server string, body DoSearchJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return newDoSearchRequestWithBody(server, "application/json", bodyReader)
}

// newDoSearchRequestWithBody generates requests for DoSearch with any type of body
func newDoSearchRequestWithBody(server string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/crm/v3/objects/quotes/search")
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

// newGetByIdQuoteRequest generates requests for GetByIdQuote
func newGetByIdQuoteRequest(server string, quoteId string, params *GetByIdQuoteParams) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "quoteId", runtime.ParamLocationPath, quoteId)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/crm/v3/objects/quotes/%s", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	queryValues := queryURL.Query()

	if params.Properties != nil {
		if queryFrag, err := runtime.StyleParamWithLocation("form", true, "properties", runtime.ParamLocationQuery, *params.Properties); err != nil {
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

	if params.Associations != nil {
		if queryFrag, err := runtime.StyleParamWithLocation("form", true, "associations", runtime.ParamLocationQuery, *params.Associations); err != nil {
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

	if params.Archived != nil {
		if queryFrag, err := runtime.StyleParamWithLocation("form", true, "archived", runtime.ParamLocationQuery, *params.Archived); err != nil {
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

	if params.IdProperty != nil {
		if queryFrag, err := runtime.StyleParamWithLocation("form", true, "idProperty", runtime.ParamLocationQuery, *params.IdProperty); err != nil {
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

// newGetAllToObjectTypeRequest generates requests for GetAllToObjectType
func newGetAllToObjectTypeRequest(server string, quoteId string, toObjectType string, params *GetAllToObjectTypeParams) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "quoteId", runtime.ParamLocationPath, quoteId)
	if err != nil {
		return nil, err
	}

	var pathParam1 string

	pathParam1, err = runtime.StyleParamWithLocation("simple", false, "toObjectType", runtime.ParamLocationPath, toObjectType)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/crm/v3/objects/quotes/%s/associations/%s", pathParam0, pathParam1)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
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
	// GetPageQuotes request
	GetPageQuotes(ctx context.Context, params *GetPageQuotesParams, reqEditors ...RequestEditorFn) (*GetPageQuotesResponse, error)

	// ReadBatch request with any body
	ReadBatchWithBody(ctx context.Context, params *ReadBatchParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*ReadBatchResponse, error)
	ReadBatch(ctx context.Context, params *ReadBatchParams, body ReadBatchJSONRequestBody, reqEditors ...RequestEditorFn) (*ReadBatchResponse, error)

	// DoSearch request with any body
	DoSearchWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*DoSearchResponse, error)
	DoSearch(ctx context.Context, body DoSearchJSONRequestBody, reqEditors ...RequestEditorFn) (*DoSearchResponse, error)

	// GetByIdQuote request
	GetByIdQuote(ctx context.Context, quoteId string, params *GetByIdQuoteParams, reqEditors ...RequestEditorFn) (*GetByIdQuoteResponse, error)

	// GetAllToObjectType request
	GetAllToObjectType(ctx context.Context, quoteId string, toObjectType string, params *GetAllToObjectTypeParams, reqEditors ...RequestEditorFn) (*GetAllToObjectTypeResponse, error)
}

type GetPageQuotesResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *CollectionResponseSimplePublicObjectWithAssociationsForwardPaging
}

// Status returns HTTPResponse.Status
func (r GetPageQuotesResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetPageQuotesResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type ReadBatchResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *BatchResponseSimplePublicObject
	JSON207      *BatchResponseSimplePublicObjectWithErrors
}

// Status returns HTTPResponse.Status
func (r ReadBatchResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r ReadBatchResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type DoSearchResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *CollectionResponseWithTotalSimplePublicObjectForwardPaging
}

// Status returns HTTPResponse.Status
func (r DoSearchResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r DoSearchResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetByIdQuoteResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *SimplePublicObjectWithAssociations
}

// Status returns HTTPResponse.Status
func (r GetByIdQuoteResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetByIdQuoteResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetAllToObjectTypeResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *CollectionResponseAssociatedIdForwardPaging
}

// Status returns HTTPResponse.Status
func (r GetAllToObjectTypeResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetAllToObjectTypeResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// GetPageQuotes request returning *GetPageQuotesResponse
func (c *Client) GetPageQuotes(ctx context.Context, params *GetPageQuotesParams, reqEditors ...RequestEditorFn) (*GetPageQuotesResponse, error) {
	rsp, err := c.doGetPageQuotes(ctx, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return parseGetPageQuotesResponse(rsp)
}

// ReadBatchWithBody request with arbitrary body returning *ReadBatchResponse
func (c *Client) ReadBatchWithBody(ctx context.Context, params *ReadBatchParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*ReadBatchResponse, error) {
	rsp, err := c.doReadBatchWithBody(ctx, params, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return parseReadBatchResponse(rsp)
}

func (c *Client) ReadBatch(ctx context.Context, params *ReadBatchParams, body ReadBatchJSONRequestBody, reqEditors ...RequestEditorFn) (*ReadBatchResponse, error) {
	rsp, err := c.doReadBatch(ctx, params, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return parseReadBatchResponse(rsp)
}

// DoSearchWithBody request with arbitrary body returning *DoSearchResponse
func (c *Client) DoSearchWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*DoSearchResponse, error) {
	rsp, err := c.doDoSearchWithBody(ctx, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return parseDoSearchResponse(rsp)
}

func (c *Client) DoSearch(ctx context.Context, body DoSearchJSONRequestBody, reqEditors ...RequestEditorFn) (*DoSearchResponse, error) {
	rsp, err := c.doDoSearch(ctx, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return parseDoSearchResponse(rsp)
}

// GetByIdQuote request returning *GetByIdQuoteResponse
func (c *Client) GetByIdQuote(ctx context.Context, quoteId string, params *GetByIdQuoteParams, reqEditors ...RequestEditorFn) (*GetByIdQuoteResponse, error) {
	rsp, err := c.doGetByIdQuote(ctx, quoteId, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return parseGetByIdQuoteResponse(rsp)
}

// GetAllToObjectType request returning *GetAllToObjectTypeResponse
func (c *Client) GetAllToObjectType(ctx context.Context, quoteId string, toObjectType string, params *GetAllToObjectTypeParams, reqEditors ...RequestEditorFn) (*GetAllToObjectTypeResponse, error) {
	rsp, err := c.doGetAllToObjectType(ctx, quoteId, toObjectType, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return parseGetAllToObjectTypeResponse(rsp)
}

// parseGetPageQuotesResponse parses an HTTP response from a GetPageQuotes call.
func parseGetPageQuotesResponse(rsp *http.Response) (*GetPageQuotesResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetPageQuotesResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest CollectionResponseSimplePublicObjectWithAssociationsForwardPaging
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest
	}

	return response, nil
}

// parseReadBatchResponse parses an HTTP response from a ReadBatch call.
func parseReadBatchResponse(rsp *http.Response) (*ReadBatchResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &ReadBatchResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest BatchResponseSimplePublicObject
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 207:
		var dest BatchResponseSimplePublicObjectWithErrors
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON207 = &dest

	}

	return response, nil
}

// parseDoSearchResponse parses an HTTP response from a DoSearch call.
func parseDoSearchResponse(rsp *http.Response) (*DoSearchResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &DoSearchResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest CollectionResponseWithTotalSimplePublicObjectForwardPaging
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest
	}

	return response, nil
}

// parseGetByIdQuoteResponse parses an HTTP response from a GetByIdQuote call.
func parseGetByIdQuoteResponse(rsp *http.Response) (*GetByIdQuoteResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetByIdQuoteResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest SimplePublicObjectWithAssociations
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest
	}

	return response, nil
}

// parseGetAllToObjectTypeResponse parses an HTTP response from a GetAllToObjectType call.
func parseGetAllToObjectTypeResponse(rsp *http.Response) (*GetAllToObjectTypeResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetAllToObjectTypeResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest CollectionResponseAssociatedIdForwardPaging
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest
	}

	return response, nil
}
