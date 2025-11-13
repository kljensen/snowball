package portuguese

import (
	"github.com/kljensen/snowball/snowballword"
)

// Preprocess handles Portuguese-specific preprocessing
// Converts a~ and o~ sequences
func preprocess(word *snowballword.SnowballWord) {
	// Replace a~ with ã and o~ with õ
	newRS := make([]rune, 0, len(word.RS))
	for i := 0; i < len(word.RS); i++ {
		if i+1 < len(word.RS) && word.RS[i+1] == '~' {
			switch word.RS[i] {
			case 'a':
				newRS = append(newRS, 'ã')
				i++ // skip the ~
				continue
			case 'o':
				newRS = append(newRS, 'õ')
				i++ // skip the ~
				continue
			}
		}
		newRS = append(newRS, word.RS[i])
	}
	word.RS = newRS

	// Calculate regions R1, R2, and RV
	r1start, r2start, rvstart := findRegions(word)
	word.R1start = r1start
	word.R2start = r2start
	word.RVstart = rvstart
}
