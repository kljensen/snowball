package english

import (
	"github.com/kljensen/snowball/stemword"
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

			// Delete if the preceding word part contains a vowel
			w.ReplaceSuffix(suffix, "", true)

			// and after the deletion...
			var (
				newSuffix string
			)

			// If the word ends "at", "bl" or "iz" add "e" 
			newSuffix = w.FirstSuffix("at", "bl", "iz")
			if newSuffix != "" {
				w.ReplaceSuffix(newSuffix, newSuffix+"e", true)
				return true
			}

			// If the word ends with a double remove the last letter
			newSuffix = w.FirstSuffix("bb", "dd", "ff", "gg", "mm", "nn", "pp", "rr", "tt")
			if newSuffix != "" {
				w.ReplaceSuffix(newSuffix, newSuffix[:1], true)
				return true
			}

			// If the word is short, add "e"
			if isShortWord(w) {
				// By definition, r1 and r2 are the empty string for
				// short words.
				w.RS = append(w.RS, []rune("e")...)
				w.R1start = len(w.RS)
				w.R2start = len(w.RS)
				return true
			}
			return true
		}

	}

	return false
}
