package dids

// DIDProperty - interface of DID nodes
type DIDProperty interface {
	Valid() bool
}

// DIDNode - did element
type DIDNode struct {
	ID   `json:"id"`
	Type `json:"type,omitempty"`
}
