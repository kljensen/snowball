package portuguese

import (
	"github.com/kljensen/snowball/snowballword"
)

// Postprocess converts back from internal representation
func postprocess(word *snowballword.SnowballWord) {
	// Replace ã with a~ and õ with o~
	newRS := make([]rune, 0, len(word.RS))
	for i := 0; i < len(word.RS); i++ {
		switch word.RS[i] {
		case 'ã':
			newRS = append(newRS, 'a', '~')
		case 'õ':
			newRS = append(newRS, 'o', '~')
		default:
			newRS = append(newRS, word.RS[i])
		}
	}
	word.RS = newRS
}
