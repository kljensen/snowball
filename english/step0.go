package english

import (
	"github.com/kljensen/snowball/stemword"
)

// Step 0 is to strip off apostrophes and "s".
//
func step0(w *stemword.Word) bool {
	suffix := w.FirstSuffix("'s'", "'s", "'")
	if suffix == "" {
		return false
	}
	w.ReplaceSuffix(suffix, "", true)
	return true
}
