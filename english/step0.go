package english

import (
	"github.com/kljensen/snowball/stemword"
)

// Step 0 is to strip off apostrophes and "s".
//
func step0(w *stemword.Word) bool {
	replaced := false
	var step0Suffixes = [3]string{"'s'", "'s", "'"}
	for _, suffix := range step0Suffixes {
		replaced = w.ReplaceSuffix(suffix, "", false)
		if replaced {
			return true
		}
	}
	return false
}
