package finnish

import (
	"github.com/kljensen/snowball/snowballword"
)

var endingRemoved bool

// Step 1: Remove particles and adverbs
func particleEtc(word *snowballword.SnowballWord) bool {
	if !hasParticleEnd(word) {
		return false
	}

	suffixes := []string{
		"kinkin", "kinkaan", "kinkään", "kinko", "kinkö", "kinhan", "kinhän", "kinpa", "kinpä",
		"kin", "kaan", "kään", "ko", "kö", "han", "hän", "pa", "pä", "sti",
	}

	for _, suffix := range suffixes {
		suffixRunes := []rune(suffix)
		if word.HasSuffixRunes(suffixRunes) {
			idx := len(word.RS) - len(suffixRunes)
			// Check if in R1 (or R2 for 'sti')
			if suffix == "sti" {
				if word.R2start <= idx {
					word.RemoveLastNRunes(len(suffixRunes))
					return true
				}
			} else if word.R1start <= idx {
				word.RemoveLastNRunes(len(suffixRunes))
				return true
			}
		}
	}
	return false
}

// Step 2: Remove possessive suffixes
func possessive(word *snowballword.SnowballWord) bool {
	suffixes := []string{
		"nsä", "nsä", "mme", "nne",
		"si", "ni",
		"an", "än", "en",
	}

	for _, suffix := range suffixes {
		suffixRunes := []rune(suffix)
		if word.HasSuffixRunes(suffixRunes) {
			idx := len(word.RS) - len(suffixRunes)
			if word.R1start <= idx {
				switch suffix {
				case "si":
					// Don't delete if preceded by 'k' (ksi = comitative)
					if idx > 0 && word.RS[idx-1] == 'k' {
						continue
					}
					word.RemoveLastNRunes(len(suffixRunes))
					return true
				case "ni":
					word.RemoveLastNRunes(len(suffixRunes))
					// Handle kseni -> ksi
					if len(word.RS) >= 3 && string(word.RS[len(word.RS)-3:]) == "kse" {
						word.RemoveLastNRunes(3)
						word.RS = append(word.RS, []rune("ksi")...)
					}
					return true
				default:
					word.RemoveLastNRunes(len(suffixRunes))
					return true
				}
			}
		}
	}
	return false
}

// Step 3: Remove case endings
func caseEnding(word *snowballword.SnowballWord) bool {
	suffixes := []string{
		"seen", "siin", "tten", "den",
		"ssa", "ssä", "sta", "stä",
		"lla", "llä", "lta", "ltä", "lle",
		"na", "nä", "ksi", "ine",
		"han", "hen", "hin", "hon", "hän", "hön",
		"tta", "ttä", "ta", "tä",
		"a", "ä", "n",
	}

	for _, suffix := range suffixes {
		suffixRunes := []rune(suffix)
		if word.HasSuffixRunes(suffixRunes) {
			idx := len(word.RS) - len(suffixRunes)
			if word.R1start <= idx {
				word.RemoveLastNRunes(len(suffixRunes))
				endingRemoved = true
				return true
			}
		}
	}
	return false
}

// Step 4: Remove other endings (comparative/superlative)
func otherEndings(word *snowballword.SnowballWord) bool {
	suffixes := []string{
		"impi", "impa", "impä", "immi", "imma", "immä",
		"mpi", "mpa", "mpä", "mmi", "mma", "mmä",
		"eja", "ejä",
	}

	for _, suffix := range suffixes {
		suffixRunes := []rune(suffix)
		if word.HasSuffixRunes(suffixRunes) {
			idx := len(word.RS) - len(suffixRunes)
			if word.R2start <= idx {
				// Check 'po' condition for some suffixes
				if suffix == "mpi" || suffix == "mpa" || suffix == "mpä" ||
					suffix == "mmi" || suffix == "mma" || suffix == "mmä" {
					if idx >= 2 && string(word.RS[idx-2:idx]) == "po" {
						continue
					}
				}
				word.RemoveLastNRunes(len(suffixRunes))
				return true
			}
		}
	}
	return false
}

// Step 5: Remove plural 'i' or 'j'
func iPlural(word *snowballword.SnowballWord) bool {
	if len(word.RS) == 0 {
		return false
	}
	lastRune := word.RS[len(word.RS)-1]
	if (lastRune == 'i' || lastRune == 'j') && word.R1start < len(word.RS) {
		word.RemoveLastNRunes(1)
		return true
	}
	return false
}

// Step 6: Remove plural 't'
func tPlural(word *snowballword.SnowballWord) bool {
	if len(word.RS) >= 2 {
		if word.RS[len(word.RS)-1] == 't' && isV1(word.RS[len(word.RS)-2]) {
			if word.R1start < len(word.RS) {
				word.RemoveLastNRunes(1)
				return true
			}
		}
	}
	return false
}

// Step 7: Tidy up
func tidy(word *snowballword.SnowballWord) {
	if len(word.RS) < 2 {
		return
	}

	// Undouble long vowels (aa, ee, ii, oo, uu, ää, öö)
	for len(word.RS) >= 2 {
		last2 := string(word.RS[len(word.RS)-2:])
		if last2 == "aa" || last2 == "ee" || last2 == "ii" ||
			last2 == "oo" || last2 == "uu" || last2 == "ää" || last2 == "öö" {
			word.RemoveLastNRunes(1)
			break
		}
		break
	}

	// Remove trailing a, ä, e, i if preceded by consonant
	if len(word.RS) >= 2 {
		lastRune := word.RS[len(word.RS)-1]
		prevRune := word.RS[len(word.RS)-2]
		if isAEI(lastRune) && isConsonant(prevRune) {
			if word.R1start < len(word.RS) {
				word.RemoveLastNRunes(1)
			}
		}
	}

	// Undouble final consonant
	if len(word.RS) >= 2 {
		last := word.RS[len(word.RS)-1]
		prev := word.RS[len(word.RS)-2]
		if last == prev && isConsonant(last) {
			word.RemoveLastNRunes(1)
		}
	}
}
