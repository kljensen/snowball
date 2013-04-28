package snowball

import "testing"

func TestConstants(t *testing.T) {
	expectedVowels := "aeiouy"
	if vowels != "aeiouy" {
		t.Errorf("Expected %v, got %v", expectedVowels, vowels)
	}
}

// Test stopWords for things we know should be true
// or false.
//
func TestStopwords(t *testing.T) {

	// Test true
	knownTrueStopwords := [...]string{
		"a",
		"for",
		"be",
		"was",
	}
	for _, word := range knownTrueStopwords {
		if stopWords[word] == false {
			t.Errorf("Expected %v, to be in stopWords", word)
		}
	}

	// Test false
	knownFalseStopwords := [...]string{
		"truck",
		"deoxyribonucleic",
		"farse",
		"bullschnizzle",
	}
	for _, word := range knownFalseStopwords {
		if stopWords[word] == true {
			t.Errorf("Expected %v, to be in stopWords", word)
		}
	}
}

// Test specialWords for things we know should be present
// and not present.
//
func TestSpecialwords(t *testing.T) {

	// Test true
	knownTrueSpecialwords := [...]string{
		"exceeding",
		"early",
		"outing",
	}
	for _, word := range knownTrueSpecialwords {
		if _, ok := specialWords[word]; !ok {
			t.Errorf("Expected %v, to be in specialWords", word)
		}
	}

	// Test false
	knownFalseSpecialwords := [...]string{
		"truck",
		"deoxyribonucleic",
		"farse",
		"bullschnizzle",
	}
	for _, word := range knownFalseSpecialwords {
		if _, ok := specialWords[word]; ok {
			t.Errorf("Expected %v, to NOT be in specialWords", word)
		}
	}
}

func TestNormalizeApostrophes(t *testing.T) {
	variants := [...]string{
		"\u2019xxx\u2019",
		"\u2018xxx\u2018",
		"\u201Bxxx\u201B",
		"’xxx’",
		"‘xxx‘",
		"‛xxx‛",
	}
	for _, v := range variants {
		normalizedVersion := normalizeApostrophes(v)
		if normalizedVersion != "'xxx'" {
			t.Errorf("Expected \"'xxx'\", not \"%v\"", normalizedVersion)
		}
	}
}

func TestIsLowerVowel(t *testing.T) {
	for _, r := range vowels {
		if isLowerVowel(r) == false {
			t.Errorf("Expected \"%v\" to be a vowel", r)
		}
	}

	consonant := "bcdfghjklmnpqrstvwxz"
	for _, r := range consonant {
		if isLowerVowel(r) == true {
			t.Errorf("Expected \"%v\" to NOT be a vowel", r)
		}
	}
}

func TestCapitalizeYs(t *testing.T) {
	word1 := "ysdcsdeysdfsysdfsdiyoyyyxyxayxey"
	word2 := capitalizeYs(word1)
	word2Expected := "YsdcsdeYsdfsysdfsdiYoYyYxyxaYxeY"
	if word2 != word2Expected {
		t.Errorf("Expected \"%v\", but got \"%v\"", word2Expected, word2)

	}
}

func TestPreprocessWord(t *testing.T) {
	var wordTests = []struct {
		wordIn  string
		wordOut string
	}{
		{"arguing", "arguing"},
		{"Arguing", "arguing"},
		{"'catty", "catty"},
		{"Kyle’s", "kyle's"},
		{"toy", "toY"},
	}
	for _, testCase := range wordTests {
		output := preprocessWord(testCase.wordIn)
		if output != testCase.wordOut {
			t.Errorf("Expected \"%v\", but got \"%v\"", testCase.wordOut, output)

		}
	}
}
