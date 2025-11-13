package polish

import (
	"github.com/kljensen/snowball/snowballword"
	"strings"
)

// Stem a Polish word. This is the only exported
// function in this package.
func Stem(word string, stemStopWords bool) string {
	word = strings.ToLower(strings.TrimSpace(word))

	if len(word) <= 2 || (stemStopWords == false && IsStopWord(word)) {
		return word
	}

	w := snowballword.New(word)

	// Mark regions
	markRegions(w)

	// Remove endings (requires at least 3 characters)
	removeEndings(w)

	return w.String()
}
