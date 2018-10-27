/*
	Herein lie all the tests of the norwegian snowball stemmer.
	TODO
*/
package norwegian

import (
	"testing"

	"github.com/kljensen/snowball/snowballword"
)

// Test stopWords for things we know should be true
// or false.
//
func Test_stopWords(t *testing.T) {

	// Test true
	knownTrueStopwords := [...]string{
		"og",
		"for",
		"mye",
		"ikke",
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

func Test_r1(t *testing.T) {
	var wordTests = []struct {
		word string
		r1   string
	}{
		{"åpnet", "et"},
		{"åpner", "er"},
		{"hvems", "s"},
		{"ørene", "ne"},
		// Special cases below
	}
	for _, testCase := range wordTests {
		w := snowballword.New(testCase.word)
		r1start := r1(w)
		w.R1start = r1start
		if w.R1String() != testCase.r1 {
			t.Errorf("Expected \"{%v}\", but got \"{%v}\"", testCase.r1, w.R1String())
		}
	}
}

type stepFunc func(*snowballword.SnowballWord) bool
type stepTest struct {
	wordIn  string
	r1start int
	wordOut string
	r1out   string
}

func runStepTest(t *testing.T, f stepFunc, tcs []stepTest) {
	for _, testCase := range tcs {
		w := snowballword.New(testCase.wordIn)
		w.R1start = testCase.r1start
		_ = f(w)
		if w.String() != testCase.wordOut || w.R1String() != testCase.r1out {
			t.Errorf("Expected \"{%v, %v}\", but got \"{%v, %v}\"", testCase.wordOut, testCase.r1out, w.String(), w.R1String())
		}
	}
}

func Test_step1(t *testing.T) {
	var testCases = []stepTest{
		{"høytidlighetene", 3, "høytidlig", "tidlig"},
		{"øyets", 3, "øyet", "t"},
		{"ørets", 3, "øret", "t"},
	}
	runStepTest(t, step1, testCases)
}

func Test_step2(t *testing.T) {
	var testCases = []stepTest{}
	runStepTest(t, step2, testCases)
}

func Test_step3(t *testing.T) {
	var testCases = []stepTest{
		{"årlig", 3, "årl", ""},
	}
	runStepTest(t, step3, testCases)
}

func Test_Stem(t *testing.T) {
	var testCases = []struct {
		in            string
		stemStopWords bool
		out           string
	}{
		{"havnedistrikt", true, "havnedistrikt"},
		{"havnedistriktene", true, "havnedistrikt"},
		{"havnedistrikter", true, "havnedistrikt"},
		{"havnedistriktets", true, "havnedistrikt"},
		{"havnedistriktets", true, "havnedistrikt"},
		{"opp", true, "opp"},
		{"oppad", true, "oppad"},
		{"opning", true, "opning"},
		{"havneinteresser", true, "havneinteress"},
		{"oppbygginga", true, "oppbygging"},
		{"oppbyggingen", true, "oppbygging"},
		{"oppdaterte", true, "oppdater"},
		{"tredjepersons", true, "tredjeperson"},
		{"uspesisfisert", true, "uspesisfiser"},
		{"voks", true, "voks"},
	}
	for _, tc := range testCases {
		stemmed := Stem(tc.in, tc.stemStopWords)
		if stemmed != tc.out {
			t.Errorf("Expected %v to stem to %v, but got %v", tc.in, tc.out, stemmed)
		}
	}

}
