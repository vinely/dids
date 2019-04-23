package ont

import (
	"errors"

	"github.com/vinely/dids"
	"github.com/vinely/ontchain/ontsdk"
)

const (
	// OntMethod - method for this type of DID
	OntMethod = "ont"
)

// ID - DID for Ontology
type ID struct {
	s *dids.DIDScheme
}

func init() {
	dt := &dids.DIDType{
		Method: OntMethod,
		Info:   "DID for Ontology",
		New:    new,
		Build:  ontid,
	}
	dids.RegisterDIDType(dt)
}

// Constructors

func ontid(scheme *dids.DIDScheme) (dids.DID, error) {
	return &ID{scheme}, nil
}

// NewDID - simplest way to construct a DID
func new() (dids.DID, error) {
	id, err := ontsdk.GenerateID()
	if err != nil {
		return nil, err
	}
	s, err := dids.Parse(id)
	if err != nil {
		return nil, err
	}
	return s, nil
}

//type DID interface {
// 	Scheme() *DIDScheme
// 	String() string
// 	VerifyID() error
// }

// Scheme - DID interface
func (o *ID) Scheme() *dids.DIDScheme {
	return o.s
}

// String - DID interface
func (o *ID) String() string {
	return o.s.String()
}

// VerifyID - DID interface
func (o *ID) VerifyID() error {
	if o.s.ID == "" {
		return errors.New("empty ontid")
	}
	if !ontsdk.VerifyID(o.s.String()) {
		return errors.New("invalid id :" + o.s.ID)
	}
	return nil
}
