package snowball

import (
	"strings"
)

var stopWords = map[string]bool{
	"a":          true,
	"about":      true,
	"above":      true,
	"after":      true,
	"again":      true,
	"against":    true,
	"all":        true,
	"am":         true,
	"an":         true,
	"and":        true,
	"any":        true,
	"are":        true,
	"as":         true,
	"at":         true,
	"be":         true,
	"because":    true,
	"been":       true,
	"before":     true,
	"being":      true,
	"below":      true,
	"between":    true,
	"both":       true,
	"but":        true,
	"by":         true,
	"can":        true,
	"did":        true,
	"do":         true,
	"does":       true,
	"doing":      true,
	"don":        true,
	"down":       true,
	"during":     true,
	"each":       true,
	"few":        true,
	"for":        true,
	"from":       true,
	"further":    true,
	"had":        true,
	"has":        true,
	"have":       true,
	"having":     true,
	"he":         true,
	"her":        true,
	"here":       true,
	"hers":       true,
	"herself":    true,
	"him":        true,
	"himself":    true,
	"his":        true,
	"how":        true,
	"i":          true,
	"if":         true,
	"in":         true,
	"into":       true,
	"is":         true,
	"it":         true,
	"its":        true,
	"itself":     true,
	"just":       true,
	"me":         true,
	"more":       true,
	"most":       true,
	"my":         true,
	"myself":     true,
	"no":         true,
	"nor":        true,
	"not":        true,
	"now":        true,
	"of":         true,
	"off":        true,
	"on":         true,
	"once":       true,
	"only":       true,
	"or":         true,
	"other":      true,
	"our":        true,
	"ours":       true,
	"ourselves":  true,
	"out":        true,
	"over":       true,
	"own":        true,
	"s":          true,
	"same":       true,
	"she":        true,
	"should":     true,
	"so":         true,
	"some":       true,
	"such":       true,
	"t":          true,
	"than":       true,
	"that":       true,
	"the":        true,
	"their":      true,
	"theirs":     true,
	"them":       true,
	"themselves": true,
	"then":       true,
	"there":      true,
	"these":      true,
	"they":       true,
	"this":       true,
	"those":      true,
	"through":    true,
	"to":         true,
	"too":        true,
	"under":      true,
	"until":      true,
	"up":         true,
	"very":       true,
	"was":        true,
	"we":         true,
	"were":       true,
	"what":       true,
	"when":       true,
	"where":      true,
	"which":      true,
	"while":      true,
	"who":        true,
	"whom":       true,
	"why":        true,
	"will":       true,
	"with":       true,
	"you":        true,
	"your":       true,
	"yours":      true,
	"yourself":   true,
	"yourselves": true,
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
	if len(word) <= 2 || stopWords[word] {
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
