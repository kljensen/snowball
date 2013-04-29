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
			isShort = true
			return
		}
	} else if numRunes >= 3 {
		r3 := runes[len(runes)-1]
		r2 := runes[len(runes)-2]
		r1 := runes[len(runes)-3]
		// w, x, Y rune codepoints = 119, 120, 89
		if !isLowerVowel(r3) && r3 != 119 && r3 != 120 && r3 != 89 && isLowerVowel(r2) && !isLowerVowel(r1) {
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
				wordOut, r1out, r2out, _ = replaceWordR1R2Suffix(wordOut, r1out, r2out, suffix, suffix+"e", true)
				return
			}

			// Check for double consonant ending
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

// for suffix in self.__step1b_suffixes:
//     if word.endswith(suffix):
//         if suffix in ("eed", "eedly"):

//             if r1.endswith(suffix):
//                 word = "".join((word[:-len(suffix)], "ee"))

//                 if len(r1) >= len(suffix):
//                     r1 = "".join((r1[:-len(suffix)], "ee"))
//                 else:
//                     r1 = ""

//                 if len(r2) >= len(suffix):
//                     r2 = "".join((r2[:-len(suffix)], "ee"))
//                 else:
//                     r2 = ""
//         else:
//             for letter in word[:-len(suffix)]:
//                 if letter in self.__vowels:
//                     step1b_vowel_found = True
//                     break

//             if step1b_vowel_found:
//                 word = word[:-len(suffix)]
//                 r1 = r1[:-len(suffix)]
//                 r2 = r2[:-len(suffix)]

//                 if word.endswith(("at", "bl", "iz")):
//                     word = "".join((word, "e"))
//                     r1 = "".join((r1, "e"))

//                     if len(word) > 5 or len(r1) >=3:
//                         r2 = "".join((r2, "e"))

//                 elif word.endswith(self.__double_consonants):
//                     word = word[:-1]
//                     r1 = r1[:-1]
//                     r2 = r2[:-1]

//                 elif ((r1 == "" and len(word) >= 3 and
//                        word[-1] not in self.__vowels and
//                        word[-1] not in "wxY" and
//                        word[-2] in self.__vowels and
//                        word[-3] not in self.__vowels)
//                       or
//                       (r1 == "" and len(word) == 2 and
//                        word[0] in self.__vowels and
//                        word[1] not in self.__vowels)):

//                     word = "".join((word, "e"))

//                     if len(r1) > 0:
//                         r1 = "".join((r1, "e"))

//                     if len(r2) > 0:
//                         r2 = "".join((r2, "e"))
//         break
