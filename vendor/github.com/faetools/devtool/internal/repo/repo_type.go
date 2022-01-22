package repo

import (
	"sort"
)

//go:generate jsonenums -type=Type

// Type is a type to repository.
type Type int

const (
	// Microservice is a repository that defines a microservice.
	Microservice Type = iota
	// App is a repository that defines an app.
	App
	// GoLibrary is a repository that defines.
	GoLibrary
	// DygoSite is a repository that defines a dygo site.
	DygoSite
	// HugoModule is a repository that defines a hugo module.
	HugoModule
	// HugoSite is a repository that defines a hugo site.
	HugoSite
	// Other is any other repository.
	Other
)

var allTypes []string

func init() {
	for t := range _TypeNameToValue {
		allTypes = append(allTypes, t)
	}

	sort.Slice(allTypes, func(i, j int) bool {
		return _TypeNameToValue[allTypes[i]] < _TypeNameToValue[allTypes[j]]
	})
}
