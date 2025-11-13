package portuguese

import (
	"github.com/kljensen/snowball/snowballword"
)

// Step 2: Verb suffix removal
func step2(word *snowballword.SnowballWord) bool {
	suffix := word.FirstSuffixIfIn(word.RVstart, len(word.RS),
		"aríamos", "eríamos", "iríamos",
		"ássemos", "êssemos", "íssemos",
		"aríeis", "eríeis", "iríeis",
		"ásseis", "ésseis", "ísseis",
		"áramos", "éramos", "íramos",
		"ávamos",
		"aremos", "eremos", "iremos",
		"ariam", "eriam", "iriam",
		"assem", "essem", "issem",
		"aram", "eram", "iram",
		"avam",
		"arem", "erem", "irem",
		"arão", "erão", "irão",
		"ando", "endo", "indo",
		"adas", "idas",
		"arás", "erás", "irás",
		"ares", "eres", "ires",
		"ássei", "ésseis", "ísseis",
		"arias", "erias", "irias",
		"asses", "esses", "isses",
		"astes", "estes", "istes",
		"áreis", "éreis", "íreis",
		"áveis",
		"ados", "idos",
		"amos",
		"arei", "erei", "irei",
		"aria", "eria", "iria",
		"asse", "esse", "isse",
		"aste", "este", "iste",
		"avas",
		"ada", "ida",
		"ara", "era", "ira",
		"ava",
		"ais", "eis",
		"áis", "éis",
		"am", "em",
		"ar", "er", "ir",
		"as", "es",
		"ás", "és",
		"ei",
		"ou",
		"ia",
		"iu",
		"eu",
	)

	if suffix == "" {
		return false
	}

	// Delete if in RV
	suffixLen := len([]rune(suffix))
	if word.RVstart <= len(word.RS)-suffixLen {
		word.RemoveLastNRunes(suffixLen)
		return true
	}

	return false
}
