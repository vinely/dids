package dids

// DIDProperty - interface of DID nodes
type DIDProperty interface {
	Valid() bool
}

// DID interface for DIDs
type DID interface {
	ID() *ID
	Scheme() *DIDScheme
	Type() *DIDType
}

// DIDNode - did element
type DIDNode struct {
	ID   `json:"id"`
	Type `json:"type,omitempty"`
}
