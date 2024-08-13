package english

import (
	"unicode/utf8"

	"github.com/kljensen/snowball/snowballword"
)

// Step 0 is to strip off apostrophes and "s".
func step0(w *snowballword.SnowballWord) bool {
	suffix := w.FirstSuffix("'s'", "'s", "'")
	if suffix == "" {
		return false
	}
	sLen := utf8.RuneCountInString(suffix)
	w.RemoveLastNRunes(sLen)
	return true
}
