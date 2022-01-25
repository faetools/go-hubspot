# go-hubspot

![Devtool version](https://img.shields.io/badge/Devtool-0.0.1-brightgreen.svg)
![Maintainer](https://img.shields.io/badge/team-firestarters-blue)
[![Go Report Card](https://goreportcard.com/badge/github.com/faetools/go-hubspot)](https://goreportcard.com/report/github.com/faetools/go-hubspot)
[![GoDoc Reference](https://img.shields.io/badge/godoc-reference-blue.svg)](https://pkg.go.dev/github.com/faetools/go-hubspot)

## Example Usage

```go
package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/faetools/go-hubspot/pkg/crm/contacts"
)

func addAPIKey(ctx context.Context, req *http.Request) error {
	q := req.URL.Query()
	q.Add("hapikey", os.Getenv("HUBSPOT_API_KEY"))
	req.URL.RawQuery = q.Encode()
	return nil
}

func main() {
	ctx := context.Background()

	c, err := contacts.NewClient(contacts.WithRequestEditorFn(addAPIKey))
	checkErr(err)

	resp, err := c.ListContacts(ctx, &contacts.ListContactsParams{
		Properties: &contacts.PropertyKeys,
	})
	checkErr(err)

	if resp.StatusCode() != http.StatusOK {
		fmt.Printf("Got status %q with body: %s\n", resp.Status(), string(resp.Body))
		os.Exit(1)
	}

	for _, res := range resp.JSON200.Results {
		fmt.Printf("%s %s (%s)\n", res.FirstName(), res.LastName(), res.Email())
	}
}

func checkErr(err error) {
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}
```

## API Key

You can get your API key at https://app.hubspot.com/api-key/12345/call-log, with `12345` being your account number.
