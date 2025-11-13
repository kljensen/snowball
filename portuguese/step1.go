package portuguese

import (
	"github.com/kljensen/snowball/snowballword"
)

// Step 1: Standard suffix removal
func step1(word *snowballword.SnowballWord) bool {
	suffix := word.FirstSuffixIfIn(word.R2start, len(word.RS),
		"amentos", "imentos", "amento", "imento",
		"adoras", "adores", "aç", "ões", "adora", "ações",
		"antes", "ância", "ante", "ador",
		"ismo", "ismos", "ável", "ível", "istas", "ista",
		"osas", "osos", "osa", "oso",
		"amente",
		"mente", "idade", "idades",
		"ivas", "ivos", "iva", "ivo",
		"ira", "iras",
	)

	switch suffix {
	case "":
		return false

	case "amente":
		// In R1, delete if preceded by os/ic/ad in R2
		if word.R1start > len(word.RS)-7 {
			return false
		}
		word.RemoveLastNRunes(6)
		// Try deleting os, ic, ad if in R2
		suffix2 := word.FirstSuffixIfIn(word.R2start, len(word.RS),
			"iv", "os", "ic", "ad",
		)
		if suffix2 == "iv" {
			word.RemoveLastNRunes(2)
			// If preceded by 'at' in R2, delete
			word.RemoveFirstSuffixIfIn(word.R2start, "at")
		} else if suffix2 != "" {
			word.RemoveLastNRunes(len(suffix2))
		}
		return true

	case "mente":
		if word.R2start > len(word.RS)-5 {
			return false
		}
		word.RemoveLastNRunes(5)
		// Try deleting ante/avel/ível in R2
		word.RemoveFirstSuffixIfIn(word.R2start, "ante", "avel", "ível")
		return true

	case "idade", "idades":
		if word.R2start > len(word.RS)-len(suffix) {
			return false
		}
		word.RemoveLastNRunes(len(suffix))
		// Try deleting abil/ic/iv in R2
		word.RemoveFirstSuffixIfIn(word.R2start, "abil", "ic", "iv")
		return true

	case "iva", "ivo", "ivas", "ivos":
		if word.R2start > len(word.RS)-len(suffix) {
			return false
		}
		word.RemoveLastNRunes(len(suffix))
		// Try deleting 'at' in R2
		word.RemoveFirstSuffixIfIn(word.R2start, "at")
		return true

	case "ira", "iras":
		// Delete if in RV and preceded by 'e'
		idx := len(word.RS) - len(suffix)
		if word.RVstart > idx {
			return false
		}
		if idx > 0 && word.RS[idx-1] == 'e' {
			word.RemoveLastNRunes(len(suffix))
			word.ReplaceSuffixRunes([]rune("e"), []rune("ir"), true)
			return true
		}
		return false

	default:
		// All other suffixes: delete if in R2
		if word.R2start > len(word.RS)-len(suffix) {
			return false
		}
		word.RemoveLastNRunes(len(suffix))
		return true
	}
}
