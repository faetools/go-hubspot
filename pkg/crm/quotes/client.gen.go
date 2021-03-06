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
	"net/http"
	"net/url"
	"strings"

	"github.com/faetools/client"
)

// operation paths

const (
	opPathGetQuoteFormat           = "./crm/v3/objects/quotes/%s"
	opPathGetAllToObjectTypeFormat = "./crm/v3/objects/quotes/%s/associations/%s"
)

var (
	opPathListQuotes = client.MustParseURL("./crm/v3/objects/quotes")
	opPathReadBatch  = client.MustParseURL("./crm/v3/objects/quotes/batch/read")
	opPathDoSearch   = client.MustParseURL("./crm/v3/objects/quotes/search")
)

// ClientInterface interface specification for the client.
type ClientInterface interface {
	// ListQuotes request
	ListQuotes(ctx context.Context, params *ListQuotesParams, reqEditors ...client.RequestEditorFn) (*ListQuotesResponse, error)

	// ReadBatch request with any body
	ReadBatchWithBody(ctx context.Context, params *ReadBatchParams, contentType string, body io.Reader, reqEditors ...client.RequestEditorFn) (*ReadBatchResponse, error)
	ReadBatch(ctx context.Context, params *ReadBatchParams, body ReadBatchJSONRequestBody, reqEditors ...client.RequestEditorFn) (*ReadBatchResponse, error)

	// DoSearch request with any body
	DoSearchWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...client.RequestEditorFn) (*DoSearchResponse, error)
	DoSearch(ctx context.Context, body DoSearchJSONRequestBody, reqEditors ...client.RequestEditorFn) (*DoSearchResponse, error)

	// GetQuote request
	GetQuote(ctx context.Context, quoteId string, params *GetQuoteParams, reqEditors ...client.RequestEditorFn) (*GetQuoteResponse, error)

	// GetAllToObjectType request
	GetAllToObjectType(ctx context.Context, quoteId string, toObjectType string, params *GetAllToObjectTypeParams, reqEditors ...client.RequestEditorFn) (*GetAllToObjectTypeResponse, error)
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

// ListQuotes: GET /crm/v3/objects/quotes

type ListQuotesResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *CollectionResponseSimplePublicObjectWithAssociationsForwardPaging
}

