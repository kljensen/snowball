package italian

import (
	"github.com/kljensen/snowball/snowballword"
)

// Step 1: Standard suffix or verb suffix removal
func step1(word *snowballword.SnowballWord) {
	// Try standard suffixes first
	if standardSuffix(word) {
		return
	}
	// If no standard suffix removed, try verb suffixes
	verbSuffix(word)
}

func standardSuffix(word *snowballword.SnowballWord) bool {
	// Ordered by length (longest first) to ensure longest match
	suffixes := []string{
		// Length 7
		"amenti", "imento", "azione", "azioni", "atrice", "atrici",
		// Length 6
		"amente", "amento", "atori", "uzione", "uzioni", "usione", "usioni",
		"logia", "logie",
		// Length 5
		"abile", "abili", "ibile", "ibili", "mente", "atore",
		"anza", "anze", "enza", "enze",
		"ismo", "ismi", "ista", "iste", "isti",
		"iche", "ichi", "istà", "istè", "istì",
		"ante", "anti",
		// Length 4
		"oso", "osi", "osa", "ose",
		"ico", "ici", "ica", "ice",
		"ivo", "ivi", "iva", "ive",
		// Length 3
		"ità",
	}

	suffix := word.FirstSuffixIfIn(word.R2start, len(word.RS), suffixes...)
	if suffix == "" {
		return false
	}

	// Handle special cases
	switch suffix {
	case "amente":
		if word.R1start > len(word.RS)-6 {
			return false
		}
		word.RemoveLastNRunes(6)
		// Try removing iv/os/ic/abil in R2
		word.RemoveFirstSuffixIfIn(word.R2start, "iv", "os", "ic", "abil")
		return true

	case "logia", "logie":
		word.ReplaceSuffixRunes([]rune(suffix), []rune("log"), true)
		return true

	case "uzione", "uzioni", "usione", "usioni":
		word.ReplaceSuffixRunes([]rune(suffix), []rune("u"), true)
		return true

	case "enza", "enze":
		word.ReplaceSuffixRunes([]rune(suffix), []rune("ente"), true)
		return true

	case "ivo", "ivi", "iva", "ive":
		word.RemoveLastNRunes(len([]rune(suffix)))
		// Try removing at in R2
		word.RemoveFirstSuffixIfIn(word.R2start, "at")
		return true

	default:
		word.RemoveLastNRunes(len([]rune(suffix)))
		return true
	}
}

func verbSuffix(word *snowballword.SnowballWord) bool {
	// Ordered by length (longest first) to ensure longest match
	verbs := []string{
		// Length 8
		"erebbero", "irebbero",
		// Length 7
		"assero", "assimo", "eranno", "erebbe", "eremmo", "ereste", "eresti",
		"essero", "iranno", "irebbe", "iremmo", "ireste", "iresti", "iscano",
		"iscono", "issero",
		// Length 6
		"arono", "avamo", "avano", "avate", "eremo", "erete", "erono",
		"evamo", "evano", "evate", "iremo", "irete", "irono", "ivamo",
		"ivano", "ivate",
		// Length 5
		"ammo", "ando", "asse", "assi", "emmo", "enda", "ende", "endi",
		"endo", "erai", "erei", "Yamo", "iamo", "immo", "irai", "irei",
		"isca", "isce", "isci", "isco",
		// Length 4
		"ano", "are", "ata", "ate", "ati", "ato", "ava", "avi", "avo",
		"erà", "ere", "erò", "ete", "eva", "evi", "evo", "irà", "ire",
		"irò", "ita", "ite", "iti", "ito", "iva", "ivi", "ivo", "ono",
		"uta", "ute", "uti", "uto",
		// Length 2
		"ar", "ir",
	}

	suffix := word.FirstSuffixIfIn(word.RVstart, len(word.RS), verbs...)
	if suffix != "" {
		word.RemoveLastNRunes(len([]rune(suffix)))
		return true
	}
	return false
}
