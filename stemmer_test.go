package snowball

import "testing"

// Stuct for holding tests where a word is transformed
// into another by the function to be tested.
//
type simpleStringTestCase struct {
	wordIn  string
	wordOut string
}

// A type representing all functions that take one
// string and return another.
type simpleStringFunction func(string) string

// Runs a series of test cases for functions that just
// transform one string into another.
//
func runSimpleStringTests(t *testing.T, f simpleStringFunction, tcs []simpleStringTestCase) {
	for _, testCase := range tcs {
		output := f(testCase.wordIn)
		if output != testCase.wordOut {
			t.Errorf("Expected \"%v\", but got \"%v\"", testCase.wordOut, output)
		}
	}
}

func Test_constants(t *testing.T) {
	expectedVowels := "aeiouy"
	if vowels != "aeiouy" {
		t.Errorf("Expected %v, got %v", expectedVowels, vowels)
	}
}

// Test stopWords for things we know should be true
// or false.
//
func Test_stopWords(t *testing.T) {

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
func Test_specialWords(t *testing.T) {

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

func Test_normalizeApostrophes(t *testing.T) {
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

func Test_isLowerVowel(t *testing.T) {
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

func Test_capitalizeYs(t *testing.T) {
	var wordTests = []simpleStringTestCase{
		{"ysdcsdeysdfsysdfsdiyoyyyxyxayxey", "YsdcsdeYsdfsysdfsdiYoYyYxyxaYxeY"},
	}
	runSimpleStringTests(t, capitalizeYs, wordTests)
}

func Test_preprocessWord(t *testing.T) {
	var wordTests = []simpleStringTestCase{
		{"arguing", "arguing"},
		{"Arguing", "arguing"},
		{"'catty", "catty"},
		{"Kyle’s", "kyle's"},
		{"toy", "toY"},
	}
	runSimpleStringTests(t, preprocessWord, wordTests)
}

func Test_vnvSuffix(t *testing.T) {
	var wordTests = []simpleStringTestCase{
		{"crepuscular", "uscular"},
		{"uscular", "cular"},
	}
	runSimpleStringTests(t, vnvSuffix, wordTests)
}

func Test_r1r2(t *testing.T) {
	var wordTests = []struct {
		word string
		r1   string
		r2   string
	}{
		{"crepuscular", "uscular", "cular"},
		{"beautiful", "iful", "ul"},
		{"beauty", "y", ""},
		{"eucharist", "harist", "ist"},
		{"animadversion", "imadversion", "adversion"},
		{"mistresses", "tresses", "ses"},
		{"sprinkled", "kled", ""},
		// Special cases below
		{"communism", "ism", "m"},
		{"arsenal", "al", ""},
		{"generalities", "alities", "ities"},
		{"embed", "bed", ""},
	}
	for _, testCase := range wordTests {
		r1, r2 := r1r2(testCase.word)
		if r1 != testCase.r1 || r2 != testCase.r2 {
			t.Errorf("Expected \"{%v, %v}\", but got \"{%v, %v}\"", testCase.r1, testCase.r2, r1, r2)
		}
	}
}

func Test_replaceWordR1R2Suffix(t *testing.T) {
	var wordTests = []struct {
		wordIn   string
		r1in     string
		r2in     string
		suffix   string
		repl     string
		known    bool
		wordOut  string
		r1out    string
		r2out    string
		replaced bool
	}{
		{"animadversion", "imadversion", "adversion", "version", "", false, "animad", "imad", "ad", true},
		{"animadversion", "imadversion", "adversion", "version", "", true, "animad", "imad", "ad", true},
		{"animadversion", "imadversion", "", "version", "", true, "animad", "imad", "", true},
		{"animadversion", "imadversion", "adversion", "version", "xx", false, "animadxx", "imadxx", "adxx", true},
		{"animadversion", "imadversion", "adversion", "versionXX", "yy", false, "animadversion", "imadversion", "adversion", false},
		{"animadversion", "imadversion", "adversion", "versionXX", "yy", false, "animadversion", "imadversion", "adversion", false},
		{"vett", "t", "", "tt", "t", true, "vet", "", "", true},
	}
	for _, testCase := range wordTests {
		wordOut, r1out, r2out, replaced := replaceWordR1R2Suffix(testCase.wordIn, testCase.r1in, testCase.r2in, testCase.suffix, testCase.repl, testCase.known)
		if wordOut != testCase.wordOut || r1out != testCase.r1out || r2out != testCase.r2out || replaced != testCase.replaced {
			t.Errorf("Expected \"{%v, %v, %v, %v}\", but got \"{%v, %v, %v, %v}\"", testCase.wordOut, testCase.r1out, testCase.r2out, testCase.replaced, wordOut, r1out, r2out, replaced)
		}
	}
}

func Test_hasVowel(t *testing.T) {
	var testCases = []struct {
		word      string
		leftSkip  int
		rightSkip int
		result    bool
	}{
		{"television", 0, 0, true},
		{"xxxaxxx", 0, 0, true},
		{"xxxaxxx", 1, 1, true},
		{"xxxaxxx", 1, 4, false},
		{"xxxaxxx", 4, 1, false},
		{"xxxaxxx", 4, 10, false},
		{"xxxaxxx", 4, 5, false},
		{"axxxxxx", 1, 0, false},
		{"xxxxxxa", 0, 1, false},
	}
	for _, testCase := range testCases {
		result := hasVowel(testCase.word, testCase.leftSkip, testCase.rightSkip)
		if result != testCase.result {
			t.Errorf("Expected %v, but got %v for \"{%v, %v, %v}\"", testCase.result, result, testCase.word, testCase.leftSkip, testCase.rightSkip)
		}
	}
}

func Test_isShortWord(t *testing.T) {
	// bed, shed and shred are short words, bead, embed, beds are not short words. 
	var testCases = []struct {
		word    string
		isShort bool
	}{
		{"bed", true},
		{"shed", true},
		{"shred", true},
		{"bead", false},
		{"embed", false},
		{"beds", false},
	}
	for _, testCase := range testCases {
		r1, _ := r1r2(testCase.word)
		isShort := isShortWord(testCase.word, r1)
		if isShort != testCase.isShort {
			t.Errorf("Expected %v, but got %v for \"{%v, %v}\"", testCase.isShort, isShort, testCase.word, r1)
		}
	}
}
func Test_firstSuffix(t *testing.T) {
	var testCases = []struct {
		word      string
		sufficies []string
		suffix    string
		found     bool
	}{
		{"mistresses", []string{"tresses", "ses"}, "tresses", true},
	}
	for _, testCase := range testCases {
		suffix, found := firstSuffix(testCase.word, testCase.sufficies...)
		if suffix != testCase.suffix || found != testCase.found {
			t.Errorf("Expected \"{%v, %v, %v}\", but got \"{%v, %v, %v}\"", testCase.suffix, testCase.found, suffix, found)
		}
	}
}

type stepFunc func(string, string, string) (string, string, string)
type stepTest struct {
	wordIn  string
	r1in    string
	r2in    string
	wordOut string
	r1out   string
	r2out   string
}

func runStepTest(t *testing.T, f stepFunc, tcs []stepTest) {
	for _, testCase := range tcs {
		wordOut, r1out, r2out := f(testCase.wordIn, testCase.r1in, testCase.r2in)
		if wordOut != testCase.wordOut || r1out != testCase.r1out || r2out != testCase.r2out {
			t.Errorf("Expected \"{%v, %v, %v}\", but got \"{%v, %v, %v}\"", testCase.wordOut, testCase.r1out, testCase.r2out, wordOut, r1out, r2out)
		}
	}
}

func Test_step0(t *testing.T) {
	var testCases = []stepTest{
		{"general's", "al's", "", "general", "al", ""},
		{"general's'", "al's'", "", "general", "al", ""},
		{"spices'", "es'", "", "spices", "es", ""},
	}
	runStepTest(t, step0, testCases)
}

func Test_step1a(t *testing.T) {
	var testCases = []stepTest{
		{"ties", "ties", "ties", "tie", "tie", "tie"},
		{"cries", "cries", "cries", "cri", "cri", "cri"},
		{"mistresses", "tresses", "ses", "mistress", "tress", ""},
	}
	runStepTest(t, step1a, testCases)
}

func Test_step1b(t *testing.T) {

	// I could find immediately conjure up true words to
	// which these cases apply; so, I made some up.

	var testCases = []stepTest{
		{"exxeedly", "xxeedly", "", "exxee", "xxee", ""},
		{"exxeed", "xxeed", "", "exxee", "xxee", ""},
		{"luxuriated", "uriated", "iated", "luxuriate", "uriate", "iate"},
		{"luxuribled", "uribled", "ibled", "luxurible", "urible", "ible"},
		{"luxuriized", "uriized", "iized", "luxuriize", "uriize", "iize"},
		{"luxuriedly", "uriedly", "iedly", "luxuri", "uri", "i"},
		{"vetted", "ted", "", "vet", "", ""},
		{"hopping", "ping", "", "hop", "", ""},
	}
	runStepTest(t, step1b, testCases)
}

func Test_step1c(t *testing.T) {
	var testCases = []stepTest{
		{"cry", "", "", "cri", "", ""},
		{"say", "", "", "say", "", ""},
		{"by", "", "", "by", "", ""},
		{"xexby", "xby", "", "xexbi", "xbi", ""},
	}
	runStepTest(t, step1c, testCases)
}
