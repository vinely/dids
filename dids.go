package dids

import (
	"errors"
	"strings"
)

const (
	// DIDHeader  header of DIDScheme
	DIDHeader = "did"
	separator = ":"
)

// ID - string of DID
type ID string

// Valid - check id validity
func (id *ID) Valid() bool {
	_, err := Parse(string(*id))
	if err != nil {
		return false
	}
	return true
}

// Scheme - parse id to didshceme
func (id *ID) Scheme() (*DIDScheme, error) {
	uri := string(*id)
	return ParseScheme(uri)
}

// String - id string
func (id *ID) String() string {
	return string(*id)
}

// DIDScheme Decentralized Identifiers
// The generic DID scheme is a URI scheme conformant with [RFC3986].
// It consists of a DID followed by an optional path and/or fragment.
// The term DID refers only to the identifier conforming to the did rule in the ABNF below;
// when used alone, it does not include a path or fragment.
// A DID that may optionally include a path and/or fragment is called a DID reference.
// did-reference      = did [ "/" did-path ] [ "#" did-fragment ]
// did                = "did:" method ":" specific-idstring
// method             = 1*methodchar
// methodchar         = %x61-7A / DIGIT
// specific-idstring  = idstring *( ":" idstring )
// idstring           = 1*idchar
// idchar             = ALPHA / DIGIT / "." / "-"
//	did:[method:][:idstring]...[/path][?query][#fragment]
type DIDScheme struct {
	Method   string
	ID       string   // major idstring
	IDExt    []string // rest idstrings
	RawPath  string   // encoded path hint (see EscapedPath method)
	Query    string   // encoded query values, without '?'
	Fragment string   // fragment for references, without '#'
}

// GetMethod - Parse method string from uri
func GetMethod(uri string) (string, error) {
	if uri[0:3] != DIDHeader || uri[3:4] != separator {
		return "", errors.New("not a did uri")
	}
	rest := uri[4:]
	method, _ := split(rest, separator)
	return method, nil
}

// ParseScheme - parse uri string to DIDScheme
func ParseScheme(uri string) (*DIDScheme, error) {
	if uri[0:3] != DIDHeader || uri[3:4] != separator {
		return nil, errors.New("not a did uri")
	}
	rest := uri[4:]
	if stringContainsInvalidByte(rest) {
		return nil, errors.New("contain invalid charactor")
	}
	did := &DIDScheme{}
	rest, did.Fragment = split(rest, "#")
	rest, did.Query = split(rest, "?")
	rest, did.RawPath = split(rest, "/")
	strs := strings.Split(rest, separator)
	if len(strs) < 2 {
		return nil, errors.New("missing method or idstring")
	}
	did.Method = strs[0]
	did.ID = strs[1]
	did.IDExt = strs[2:]
	return did, nil
}

// Base - Uri short string of DID
func (did *DIDScheme) Base() string {
	var buf strings.Builder

	buf.WriteString(DIDHeader)
	buf.WriteString(separator)
	buf.WriteString(did.Method)
	buf.WriteString(separator)
	buf.WriteString(did.ID)
	return buf.String()
}

// String - Uri string of DID
func (did *DIDScheme) String() string {
	var buf strings.Builder

	buf.WriteString(DIDHeader)
	buf.WriteString(separator)
	buf.WriteString(did.Method)
	buf.WriteString(separator)
	buf.WriteString(did.ID)

	for _, s := range did.IDExt {
		buf.WriteString(separator)
		buf.WriteString(s)
	}
	if did.RawPath != "" {
		buf.WriteString("/")
		buf.WriteString(did.RawPath)
	}
	if did.Query != "" {
		buf.WriteString("?")
		buf.WriteString(did.Query)
	}
	if did.Fragment != "" {
		buf.WriteString("#")
		buf.WriteString(did.Fragment)
	}

	return buf.String()
}

var (
	didTypes = make(map[string]DIDType)
)

// DIDType -
type DIDType struct {
	Method string                        // method in did is the type sign of did
	Info   string                        // information description of did type
	New    func() (DID, error)           // New a did
	Parse  func(uri string) (DID, error) // create did from scheme
}

// RegisterDIDType - Register a new DID type
func RegisterDIDType(d *DIDType) error {
	_, ok := didTypes[d.Method]
	if ok {
		return errors.New("DID type :" + d.Method + " already existed")
	}
	didTypes[d.Method] = *d
	return nil
}

// RevokeDIDType - Revoke a new DID type
func RevokeDIDType(method string) {
	delete(didTypes, method)
}

// GetDIDTypeFromMethod -
func GetDIDTypeFromMethod(method string) (*DIDType, error) {
	dt, ok := didTypes[method]
	if !ok {
		return nil, errors.New("did not have this method: " + method)
	}
	return &dt, nil
}

// Parse - parse uri string to DID struct
func Parse(uri string) (DID, error) {
	method, err := GetMethod(uri)
	if err != nil {
		return nil, err
	}
	dt, err := GetDIDTypeFromMethod(method)
	if err != nil {
		return nil, err
	}
	did, err := dt.Parse(uri)
	if err != nil {
		return nil, err
	}
	return did, nil
}

// CreateID - convenience way to create a specifical DID
func CreateID(method string) (DID, error) {
	dt, err := GetDIDTypeFromMethod(method)
	if err != nil {
		return nil, err
	}
	did, err := dt.New()
	if err != nil {
		return nil, err
	}
	return did, nil
}
