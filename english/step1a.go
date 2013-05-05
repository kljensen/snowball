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
		w.ReplaceSuffix(suffix, "ss", true)
		return true

	case "ies", "ied":
		var repl string
		if len(w.RS) == 4 {
			repl = "ie"
		} else {
			repl = "i"
		}
		w.ReplaceSuffix(suffix, repl, true)
		return true

	case "us", "ss":
		return false

	case "s":
		for i := 0; i < len(w.RS)-2; i++ {
			if isLowerVowel(w.RS[i]) {
				w.ReplaceSuffix(suffix, "", true)
				return true
			}
		}
	}
	return false
}
