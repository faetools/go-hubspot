package client

import (
	"net/url"
	"strings"
)

// Option allows setting custom parameters during construction.
type Option func(*Client) error

// WithHTTPClient allows overriding the default Doer, which is
// automatically created using http.Client. This is useful for tests.
func WithHTTPClient(doer HTTPRequestDoer) Option {
	return func(c *Client) error {
		c.Client = doer
		return nil
	}
}

// WithRequestEditorFn allows setting up a callback function, which will be
// called right before sending the request. This can be used to mutate the request.
func WithRequestEditorFn(fn RequestEditorFn) Option {
	return func(c *Client) error {
		c.RequestEditors = append(c.RequestEditors, fn)
		return nil
	}
}

// WithBaseURL overrides the baseURL.
func WithBaseURL(baseURL string) Option {
	return func(c *Client) error {
		// ensure the server URL always has a trailing slash
		if !strings.HasSuffix(baseURL, "/") {
			baseURL += "/"
		}

		newBaseURL, err := url.Parse(baseURL)
		if err != nil {
			return err
		}

		c.BaseURL = newBaseURL
		return nil
	}
}
