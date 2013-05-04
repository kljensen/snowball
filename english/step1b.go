package snowball

import (
	"github.com/kljensen/snowball/stemword"
)

func step1b(w *stemword.Word) (didReplacement bool) {

	suffix := w.FirstSuffix("eedly", "eed")
	if suffix != "" {

		// Notice that, the original algorithm is oddly
		// articulated at this step and says, if we found
		// one of these sufficies, to "replace by ee if in R1".
		// The NLTK implementation replaces by "ee" in each
		// of `wordOut`, `r1out`, `r2out`, which is what we've
		// done here.
		didReplacement = w.ReplaceSuffix(suffix, "ee")
		return
	}

	suffix = w.FirstSuffix("ed", "edly", "ing", "ingly")
	if suffix != "" {
		didReplacement = true
		for i := 0; i < len(w.RS)-len(suffix); i++ {
			if isLowerVowel(w.RS[i]) {
				w.ReplaceSuffix(suffix, "")

				var (
					newSuffix string
				)

				// Check for special ending
				newSuffix = w.FirstSuffix("at", "bl", "iz")
				if newSuffix != "" {
					w.ReplaceSuffix(newSuffix, newSuffix+"e")
					return
				}

				// Check for double consonant ending.  Note that, the original algorithm
				// implies that all double consonant endings should be removed; however,
				// the NLTK implementation only removes the following sufficies.
				//
				newSuffix = w.FirstSuffix("bb", "dd", "ff", "gg", "mm", "nn", "pp", "rr", "tt")
				if newSuffix != "" {
					w.ReplaceSuffix(newSuffix, newSuffix[:1])
					return
				}

				// Check for a short word
				if isShortWord(w) {
					// By definition, r1 and r2 are the empty string for
					// short words.
					w.RS = append(w.RS, []rune("e")...)
					w.R1start = len(w.RS)
					w.R2start = len(w.RS)
					return
				}

				break
			}
		}

		// return
	}
	return
}
