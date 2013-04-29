package snowball

import (
	"strings"
)

// Replaces a `suffix` on each of `wordIn`, `r1in`, and `r2in`,
// with `repl`. To indicate that `wordIn` is known to end in `suffix`,
// set `known` to true.  Here, we assume that `r2in` is a suffix of
// `r1in`, and both are sufficies of `wordIn`.  If that is not the 
// case, you will not get the results you intend.
//
func replaceWordR1R2Suffix(wordIn, r1in, r2in, suffix, repl string, known bool) (wordOut, r1out, r2out string, replaced bool) {
	wordOut = wordIn
	r1out = r1in
	r2out = r2in
	if known || strings.HasSuffix(wordIn, suffix) {
		wordOut = wordIn[:len(wordIn)-len(suffix)] + repl
		if strings.HasSuffix(r1in, suffix) {
			r1out = r1in[:len(r1in)-len(suffix)] + repl
			if strings.HasSuffix(r2in, suffix) {
				r2out = r2in[:len(r2in)-len(suffix)] + repl
			} else {
				r2out = ""
			}
		} else {
			r1out = ""
			r2out = ""
		}
		replaced = true
	}
	return
}

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
