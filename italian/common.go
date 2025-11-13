package italian

import (
	"github.com/kljensen/snowball/snowballword"
)

// Italian vowels: a e i o u à è ì ò ù
func isLowerVowel(r rune) bool {
	switch r {
	case 97, 101, 105, 111, 117: // a e i o u
		return true
	case 224, 232, 236, 242, 249: // à è ì ò ù
		return true
	}
	return false
}

// Find the starting point of regions R1, R2, and RV for Italian
func findRegions(word *snowballword.SnowballWord) (r1start, r2start, rvstart int) {
	// R1: region after first non-vowel following a vowel
	r1start = len(word.RS)
	for i := 0; i < len(word.RS)-1; i++ {
		if isLowerVowel(word.RS[i]) && !isLowerVowel(word.RS[i+1]) {
			r1start = i + 2
			break
		}
	}

	// R2: region after first non-vowel following a vowel in R1
	r2start = len(word.RS)
	for i := r1start; i < len(word.RS)-1; i++ {
		if isLowerVowel(word.RS[i]) && !isLowerVowel(word.RS[i+1]) {
			r2start = i + 2
			break
		}
	}

	// RV: Special handling for Italian
	// If second letter is consonant, RV is region after next following vowel
	// If first two letters are vowels, RV is region after next consonant
	// Otherwise RV is region after third letter
	// Special case: "divano" -> RV should not make "div" collide with "diva"
	rvstart = len(word.RS)
	if len(word.RS) >= 3 {
		// Check for special case "divan" - handled in mark_regions
		str := string(word.RS)
		if len(str) >= 5 && str[:5] == "divan" {
			rvstart = 5
		} else {
			switch {
			case !isLowerVowel(word.RS[1]):
				// Second letter is consonant
				for i := 2; i < len(word.RS); i++ {
					if isLowerVowel(word.RS[i]) {
						rvstart = i + 1
						break
					}
				}
			case isLowerVowel(word.RS[0]) && isLowerVowel(word.RS[1]):
				// First two are vowels
				for i := 2; i < len(word.RS); i++ {
					if !isLowerVowel(word.RS[i]) {
						rvstart = i + 1
						break
					}
				}
			default:
				// Consonant-vowel case
				rvstart = 3
			}
		}
	}

	return
}

// Return true if the input word is an Italian stop word.
func IsStopWord(word string) bool {
	switch word {
	case "a", "al", "all", "alla", "alle", "allo", "agli", "anche",
		"come", "con", "cosa", "cui",
		"da", "dal", "dall", "dalla", "dalle", "dallo", "dagli", "dagl", "dei", "del",
		"dell", "della", "delle", "dello", "degli", "degl", "di", "dove",
		"e", "è", "ed",
		"fu",
		"gli", "glie",
		"ha", "hai", "hanno", "ho",
		"i", "il", "in",
		"io",
		"l", "la", "le", "lei", "li", "lo", "loro", "lui",
		"ma", "me", "mi", "mia", "mie", "miei", "mio",
		"ne", "nel", "nell", "nella", "nelle", "nello", "negli", "negl",
		"noi", "non", "nostra", "nostre", "nostri", "nostro",
		"o", "od",
		"per", "perché", "più",
		"qual", "quale", "quanta", "quante", "quanti", "quanto", "quella", "quelle",
		"quelli", "quello", "questa", "queste", "questi", "questo",
		"se", "sei", "si", "sia", "siamo", "siano", "siete", "sono", "sta", "stai",
		"stare", "stato", "sto", "su", "sua", "sue", "sui", "suo", "suoi",
		"te", "ti", "tra", "tu", "tua", "tue", "tuo", "tuoi",
		"un", "una", "uno",
		"vi", "voi", "vostra", "vostre", "vostri", "vostro":
		return true
	}
	return false
}
