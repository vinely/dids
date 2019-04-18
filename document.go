package dids

// A DID points to a DID Document.
// DID Documents are the serialization of the § 4. Data Model.
// The following sections define the properties of the DID Document,
// including whether these properties are required or optional.

// DIDDocment -
type DIDDocment struct {
	Context   string `json:"@context,omitempty"`
	ID        string `json:"id,omitempty"`
	PublicKey string `json:"publicKey,omitempty"`
	// Authentication string
	// Service        string
}
