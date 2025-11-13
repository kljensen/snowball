package german

import (
	"github.com/kljensen/snowball/snowballword"
)

// Preprocess performs the German-specific preprocessing
// which includes:
// 1. Convert u and y between vowels to uppercase U and Y
// 2. Normalize special characters (ß -> ss, ae -> ä, oe -> ö, ue -> ü, but not que)
// 3. Calculate R1 and R2 regions
func preprocess(word *snowballword.SnowballWord) {

	// First pass: convert u and y between vowels to uppercase
	for i := 1; i < len(word.RS)-1; i++ {
		if (word.RS[i] == 117 || word.RS[i] == 121) && // u or y
			isLowerVowel(word.RS[i-1]) &&
			isLowerVowel(word.RS[i+1]) {
			word.RS[i] -= 32 // Convert to uppercase (U or Y)
		}
	}

	// Second pass: normalize special characters
	// We need to handle ß -> ss and ae/oe/ue -> umlaut conversions
	newRS := make([]rune, 0, len(word.RS))
	for i := 0; i < len(word.RS); i++ {
		switch word.RS[i] {
		case 223: // ß
			newRS = append(newRS, 115, 115) // ss
		case 97: // a
			if i+1 < len(word.RS) && word.RS[i+1] == 101 { // ae
				newRS = append(newRS, 228) // ä
				i++ // skip the 'e'
			} else {
				newRS = append(newRS, word.RS[i])
			}
		case 111: // o
			if i+1 < len(word.RS) && word.RS[i+1] == 101 { // oe
				newRS = append(newRS, 246) // ö
				i++ // skip the 'e'
			} else {
				newRS = append(newRS, word.RS[i])
			}
		case 117: // u
			if i+1 < len(word.RS) && word.RS[i+1] == 101 { // ue
				newRS = append(newRS, 252) // ü
				i++ // skip the 'e'
			} else {
				newRS = append(newRS, word.RS[i])
			}
		case 113: // q
			// Handle qu - don't convert ue after q
			newRS = append(newRS, word.RS[i])
			if i+1 < len(word.RS) && word.RS[i+1] == 117 { // u
				newRS = append(newRS, word.RS[i+1])
				i++ // skip conversion for this u
			}
		default:
			newRS = append(newRS, word.RS[i])
		}
	}
	word.RS = newRS

	// Calculate regions
	r1start, r2start := findRegions(word)
	word.R1start = r1start
	word.R2start = r2start
}
