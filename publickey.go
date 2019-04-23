package dids

// The rules for public keys are:
// A DID Document MAY include a publicKey property.
// The value of the publicKey property MUST be an array of public keys.
// Each public key MUST include id and type properties, and exactly one value property.
// The array of public keys SHOULD NOT contain duplicate entries with the same id and different value properties with different formats.
// Each public key MUST include a controller property, which identifies the controller of the corresponding private key.
// The value property of a public key MUST be exactly one of publicKeyPem, publicKeyJwk, publicKeyHex, publicKeyBase64, publicKeyBase58, publicKeyMultibase, depending on the format and encoding of the public key.
// A registry of key types and formats is available in Appendix § A. Registries.

// PublicKey - interface of did document publickey
type PublicKey interface {
	Valid() bool
}

// PublicKeyValue - Value is a string
// The value property of a public key MUST be exactly one of publicKeyPem, publicKeyJwk, publicKeyHex, publicKeyBase64, publicKeyBase58, publicKeyMultibase, depending on the format and encoding of the public key.
type PublicKeyValue string

// PublicKeyPem -
type PublicKeyPem struct {
	Value PublicKeyValue `json:"publicKeyPem"`
}

// PublicKeyJwk -
type PublicKeyJwk struct {
	Value PublicKeyValue `json:"publicKeyJwk"`
}

// PublicKeyHex -
type PublicKeyHex struct {
	Value PublicKeyValue `json:"publicKeyHex"`
}

// PublicKeyBase64 -
type PublicKeyBase64 struct {
	Value PublicKeyValue `json:"publicKeyBase64"`
}

// PublicKeyBase58 -
type PublicKeyBase58 struct {
	Value PublicKeyValue `json:"publicKeyBase58"`
}

// PublicKeyMultibase -
type PublicKeyMultibase struct {
	Value PublicKeyValue `json:"publicKeyMultibase"`
}

// BasePublicKey - Base public key
type BasePublicKey struct {
	DIDNode
	Controller ID `json:"controller"` // identity of controller
}

// examples from https://w3c-ccg.github.io/ld-cryptosuite-registry/

// https://w3c-ccg.github.io/ld-cryptosuite-registry/#Ed25519Signature2018
// Summary
// Identifiers	Ed25519Signature2018, Ed25519VerificationKey2018
// Status	PROVISIONAL
// Authors	Markus Sabadello
// Specification	Ed25519 Signature Suite 2018

// Ed25519Key -
type Ed25519Key struct {
	BasePublicKey
	PublicKeyBase58
}

const (
	// Ed25519KeyType - Ed25519Key Type Category
	Ed25519KeyType Type = "Ed25519Key"
	// Ed25519Signature2018 -
	Ed25519Signature2018 Type = "Ed25519Signature2018"
	// Ed25519VerificationKey2018 -
	Ed25519VerificationKey2018 Type = "Ed25519VerificationKey2018"
)

func init() {
	RegisterType(Ed25519Signature2018, Ed25519KeyType)
	RegisterType(Ed25519VerificationKey2018, Ed25519KeyType)
}

// https://w3c-ccg.github.io/ld-cryptosuite-registry/#RsaSignature2018

// Summary
// Identifiers	RsaSignature2018, RsaVerificationKey2018
// Status	PROVISIONAL
// Authors	Dave Longley, Manu Sporny
// Specification	RSA Signature Suite 2018

// RsaKey -
type RsaKey struct {
	BasePublicKey
	PublicKeyPem
}

const (
	// RsaKeyType -
	RsaKeyType Type = "RsaKey"
	// RsaSignature2018 -
	RsaSignature2018 Type = "RsaSignature2018"
	// RsaVerificationKey2018 -
	RsaVerificationKey2018 Type = "RsaVerificationKey2018"
)

func init() {
	RegisterType(RsaSignature2018, RsaKeyType)
	RegisterType(RsaVerificationKey2018, RsaKeyType)
}

// https://w3c-ccg.github.io/ld-cryptosuite-registry/#EdDsaSASignatureSecp256k1

// Summary
// Identifiers	EdDsaSASignatureSecp256k1, EdDsaSAPublicKeySecp256k1
// Status	PROVISIONAL
// Authors	Harlan Wood, Manu Sporny
// Specification	Koblitz Signature Suite 2016

// EdDsaSAKey -
type EdDsaSAKey struct {
	BasePublicKey
	PublicKeyHex
}

const (
	// EdDsaSAKeyType -
	EdDsaSAKeyType Type = "EdDsaSAKey"
	// EdDsaSASignatureSecp256k1 -
	EdDsaSASignatureSecp256k1 Type = "EdDsaSASignatureSecp256k1"
	// EdDsaSAPublicKeySecp256k1 -
	EdDsaSAPublicKeySecp256k1 Type = "EdDsaSAPublicKeySecp256k1"
)

func init() {
	RegisterType(EdDsaSASignatureSecp256k1, EdDsaSAKeyType)
	RegisterType(EdDsaSAPublicKeySecp256k1, EdDsaSAKeyType)
}
