package portuguese

import (
	"github.com/kljensen/snowball/snowballword"
)

// Portuguese vowels: a e i o u á é í ó ú â ê ô
func isLowerVowel(r rune) bool {
	switch r {
	case 97, 101, 105, 111, 117: // a e i o u
		return true
	case 225, 233, 237, 243, 250: // á é í ó ú
		return true
	case 226, 234, 244: // â ê ô
		return true
	}
	return false
}

// Find the starting point of regions R1, R2, and RV
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

	// RV: If the second letter is a consonant, RV is the region after the next
	// following vowel, or if the first two letters are vowels, RV is the region
	// after the next consonant, and otherwise (consonant-vowel case) RV is the
	// region after the third letter.
	rvstart = len(word.RS)
	if len(word.RS) >= 3 {
		switch {
		case !isLowerVowel(word.RS[1]):
			// Second letter is consonant, find next vowel
			for i := 2; i < len(word.RS); i++ {
				if isLowerVowel(word.RS[i]) {
					rvstart = i + 1
					break
				}
			}
		case isLowerVowel(word.RS[0]) && isLowerVowel(word.RS[1]):
			// First two are vowels, find next consonant
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

	return
}

// Return true if the input word is a Portuguese stop word.
func IsStopWord(word string) bool {
	switch word {
	case "o", "a", "os", "as", "um", "uma", "uns", "umas",
		"de", "do", "da", "dos", "das", "em", "no", "na", "nos", "nas",
		"por", "pelo", "pela", "pelos", "pelas",
		"ao", "aos", "à", "às",
		"e", "é", "ou",
		"mas", "mais", "menos",
		"para", "com", "sem", "sob", "sobre",
		"entre", "até", "desde", "após", "antes",
		"que", "qual", "quais", "quanto", "quantos", "quanta", "quantas",
		"este", "esta", "estes", "estas", "esse", "essa", "esses", "essas",
		"aquele", "aquela", "aqueles", "aquelas",
		"isto", "isso", "aquilo",
		"eu", "tu", "ele", "ela", "nós", "vós", "eles", "elas",
		"me", "te", "se", "lhe", "lhes", "vos",
		"meu", "minha", "meus", "minhas",
		"teu", "tua", "teus", "tuas",
		"seu", "sua", "seus", "suas",
		"nosso", "nossa", "nossos", "nossas",
		"vosso", "vossa", "vossos", "vossas",
		"não", "nem", "nunca", "jamais", "também", "só", "já", "ainda",
		"quando", "onde", "como", "porque", "porquê",
		"muito", "muita", "muitos", "muitas",
		"pouco", "pouca", "poucos", "poucas",
		"todo", "toda", "todos", "todas",
		"outro", "outra", "outros", "outras":
		return true
	}
	return false
}
