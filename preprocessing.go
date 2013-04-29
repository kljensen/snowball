package snowball

import (
	// "errors"
	// "log"
	"strings"
	"unicode/utf8"
)

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

// Test if a string has a rune, skipping parts of the string
// that are less than `leftSkip` of the beginning and `rightSkip`
// of the end.
//
func hasRune(word string, leftSkip int, rightSkip int, testRunes ...rune) bool {
	leftMin := leftSkip
	rightMax := utf8.RuneCountInString(word) - rightSkip
	for i, r := range word {
		if i < leftMin {
			continue
		} else if i >= rightMax {
			break
		}
		for _, tr := range testRunes {
			if r == tr {
				return true
			}
		}
	}
	return false
}

// Test if a string has a vowel, skipping parts of the string
// that are less than `leftSkip` of the beginning and `rightSkip`
// of the end.  (All counts in runes.)
//
func hasVowel(word string, leftSkip int, rightSkip int) bool {
	return hasRune(word, leftSkip, rightSkip, 97, 101, 105, 111, 117, 121)
}

// Checks if a rune is a lowercase English vowel.
//
func isLowerVowel(r rune) bool {
	switch r {
	case 97, 101, 105, 111, 117, 121:
		return true
	}
	return false
}

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

// Takes an `inputWord` and applies various transformations
// necessary for the other, subsequent stemming steps.
//
func preprocessWord(word string) string {
	word = strings.ToLower(word)

	// Return small words and stop words
	if len(word) <= 2 || stopWords[word] {
		return word
	}

	// Return special words
	if specialVersion, ok := specialWords[word]; ok {
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

// Finds the region after the first non-vowel following a vowel,
// or is the null region at the end of the word if there is no
// such non-vowel.
//
func vnvSuffix(word string) string {
	runes := []rune(word)
	// uscular
	for i := 1; i < len(runes); i++ {
		if isLowerVowel(runes[i-1]) && !isLowerVowel(runes[i]) {
			return string(runes[i+1:])
		}
	}
	return ""
}

// R1 is the region after the first non-vowel following a vowel,
// or is the null region at the end of the word if there is no
// such non-vowel.
//
// R2 is the region after the first non-vowel following a vowel
// in R1, or is the null region at the end of the word if there
// is no such non-vowel.
//
// See http://snowball.tartarus.org/texts/r1r2.html
//
func r1r2(word string) (r1, r2 string) {

	specialPrefixes := []string{"gener", "commun", "arsen"}
	hasSpecialPrefix := false
	specialPrefix := ""
	for _, specialPrefix = range specialPrefixes {
		if strings.HasPrefix(word, specialPrefix) {
			hasSpecialPrefix = true
			break
		}
	}

	if hasSpecialPrefix {
		if specialPrefix == "commun" {
			r1 = word[6:]
		} else {
			r1 = word[5:]
		}

	} else {
		r1 = vnvSuffix(word)
	}
	r2 = vnvSuffix(r1)
	return
}
