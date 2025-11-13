package finnish

import (
	"github.com/kljensen/snowball/snowballword"
	"strings"
)

// Stem a Finnish word. This is the only exported
// function in this package.
func Stem(word string, stemStopWords bool) string {
	word = strings.ToLower(strings.TrimSpace(word))

	if len(word) <= 2 || (!stemStopWords && IsStopWord(word)) {
		return word
	}

	w := snowballword.New(word)

	// Mark regions
	markRegions(w)

	// Apply stemming steps
	endingRemoved = false
	particleEtc(w)
	possessive(w)
	caseEnding(w)
	otherEndings(w)

	if endingRemoved {
		iPlural(w)
	} else {
		tPlural(w)
	}

	tidy(w)

	return w.String()
}
