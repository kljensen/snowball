package snowball

import (
	"strings"
)

// Step 1a is noralization of various special "s"-endings.
//
func step1a(wordIn, r1in, r2in string) (wordOut, r1out, r2out string) {
	wordOut = wordIn
	r1out = r1in
	r2out = r2in

	switch {
	case strings.HasSuffix(wordIn, "sses"):
		wordOut, r1out, r2out, _ = replaceWordR1R2Suffix(wordIn, r1in, r2in, "sses", "ss", true)

	case strings.HasSuffix(wordIn, "ied"):
	case strings.HasSuffix(wordIn, "ies"):
		repl := "i"
		if len(wordIn) == 4 {
			repl = "ie"
		}
		var suffix string
		if wordIn[len(wordIn)-1:] == "d" {
			suffix = "ied"
		} else {
			suffix = "ies"
		}
		wordOut, r1out, r2out, _ = replaceWordR1R2Suffix(wordIn, r1in, r2in, suffix, repl, true)

	case strings.HasSuffix(wordIn, "s"):
		runes := []rune(wordIn[:len(wordIn)-1])
		hadVowel := false
		for i := 0; i < len(runes)-2; i++ {
			if isLowerVowel(runes[i]) {
				hadVowel = true
			}
		}
		if hadVowel {
			wordOut, r1out, r2out, _ = replaceWordR1R2Suffix(wordIn, r1in, r2in, "s", "", true)
		}
	}
	return
}
