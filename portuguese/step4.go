package portuguese

import (
	"github.com/kljensen/snowball/snowballword"
)

// Step 4: Residual form
func step4(word *snowballword.SnowballWord) {
	suffix := word.FirstSuffixIfIn(word.RVstart, len(word.RS),
		"é", "ê", "e",
	)

	if suffix != "" {
		word.RemoveLastNRunes(len(suffix))

		// If removed 'e' or 'ê' or 'é', and preceded by 'gu' or 'ci', remove the 'u' or 'i'
		if len(word.RS) >= 2 {
			last2 := string(word.RS[len(word.RS)-2:])
			if last2 == "gu" || last2 == "ci" {
				if word.RVstart <= len(word.RS)-1 {
					word.RemoveLastNRunes(1)
				}
			}
		}
	}

	// Also try ç -> c
	if len(word.RS) > 0 && word.RS[len(word.RS)-1] == 'ç' {
		word.RS[len(word.RS)-1] = 'c'
	}
}
