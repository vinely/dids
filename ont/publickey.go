package ont

import (
	"github.com/vinely/dids"
	sdk "github.com/vinely/ontchain/ontsdk"
)

const (
	// PublicKeyFragmentHeader - publickey label will be fragment of did scheme.
	// The header will be prefix of the fragment
	PublicKeyFragmentHeader = "Key_"
)

// PublicKey for ont
type PublicKey struct {
	dids.BasePublicKey
	index int
	dids.PublicKeyHex
}

func getPublicKeyFromSDK(id *sdk.Identity, i int) (*PublicKey, error) {
	c, err := id.GetControllerDataByIndex(i)
	if err != nil {
		return nil, err
	}
	pk := &PublicKey{}
	pk.ID = dids.ID(id.ID + "#" + PublicKeyFragmentHeader + c.ID)
	pk.Controller = dids.ID(id.ID)
	pk.Type = ""
	pk.index = i
	pk.PublicKeyHex = dids.PublicKeyHex{Value: dids.PublicKeyValue(c.Public)}
	return pk, nil
}

// Key - key value of public key
func (pk *PublicKey) Key() string {
	return string(pk.PublicKeyHex.Value)
}

// Valid - for interface PublicKey
func (pk *PublicKey) Valid() bool {
	// TODO not implement
	return true
}
