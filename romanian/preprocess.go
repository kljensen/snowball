package romanian

import (
	"github.com/kljensen/snowball/snowballword"
)

func preprocess(word *snowballword.SnowballWord) {
	// Normalize cedilla to comma-below forms
	for i, r := range word.RS {
		switch r {
		case 0x015F: // ş cedilla -> ș comma
			word.RS[i] = 0x0219
		case 0x0163: // ţ cedilla -> ț comma
			word.RS[i] = 0x021B
		}
	}

	// Mark u/i between vowels as U/I
	for i := 1; i < len(word.RS)-1; i++ {
		if word.RS[i] == 'u' && isVowel(word.RS[i-1]) && isVowel(word.RS[i+1]) {
			word.RS[i] = 'U'
		} else if word.RS[i] == 'i' && isVowel(word.RS[i-1]) && isVowel(word.RS[i+1]) {
			word.RS[i] = 'I'
		}
	}

	// Mark regions
	r1, r2, rv := findRegions(word)
	word.R1start = r1
	word.R2start = r2
	word.RVstart = rv
}
