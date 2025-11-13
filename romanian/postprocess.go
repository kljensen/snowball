package romanian

import (
	"github.com/kljensen/snowball/snowballword"
)

func postprocess(word *snowballword.SnowballWord) {
	// Convert U/I back to u/i
	for i, r := range word.RS {
		switch r {
		case 'U':
			word.RS[i] = 'u'
		case 'I':
			word.RS[i] = 'i'
		}
	}
}
