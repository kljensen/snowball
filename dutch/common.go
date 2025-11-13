package dutch

import (
	"github.com/kljensen/snowball/snowballword"
)

// Dutch vowels: a e i o u y + accented variants
func isVowel(r rune) bool {
	switch r {
	case 'a', 'e', 'i', 'o', 'u', 'y':
		return true
	case 0x00E0, 0x00E1, 0x00E2, 0x00E4: // à á â ä
		return true
	case 0x00E8, 0x00E9, 0x00EA, 0x00EB: // è é ê ë
		return true
	case 0x00EC, 0x00ED, 0x00EE, 0x00EF: // ì í î ï
		return true
	case 0x00F2, 0x00F3, 0x00F4, 0x00F6: // ò ó ô ö
		return true
	case 0x00F9, 0x00FA, 0x00FB, 0x00FC: // ù ú û ü
		return true
	}
	return false
}

// Mark regions R1 and R2
func markRegions(word *snowballword.SnowballWord) {
	// R1: after first non-vowel following a vowel, or after first 3 letters
	r1 := len(word.RS)
	for i := 0; i < len(word.RS)-1; i++ {
		if isVowel(word.RS[i]) && !isVowel(word.RS[i+1]) {
			r1 = i + 2
			break
		}
	}
	if r1 < 3 {
		r1 = 3
	}

	// R2: after first non-vowel following a vowel in R1
	r2 := len(word.RS)
	for i := r1; i < len(word.RS)-1; i++ {
		if isVowel(word.RS[i]) && !isVowel(word.RS[i+1]) {
			r2 = i + 2
			break
		}
	}

	word.R1start = r1
	word.R2start = r2
}

// IsStopWord returns true if the word is a Dutch stop word
func IsStopWord(word string) bool {
	switch word {
	case "de", "het", "een", "en", "van", "op", "in", "dat", "die", "voor",
		"te", "aan", "met", "is", "als", "zijn", "wordt", "uit", "naar", "door",
		"om", "over", "er", "meer", "deze", "hij", "kan", "ook", "dan", "bij",
		"tot", "moet", "nog", "nu", "maar", "al", "niet", "wat", "wel", "worden",
		"heeft", "hebben", "was", "waren", "hun", "ze", "zij", "haar", "hem":
		return true
	}
	return false
}
