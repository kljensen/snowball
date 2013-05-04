package english

import (
	"github.com/kljensen/snowball/stemword"
	"strings"
	"unicode/utf8"
)

// Finds the region after the first non-vowel following a vowel,
// or a the null region at the end of the word if there is no
// such non-vowel.  Returns the index in the Word where the 
// region starts; optionally skips the first `start` characters.
//
func vnvSuffix(word *stemword.Word, start int) int {
	for i := 1; i < len(word.RS[start:]); i++ {
		j := start + i
		if isLowerVowel(word.RS[j-1]) && !isLowerVowel(word.RS[j]) {
			return j + 1
		}
	}
	return len(word.RS)
}

// Find the starting point of the two regions R1 & R2.
//
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
func r1r2(word *stemword.Word) (r1start, r2start int) {

	specialPrefix := word.FirstPrefix("gener", "commun", "arsen")

	if specialPrefix != "" {
		r1start = len(specialPrefix)
	} else {
		r1start = vnvSuffix(word, 0)
	}
	r2start = vnvSuffix(word, r1start)
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

// Returns the first matching suffix
//
func firstSuffix(word string, sufficies ...string) (suffix string, found bool) {
	for _, suffix = range sufficies {
		if strings.HasSuffix(word, suffix) {
			found = true
			return
		}
	}
	suffix = ""
	return
}

// Replaces a `suffix` on each of `wordIn`, `r1in`, and `r2in`,
// with `repl`. To indicate that `wordIn` is known to end in `suffix`,
// set `known` to true.  Here, we assume that `r2in` is a suffix of
// `r1in`, and both are sufficies of `wordIn`.  If that is not the 
// case, you will not get the results you intend.
//
func replaceWordR1R2Suffix(wordIn, r1in, r2in, suffix, repl string, known bool) (wordOut, r1out, r2out string, replaced bool) {
	wordOut = wordIn
	r1out = r1in
	r2out = r2in
	suffixLen := len(suffix)
	if known || strings.HasSuffix(wordIn, suffix) {
		wordOut = wordIn[:len(wordIn)-suffixLen] + repl
		r1len := len(r1in)
		if suffixLen <= r1len {
			r1out = r1in[:r1len-suffixLen] + repl
			r2len := len(r2in)
			if suffixLen <= r2len {
				r2out = r2in[:r2len-suffixLen] + repl
			} else {
				r2out = ""
			}
		} else {
			r1out = ""
			r2out = ""
		}
		replaced = true
	}
	return
}

// A word is called short if it ends in a short syllable, and if R1 is null. 
// Define a short syllable in a word as either
//  (a) a vowel followed by a non-vowel other than w, x or Y
//      and preceded by a non-vowel, or
//  (b) a vowel at the beginning of the word followed by a non-vowel. 
//
func isShortWord(w *stemword.Word) (isShort bool) {

	// If r1 is not empty, not short
	if w.R1start < len(w.RS) {
		return
	}

	if len(w.RS) == 2 {
		if isLowerVowel(w.RS[0]) && !isLowerVowel(w.RS[1]) {

			// The word is just two letters, starting with a 
			// vowel and ending with a non-vowel.

			isShort = true
			return
		}
	} else if len(w.RS) >= 3 {

		s1 := w.RS[len(w.RS)-1]
		s2 := w.RS[len(w.RS)-2]
		s3 := w.RS[len(w.RS)-3]
		// w, x, Y rune codepoints = 119, 120, 89
		if !isLowerVowel(s1) && s1 != 119 && s1 != 120 && s1 != 89 && isLowerVowel(s2) && !isLowerVowel(s3) {

			// The word ends in non-vowel, vowel, non-vowel not in wXY
			isShort = true
			return
		}
	}
	return
}
