package spanish

import (
	"github.com/kljensen/snowball/snowballword"
	"testing"
)

// Test stopWords for things we know should be true
// or false.
//
func Test_stopWords(t *testing.T) {
	testCases := []struct {
		word   string
		result bool
	}{
		{"el", true},
		{"queso", false},
	}
	for _, testCase := range testCases {
		result := isStopWord(testCase.word)
		if result != testCase.result {
			t.Errorf("Expect isStopWord(\"%v\") = %v, but got %v",
				testCase.word, testCase.result, result,
			)
		}
	}
}

// Test isLowerVowel for things we know should be true
// or false.
//
func Test_isLowerVowel(t *testing.T) {
	testCases := []struct {
		chars  string
		result bool
	}{
		// These are all vowels.
		{"aeiouáéíóúü", true},
		// None of these are vowels.
		{"cbfqhkl", false},
	}
	for _, testCase := range testCases {
		for _, r := range testCase.chars {
			result := isLowerVowel(r)
			if result != testCase.result {
				t.Errorf("Expect isLowerVowel(\"%v\") = %v, but got %v",
					r, testCase.result, result,
				)
			}

		}
	}
}

// Test isLowerVowel for things we know should be true
// or false.
//
func Test_findRegions(t *testing.T) {
	testCases := []struct {
		word    string
		r1start int
		r2start int
		rvstart int
	}{
		{"macho", 3, 5, 4},
		{"olivia", 2, 4, 3},
		{"trabajo", 4, 6, 3},
		{"áureo", 3, 5, 3},
	}
	for _, testCase := range testCases {
		w := snowballword.New(testCase.word)
		r1start, r2start, rvstart := findRegions(w)
		if r1start != testCase.r1start || r2start != testCase.r2start || rvstart != testCase.rvstart {
			t.Errorf("Expect findRegions(\"%v\") = %v, %v, %v, but got %v, %v, %v",
				testCase.word, testCase.r1start, testCase.r2start, testCase.rvstart,
				r1start, r2start, rvstart,
			)
		}

	}
}
