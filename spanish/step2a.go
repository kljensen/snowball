package spanish

import (
	"unicode/utf8"

	"github.com/kljensen/snowball/snowballword"
)

// Step 2a is the removal of verb suffixes beginning y,
// Search for the longest among the following suffixes
// in RV, and if found, delete if preceded by u.
func step2a(word *snowballword.SnowballWord) bool {
	suffix := word.FirstSuffixIn(word.RVstart, len(word.RS), "ya", "ye", "yan", "yen", "yeron", "yendo", "yo", "yó", "yas", "yes", "yais", "yamos")
	if suffix != "" {
		suffixLength := utf8.RuneCountInString(suffix)
		idx := len(word.RS) - suffixLength - 1
		if idx >= 0 && word.RS[idx] == 117 {
			word.RemoveLastNRunes(suffixLength)
			return true
		}
	}
	return false
}
