package snowball

import (
	"strings"
)

// Step 0 is to strip off apostrophes and "s".
//
func step0(wordIn, r1in, r2in string) (wordOut, r1out, r2out string) {
	wordOut = wordIn
	r1out = r1in
	r2out = r2in
	var step0Suffixes = [3]string{"'s'", "'s", "'"}
	for _, suffix := range step0Suffixes {
		if strings.HasSuffix(wordIn, suffix) {
			wordOut = wordIn[:len(wordIn)-len(suffix)]
			if strings.HasSuffix(r1in, suffix) {
				r1out = r1in[:len(r1in)-len(suffix)]
				if strings.HasSuffix(r2in, suffix) {
					r2out = r2in[:len(r2in)-len(suffix)]
				}
			}
			return
		}
	}
	return
}
