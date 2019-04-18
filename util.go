package dids

import "strings"

// some util functions

// split string with separator string "sep"
// return two part of string
func split(src string, sep string) (string, string) {
	i := strings.Index(src, sep)
	if i < 0 {
		return src, ""
	}
	return src[:i], src[i+len(sep):]
}

func stringContainsInvalidByte(s string) bool {
	for i := 0; i < len(s); i++ {
		b := s[i]
		if b < ' ' || b >= 0x7f {
			return true
		}
	}
	return false
}
