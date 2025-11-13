package dutch

import (
	"github.com/kljensen/snowball/snowballword"
)

// Step 1: Remove common Dutch suffixes
func step1(word *snowballword.SnowballWord) bool {
	suffixes := []string{
		"heden", "eden",
		"heid", "heid",
		"tje", "tjes",
		"ster", "sters",
		"ing", "ingen",
		"aar", "aars",
		"end", "enden",
		"ig", "ige",
		"lijk", "elijke",
		"baar", "bare",
		"en", "e", "s",
	}

	for _, suffix := range suffixes {
		suffixRunes := []rune(suffix)
		if word.HasSuffixRunes(suffixRunes) {
			idx := len(word.RS) - len(suffixRunes)
			if word.R1start <= idx {
				word.RemoveLastNRunes(len(suffixRunes))
				return true
			}
		}
	}
	return false
}

// Step 2: Undouble consonants
func step2(word *snowballword.SnowballWord) {
	if len(word.RS) >= 2 {
		last := word.RS[len(word.RS)-1]
		prev := word.RS[len(word.RS)-2]
		if last == prev && !isVowel(last) {
			word.RemoveLastNRunes(1)
		}
	}
}

// Step 3a: Remove e if preceded by non-vowel
func step3a(word *snowballword.SnowballWord) {
	if len(word.RS) >= 2 && word.RS[len(word.RS)-1] == 'e' {
		if !isVowel(word.RS[len(word.RS)-2]) {
			idx := len(word.RS) - 1
			if word.R1start <= idx {
				word.RemoveLastNRunes(1)
			}
		}
	}
}

// Step 3b: Undouble consonants again
func step3b(word *snowballword.SnowballWord) {
	step2(word)
}
