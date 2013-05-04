package english

import (
	"github.com/kljensen/snowball/stemword"
	"log"
)

func step2(w *stemword.Word) bool {

	// Ending for which to check; the order is important.
	endings := [24]string{
		"tional",
		"enci",
		"anci",
		"abli",
		"entli",
		"izer", "ization",
		"ational", "ation", "ator",
		"alism", "aliti", "alli",
		"fulness",
		"ousli", "ousness",
		"iveness", "iviti",
		"biliti", "bli",
		"ogi",
		"fulli",
		"lessli",
		"li",
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
	log.Println(possibleR1Endings[0])

	// Find all endings in R1
	suffix := w.FirstSuffix(possibleR1Endings...)
	log.Println("--", suffix, "--")
	log.Println(w)

	// Handle special cases
	switch suffix {
	case "":
		// No special suffix found
		return false

	case "li":
		// Delete if preceded by a valid li-ending
		// Valid li-endings: cdeghkmnrt
		// cdeghkmnrt = 99 100 101 103 104 107 109 110 114 116
		rsLen := len(w.RS)
		if rsLen >= 3 {
			switch w.RS[rsLen-3] {
			case 99, 100, 101, 103, 104, 107, 109, 110, 114, 116:
				w.ReplaceSuffix(suffix, "", true)
				return true
			}
		}
		return false

	case "ogi":
		// Replace by og if preceded by l
		// l = 108
		rsLen := len(w.RS)
		if rsLen >= 4 && w.RS[rsLen-4] == 108 {
			w.ReplaceSuffix(suffix, "og", true)
		}
		return true
	}

	// Handle basic replacements
	var repl string
	switch suffix {
	case "tional":
		repl = "tion"
	case "enci":
		repl = "ence"
	case "anci":
		repl = "ance"
	case "abli":
		repl = "able"
	case "entli":
		repl = "ent"
	case "izer", "ization":
		repl = "ize"
	case "ational", "ation", "ator":
		repl = "ate"
	case "alism", "aliti", "alli":
		repl = "al"
	case "fulness":
		repl = "ful"
	case "ousli", "ousness":
		repl = "ous"
	case "iveness", "iviti":
		repl = "ive"
	case "biliti", "bli":
		repl = "ble"
	case "fulli":
		repl = "ful"
	case "lessli":
		repl = "less"
	}
	w.ReplaceSuffix(suffix, repl, true)
	return true

}
