// Package visitoridentification provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/faetools/devtool version (devel) DO NOT EDIT.
package visitoridentification

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/faetools/client"
)

// operation paths

var opPathGenerateTokenCreate = client.MustParseURL("./conversations/v3/visitor-identification/tokens/create")

// ClientInterface interface specification for the client.
type ClientInterface interface {
	// GenerateTokenCreate request with any body
	GenerateTokenCreateWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...client.RequestEditorFn) (*GenerateTokenCreateResponse, error)
	GenerateTokenCreate(ctx context.Context, body GenerateTokenCreateJSONRequestBody, reqEditors ...client.RequestEditorFn) (*GenerateTokenCreateResponse, error)
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

// GenerateTokenCreate: POST /conversations/v3/visitor-identification/tokens/create

type GenerateTokenCreateResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *IdentificationTokenResponse
}

// Status returns HTTPResponse.Status
func (r GenerateTokenCreateResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GenerateTokenCreateResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// newGenerateTokenCreateRequestWithBody generates requests for GenerateTokenCreate with any type of body
func newGenerateTokenCreateRequestWithBody(baseURL *url.URL, contentType string, body io.Reader) (*http.Request, error) {
	queryURL := baseURL.ResolveReference(opPathGenerateTokenCreate)

	req, err := http.NewRequest(http.MethodPost, queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add(client.ContentType, contentType)

	return req, nil
}

// GenerateTokenCreateWithBody request with arbitrary body returning *GenerateTokenCreateResponse
func (c *Client) GenerateTokenCreateWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...client.RequestEditorFn) (*GenerateTokenCreateResponse, error) {
	rsp, err := c.doGenerateTokenCreateWithBody(ctx, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}

	return parseGenerateTokenCreateResponse(rsp)
}

func (c *Client) doGenerateTokenCreateWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...client.RequestEditorFn) (*http.Response, error) {
	req, err := newGenerateTokenCreateRequestWithBody(c.BaseURL, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}

	return c.Client.Do(req)
}

func (c *Client) GenerateTokenCreate(ctx context.Context, body GenerateTokenCreateJSONRequestBody, reqEditors ...client.RequestEditorFn) (*GenerateTokenCreateResponse, error) {
	rsp, err := c.doGenerateTokenCreate(ctx, body, reqEditors...)
	if err != nil {
		return nil, err
	}

	return parseGenerateTokenCreateResponse(rsp)
}

// newGenerateTokenCreateRequest calls the generic GenerateTokenCreate builder with application/json body.
func newGenerateTokenCreateRequest(baseURL *url.URL, body GenerateTokenCreateJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return newGenerateTokenCreateRequestWithBody(baseURL, client.MIMEApplicationJSON, bodyReader)
}

func (c *Client) doGenerateTokenCreate(ctx context.Context, body GenerateTokenCreateJSONRequestBody, reqEditors ...client.RequestEditorFn) (*http.Response, error) {
	req, err := newGenerateTokenCreateRequest(c.BaseURL, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}

	return c.Client.Do(req)
}

// parseGenerateTokenCreateResponse parses an HTTP response from a GenerateTokenCreate call.
func parseGenerateTokenCreateResponse(rsp *http.Response) (*GenerateTokenCreateResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	if err != nil {
		return nil, err
	}
	defer rsp.Body.Close()

	response := &GenerateTokenCreateResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest IdentificationTokenResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest
	}

	return response, nil
}
