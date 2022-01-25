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
	// parameter styles
	form   = "form"
	simple = "simple"

	contentType = "Content-Type"
	json        = "json"

	// MIMEApplicationJSON defines standard type "application/json"
	MIMEApplicationJSON = "application/json"

	// ContentType is the header key to define the type of content.
	ContentType = "Content-Type"
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

// GetPathParam returns the path parameter value.
func GetPathParam(paramName string, value interface{}) (string, error) {
	return runtime.StyleParamWithLocation(simple, false, paramName, runtime.ParamLocationPath, value)
}
