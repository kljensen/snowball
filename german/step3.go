package german

import (
	"github.com/kljensen/snowball/snowballword"
)

// Step 3: Third do block - Remove derivational suffixes in R2
// Removes: end, ung, ig, ik, isch, lich, heit, keit
func step3(word *snowballword.SnowballWord) {

	suffix := word.FirstSuffixIfIn(word.R2start, len(word.RS),
		"isch", "lich", "heit", "keit", "end", "ung", "ig", "ik",
	)

	switch suffix {
	case "end", "ung":
		word.RemoveLastNRunes(len(suffix))
		// Try to remove 'ig' suffix if not preceded by 'e' and in R2
		if len(word.RS) >= 2 && word.R2start <= len(word.RS)-2 {
			if word.RS[len(word.RS)-2] == 105 && // i
				word.RS[len(word.RS)-1] == 103 { // g
				// Check not preceded by 'e'
				if len(word.RS) < 3 || word.RS[len(word.RS)-3] != 101 {
					word.RemoveLastNRunes(2)
				}
			}
		}
		return

	case "ig", "ik", "isch":
		// Must not be preceded by 'e'
		idx := len(word.RS) - len(suffix)
		if idx > 0 && word.RS[idx-1] == 101 { // e
			return
		}
		word.RemoveLastNRunes(len(suffix))
		return

	case "lich", "heit":
		word.RemoveLastNRunes(len(suffix))
		// Try to remove 'er' or 'en' in R1
		if len(word.RS) >= 2 && word.R1start <= len(word.RS)-2 {
			if (word.RS[len(word.RS)-2] == 101 && word.RS[len(word.RS)-1] == 114) || // er
				(word.RS[len(word.RS)-2] == 101 && word.RS[len(word.RS)-1] == 110) { // en
				word.RemoveLastNRunes(2)
			}
		}
		return

	case "keit":
		word.RemoveLastNRunes(4)
		// Try to remove 'lich' or 'ig' in R2
		if len(word.RS) >= 4 && word.R2start <= len(word.RS)-4 {
			if word.RS[len(word.RS)-4] == 108 && // l
				word.RS[len(word.RS)-3] == 105 && // i
				word.RS[len(word.RS)-2] == 99 && // c
				word.RS[len(word.RS)-1] == 104 { // h (lich)
				word.RemoveLastNRunes(4)
				return
			}
		}
		if len(word.RS) >= 2 && word.R2start <= len(word.RS)-2 {
			if word.RS[len(word.RS)-2] == 105 && // i
				word.RS[len(word.RS)-1] == 103 { // g
				word.RemoveLastNRunes(2)
			}
		}
		return
	}
}
