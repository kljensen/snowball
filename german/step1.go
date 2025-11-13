package german

import (
	"github.com/kljensen/snowball/snowballword"
)

// Step 1: First do block - Remove common suffixes in R1
// Removes: erinnen, erin, em, ern, er, en, es, e, s, ln, lns
func step1(word *snowballword.SnowballWord) {

	suffix := word.FirstSuffixIfIn(word.R1start, len(word.RS),
		"erinnen", "erin", "ern", "er", "em", "en", "es", "e", "lns", "ln", "s",
	)

	switch suffix {
	case "erinnen", "erin", "ern", "er":
		// Conflate female versions of nouns
		word.RemoveLastNRunes(len(suffix))
		return

	case "em":
		// Don't remove -em from words ending -system
		idx := len(word.RS) - 2
		if idx >= 4 {
			// Check if preceded by "syst"
			if word.RS[idx-3] == 115 && // s
				word.RS[idx-2] == 121 && // y
				word.RS[idx-1] == 115 && // s
				word.RS[idx] == 116 { // t
				return
			}
		}
		word.RemoveLastNRunes(2)
		return

	case "en", "es", "e":
		word.RemoveLastNRunes(len(suffix))
		// Try to remove 's' from 'nis' suffix
		if len(word.RS) >= 3 {
			idx := len(word.RS)
			if word.RS[idx-1] == 115 && // s
				word.RS[idx-2] == 105 && // i
				word.RS[idx-3] == 110 { // n
				word.RemoveLastNRunes(1)
			}
		}
		return

	case "s":
		// Valid s-ending: must be preceded by bdfghklmnrt
		if isValidSEnding(word, len(word.RS)-1) {
			word.RemoveLastNRunes(1)
		}
		return

	case "lns":
		// Replace lns with l
		word.RemoveLastNRunes(3)
		word.RS = append(word.RS, 108) // l
		return

	case "ln":
		// Replace ln with l
		word.RemoveLastNRunes(2)
		word.RS = append(word.RS, 108) // l
		return
	}
}
