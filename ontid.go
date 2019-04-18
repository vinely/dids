package dids

import (
	"errors"

	"github.com/vinely/ontchain/ontsdk"
)

const (
	// OntMethod - method for this type of DID
	OntMethod = "ont"
)

// OntID - DID for Ontology
type OntID struct {
	s *DIDScheme
}

func init() {
	dt := &DIDType{
		Method: OntMethod,
		Info:   "DID for Ontology",
		New:    new,
		Build:  ontid,
	}
	RegisterDIDType(dt)
}

// Constructors

func ontid(scheme *DIDScheme) (DID, error) {
	return &OntID{scheme}, nil
}

// NewDID - simplest way to construct a DID
func new() (DID, error) {
	id, err := ontsdk.GenerateID()
	if err != nil {
		return nil, err
	}
	s := &DIDScheme{
		Method: OntMethod,
		ID:     id,
	}
	return &OntID{s}, nil
}

//type DID interface {
// 	Scheme() *DIDScheme
// 	String() string
// 	VerifyID() error
// }

// Scheme - DID interface
func (o *OntID) Scheme() *DIDScheme {
	return o.s
}

// String - DID interface
func (o *OntID) String() string {
	return o.s.String()
}

// VerifyID - DID interface
func (o *OntID) VerifyID() error {
	if o.s.ID == "" {
		return errors.New("empty ontid")
	}
	if !ontsdk.VerifyID(o.s.ID) {
		return errors.New("invalid id :" + o.s.ID)
	}
	return nil
}
