package english

import (
	"github.com/kljensen/snowball/stemword"
	// "log"
)

func step1b(w *stemword.Word) bool {

	suffix := w.FirstSuffix("eedly", "ingly", "edly", "ing", "eed", "ed")

	switch suffix {
	case "":
		return false
	case "eed", "eedly":
		// Replace by ee if in R1 
		if len(suffix) <= len(w.RS)-w.R1start {
			w.ReplaceSuffix(suffix, "ee", true)
		}
		return true
	case "ed", "edly", "ing", "ingly":
		hasLowerVowel := false
		for i := 0; i < len(w.RS)-len(suffix); i++ {
			if isLowerVowel(w.RS[i]) {
				hasLowerVowel = true
				break
			}
		}
		if hasLowerVowel {

			originalR1start := w.R1start
			originalR2start := w.R2start

			// Delete if the preceding word part contains a vowel
			w.ReplaceSuffix(suffix, "", true)

			// and after the deletion...
			var (
				newSuffix string
			)

			// If the word ends "at", "bl" or "iz" add "e" 
			newSuffix = w.FirstSuffix("at", "bl", "iz", "bb", "dd", "ff", "gg", "mm", "nn", "pp", "rr", "tt")
			switch newSuffix {
			case "":
				// If the word is short, add "e"
				if isShortWord(w) {
					// By definition, r1 and r2 are the empty string for
					// short words.
					w.RS = append(w.RS, []rune("e")...)
					w.R1start = len(w.RS)
					w.R2start = len(w.RS)
					return true
				}
			case "at", "bl", "iz":
				w.ReplaceSuffix(newSuffix, newSuffix+"e", true)

			case "bb", "dd", "ff", "gg", "mm", "nn", "pp", "rr", "tt":
				w.ReplaceSuffix(newSuffix, newSuffix[:1], true)
			}

			// Because we did a double replacement,
			// we need to fix R1 and R2 manually.
			rsLen := len(w.RS)
			if originalR1start < rsLen {
				w.R1start = originalR1start
			} else {
				w.R1start = rsLen
			}
			if originalR2start < rsLen {
				w.R2start = originalR2start
			} else {
				w.R2start = rsLen
			}

			return true
		}

	}

	return false
}
