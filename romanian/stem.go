package romanian

import (
	"github.com/kljensen/snowball/snowballword"
	"strings"
)

// Stem a Romanian word. This is the only exported
// function in this package.
func Stem(word string, stemStopWords bool) string {
	word = strings.ToLower(strings.TrimSpace(word))

	if len(word) <= 2 || (stemStopWords == false && IsStopWord(word)) {
		return word
	}

	w := snowballword.New(word)

	// Stemming steps
	preprocess(w)
	step0(w)
	step1(w)
	postprocess(w)

	return w.String()
}
