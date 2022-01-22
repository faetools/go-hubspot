package service

// Definition defines the service.
type Definition struct {
	HTTPS *HTTPS `json:"https,omitempty"`
	Tier  string `json:"tier,omitempty"`

	// ProductionTraffic signifies that the service is production ready and stricter rules apply. It is `false` by default for all new services and must be set to `true` manually once all the criteria for production-ready traffic are met.
	ProductionTraffic bool `json:"productionTraffic,omitempty"`
}

// IsHTTPService returns whether or not the service has HTTP endpoints.
func (d *Definition) IsHTTPService() bool {
	return d != nil && d.HTTPS.isEnabled()
}

func (h *HTTPS) isEnabled() bool {
	return h != nil && h.Enabled
}

// IsProductionTraffic returns whether or not the service is production ready.
func (d *Definition) IsProductionTraffic() bool {
	return d != nil && d.ProductionTraffic
}

var allTiers = []string{
	"Undefined",
	"Tier 1 - Critical",
	"Tier 2 - Important",
	"Tier 3 - Ancillary",
	"Tier 4 - Internal",
}

// HTTPS defines how the.
type HTTPS struct {
	Enabled bool `json:"enabled,omitempty"`
	// Type               string `json:"type,omitempty"`
	// AllowedPayloadSize int    `json:"allowed_payload_size,omitempty"`
	// RateLimitSecond    int    `json:"rate_limit_second,omitempty"`
	// TimeoutSecond      int    `json:"timeout_second,omitempty"`.
}
