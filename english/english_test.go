package english

import (
	"github.com/kljensen/snowball/stemword"
	"testing"
)

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
		if isStopWord(word) == false {
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
		if isStopWord(word) == true {
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
		if stemmed := stemSpecialWord(word); stemmed == "" {
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
		if stemmed := stemSpecialWord(word); stemmed != "" {
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
	var wordTests = []struct {
		word  string
		start int
		pos   int
	}{
		{"crepuscular", 0, 4},
		{"uscular", 0, 2},
	}
	for _, tc := range wordTests {
		w := stemword.New(tc.word)
		pos := vnvSuffix(w, tc.start)
		if pos != tc.pos {
			t.Errorf("Expected %v, but got %v", tc.pos, pos)
		}
	}
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
		w := stemword.New(testCase.word)
		r1start, r2start := r1r2(w)
		w.R1start = r1start
		w.R2start = r2start
		if w.R1String() != testCase.r1 || w.R2String() != testCase.r2 {
			t.Errorf("Expected \"{%v, %v}\", but got \"{%v, %v}\"", testCase.r1, testCase.r2, w.R1String(), w.R2String())
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
		w := stemword.New(testCase.word)
		r1start, r2start := r1r2(w)
		w.R1start = r1start
		w.R2start = r2start
		isShort := isShortWord(w)
		if isShort != testCase.isShort {
			t.Errorf("Expected %v, but got %v for \"{%v, %v}\"", testCase.isShort, isShort, testCase.word, w.R1String())
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

type stepFunc func(*stemword.Word) bool
type stepTest struct {
	wordIn  string
	r1start int
	r2start int
	wordOut string
	r1out   string
	r2out   string
}

func runStepTest(t *testing.T, f stepFunc, tcs []stepTest) {
	for _, testCase := range tcs {
		w := stemword.New(testCase.wordIn)
		w.R1start = testCase.r1start
		w.R2start = testCase.r2start
		_ = f(w)
		// t.Log("Replaced =", replaced)
		// t.Log("w.R1start =", w.R1start)
		// t.Log("w.R2start =", w.R2start)
		// t.Log("w.R2start =", w.R2start)
		// t.Log("w.RS =", w.RS)
		if w.String() != testCase.wordOut || w.R1String() != testCase.r1out || w.R2String() != testCase.r2out {
			t.Errorf("Expected \"{%v, %v, %v}\", but got \"{%v, %v, %v}\"", testCase.wordOut, testCase.r1out, testCase.r2out, w.String(), w.R1String(), w.R2String())
		}
	}
}

func Test_step0(t *testing.T) {
	var testCases = []stepTest{
		{"general's", 5, 9, "general", "al", ""},
		{"general's'", 5, 10, "general", "al", ""},
		{"spices'", 4, 7, "spices", "es", ""},
	}
	runStepTest(t, step0, testCases)
}

func Test_step1a(t *testing.T) {
	var testCases = []stepTest{
		{"ties", 0, 0, "tie", "tie", "tie"},
		{"cries", 0, 0, "cri", "cri", "cri"},
		{"mistresses", 3, 7, "mistress", "tress", ""},
	}
	runStepTest(t, step1a, testCases)
}

func Test_step1b(t *testing.T) {

	// I could find immediately conjure up true words to
	// which these cases apply; so, I made some up.

	var testCases = []stepTest{
		{"exxeedly", 1, 8, "exxee", "xxee", ""},
		{"exxeed", 1, 7, "exxee", "xxee", ""},
		{"luxuriated", 3, 5, "luxuriate", "uriate", "iate"},
		{"luxuribled", 3, 5, "luxurible", "urible", "ible"},
		{"luxuriized", 3, 5, "luxuriize", "uriize", "iize"},
		{"luxuriedly", 3, 5, "luxuri", "uri", "i"},
		{"vetted", 3, 6, "vet", "", ""},
		{"hopping", 3, 7, "hop", "", ""},
	}
	runStepTest(t, step1b, testCases)
}

func Test_step1c(t *testing.T) {
	var testCases = []stepTest{
		{"cry", 3, 3, "cri", "", ""},
		{"say", 3, 3, "say", "", ""},
		{"by", 2, 2, "by", "", ""},
		{"xexby", 2, 5, "xexbi", "xbi", ""},
	}
	runStepTest(t, step1c, testCases)
}

func Test_step2(t *testing.T) {
	// Here I've faked R1 & R2 for simplicity
	var testCases = []stepTest{
		// Test "tional"
		{"xxxtional", 3, 5, "xxxtion", "tion", ""},
		// Test when "tional" doesn't fit in R1
		{"xxxtional", 4, 5, "xxxtional", "ional", "onal"},
		// Test "li"
		{"xxxcli", 3, 6, "xxxc", "c", ""},
		// Test "li", non-valid li letter preceeding
		{"xxxxli", 3, 6, "xxxxli", "xli", ""},
		// Test "ogi"
		{"xxlogi", 2, 6, "xxlog", "log", ""},
		// Test "ogi", not preceeded by "l"
		{"xxxogi", 2, 6, "xxxogi", "xogi", ""},
		// Test the others, which are simple replacements
		{"xxxxenci", 3, 7, "xxxxence", "xence", ""},
		{"xxxxanci", 3, 7, "xxxxance", "xance", ""},
		{"xxxxabli", 3, 7, "xxxxable", "xable", ""},
		{"xxxxentli", 3, 8, "xxxxent", "xent", ""},
		{"xxxxizer", 3, 7, "xxxxize", "xize", ""},
		{"xxxxization", 3, 10, "xxxxize", "xize", ""},
		{"xxxxational", 3, 10, "xxxxate", "xate", ""},
		{"xxxxation", 3, 8, "xxxxate", "xate", ""},
		{"xxxxator", 3, 7, "xxxxate", "xate", ""},
		{"xxxxalism", 3, 8, "xxxxal", "xal", ""},
		{"xxxxaliti", 3, 8, "xxxxal", "xal", ""},
		{"xxxxalli", 3, 7, "xxxxal", "xal", ""},
		{"xxxxfulness", 3, 10, "xxxxful", "xful", ""},
		{"xxxxousli", 3, 8, "xxxxous", "xous", ""},
		{"xxxxousness", 3, 10, "xxxxous", "xous", ""},
		{"xxxxiveness", 3, 10, "xxxxive", "xive", ""},
		{"xxxxiviti", 3, 8, "xxxxive", "xive", ""},
		{"xxxxbiliti", 3, 9, "xxxxble", "xble", ""},
		{"xxxxbli", 3, 6, "xxxxble", "xble", ""},
		{"xxxxfulli", 3, 8, "xxxxful", "xful", ""},
		{"xxxxlessli", 3, 8, "xxxxless", "xless", ""},
		// Some of the same words, this time not in our fake R1
		{"xxxxenci", 8, 8, "xxxxenci", "", ""},
		{"xxxxanci", 8, 8, "xxxxanci", "", ""},
		{"xxxxabli", 8, 8, "xxxxabli", "", ""},
		{"xxxxentli", 9, 9, "xxxxentli", "", ""},
		{"xxxxizer", 8, 8, "xxxxizer", "", ""},
		{"xxxxization", 11, 11, "xxxxization", "", ""},
		{"xxxxational", 11, 11, "xxxxational", "", ""},
		{"xxxxation", 9, 9, "xxxxation", "", ""},
		{"xxxxator", 8, 8, "xxxxator", "", ""},
	}
	runStepTest(t, step2, testCases)
}
