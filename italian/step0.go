package italian

import (
	"github.com/kljensen/snowball/snowballword"
)

// Step 0: Attached pronoun removal
func step0(word *snowballword.SnowballWord) {
	// Pronouns: ci gli la le li lo mi ne si ti vi
	// Compound: sene gliela gliele glieli glielo gliene mela mele meli melo mene
	//           tela tele teli telo tene cela cele celi celo cene vela vele veli velo vene
	pronouns := []string{
		"sene", "gliela", "gliele", "glieli", "glielo", "gliene",
		"mela", "mele", "meli", "melo", "mene",
		"tela", "tele", "teli", "telo", "tene",
		"cela", "cele", "celi", "celo", "cene",
		"vela", "vele", "veli", "velo", "vene",
		"ci", "gli", "la", "le", "li", "lo", "mi", "ne", "si", "ti", "vi",
	}

	suffix := word.FirstSuffixIfIn(word.RVstart, len(word.RS), pronouns...)
	if suffix == "" {
		return
	}

	// Check if preceded by ando/endo (delete) or ar/er/ir (replace with e)
	idx := len(word.RS) - len(suffix)
	if idx >= 4 {
		prev4 := string(word.RS[idx-4 : idx])
		if prev4 == "ando" || prev4 == "endo" {
			if word.RVstart <= idx-4 {
				word.RemoveLastNRunes(len(suffix))
			}
			return
		}
	}
	if idx >= 2 {
		prev2 := string(word.RS[idx-2 : idx])
		if prev2 == "ar" || prev2 == "er" || prev2 == "ir" {
			if word.RVstart <= idx-2 {
				word.RemoveLastNRunes(len(suffix))
				word.RS[len(word.RS)-2] = 'e'
				word.RemoveLastNRunes(1)
			}
		}
	}
}
