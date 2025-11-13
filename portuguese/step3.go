package portuguese

import (
	"github.com/kljensen/snowball/snowballword"
)

// Step 3: Residual suffix removal
func step3(word *snowballword.SnowballWord) {
	suffix := word.FirstSuffixIfIn(word.RVstart, len(word.RS),
		"os", "a", "i", "o", "á", "í", "ó",
	)

	if suffix != "" {
		word.RemoveLastNRunes(len([]rune(suffix)))
	}
}
