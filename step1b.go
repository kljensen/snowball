package snowball

// A word is called short if it ends in a short syllable, and if R1 is null. 
// Define a short syllable in a word as either
//  (a) a vowel followed by a non-vowel other than w, x or Y
//      and preceded by a non-vowel, or
//  (b) a vowel at the beginning of the word followed by a non-vowel. 
//
func isShortWord(word, r1 string) (isShort bool) {

	// If r1 is not empty, not short
	if r1 != "" {
		return
	}
	runes := []rune(word)
	numRunes := len(runes)

	if numRunes == 2 {
		if isLowerVowel(runes[0]) && !isLowerVowel(runes[1]) {

			// The word is just two letters, starting with a 
			// vowel and ending with a non-vowel.

			isShort = true
			return
		}
	} else if numRunes >= 3 {

		r3 := runes[len(runes)-1]
		r2 := runes[len(runes)-2]
		r1 := runes[len(runes)-3]
		// w, x, Y rune codepoints = 119, 120, 89
		if !isLowerVowel(r3) && r3 != 119 && r3 != 120 && r3 != 89 && isLowerVowel(r2) && !isLowerVowel(r1) {

			// The word ends in non-vowel, vowel, non-vowel not in wXY
			isShort = true
			return
		}
	}
	return
}

func step1b(wordIn, r1in, r2in string) (wordOut, r1out, r2out string) {
	wordOut = wordIn
	r1out = r1in
	r2out = r2in

	var (
		suffix string
		found  bool
	)

	suffix, found = firstSuffix(wordIn, "eedly", "eed")
	if found {

		// Notice that, the original algorithm is oddly
		// articulated at this step and says, if we found
		// one of these sufficies, to "replace by ee if in R1".
		// The NLTK implementation replaces by "ee" in each
		// of `wordOut`, `r1out`, `r2out`, which is what we've
		// done here.

		wordOut, r1out, r2out, _ = replaceWordR1R2Suffix(wordOut, r1out, r2out, suffix, "ee", true)
		return
	}

	suffix, found = firstSuffix(wordIn, "ed", "edly", "ing", "ingly")
	if found {
		if hasVowel(wordIn, 0, len(suffix)) {
			wordOut, r1out, r2out, _ = replaceWordR1R2Suffix(wordOut, r1out, r2out, suffix, "", true)
			var (
				newSuffix      string
				newSuffixFound bool
			)

			// Check for special ending
			newSuffix, newSuffixFound = firstSuffix(wordOut, "at", "bl", "iz")
			if newSuffixFound {
				wordOut, r1out, r2out, _ = replaceWordR1R2Suffix(wordOut, r1out, r2out, newSuffix, newSuffix+"e", true)
				return
			}

			// Check for double consonant ending.  Note that, the original algorithm
			// implies that all double consonant endings should be removed; however,
			// the NLTK implementation only removes the following sufficies.
			//
			newSuffix, newSuffixFound = firstSuffix(wordOut, "bb", "dd", "ff", "gg",
				"mm", "nn", "pp", "rr", "tt")
			if newSuffixFound {
				wordOut, r1out, r2out, _ = replaceWordR1R2Suffix(wordOut, r1out, r2out, newSuffix, newSuffix[:1], true)
				return
			}

			// Check for a short word
			if isShortWord(wordOut, r1out) {
				// By definition, r1 and r2 are the empty string for
				// short words.
				wordOut = wordOut + "e"
				return
			}
		}

		// return
	}
	return
}
