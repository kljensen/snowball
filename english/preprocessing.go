package snowball

import (
	"strings"
)

// Capitalize all 'Y's preceded by vowels or starting a word
//
func capitalizeYs(word string) string {
	runes := []rune(word)
	for i, r := range runes {
		if r == 'y' && (i == 0 || isLowerVowel(runes[i-1])) {
			runes[i] = 'Y'
		}
	}
	return string(runes)
}

// Replaces all different kinds of apostrophes with a single
// kind: "\x27".
//
func normalizeApostrophes(inputWord string) string {
	var apostrophes = [3]string{"\u2019", "\u2018", "\u201B"}
	outputWord := inputWord
	for _, apostrophe := range apostrophes {
		outputWord = strings.Replace(outputWord, apostrophe, "\x27", -1)
	}
	return outputWord
}

// Takes an `inputWord` and applies various transformations
// necessary for the other, subsequent stemming steps.
//
func preprocessWord(word string) string {
	word = strings.ToLower(word)

	// Return small words and stop words
	if len(word) <= 2 || isStopWord(word) {
		return word
	}

	// Return special words
	if specialVersion := stemSpecialWord(word); specialVersion != "" {
		word = specialVersion
		return word
	}

	// Normalize all possible apostrophe variations
	word = normalizeApostrophes(word)

	// Trim off leading apostropes.  (Slight variation from
	// NLTK implementation here, in which only the first is removed.)
	word = strings.TrimLeft(word, "\x27")

	// Capitalize all 'Y's preceded by vowels or starting a word
	word = capitalizeYs(word)

	return word
}
