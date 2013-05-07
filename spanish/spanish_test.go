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
