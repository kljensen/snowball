package german

import (
	"github.com/kljensen/snowball/snowballword"
)

// Step 2: Second do block - Remove additional suffixes in R1
// Removes: est, en, er, st, et (with special conditions)
func step2(word *snowballword.SnowballWord) {

	suffix := word.FirstSuffixIfIn(word.R1start, len(word.RS),
		"est", "st", "en", "er", "et",
	)

	switch suffix {
	case "est", "en", "er":
		word.RemoveLastNRunes(len(suffix))
		return

	case "st":
		// Must be preceded by valid st-ending and have at least 3 letters before
		if isValidStEnding(word, len(word.RS)-2) {
			word.RemoveLastNRunes(2)
		}
		return

	case "et":
		// Must be preceded by valid et-ending character
		idx := len(word.RS) - 2
		if !isValidEtEnding(word, idx) {
			return
		}

		// Check for exceptions: geordn, intern, plan, tick
		// Note: 'tr' is also in the list but with comment "Still conflate"
		// which means tr should be removed
		if idx >= 5 {
			// Check for "geordn" (the 6 chars before 'et')
			if word.RS[idx-5] == 103 && // g
				word.RS[idx-4] == 101 && // e
				word.RS[idx-3] == 111 && // o
				word.RS[idx-2] == 114 && // r
				word.RS[idx-1] == 100 && // d
				word.RS[idx] == 110 { // n
				return
			}
			// Check for "intern" (the 6 chars before 'et')
			if word.RS[idx-5] == 105 && // i
				word.RS[idx-4] == 110 && // n
				word.RS[idx-3] == 116 && // t
				word.RS[idx-2] == 101 && // e
				word.RS[idx-1] == 114 && // r
				word.RS[idx] == 110 { // n
				return
			}
		}
		if idx >= 3 {
			// Check for "plan" (the 4 chars before 'et')
			if word.RS[idx-3] == 112 && // p
				word.RS[idx-2] == 108 && // l
				word.RS[idx-1] == 97 && // a
				word.RS[idx] == 110 { // n
				return
			}
			// Check for "tick" (the 4 chars before 'et')
			if word.RS[idx-3] == 116 && // t
				word.RS[idx-2] == 105 && // i
				word.RS[idx-1] == 99 && // c
				word.RS[idx] == 107 { // k
				return
			}
		}

		word.RemoveLastNRunes(2)
		return
	}
}
