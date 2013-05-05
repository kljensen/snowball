package english

import (
	"github.com/kljensen/snowball/stemword"
)

// Step 4:
// Search for the longest among the following suffixes,
// and, if found and in R2, perform the action indicated. 

// al, ance, ence, er, ic, able, ible, ant, ement, ment,
// ent, ism, ate, iti, ous, ive, ize
// delete
//
// ion
// delete if preceded by s or t
func step4(w *stemword.Word) bool {

	// Ending for which to check, longest first.
	endings := [18]string{
		"ement",
		"ance",
		"ence",
		"able",
		"ible",
		"ment",
		"ent",
		"ant",
		"ism",
		"ate",
		"iti",
		"ous",
		"ive",
		"ize",
		"ion",
		"al",
		"er",
		"ic",
	}

	// Filter out those endings that are too long to be in R1
	r2Len := len(w.RS) - w.R2start
	possibleR2Endings := make([]string, 0, len(endings))
	for _, ending := range endings {
		if len(ending) <= r2Len {
			possibleR2Endings = append(possibleR2Endings, ending)
		}
	}
	if len(possibleR2Endings) == 0 {
		return false
	}

	// Find all endings in R1
	suffix := w.FirstSuffix(possibleR2Endings...)

	// Handle special cases
	switch suffix {
	case "":
		return false
	case "ion":
		// Replace by og if preceded by l
		// l = 108
		rsLen := len(w.RS)
		if rsLen >= 4 {
			switch w.RS[rsLen-4] {
			case 115, 116:
				w.ReplaceSuffix(suffix, "", true)
				return true
			}

		}
		return false
	}

	// Handle basic replacements
	w.ReplaceSuffix(suffix, "", true)
	return true

}
