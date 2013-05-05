package english

import (
	"github.com/kljensen/snowball/stemword"
)

// Search for the longest among the following suffixes,
// and, if found and in R1, perform the action indicated. 
// 
// tional:   replace by tion
// ational:   replace by ate
// alize:   replace by al
// icate   iciti   ical:   replace by ic
// ful   ness:   delete
// 
func step3(w *stemword.Word) bool {

	// Ending for which to check, longest first.
	endings := [7]string{
		"ational",
		"tional",
		"alize",
		"icate",
		"iciti",
		"ical",
		"ful",
	}

	// Filter out those endings that are too long to be in R1
	r1Len := len(w.RS) - w.R1start
	possibleR1Endings := make([]string, 0, len(endings))
	for _, ending := range endings {
		if len(ending) <= r1Len {
			possibleR1Endings = append(possibleR1Endings, ending)
		}
	}
	if len(possibleR1Endings) == 0 {
		return false
	}

	// Find all endings in R1
	suffix := w.FirstSuffix(possibleR1Endings...)

	// Handle special cases
	if suffix == "" {
		return false
	}

	// Handle basic replacements
	var repl string
	switch suffix {
	case "ational":
		repl = "ate"
	case "tional":
		repl = "tion"
	case "alize":
		repl = "al"
	case "icate", "iciti", "ical":
		repl = "ic"
	case "ful", "ness":
		repl = ""
	}
	w.ReplaceSuffix(suffix, repl, true)
	return true

}
