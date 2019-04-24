package ont

import (
	"github.com/vinely/dids"
	"github.com/vinely/ontchain/ontsdk"
)

const (
	// OntMethod - method for this type of DID
	OntMethod = "ont"
)

var (
	didType *dids.DIDType
)

// DID - DID for Ontology
type DID struct {
	id *dids.ID
	t  *dids.DIDType
	s  *dids.DIDScheme
}

func init() {
	didType := &dids.DIDType{
		Method: OntMethod,
		Info:   "DID for Ontology",
		New:    NewDID,
		Parse:  ParseDID,
	}
	dids.RegisterDIDType(didType)
}

// Constructors

// ParseDID - parse of this type
func ParseDID(uri string) (dids.DID, error) {
	s, err := dids.ParseScheme(uri)
	if err != nil {
		return nil, err
	}
	i := dids.ID(uri)
	return &DID{
		id: &i,
		t:  didType,
		s:  s,
	}, nil
}

// NewDID - simplest way to construct a DID
func NewDID() (dids.DID, error) {
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

// type DID interface {
// 	ID() *ID
// 	Scheme() *DIDScheme
// 	Type() *DIDType
// }

// Scheme - DID interface
func (o *DID) Scheme() *dids.DIDScheme {
	return o.s
}

// ID - DID interface
func (o *DID) ID() *dids.ID {
	return o.id
}

// Type - DID interface
func (o *DID) Type() *dids.DIDType {
	return o.t
}
