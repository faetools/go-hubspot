package client

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/deepmap/oapi-codegen/pkg/runtime"
)

const (
	form        = "form" // parameter style
	contentType = "Content-Type"
	json        = "json"
)

// RequestEditorFn  is the function signature for the RequestEditor callback function
type RequestEditorFn func(ctx context.Context, req *http.Request) error

// HTTPRequestDoer performs HTTP requests.
//
// The standard http.Client implements this interface.
type HTTPRequestDoer interface {
	Do(req *http.Request) (*http.Response, error)
}

// Client is an interface for all clients generated using the faetools go client library generator.
type Client interface {
	// SetClient sets the underlying client.
	SetClient(HTTPRequestDoer)

	// AddRequestEditor adds a request editor to the client.
	AddRequestEditor(RequestEditorFn)

	// SetBaseURL overrides the baseURL.
	SetBaseURL(*url.URL)
}

// Option allows setting custom parameters during construction.
type Option func(Client) error

// WithHTTPClient allows overriding the default Doer, which is
// automatically created using http.Client. This is useful for tests.
func WithHTTPClient(doer HTTPRequestDoer) Option {
	return func(c Client) error {
		c.SetClient(doer)
		return nil
	}
}

// WithRequestEditorFn allows setting up a callback function, which will be
// called right before sending the request. This can be used to mutate the request.
func WithRequestEditorFn(fn RequestEditorFn) Option {
	return func(c Client) error {
		c.AddRequestEditor(fn)
		return nil
	}
}

// WithBaseURL overrides the baseURL.
func WithBaseURL(baseURL string) Option {
	return func(c Client) error {
		// ensure the server URL always has a trailing slash
		if !strings.HasSuffix(baseURL, "/") {
			baseURL += "/"
		}

		newBaseURL, err := url.Parse(baseURL)
		if err != nil {
			return err
		}

		c.SetBaseURL(newBaseURL)
		return nil
	}
}

// MustParseURL parses the URL. It panics if there is an error.
func MustParseURL(rawURL string) *url.URL {
	u, err := url.Parse(rawURL)
	if err != nil {
		panic(fmt.Sprintf("OpenAPI contains unparsable operation path: %s", err))
	}

	return u
}

// AddQueryParam adds a certain parameter with its value to the query.
func AddQueryParam(query url.Values, paramName string, value interface{}) error {
	if value == nil {
		return nil
	}

	queryFrag, err := runtime.StyleParamWithLocation(form, true, paramName, runtime.ParamLocationQuery, value)
	if err != nil {
		return err
	}

	parsed, err := url.ParseQuery(queryFrag)
	if err != nil {
		return err
	}

	for k, v := range parsed {
		for _, v2 := range v {
			query.Add(k, v2)
		}
	}

	return nil
}

// IsJSONResponseWithCode checks whether the response is a JSON response with the given code.
func IsJSONResponseWithCode(rsp *http.Response, code int) bool {
	return rsp.StatusCode == code && strings.Contains(rsp.Header.Get(contentType), json)
}

// TODO:
// - make static paths constant
// operationPath := fmt.Sprintf("/crm/v3/objects/contacts")
// 	if operationPath[0] == '/' {
// 		operationPath = "." + operationPath
// 	}
// - replace serverURL.Parse with serverURL.ResolveReference
// - maybe even keep queryURL on hand
// - method to http.Method
