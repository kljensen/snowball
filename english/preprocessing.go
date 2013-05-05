package english

import (
	"github.com/kljensen/snowball/stemword"
)

// Takes an `inputWord` and applies various transformations
// necessary for the other, subsequent stemming steps.
//
func preprocess(word *stemword.Word) {
	normalizeApostrophes(word)
	trimLeftApostrophes(word)
	capitalizeYs(word)
}
