package snowball

// Replace suffix y or Y by i if preceded by a non-vowel which is not
// the first letter of the word (so cry -> cri, by -> by, say -> say)
//
func step1c(wordIn, r1in, r2in string) (wordOut, r1out, r2out string) {
	wordOut = wordIn
	r1out = r1in
	r2out = r2in

	runes := []rune(wordOut)
	runesLen := len(runes)

	// y = 121
	// Y = 89
	if len(wordIn) > 2 && (runes[runesLen-1] == 121 || runes[runesLen-1] == 89) && !isLowerVowel(runes[runesLen-2]) {
		wordOut, r1out, r2out, _ = replaceWordR1R2Suffix(wordOut, r1out, r2out, string(runes[runesLen-1]), "i", true)
	}
	return
}
