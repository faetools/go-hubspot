package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/faetools/devtool/cgtools"
	"github.com/pkg/errors"
	"golang.org/x/sync/errgroup"

	"github.com/spf13/afero"
)

// NOTE: The code below is just the code to generate the openapis.

const (
	apiCatalogEndpoint = "https://api.hubspot.com/api-catalog-public/v1/apis"

	stagePreview = "DEVELOPER_PREVIEW"
	stageLatest  = "LATEST"

	mode = 0o644
)

func main() {
	ctx := context.Background()
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, apiCatalogEndpoint, nil)
	checkErr(err)

	resp, err := http.DefaultClient.Do(req)
	checkErr(err)
	defer resp.Body.Close()

	r := response{}
	checkErr(json.NewDecoder(resp.Body).Decode(&r))

	g := cgtools.Generator{Fs: afero.NewOsFs()}

	eg := errgroup.Group{}

	for _, res := range r.Results {
		res := res
	FEATURES:
		for featureName, f := range res.Features {
			featureName, f := featureName, f

			switch f.Stage {
			case stageLatest: // what we want
			case stagePreview: // skip
				continue FEATURES
			default:
				fmt.Printf("unknown stage %s for feature %s of %s\n",
					f.Stage, featureName, res.Name)
			}

			eg.Go(func() error {
				return addAPI(ctx, g, res.Name, featureName, f)
			})
		}
	}

	checkErr(eg.Wait())
	fmt.Println("Done.")
}

func addAPI(ctx context.Context, g cgtools.Generator,
	name, featureName string, f feature) error {
	errExpl := fmt.Sprintf("getting feature %s of %s", featureName, name)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, f.OpenAPI, nil)
	if err != nil {
		return errors.Wrap(err, errExpl)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return errors.Wrap(err, errExpl)
	}
	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return errors.Wrap(err, errExpl)
	}

	resName := strings.Title(strings.ToLower(name))

	filename := fmt.Sprintf("api/%s/%s/openapi.json", resName, featureName)
	if resName == featureName {
		filename = fmt.Sprintf("api/%s/openapi.json", resName)
	}

	return errors.Wrap(g.WriteBytes(filename, b, mode), errExpl)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

type response struct {
	Results results `json:"results"`
}

type results []result

type result struct {
	Name     string             `json:"name"`
	Features map[string]feature `json:"features,omitempty"`
}

type feature struct {
	OpenAPI string `json:"openAPI"`
	Stage   string `json:"stage"`
}
