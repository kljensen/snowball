package german

import (
	"testing"
)

func TestStem(t *testing.T) {
	testCases := []struct {
		word     string
		expected string
	}{
		{"lesen", "les"},
		{"bergen", "berg"},
		{"schön", "schon"},
		{"äpfel", "apfel"},
		{"systeme", "system"},
		{"abends", "abend"},
		{"abenteuer", "abenteu"},
		{"abenteuerlich", "abenteu"},
	}

	for _, tc := range testCases {
		result := Stem(tc.word, true)
		if result != tc.expected {
			t.Errorf("Stem(%q) = %q, expected %q", tc.word, result, tc.expected)
		}
	}
}

func TestIsStopWord(t *testing.T) {
	stopWords := []string{"der", "die", "das", "und", "oder", "aber"}
	for _, word := range stopWords {
		if !IsStopWord(word) {
			t.Errorf("IsStopWord(%q) = false, expected true", word)
		}
	}

	nonStopWords := []string{"haus", "baum", "katze"}
	for _, word := range nonStopWords {
		if IsStopWord(word) {
			t.Errorf("IsStopWord(%q) = true, expected false", word)
		}
	}
}
