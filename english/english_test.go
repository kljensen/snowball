/*
	Herein lie all the tests of the Snowball English stemmer.

	Many of the tests are drawn from cases where this implementation
	did not match the results of the Python NLTK implementation.
*/
package english

import (
	"testing"

	"github.com/kljensen/snowball/romance"
	"github.com/kljensen/snowball/snowballword"
)

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
		w := snowballword.New(v)
		normalizeApostrophes(w)
		if w.String() != "'xxx'" {
			t.Errorf("Expected \"'xxx'\", not \"%v\"", w.String())
		}
	}
}

func Test_capitalizeYs(t *testing.T) {
	var wordTests = []struct {
		in  string
		out string
	}{
		{"ysdcsdeysdfsysdfsdiyoyyyxyxayxey", "YsdcsdeYsdfsysdfsdiYoYyYxyxaYxeY"},
	}
	for _, wt := range wordTests {
		w := snowballword.New(wt.in)
		capitalizeYs(w)
		if w.String() != wt.out {
			t.Errorf("Expected \"%v\", not \"%v\"", wt.out, w.String())
		}
	}
}
func Test_preprocess(t *testing.T) {
	var wordTests = []struct {
		in  string
		out string
	}{
		{"arguing", "arguing"},
		{"'catty", "catty"},
		{"kyle’s", "kyle's"},
		{"toy", "toY"},
	}
	for _, wt := range wordTests {
		w := snowballword.New(wt.in)
		preprocess(w)
		if w.String() != wt.out {
			t.Errorf("Expected \"%v\", not \"%v\"", wt.out, w.String())
		}
	}
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
		w := snowballword.New(tc.word)
		pos := romance.VnvSuffix(w, isLowerVowel, tc.start)
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
		w := snowballword.New(testCase.word)
		r1start, r2start := r1r2(w)
		w.R1start = r1start
		w.R2start = r2start
		if w.R1String() != testCase.r1 || w.R2String() != testCase.r2 {
			t.Errorf("Expected \"{%v, %v}\", but got \"{%v, %v}\"", testCase.r1, testCase.r2, w.R1String(), w.R2String())
		}
	}
}

func Test_isShortWord(t *testing.T) {
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
		w := snowballword.New(testCase.word)
		r1start, r2start := r1r2(w)
		w.R1start = r1start
		w.R2start = r2start
		isShort := isShortWord(w)
		if isShort != testCase.isShort {
			t.Errorf("Expected %v, but got %v for \"{%v, %v}\"", testCase.isShort, isShort, testCase.word, w.R1String())
		}
	}
}

func Test_endsShortSyllable(t *testing.T) {
	var testCases = []struct {
		word   string
		pos    int
		result bool
	}{
		{"absolute", 7, true},
		{"ape", 2, true},
		{"rap", 3, true},
		{"trap", 4, true},
		{"entrap", 6, true},
		{"uproot", 6, false},
		{"bestow", 6, false},
		{"disturb", 7, false},
	}
	for _, testCase := range testCases {
		w := snowballword.New(testCase.word)
		result := endsShortSyllable(w, testCase.pos)
		if result != testCase.result {
			t.Errorf("Expected endsShortSyllable(%v, %v) to return %v, not %v", testCase.word, testCase.pos, testCase.result, result)
		}
	}

}

type stepFunc func(*snowballword.SnowballWord) bool
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
		w := snowballword.New(testCase.wordIn)
		w.R1start = testCase.r1start
		w.R2start = testCase.r2start
		_ = f(w)
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
		{"mistresses", 3, 7, "mistress", "tress", "s"},
		{"ied", 3, 3, "ie", "", ""},
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
		{"breed", 5, 5, "breed", "", ""},
		{"skating", 4, 6, "skate", "e", ""},
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
		{"fluentli", 5, 8, "fluentli", "tli", ""},
		// Test "tional"
		{"xxxtional", 3, 5, "xxxtion", "tion", "on"},
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
		{"xxxxenci", 3, 7, "xxxxence", "xence", "e"},
		{"xxxxanci", 3, 7, "xxxxance", "xance", "e"},
		{"xxxxabli", 3, 7, "xxxxable", "xable", "e"},
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
		{"xxxxbli", 3, 6, "xxxxble", "xble", "e"},
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

func Test_step4(t *testing.T) {
	var testCases = []stepTest{
		{"accumulate", 2, 5, "accumul", "cumul", "ul"},
		{"agreement", 2, 6, "agreement", "reement", "ent"},
	}
	runStepTest(t, step4, testCases)
}
func Test_step5(t *testing.T) {
	var testCases = []stepTest{
		{"skate", 4, 5, "skate", "e", ""},
	}
	runStepTest(t, step5, testCases)
}

func Test_Stem(t *testing.T) {
	var testCases = []struct {
		in            string
		stemStopWords bool
		out           string
	}{
		{"aberration", true, "aberr"},
		{"abruptness", true, "abrupt"},
		{"absolute", true, "absolut"},
		{"abated", true, "abat"},
		{"acclivity", true, "accliv"},
		{"accumulations", true, "accumul"},
		{"agreement", true, "agreement"},
		{"breed", true, "breed"},
		{"ape", true, "ape"},
		{"skating", true, "skate"},
		{"fluently", true, "fluentli"},
		{"ied", true, "ie"},
		{"ies", true, "ie"},
		// Stop words
		{"because", true, "becaus"},
		{"because", false, "because"},
		{"above", true, "abov"},
		{"above", false, "above"},
	}
	for _, tc := range testCases {
		stemmed := Stem(tc.in, tc.stemStopWords)
		if stemmed != tc.out {
			t.Errorf("Expected %v to stem to %v, but got %v", tc.in, tc.out, stemmed)
		}
	}

}
