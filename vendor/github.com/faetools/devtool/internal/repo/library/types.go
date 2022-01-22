package library

// Definition defines the library.
type Definition struct {
	SubType SubType `json:"subType"`
}

// IsThirdPartyAPI returns whether or not the library is a third party API.
func (d *Definition) IsThirdPartyAPI() bool {
	return d != nil && d.SubType == ThirdPartyAPI
}
