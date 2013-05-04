package snowball

import (
	"github.com/kljensen/snowball/stemword"
)

// Step 1a is noralization of various special "s"-endings.
//
func step1a(w *stemword.Word) (didReplacement bool) {
	suffix := w.FirstSuffix("sses", "ied", "ies", "s")
	switch suffix {
	case "sses":
		didReplacement = w.ReplaceSuffix(suffix, "ss")

	case "ied":
	case "ies":
		var repl string
		if len(w.RS) == 4 {
			repl = "ie"
		} else {
			repl = "i"
		}
		didReplacement = w.ReplaceSuffix(suffix, repl)
	case "s":
		for i := 0; i < 2 && i < len(w.RS); i++ {
			if isLowerVowel(w.RS[i]) {
				didReplacement = w.ReplaceSuffix(suffix, "")
			}
		}
	}
	return didReplacement
}
