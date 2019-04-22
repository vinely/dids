package dids

// The rules for service endpoints are:

// A DID Document MAY include a service property.
// The value of the service property should be an array of service endpoints.
// Each service endpoint must include id, type, and serviceEndpoint properties, and MAY include additional properties.
// The service endpoint protocol SHOULD be published in an open standard specification.
// The value of the serviceEndpoint property MUST be a JSON-LD object or a valid URI conforming to [RFC3986]
// and normalized according to the rules in section 6 of [RFC3986] and to any normalization rules in its applicable URI scheme specification.

// Service - service of did document
type Service interface {
}

// BaseService - service endpoint
type BaseService struct {
	DIDNode
	ServiceEndpoint string `json:"serviceEndpoint"`
	// and MAY include additional properties.
}
