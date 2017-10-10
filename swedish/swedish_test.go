/*
	Herein lie all the tests of the Swedish snowball stemmer.

*/
package swedish

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
		"och",
		"för",
		"att",
		"inte",
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
		{"öppnade", "nade"},
		{"örnar", "ar"},
		{"vems", "s"},
		{"årorna", "rna"},
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
		{"högtidligheterna", 3, "högtidlig", "tidlig"},
		{"ögats", 3, "ögat", "t"},
		{"ärade", 3, "ärad", "d"},
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
		{"jaktkarlar", true, "jaktkarl"},
		{"jaktkarlarne", true, "jaktkarl"},
		{"klokaste", true, "klok"},
		{"klokheten", true, "klok"},
		{"friskt", true, "frisk"},
		{"fröken", true, "frök"},
		{"kloliknande", true, "klolikn"},
		{"hopplöst", true, "hopplös"},
		{"hopplöshet", true, "hopplös"},
		{"årorna", true, "årorn"},
		// {"skating", true, "skate"},
		// {"fluently", true, "fluentli"},
		// {"ied", true, "ie"},
		// {"ies", true, "ie"},
		// Stop words
		{"vilkas", true, "vilk"},
		{"vilkas", false, "vilkas"},
		// {"above", true, "abov"},
		// {"above", false, "above"},
	}
	for _, tc := range testCases {
		stemmed := Stem(tc.in, tc.stemStopWords)
		if stemmed != tc.out {
			t.Errorf("Expected %v to stem to %v, but got %v", tc.in, tc.out, stemmed)
		}
	}

}
