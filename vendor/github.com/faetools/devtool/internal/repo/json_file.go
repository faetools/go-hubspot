package repo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/AlecAivazis/survey/v2"
	"github.com/spf13/cobra"
)

var cache *Definition

// infoName returns the name of the info file.
// For faetools or personal repos, it's "info.json", otherwise it's hidden.
func infoName() string {
	wd, err := os.Getwd()
	cobra.CheckErr(err)
	wd = strings.ToLower(wd)

	if strings.Contains(wd, "faetools") || strings.Contains(wd, "markrosemaker") {
		return "info.json"
	}

	return ".info.json"
}

var mu = sync.Mutex{}

// Load returns the repo definition.
func Load() *Definition {
	mu.Lock()
	defer mu.Unlock()

	if cache == nil {
		RefreshCache()
	}

	return cache
}

// RefreshCache refreshes the repo definition cache.
func RefreshCache() {
	b, err := os.ReadFile(infoName())
	if err != nil {
		if !os.IsNotExist(err) {
			cobra.CheckErr(err)
		}

		cont := false
		err := survey.AskOne(&survey.Confirm{
			Default: true,
			Message: fmt.Sprintf("%s could not be found. Would you like to create one?", infoName()),
		}, &cont)
		cobra.CheckErr(err)

		if !cont {
			os.Exit(0)
		}

		Write(AskUser())
		return
	}

	cache = &Definition{}
	cobra.CheckErr(json.Unmarshal(b, cache))

	if cache.Service != nil && cache.Service.HTTPS != nil &&
		!cache.Service.HTTPS.Enabled {
		cache.Service.HTTPS = nil
	}

	switch {
	case cache.RepoType != Microservice:
		cache.Service = nil
	}
}

// Write writes a definition file according to the definition.
func Write(def *Definition) {
	b := &bytes.Buffer{}
	enc := json.NewEncoder(b)
	enc.SetEscapeHTML(false)
	enc.SetIndent("", "  ")

	cobra.CheckErr(enc.Encode(def))

	// Save the generated definition file.
	loc := infoName()
	_ = os.MkdirAll(filepath.Dir(loc), os.ModePerm)
	cobra.CheckErr(os.WriteFile(loc, b.Bytes(), 0o600))
	cache = def
}
