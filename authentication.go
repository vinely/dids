package dids

// The rules for Authentication are:

// A DID Document MAY include an authentication property.
// The value of the authentication property should be an array of verification methods.
// Each verification method MAY be embedded or referenced. An example of a verification method is a public key (see Section § 6.3 Public Keys).

// Authentication - authentication of did document
type Authentication interface {
}

// BaseAuthentication - simplest authentication is one line of id
type BaseAuthentication ID

// BaseAuthentication2 - authentication node
type BaseAuthentication2 struct {
	DIDNode
}
