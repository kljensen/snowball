package romanian

import (
	"github.com/kljensen/snowball/snowballword"
)

func postprocess(word *snowballword.SnowballWord) {
	// Convert U/I back to u/i
	for i, r := range word.RS {
		if r == 'U' {
			word.RS[i] = 'u'
		} else if r == 'I' {
			word.RS[i] = 'i'
		}
	}
}
