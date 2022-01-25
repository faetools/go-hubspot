// Package schemas provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/faetools/devtool version (devel) DO NOT EDIT.
package schemas

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

// ClientOption allows setting custom parameters during construction.
type ClientOption func(*Client) error

func (c *Client) doGetAllSchemas(ctx context.Context, params *GetAllSchemasParams, reqEditors ...client.RequestEditorFn) (*http.Response, error) {
	req, err := newGetAllSchemasRequest(c.Server, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.client.Do(req)
}

func (c *Client) doCreateSchemasWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...client.RequestEditorFn) (*http.Response, error) {
	req, err := newCreateSchemasRequestWithBody(c.Server, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.client.Do(req)
}

func (c *Client) doCreateSchemas(ctx context.Context, body CreateSchemasJSONRequestBody, reqEditors ...client.RequestEditorFn) (*http.Response, error) {
	req, err := newCreateSchemasRequest(c.Server, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.client.Do(req)
}

func (c *Client) doArchiveObjectType(ctx context.Context, objectType string, params *ArchiveObjectTypeParams, reqEditors ...client.RequestEditorFn) (*http.Response, error) {
	req, err := newArchiveObjectTypeRequest(c.Server, objectType, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.client.Do(req)
}

func (c *Client) doGetObjectType(ctx context.Context, objectType string, reqEditors ...client.RequestEditorFn) (*http.Response, error) {
	req, err := newGetObjectTypeRequest(c.Server, objectType)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.client.Do(req)
}

func (c *Client) doUpdateObjectTypeWithBody(ctx context.Context, objectType string, contentType string, body io.Reader, reqEditors ...client.RequestEditorFn) (*http.Response, error) {
	req, err := newUpdateObjectTypeRequestWithBody(c.Server, objectType, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.client.Do(req)
}

func (c *Client) doUpdateObjectType(ctx context.Context, objectType string, body UpdateObjectTypeJSONRequestBody, reqEditors ...client.RequestEditorFn) (*http.Response, error) {
	req, err := newUpdateObjectTypeRequest(c.Server, objectType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.client.Do(req)
}

func (c *Client) doCreateAssociationAssociationsWithBody(ctx context.Context, objectType string, contentType string, body io.Reader, reqEditors ...client.RequestEditorFn) (*http.Response, error) {
	req, err := newCreateAssociationAssociationsRequestWithBody(c.Server, objectType, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.client.Do(req)
}

func (c *Client) doCreateAssociationAssociations(ctx context.Context, objectType string, body CreateAssociationAssociationsJSONRequestBody, reqEditors ...client.RequestEditorFn) (*http.Response, error) {
	req, err := newCreateAssociationAssociationsRequest(c.Server, objectType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.client.Do(req)
}

func (c *Client) doArchiveAssociationAssociationIdentifier(ctx context.Context, objectType string, associationIdentifier string, reqEditors ...client.RequestEditorFn) (*http.Response, error) {
	req, err := newArchiveAssociationAssociationIdentifierRequest(c.Server, objectType, associationIdentifier)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.client.Do(req)
}

func (c *Client) doPurgeObjectType(ctx context.Context, objectType string, reqEditors ...client.RequestEditorFn) (*http.Response, error) {
	req, err := newPurgeObjectTypeRequest(c.Server, objectType)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.client.Do(req)
}

// newGetAllSchemasRequest generates requests for GetAllSchemas
func newGetAllSchemasRequest(server string, params *GetAllSchemasParams) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/crm/v3/schemas")
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

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// newCreateSchemasRequest calls the generic CreateSchemas builder with application/json body.
func newCreateSchemasRequest(server string, body CreateSchemasJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return newCreateSchemasRequestWithBody(server, "application/json", bodyReader)
}

// newCreateSchemasRequestWithBody generates requests for CreateSchemas with any type of body
func newCreateSchemasRequestWithBody(server string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/crm/v3/schemas")
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

// newArchiveObjectTypeRequest generates requests for ArchiveObjectType
func newArchiveObjectTypeRequest(server string, objectType string, params *ArchiveObjectTypeParams) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "objectType", runtime.ParamLocationPath, objectType)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/crm/v3/schemas/%s", pathParam0)
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

	req, err := http.NewRequest("DELETE", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// newGetObjectTypeRequest generates requests for GetObjectType
func newGetObjectTypeRequest(server string, objectType string) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "objectType", runtime.ParamLocationPath, objectType)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/crm/v3/schemas/%s", pathParam0)
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

// newUpdateObjectTypeRequest calls the generic UpdateObjectType builder with application/json body.
func newUpdateObjectTypeRequest(server string, objectType string, body UpdateObjectTypeJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return newUpdateObjectTypeRequestWithBody(server, objectType, "application/json", bodyReader)
}

// newUpdateObjectTypeRequestWithBody generates requests for UpdateObjectType with any type of body
func newUpdateObjectTypeRequestWithBody(server string, objectType string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "objectType", runtime.ParamLocationPath, objectType)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/crm/v3/schemas/%s", pathParam0)
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

// newCreateAssociationAssociationsRequest calls the generic CreateAssociationAssociations builder with application/json body.
func newCreateAssociationAssociationsRequest(server string, objectType string, body CreateAssociationAssociationsJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return newCreateAssociationAssociationsRequestWithBody(server, objectType, "application/json", bodyReader)
}

// newCreateAssociationAssociationsRequestWithBody generates requests for CreateAssociationAssociations with any type of body
func newCreateAssociationAssociationsRequestWithBody(server string, objectType string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "objectType", runtime.ParamLocationPath, objectType)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/crm/v3/schemas/%s/associations", pathParam0)
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

// newArchiveAssociationAssociationIdentifierRequest generates requests for ArchiveAssociationAssociationIdentifier
func newArchiveAssociationAssociationIdentifierRequest(server string, objectType string, associationIdentifier string) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "objectType", runtime.ParamLocationPath, objectType)
	if err != nil {
		return nil, err
	}

	var pathParam1 string

	pathParam1, err = runtime.StyleParamWithLocation("simple", false, "associationIdentifier", runtime.ParamLocationPath, associationIdentifier)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/crm/v3/schemas/%s/associations/%s", pathParam0, pathParam1)
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

// newPurgeObjectTypeRequest generates requests for PurgeObjectType
func newPurgeObjectTypeRequest(server string, objectType string) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "objectType", runtime.ParamLocationPath, objectType)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/crm/v3/schemas/%s/purge", pathParam0)
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
	Server string

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
	if client.client == nil {
		client.client = &http.Client{}
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
	client.Client
	// GetAllSchemas request
	GetAllSchemas(ctx context.Context, params *GetAllSchemasParams, reqEditors ...client.RequestEditorFn) (*GetAllSchemasResponse, error)

	// CreateSchemas request with any body
	CreateSchemasWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...client.RequestEditorFn) (*CreateSchemasResponse, error)
	CreateSchemas(ctx context.Context, body CreateSchemasJSONRequestBody, reqEditors ...client.RequestEditorFn) (*CreateSchemasResponse, error)

	// ArchiveObjectType request
	ArchiveObjectType(ctx context.Context, objectType string, params *ArchiveObjectTypeParams, reqEditors ...client.RequestEditorFn) (*ArchiveObjectTypeResponse, error)

	// GetObjectType request
	GetObjectType(ctx context.Context, objectType string, reqEditors ...client.RequestEditorFn) (*GetObjectTypeResponse, error)

	// UpdateObjectType request with any body
	UpdateObjectTypeWithBody(ctx context.Context, objectType string, contentType string, body io.Reader, reqEditors ...client.RequestEditorFn) (*UpdateObjectTypeResponse, error)
	UpdateObjectType(ctx context.Context, objectType string, body UpdateObjectTypeJSONRequestBody, reqEditors ...client.RequestEditorFn) (*UpdateObjectTypeResponse, error)

	// CreateAssociationAssociations request with any body
	CreateAssociationAssociationsWithBody(ctx context.Context, objectType string, contentType string, body io.Reader, reqEditors ...client.RequestEditorFn) (*CreateAssociationAssociationsResponse, error)
	CreateAssociationAssociations(ctx context.Context, objectType string, body CreateAssociationAssociationsJSONRequestBody, reqEditors ...client.RequestEditorFn) (*CreateAssociationAssociationsResponse, error)

	// ArchiveAssociationAssociationIdentifier request
	ArchiveAssociationAssociationIdentifier(ctx context.Context, objectType string, associationIdentifier string, reqEditors ...client.RequestEditorFn) (*ArchiveAssociationAssociationIdentifierResponse, error)

	// PurgeObjectType request
	PurgeObjectType(ctx context.Context, objectType string, reqEditors ...client.RequestEditorFn) (*PurgeObjectTypeResponse, error)
}

type GetAllSchemasResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *CollectionResponseObjectSchemaNoPaging
}

// Status returns HTTPResponse.Status
func (r GetAllSchemasResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetAllSchemasResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type CreateSchemasResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON201      *ObjectSchema
}

// Status returns HTTPResponse.Status
func (r CreateSchemasResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r CreateSchemasResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type ArchiveObjectTypeResponse struct {
	Body         []byte
	HTTPResponse *http.Response
}

// Status returns HTTPResponse.Status
func (r ArchiveObjectTypeResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r ArchiveObjectTypeResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetObjectTypeResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *ObjectSchema
}

// Status returns HTTPResponse.Status
func (r GetObjectTypeResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetObjectTypeResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type UpdateObjectTypeResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *ObjectTypeDefinition
}

// Status returns HTTPResponse.Status
func (r UpdateObjectTypeResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r UpdateObjectTypeResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type CreateAssociationAssociationsResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON201      *AssociationDefinition
}

// Status returns HTTPResponse.Status
func (r CreateAssociationAssociationsResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r CreateAssociationAssociationsResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type ArchiveAssociationAssociationIdentifierResponse struct {
	Body         []byte
	HTTPResponse *http.Response
}

// Status returns HTTPResponse.Status
func (r ArchiveAssociationAssociationIdentifierResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r ArchiveAssociationAssociationIdentifierResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type PurgeObjectTypeResponse struct {
	Body         []byte
	HTTPResponse *http.Response
}

// Status returns HTTPResponse.Status
func (r PurgeObjectTypeResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r PurgeObjectTypeResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// GetAllSchemas request returning *GetAllSchemasResponse
func (c *Client) GetAllSchemas(ctx context.Context, params *GetAllSchemasParams, reqEditors ...client.RequestEditorFn) (*GetAllSchemasResponse, error) {
	rsp, err := c.doGetAllSchemas(ctx, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return parseGetAllSchemasResponse(rsp)
}

// CreateSchemasWithBody request with arbitrary body returning *CreateSchemasResponse
func (c *Client) CreateSchemasWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...client.RequestEditorFn) (*CreateSchemasResponse, error) {
	rsp, err := c.doCreateSchemasWithBody(ctx, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return parseCreateSchemasResponse(rsp)
}

func (c *Client) CreateSchemas(ctx context.Context, body CreateSchemasJSONRequestBody, reqEditors ...client.RequestEditorFn) (*CreateSchemasResponse, error) {
	rsp, err := c.doCreateSchemas(ctx, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return parseCreateSchemasResponse(rsp)
}

// ArchiveObjectType request returning *ArchiveObjectTypeResponse
func (c *Client) ArchiveObjectType(ctx context.Context, objectType string, params *ArchiveObjectTypeParams, reqEditors ...client.RequestEditorFn) (*ArchiveObjectTypeResponse, error) {
	rsp, err := c.doArchiveObjectType(ctx, objectType, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return parseArchiveObjectTypeResponse(rsp)
}

// GetObjectType request returning *GetObjectTypeResponse
func (c *Client) GetObjectType(ctx context.Context, objectType string, reqEditors ...client.RequestEditorFn) (*GetObjectTypeResponse, error) {
	rsp, err := c.doGetObjectType(ctx, objectType, reqEditors...)
	if err != nil {
		return nil, err
	}
	return parseGetObjectTypeResponse(rsp)
}

// UpdateObjectTypeWithBody request with arbitrary body returning *UpdateObjectTypeResponse
func (c *Client) UpdateObjectTypeWithBody(ctx context.Context, objectType string, contentType string, body io.Reader, reqEditors ...client.RequestEditorFn) (*UpdateObjectTypeResponse, error) {
	rsp, err := c.doUpdateObjectTypeWithBody(ctx, objectType, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return parseUpdateObjectTypeResponse(rsp)
}

func (c *Client) UpdateObjectType(ctx context.Context, objectType string, body UpdateObjectTypeJSONRequestBody, reqEditors ...client.RequestEditorFn) (*UpdateObjectTypeResponse, error) {
	rsp, err := c.doUpdateObjectType(ctx, objectType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return parseUpdateObjectTypeResponse(rsp)
}

// CreateAssociationAssociationsWithBody request with arbitrary body returning *CreateAssociationAssociationsResponse
func (c *Client) CreateAssociationAssociationsWithBody(ctx context.Context, objectType string, contentType string, body io.Reader, reqEditors ...client.RequestEditorFn) (*CreateAssociationAssociationsResponse, error) {
	rsp, err := c.doCreateAssociationAssociationsWithBody(ctx, objectType, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return parseCreateAssociationAssociationsResponse(rsp)
}

func (c *Client) CreateAssociationAssociations(ctx context.Context, objectType string, body CreateAssociationAssociationsJSONRequestBody, reqEditors ...client.RequestEditorFn) (*CreateAssociationAssociationsResponse, error) {
	rsp, err := c.doCreateAssociationAssociations(ctx, objectType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return parseCreateAssociationAssociationsResponse(rsp)
}

// ArchiveAssociationAssociationIdentifier request returning *ArchiveAssociationAssociationIdentifierResponse
func (c *Client) ArchiveAssociationAssociationIdentifier(ctx context.Context, objectType string, associationIdentifier string, reqEditors ...client.RequestEditorFn) (*ArchiveAssociationAssociationIdentifierResponse, error) {
	rsp, err := c.doArchiveAssociationAssociationIdentifier(ctx, objectType, associationIdentifier, reqEditors...)
	if err != nil {
		return nil, err
	}
	return parseArchiveAssociationAssociationIdentifierResponse(rsp)
}

// PurgeObjectType request returning *PurgeObjectTypeResponse
func (c *Client) PurgeObjectType(ctx context.Context, objectType string, reqEditors ...client.RequestEditorFn) (*PurgeObjectTypeResponse, error) {
	rsp, err := c.doPurgeObjectType(ctx, objectType, reqEditors...)
	if err != nil {
		return nil, err
	}
	return parsePurgeObjectTypeResponse(rsp)
}

// parseGetAllSchemasResponse parses an HTTP response from a GetAllSchemas call.
func parseGetAllSchemasResponse(rsp *http.Response) (*GetAllSchemasResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetAllSchemasResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest CollectionResponseObjectSchemaNoPaging
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest
	}

	return response, nil
}

// parseCreateSchemasResponse parses an HTTP response from a CreateSchemas call.
func parseCreateSchemasResponse(rsp *http.Response) (*CreateSchemasResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &CreateSchemasResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 201:
		var dest ObjectSchema
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON201 = &dest
	}

	return response, nil
}

// parseArchiveObjectTypeResponse parses an HTTP response from a ArchiveObjectType call.
func parseArchiveObjectTypeResponse(rsp *http.Response) (*ArchiveObjectTypeResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &ArchiveObjectTypeResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	return response, nil
}

// parseGetObjectTypeResponse parses an HTTP response from a GetObjectType call.
func parseGetObjectTypeResponse(rsp *http.Response) (*GetObjectTypeResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetObjectTypeResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest ObjectSchema
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest
	}

	return response, nil
}

// parseUpdateObjectTypeResponse parses an HTTP response from a UpdateObjectType call.
func parseUpdateObjectTypeResponse(rsp *http.Response) (*UpdateObjectTypeResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &UpdateObjectTypeResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest ObjectTypeDefinition
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest
	}

	return response, nil
}

// parseCreateAssociationAssociationsResponse parses an HTTP response from a CreateAssociationAssociations call.
func parseCreateAssociationAssociationsResponse(rsp *http.Response) (*CreateAssociationAssociationsResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &CreateAssociationAssociationsResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 201:
		var dest AssociationDefinition
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON201 = &dest
	}

	return response, nil
}

// parseArchiveAssociationAssociationIdentifierResponse parses an HTTP response from a ArchiveAssociationAssociationIdentifier call.
func parseArchiveAssociationAssociationIdentifierResponse(rsp *http.Response) (*ArchiveAssociationAssociationIdentifierResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &ArchiveAssociationAssociationIdentifierResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	return response, nil
}

// parsePurgeObjectTypeResponse parses an HTTP response from a PurgeObjectType call.
func parsePurgeObjectTypeResponse(rsp *http.Response) (*PurgeObjectTypeResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &PurgeObjectTypeResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	return response, nil
}
