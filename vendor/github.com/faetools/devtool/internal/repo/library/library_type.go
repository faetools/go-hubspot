package library

import "sort"

//go:generate jsonenums -type=SubType

// SubType is a type of library.
type SubType int

const (
	// ThirdPartyAPI is a library for a third party API.
	ThirdPartyAPI SubType = iota

	// Other is any other repository.
	Other
)

var allTypes []string

func init() {
	for t := range _SubTypeNameToValue {
		allTypes = append(allTypes, t)
	}

	sort.Slice(allTypes, func(i, j int) bool {
		return _SubTypeNameToValue[allTypes[i]] < _SubTypeNameToValue[allTypes[j]]
	})
}
