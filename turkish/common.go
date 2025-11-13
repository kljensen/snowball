package turkish

import (
	"github.com/kljensen/snowball/snowballword"
)

// Turkish vowels: a e i o u ı ö ü
func isVowel(r rune) bool {
	switch r {
	case 'a', 'e', 'i', 'o', 'u':
		return true
	case 0x0131, 0x00F6, 0x00FC: // ı ö ü
		return true
	}
	return false
}

// Mark regions R1 and R2
func markRegions(word *snowballword.SnowballWord) {
	// R1: after first non-vowel following a vowel
	r1 := len(word.RS)
	for i := 0; i < len(word.RS)-1; i++ {
		if isVowel(word.RS[i]) && !isVowel(word.RS[i+1]) {
			r1 = i + 2
			break
		}
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

// IsStopWord returns true if the word is a Turkish stop word
func IsStopWord(word string) bool {
	switch word {
	case "acaba", "ama", "bana", "bazı", "belki", "ben", "benden", "beni", "benim",
		"bir", "biri", "birkaç", "birkez", "birşey", "birşeyi", "biz", "bize",
		"bizden", "bizi", "bizim", "bu", "buna", "bunda", "bundan", "bunlar",
		"bunları", "bunların", "bunu", "bunun", "burada", "çok", "çünkü",
		"da", "daha", "de", "den", "diye", "gibi", "için", "ile", "ise",
		"kadar", "ki", "kim", "mu", "mü", "mı", "nasıl", "ne", "neden", "nerde",
		"nerede", "nereye", "niçin", "niye", "o", "olan", "onlar", "onları",
		"onların", "onu", "onun", "öyle", "şey", "şeyi", "şeyler", "şu", "şuna",
		"şunda", "şundan", "şunlar", "şunu", "tüm", "var", "ve", "veya", "ya":
		return true
	}
	return false
}
