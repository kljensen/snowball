package italian

import (
	"github.com/kljensen/snowball/snowballword"
)

// Preprocess handles Italian-specific preprocessing
func preprocess(word *snowballword.SnowballWord) {
	// Convert acute accents to grave: á->à, é->è, í->ì, ó->ò, ú->ù
	for i := 0; i < len(word.RS); i++ {
		switch word.RS[i] {
		case 225: // á
			word.RS[i] = 224 // à
		case 233: // é
			word.RS[i] = 232 // è
		case 237: // í
			word.RS[i] = 236 // ì
		case 243: // ó
			word.RS[i] = 242 // ò
		case 250: // ú
			word.RS[i] = 249 // ù
		}
	}

	// Convert 'qu' to 'qU', and 'u' or 'i' between vowels to 'U' or 'I'
	newRS := make([]rune, 0, len(word.RS))
	for i := 0; i < len(word.RS); i++ {
		if i+1 < len(word.RS) && word.RS[i] == 'q' && word.RS[i+1] == 'u' {
			newRS = append(newRS, 'q', 'U')
			i++ // skip the 'u'
			continue
		}

		// Convert u/i between vowels to U/I
		if i > 0 && i+1 < len(word.RS) {
			if (word.RS[i] == 'u' || word.RS[i] == 'i') &&
				isLowerVowel(word.RS[i-1]) &&
				isLowerVowel(word.RS[i+1]) {
				if word.RS[i] == 'u' {
					newRS = append(newRS, 'U')
				} else {
					newRS = append(newRS, 'I')
				}
				continue
			}
		}
		newRS = append(newRS, word.RS[i])
	}
	word.RS = newRS

	// Calculate regions
	r1start, r2start, rvstart := findRegions(word)
	word.R1start = r1start
	word.R2start = r2start
	word.RVstart = rvstart
}
