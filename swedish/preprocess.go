package swedish

import (
	"github.com/kljensen/snowball/snowballword"
)

// Applies various transformations necessary for the
// other, subsequent stemming steps.  Most important
// of which is defining the two regions R1 & R2.
//
func preprocess(word *snowballword.SnowballWord) {

	// Clean up apostrophes
	// normalizeApostrophes(word)
	// trimLeftApostrophes(word)

	// Find the region R1. R2 is not used
	r1start := r1(word)
	word.R1start = r1start
}
