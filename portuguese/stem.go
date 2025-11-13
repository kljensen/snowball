package portuguese

import (
	"github.com/kljensen/snowball/snowballword"
	"strings"
)

// Stem a Portuguese word. This is the only exported
// function in this package.
func Stem(word string, stemStopWords bool) string {
	word = strings.ToLower(strings.TrimSpace(word))

	if len(word) <= 2 || (stemStopWords == false && IsStopWord(word)) {
		return word
	}

	w := snowballword.New(word)

	// Stem the word. Note, each of these
	// steps will alter `w` in place.
	preprocess(w)

	// Standard suffix or verb suffix
	changed := step1(w)
	if !changed {
		step2(w)
	}

	step3(w)
	step4(w)

	postprocess(w)

	return w.String()
}
