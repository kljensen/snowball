package romanian

import (
	"github.com/kljensen/snowball/snowballword"
)

// Step 0: Remove common Romanian endings
func step0(word *snowballword.SnowballWord) bool {
	suffixes := []string{
		"ului", "ul", "aua", "ele", "elor", "ea",
		"iile", "iilor", "ilor", "ile", "iua", "iei", "ii",
		"atei", "ației", "ația",
	}

	suffix := word.FirstSuffixIfIn(word.R1start, len(word.RS), suffixes...)
	if suffix == "" {
		return false
	}

	switch suffix {
	case "ul", "ului":
		word.RemoveLastNRunes(len([]rune(suffix)))
		return true
	case "aua":
		word.ReplaceSuffixRunes([]rune(suffix), []rune("a"), true)
		return true
	case "ea", "ele", "elor":
		word.ReplaceSuffixRunes([]rune(suffix), []rune("e"), true)
		return true
	case "ii", "iua", "iei", "iile", "iilor", "ilor":
		word.ReplaceSuffixRunes([]rune(suffix), []rune("i"), true)
		return true
	case "ile":
		// Check if not preceded by 'ab'
		idx := len(word.RS) - len([]rune(suffix))
		if idx >= 2 && string(word.RS[idx-2:idx]) != "ab" {
			word.ReplaceSuffixRunes([]rune(suffix), []rune("i"), true)
			return true
		}
	case "atei":
		word.ReplaceSuffixRunes([]rune(suffix), []rune("at"), true)
		return true
	case "ației", "ația":
		word.ReplaceSuffixRunes([]rune(suffix), []rune("ați"), true)
		return true
	}

	return false
}
