package italian

import (
	"github.com/kljensen/snowball/snowballword"
	"strings"
)

// Stem a Italian word. This is the only exported
// function in this package.
func Stem(word string, stemStopWords bool) string {
	word = strings.ToLower(strings.TrimSpace(word))

	if len(word) <= 2 || (stemStopWords == false && IsStopWord(word)) {
		return word
	}

	w := snowballword.New(word)
	
	// TODO: Implement stemming steps
	_ = w
	
	return word
}
