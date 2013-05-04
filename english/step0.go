package snowball

// Step 0 is to strip off apostrophes and "s".
//
func step0(wordIn, r1in, r2in string) (wordOut, r1out, r2out string) {
	wordOut = wordIn
	r1out = r1in
	r2out = r2in
	replaced := false
	var step0Suffixes = [3]string{"'s'", "'s", "'"}
	for _, suffix := range step0Suffixes {
		wordOut, r1out, r2out, replaced = replaceWordR1R2Suffix(wordIn, r1in, r2in, suffix, "", false)
		if replaced {
			return
		}
	}
	return
}
