package dids

import "errors"

// Type - type of element
type Type string

var typeList = make(map[Type][]Type)

// CheckCategory - check if type belongs to category
func (t *Type) CheckCategory(cat Type) bool {
	chain, ok := typeList[*t]
	if !ok {
		return false
	}
	if cat == "" {
		return true
	}
	for _, c := range chain {
		if c == cat {
			return true
		}
	}
	return false
}

// Valid - check Type's validation in system
func (t *Type) Valid() bool {
	return t.CheckCategory("")
}

// RegisterType - register a type to system
func RegisterType(t ...Type) error {
	if len(t) == 0 || t[0] == "" {
		return errors.New("empty parameter")
	}
	typeList[t[0]] = t[1:]
	return nil
}

// CancelType - cancel a type
func CancelType(t Type) {
	delete(typeList, t)
}
