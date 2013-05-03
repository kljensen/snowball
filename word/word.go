/*
	This package implements various string-like methods
	over slices of runes.  It is similar to the
	exp/utf8string package.
*/
package runeslice

type Word struct {
	rs      []rune
	r1start int
	r2start int
}

func New(in string) (word *Word) {
	word = &Word{rs: []rune(in)}
	return
}

func (w *Word) String() string {
	return string(w.rs)
}

// Returns the first prefix found or the empty string.
func (w *Word) FirstPrefix(prefixes ...string) string {
	found := false
	rsLen := len(w.rs)

	for _, prefix := range prefixes {
		found = true
		for i, r := range prefix {
			if i > rsLen-1 || (w.rs)[i] != r {
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
func (w *Word) FirstSuffix(sufficies ...string) (suffix string) {
	rsLen := len(w.rs)
	for _, suffix := range sufficies {
		numMatching := 0
		suffixRunes := []rune(suffix)
		suffixLen := len(suffixRunes)
		for i := 0; i < rsLen && i < suffixLen; i++ {
			if w.rs[rsLen-i-1] != suffixRunes[suffixLen-i-1] {
				break
			} else {
				numMatching += 1
			}
		}
		if numMatching == suffixLen {
			return suffix
		}
	}
	return ""
}
