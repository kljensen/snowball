package english

import (
	"github.com/kljensen/snowball/stemword"
)

// Search for the the following suffixes, and,
// if found, perform the action indicated. 
// 
// e
// delete if in R2, or in R1 and not preceded by a short syllable
// l
// delete if in R2 and preceded by l
func step5(w *stemword.Word) bool {

	// Last rune index = `lri`
	lri := len(w.RS) - 1

	// If R1 is emtpy, R2 is also empty, and we
	// need not do anything in step 5.
	if w.R1start > lri {
		return false
	}

	if w.RS[lri] == 101 {

		// Delete "e" suffix if in R2, or in R1 and not preceded
		// by a short syllable.
		if w.R2start <= lri || !endsShortSyllable(w, lri) {
			w.ReplaceSuffix("e", "", true)
			return true
		}

	} else if w.R2start <= lri && w.RS[lri] == 108 {

		// Delete "l" suffix if in R2 and preceded by "l"
		// l = 108
		w.ReplaceSuffix("l", "", true)
		return true

	}
	return false
}
