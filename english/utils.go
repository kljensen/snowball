package snowball

import (
	"strings"
	"unicode/utf8"
)

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

// Returns the stemmed version of a word if it is a special
// case, otherwise returns the empty string.
//
func stemSpecialWord(word string) (stemmed string) {
	switch word {
	case "skis":
		stemmed = "ski"
	case "skies":
		stemmed = "sky"
	case "dying":
		stemmed = "die"
	case "lying":
		stemmed = "lie"
	case "tying":
		stemmed = "tie"
	case "idly":
		stemmed = "idl"
	case "gently":
		stemmed = "gentl"
	case "ugly":
		stemmed = "ugli"
	case "early":
		stemmed = "earli"
	case "only":
		stemmed = "onli"
	case "singly":
		stemmed = "singl"
	case "sky":
		stemmed = "sky"
	case "news":
		stemmed = "news"
	case "howe":
		stemmed = "howe"
	case "atlas":
		stemmed = "atlas"
	case "cosmos":
		stemmed = "cosmos"
	case "bias":
		stemmed = "bias"
	case "andes":
		stemmed = "andes"
	case "inning":
		stemmed = "inning"
	case "innings":
		stemmed = "inning"
	case "outing":
		stemmed = "outing"
	case "outings":
		stemmed = "outing"
	case "canning":
		stemmed = "canning"
	case "cannings":
		stemmed = "canning"
	case "herring":
		stemmed = "herring"
	case "herrings":
		stemmed = "herring"
	case "earring":
		stemmed = "earring"
	case "earrings":
		stemmed = "earring"
	case "proceed":
		stemmed = "proceed"
	case "proceeds":
		stemmed = "proceed"
	case "proceeded":
		stemmed = "proceed"
	case "proceeding":
		stemmed = "proceed"
	case "exceed":
		stemmed = "exceed"
	case "exceeds":
		stemmed = "exceed"
	case "exceeded":
		stemmed = "exceed"
	case "exceeding":
		stemmed = "exceed"
	case "succeed":
		stemmed = "succeed"
	case "succeeds":
		stemmed = "succeed"
	case "succeeded":
		stemmed = "succeed"
	case "succeeding":
		stemmed = "succeed"
	}
	return
}

func isStopWord(word string) bool {
	switch word {
	case "a", "about", "above", "after", "again", "against", "all", "am", "an",
		"and", "any", "are", "as", "at", "be", "because", "been", "before",
		"being", "below", "between", "both", "but", "by", "can", "did", "do",
		"does", "doing", "don", "down", "during", "each", "few", "for", "from",
		"further", "had", "has", "have", "having", "he", "her", "here", "hers",
		"herself", "him", "himself", "his", "how", "i", "if", "in", "into", "is",
		"it", "its", "itself", "just", "me", "more", "most", "my", "myself",
		"no", "nor", "not", "now", "of", "off", "on", "once", "only", "or",
		"other", "our", "ours", "ourselves", "out", "over", "own", "s", "same",
		"she", "should", "so", "some", "such", "t", "than", "that", "the", "their",
		"theirs", "them", "themselves", "then", "there", "these", "they",
		"this", "those", "through", "to", "too", "under", "until", "up",
		"very", "was", "we", "were", "what", "when", "where", "which", "while",
		"who", "whom", "why", "will", "with", "you", "your", "yours", "yourself",
		"yourselves":
		return true
	}
	return false
}
