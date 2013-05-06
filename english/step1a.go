package english

import (
	"github.com/kljensen/snowball/stemword"
)

// Step 1a is noralization of various special "s"-endings.
//
func step1a(w *stemword.Word) bool {

	suffix := w.FirstSuffix("sses", "ied", "ies", "us", "ss", "s")
	switch suffix {

	case "sses":

		// Replace by ss 
		w.ReplaceSuffix(suffix, "ss", true)
		return true

	case "ies", "ied":

		// Replace by i if preceded by more than one letter,
		// otherwise by ie (so ties -> tie, cries -> cri).

		var repl string
		if len(w.RS) > 4 {
			repl = "i"
		} else {
			repl = "ie"
		}
		w.ReplaceSuffix(suffix, repl, true)
		return true

	case "us", "ss":

		// Do nothing
		return false

	case "s":

		// Delete if the preceding word part contains a vowel
		// not immediately before the s (so gas and this retain
		// the s, gaps and kiwis lose it) 
		//
		for i := 0; i < len(w.RS)-2; i++ {
			if isLowerVowel(w.RS[i]) {
				w.ReplaceSuffix(suffix, "", true)
				return true
			}
		}
	}
	return false
}
