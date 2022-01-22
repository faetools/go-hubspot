package repo

import (
	"github.com/faetools/devtool/internal/repo/library"
	"github.com/faetools/devtool/internal/repo/service"
)

// Definition defines the repository.
type Definition struct {
	Name           string              `json:"name"`
	Version        string              `json:"version"`
	RepoType       Type                `json:"repoType"`
	Service        *service.Definition `json:"service,omitempty"`
	Library        *library.Definition `json:"library,omitempty"`
	Team           string              `json:"team"`
	DevToolVersion string              `json:"devToolVersion"`
	Coverage       float64             `json:"coverage,omitempty"`
	QualityGate    string              `json:"qualityGate,omitempty"`
}

var allTeams = []string{
	"firestarters",
}

// IsHTTPService returns whether or not the repo defines a service with HTTP endpoints.
func IsHTTPService() bool { return Load().IsHTTPService() }

// IsHTTPService returns whether or not the repo defines a service with HTTP endpoints.
func (d Definition) IsHTTPService() bool { return d.Service.IsHTTPService() }

// IsPublic returns whether or not the repo is public.
func (d Definition) IsPublic() bool { return d.IsThirdPartyAPI() }

// IsThirdPartyAPI returns whether or not the repo is a third party API.
func IsThirdPartyAPI() bool { return Load().IsThirdPartyAPI() }

// IsThirdPartyAPI returns whether or not the repo is a third party API.
func (d Definition) IsThirdPartyAPI() bool { return d.Library.IsThirdPartyAPI() }

// IsProductionTraffic returns whether or not the repo defines a service that is production ready.
func (d Definition) IsProductionTraffic() bool { return d.Service.IsProductionTraffic() }

// IsMicroservice returns whether or not the repo is a microservice.
func IsMicroservice() bool { return Load().IsMicroservice() }

// IsMicroservice returns whether or not the repo is a microservice.
func (d Definition) IsMicroservice() bool { return d.RepoType == Microservice }

// IsProgram returns whether or not the repo is buildable.
func IsProgram() bool { return Load().IsProgram() }

// IsProgram returns whether or not the repo is buildable.
func (d Definition) IsProgram() bool {
	switch d.RepoType {
	case GoLibrary, Other:
		return false
	default:
		return true
	}
}

// UsesGo returns whether or not the repo uses go.
func UsesGo() bool { return Load().UsesGo() }

// UsesGo returns whether or not the repo uses go.
func (d Definition) UsesGo() bool { return d.RepoType != Other }
