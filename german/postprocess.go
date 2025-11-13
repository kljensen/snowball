package german

import (
	"github.com/kljensen/snowball/snowballword"
)

// Postprocess converts uppercase letters back to lowercase
// Y -> y, U -> u, and umlauts to base letters: ä -> a, ö -> o, ü -> u
func postprocess(word *snowballword.SnowballWord) {
	for i := 0; i < len(word.RS); i++ {
		switch word.RS[i] {
		case 89: // Y
			word.RS[i] = 121 // y
		case 85: // U
			word.RS[i] = 117 // u
		case 228: // ä
			word.RS[i] = 97 // a
		case 246: // ö
			word.RS[i] = 111 // o
		case 252: // ü
			word.RS[i] = 117 // u
		}
	}
}
