package spanish

import (
	// "github.com/kljensen/snowball/snowballword"
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
