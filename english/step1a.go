package snowball

// Step 1a is noralization of various special "s"-endings.
//
func step1a(wordIn, r1in, r2in string) (wordOut, r1out, r2out string) {
	wordOut = wordIn
	r1out = r1in
	r2out = r2in

	suffix, found := firstSuffix(wordIn, "sses", "ied", "ies", "s")
	if found == false {
		return
	}
	switch suffix {
	case "sses":
		wordOut, r1out, r2out, _ = replaceWordR1R2Suffix(wordOut, r1out, r2out, suffix, "ss", true)

	case "ied":
	case "ies":
		var repl string
		if len(wordIn) == 4 {
			repl = "ie"
		} else {
			repl = "i"
		}
		wordOut, r1out, r2out, _ = replaceWordR1R2Suffix(wordOut, r1out, r2out, suffix, repl, true)

	case "s":
		if hasVowel(wordIn, 0, 2) {
			wordOut, r1out, r2out, _ = replaceWordR1R2Suffix(wordOut, r1out, r2out, "s", "", true)
		}
	}
	return
}
