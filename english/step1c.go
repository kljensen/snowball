package english

import (
	"github.com/kljensen/snowball/stemword"
)

// Replace suffix y or Y by i if preceded by a non-vowel which is not
// the first letter of the word (so cry -> cri, by -> by, say -> say)
//
func step1c(w *stemword.Word) bool {

	rsLen := len(w.RS)
	// y = 121
	// Y = 89
	// i = 105
	if len(w.RS) > 2 && (w.RS[rsLen-1] == 121 || w.RS[rsLen-1] == 89) && !isLowerVowel(w.RS[rsLen-2]) {
		w.RS[rsLen-1] = 105
		return true
	}
	return false
}
