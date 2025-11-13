package italian

import (
	"github.com/kljensen/snowball/snowballword"
)

// Postprocess converts back from internal representation
func postprocess(word *snowballword.SnowballWord) {
	// Convert I and U back to lowercase
	for i := 0; i < len(word.RS); i++ {
		switch word.RS[i] {
		case 'I':
			word.RS[i] = 'i'
		case 'U':
			word.RS[i] = 'u'
		}
	}
}
