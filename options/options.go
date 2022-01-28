package options

import (
	"context"
	"net/http"

	"github.com/faetools/client"
)

const hapikey = "hapikey"

// WithAPIKey adds a given hubspot API key to requests.
func WithAPIKey(apiKey string) client.Option {
	return client.WithRequestEditorFn(
		func(_ context.Context, req *http.Request) error {
			q := req.URL.Query()
			q.Add(hapikey, apiKey)
			req.URL.RawQuery = q.Encode()
			return nil
		})
}
