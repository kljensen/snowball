package english

import (
	"github.com/kljensen/snowball/stemword"
	"log"
	"strings"
)

func logStep(name string, w *stemword.Word) {
	log.Printf("After %v -> %v (%v, %v)", name, w.String(), w.R1String(), w.R2String())
}

func Stem(word string) string {

	word = strings.ToLower(strings.TrimSpace(word))

	// Return small words and stop words
	if len(word) <= 2 || isStopWord(word) {
		return word
	}

	// Return special words
	if specialVersion := stemSpecialWord(word); specialVersion != "" {
		word = specialVersion
		return word
	}

	w := stemword.New(word)
	preprocess(w)
	logStep("preprocess", w)

	r1start, r2start := r1r2(w)
	w.R1start = r1start
	w.R2start = r2start

	_ = step0(w)
	logStep("step 0", w)
	_ = step1a(w)
	logStep("step 1a", w)
	_ = step1b(w)
	logStep("step 1b", w)
	_ = step1c(w)
	logStep("step 1c", w)
	_ = step2(w)
	logStep("step 2", w)
	_ = step3(w)
	logStep("step 3", w)
	_ = step4(w)
	logStep("step 4", w)
	_ = step5(w)
	logStep("step 5", w)

	uncapitalizeYs(w)
	return w.String()

}
