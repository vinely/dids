package ont

import (
	"crypto/rand"
	"encoding/base64"

	"github.com/vinely/dids"
	chain "github.com/vinely/ontchain"
)

// Document - my type of document
type Document struct {
	dids.DIDNode
	PublicKey []PublicKey `json:"publicKey,omitempty"`
}

func createPassword() string {
	var buf [32]byte
	_, err := rand.Read(buf[:])
	if err != nil {
		return ""
	}
	passwd := make([]byte, base64.StdEncoding.EncodedLen(len(buf)))
	base64.StdEncoding.Encode(passwd, buf[:])
	return string(buf[:])
}

// New - create an document . also return an identity for new
func New() (*Document, *chain.ManagedIdentity, error) {
	id, err := dids.CreateID(OntMethod)
	if err != nil {
		return nil, nil, err
	}

	return DocumentFromID(id.ID())
}

// DocumentFromID - create document from ID(DIDs)
func DocumentFromID(id *dids.ID) (*Document, *chain.ManagedIdentity, error) {
	passwd := createPassword()

	identity, err := chain.GetIdentityFromID(string(*id), []byte(passwd))
	if err != nil {
		return nil, nil, err
	}
	doc, err := DocumentFromIdentity(identity)
	if err != nil {
		return nil, nil, err
	}
	return doc, identity, nil
}

// DocumentFromIdentity - create document from managed identity
func DocumentFromIdentity(id *chain.ManagedIdentity) (*Document, error) {
	pk, err := getPublicKeyFromSDK(&id.Identity, 1)
	if err != nil {
		return nil, err
	}
	doc := &Document{}
	doc.ID = dids.ID(id.ID)
	doc.PublicKey = []PublicKey{*pk}
	return doc, nil
}
