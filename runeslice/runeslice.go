/*
	This package implements various string-like methods
	over slices of runes.  It is similar to the
	exp/utf8string package.
*/
package runeslice

import (
	"log"
)

type RuneSlice []rune

func (rs *RuneSlice) String() string {
	return string(*rs)
}

// Returns the first prefix found or the empty string.
func (rs *RuneSlice) FirstPrefix(prefixes ...string) string {
	found := false
	rsLen := len(*rs)

	for _, prefix := range prefixes {
		found = true
		for i, r := range prefix {
			if i > rsLen-1 || (*rs)[i] != r {
				found = false
				break
			}
		}
		if found {
			return prefix
		}
	}
	return ""
}

// Returns the first suffix found or the empty string.
func (rs *RuneSlice) FirstSuffix(sufficies ...string) (suffix string) {
	rsLen := len(*rs)
	for _, suffix := range sufficies {
		numMatching := 0
		suffixRunes := []rune(suffix)
		suffixLen := len(suffixRunes)
		for i := 0; i < rsLen && i < suffixLen; i++ {
			if (*rs)[rsLen-i-1] != suffixRunes[suffixLen-i-1] {
				break
			} else {
				numMatching += 1
			}
		}
		log.Println(suffix, numMatching)
		if numMatching == suffixLen {
			return suffix
		}
	}
	return ""
}