// Status returns HTTPResponse.Status
func (r ListQuotesResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r ListQuotesResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// newListQuotesRequest generates requests for ListQuotes
func newListQuotesRequest(baseURL *url.URL, params *ListQuotesParams) (*http.Request, error) {
	queryURL := baseURL.ResolveReference(opPathListQuotes)

	q := queryURL.Query()

	if params.Limit != nil {
		if err := client.AddQueryParam(q, "limit", *params.Limit); err != nil {
			return nil, err
		}
	}

	if params.After != nil {
		if err := client.AddQueryParam(q, "after", *params.After); err != nil {
			return nil, err
		}
	}

	if params.Properties != nil {
		if err := client.AddQueryParam(q, "properties", *params.Properties); err != nil {
			return nil, err
		}
	}

	if params.Associations != nil {
		if err := client.AddQueryParam(q, "associations", *params.Associations); err != nil {
			return nil, err
		}
	}

	if params.Archived != nil {
		if err := client.AddQueryParam(q, "archived", *params.Archived); err != nil {
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

// ListQuotes request returning *ListQuotesResponse
func (c *Client) ListQuotes(ctx context.Context, params *ListQuotesParams, reqEditors ...client.RequestEditorFn) (*ListQuotesResponse, error) {
	req, err := newListQuotesRequest(c.BaseURL, params)
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

	response := &ListQuotesResponse{
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

// ReadBatch: POST /crm/v3/objects/quotes/batch/read

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

// newReadBatchRequestWithBody generates requests for ReadBatch with any type of body
func newReadBatchRequestWithBody(baseURL *url.URL, params *ReadBatchParams, contentType string, body io.Reader) (*http.Request, error) {
	queryURL := baseURL.ResolveReference(opPathReadBatch)

	q := queryURL.Query()

	if params.Archived != nil {
		if err := client.AddQueryParam(q, "archived", *params.Archived); err != nil {
			return nil, err
		}
	}

	queryURL.RawQuery = q.Encode()

	req, err := http.NewRequest(http.MethodPost, queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add(client.ContentType, contentType)

	return req, nil
}

// ReadBatchWithBody request with arbitrary body returning *ReadBatchResponse
func (c *Client) ReadBatchWithBody(ctx context.Context, params *ReadBatchParams, contentType string, body io.Reader, reqEditors ...client.RequestEditorFn) (*ReadBatchResponse, error) {
	rsp, err := c.doReadBatchWithBody(ctx, params, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}

	return parseReadBatchResponse(rsp)
}

func (c *Client) doReadBatchWithBody(ctx context.Context, params *ReadBatchParams, contentType string, body io.Reader, reqEditors ...client.RequestEditorFn) (*http.Response, error) {
	req, err := newReadBatchRequestWithBody(c.BaseURL, params, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}

	return c.Client.Do(req)
}

func (c *Client) ReadBatch(ctx context.Context, params *ReadBatchParams, body ReadBatchJSONRequestBody, reqEditors ...client.RequestEditorFn) (*ReadBatchResponse, error) {
	rsp, err := c.doReadBatch(ctx, params, body, reqEditors...)
	if err != nil {
		return nil, err
	}

	return parseReadBatchResponse(rsp)
}

// newReadBatchRequest calls the generic ReadBatch builder with application/json body.
func newReadBatchRequest(baseURL *url.URL, params *ReadBatchParams, body ReadBatchJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return newReadBatchRequestWithBody(baseURL, params, client.MIMEApplicationJSON, bodyReader)
}

func (c *Client) doReadBatch(ctx context.Context, params *ReadBatchParams, body ReadBatchJSONRequestBody, reqEditors ...client.RequestEditorFn) (*http.Response, error) {
	req, err := newReadBatchRequest(c.BaseURL, params, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}

	return c.Client.Do(req)
}

// parseReadBatchResponse parses an HTTP response from a ReadBatch call.
func parseReadBatchResponse(rsp *http.Response) (*ReadBatchResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	if err != nil {
		return nil, err
	}
	defer rsp.Body.Close()

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

// DoSearch: POST /crm/v3/objects/quotes/search

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

// newDoSearchRequestWithBody generates requests for DoSearch with any type of body
func newDoSearchRequestWithBody(baseURL *url.URL, contentType string, body io.Reader) (*http.Request, error) {
	queryURL := baseURL.ResolveReference(opPathDoSearch)

	req, err := http.NewRequest(http.MethodPost, queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add(client.ContentType, contentType)

	return req, nil
}

// DoSearchWithBody request with arbitrary body returning *DoSearchResponse
func (c *Client) DoSearchWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...client.RequestEditorFn) (*DoSearchResponse, error) {
	rsp, err := c.doDoSearchWithBody(ctx, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}

	return parseDoSearchResponse(rsp)
}

func (c *Client) doDoSearchWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...client.RequestEditorFn) (*http.Response, error) {
	req, err := newDoSearchRequestWithBody(c.BaseURL, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}

	return c.Client.Do(req)
}

func (c *Client) DoSearch(ctx context.Context, body DoSearchJSONRequestBody, reqEditors ...client.RequestEditorFn) (*DoSearchResponse, error) {
	rsp, err := c.doDoSearch(ctx, body, reqEditors...)
	if err != nil {
		return nil, err
	}

	return parseDoSearchResponse(rsp)
}

// newDoSearchRequest calls the generic DoSearch builder with application/json body.
func newDoSearchRequest(baseURL *url.URL, body DoSearchJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return newDoSearchRequestWithBody(baseURL, client.MIMEApplicationJSON, bodyReader)
}

func (c *Client) doDoSearch(ctx context.Context, body DoSearchJSONRequestBody, reqEditors ...client.RequestEditorFn) (*http.Response, error) {
	req, err := newDoSearchRequest(c.BaseURL, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}

	return c.Client.Do(req)
}

// parseDoSearchResponse parses an HTTP response from a DoSearch call.
func parseDoSearchResponse(rsp *http.Response) (*DoSearchResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	if err != nil {
		return nil, err
	}
	defer rsp.Body.Close()

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

// GetQuote: GET /crm/v3/objects/quotes/{quoteId}

type GetQuoteResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *SimplePublicObjectWithAssociations
}

// Status returns HTTPResponse.Status
func (r GetQuoteResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetQuoteResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// newGetQuoteRequest generates requests for GetQuote
func newGetQuoteRequest(baseURL *url.URL, quoteId string, params *GetQuoteParams) (*http.Request, error) {
	pathParam0, err := client.GetPathParam("quoteId", quoteId)
	if err != nil {
		return nil, err
	}

	opPath := fmt.Sprintf(opPathGetQuoteFormat, pathParam0)

	queryURL, err := baseURL.Parse(opPath)
	if err != nil {
		return nil, err
	}

	q := queryURL.Query()

	if params.Properties != nil {
		if err := client.AddQueryParam(q, "properties", *params.Properties); err != nil {
			return nil, err
		}
	}

	if params.Associations != nil {
		if err := client.AddQueryParam(q, "associations", *params.Associations); err != nil {
			return nil, err
		}
	}

	if params.Archived != nil {
		if err := client.AddQueryParam(q, "archived", *params.Archived); err != nil {
			return nil, err
		}
	}

	if params.IdProperty != nil {
		if err := client.AddQueryParam(q, "idProperty", *params.IdProperty); err != nil {
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

// GetQuote request returning *GetQuoteResponse
func (c *Client) GetQuote(ctx context.Context, quoteId string, params *GetQuoteParams, reqEditors ...client.RequestEditorFn) (*GetQuoteResponse, error) {
	req, err := newGetQuoteRequest(c.BaseURL, quoteId, params)
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

	response := &GetQuoteResponse{
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

// GetAllToObjectType: GET /crm/v3/objects/quotes/{quoteId}/associations/{toObjectType}

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

// newGetAllToObjectTypeRequest generates requests for GetAllToObjectType
func newGetAllToObjectTypeRequest(baseURL *url.URL, quoteId string, toObjectType string, params *GetAllToObjectTypeParams) (*http.Request, error) {
	pathParam0, err := client.GetPathParam("quoteId", quoteId)
	if err != nil {
		return nil, err
	}

	pathParam1, err := client.GetPathParam("toObjectType", toObjectType)
	if err != nil {
		return nil, err
	}

	opPath := fmt.Sprintf(opPathGetAllToObjectTypeFormat, pathParam0, pathParam1)

	queryURL, err := baseURL.Parse(opPath)
	if err != nil {
		return nil, err
	}

	q := queryURL.Query()

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

// GetAllToObjectType request returning *GetAllToObjectTypeResponse
func (c *Client) GetAllToObjectType(ctx context.Context, quoteId string, toObjectType string, params *GetAllToObjectTypeParams, reqEditors ...client.RequestEditorFn) (*GetAllToObjectTypeResponse, error) {
	req, err := newGetAllToObjectTypeRequest(c.BaseURL, quoteId, toObjectType, params)
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
