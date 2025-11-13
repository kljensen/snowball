package italian

import (
	"github.com/kljensen/snowball/snowballword"
)

// Step 2: Final vowel suffix removal
func step2(word *snowballword.SnowballWord) {
	// Try removing final vowels a e i o à è ì ò
	if len(word.RS) == 0 {
		return
	}

	lastRune := word.RS[len(word.RS)-1]
	// Check if last rune is a/e/i/o or accented versions in RV
	if word.RVstart < len(word.RS) {
		switch lastRune {
		case 'a', 'e', 'i', 'o', 'à', 'è', 'ì', 'ò':
			word.RemoveLastNRunes(1)
			// If preceded by 'i', also remove the 'i' if in RV
			if len(word.RS) > 0 && word.RS[len(word.RS)-1] == 'i' &&
				word.RVstart < len(word.RS) {
				word.RemoveLastNRunes(1)
			}
		}
	}

	// Also try removing final 'h' if preceded by 'c' or 'g' in RV
	if len(word.RS) >= 2 && word.RS[len(word.RS)-1] == 'h' &&
		word.RVstart <= len(word.RS)-2 {
		prevChar := word.RS[len(word.RS)-2]
		if prevChar == 'c' || prevChar == 'g' {
			word.RemoveLastNRunes(1)
		}
	}
}
