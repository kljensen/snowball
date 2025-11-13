package romanian

import (
	"github.com/kljensen/snowball/snowballword"
)

// Step 1: Standard suffix removal
func step1(word *snowballword.SnowballWord) bool {
	// Ordered by length
	suffixes := []string{
		// Length 9+
		"abilitate", "abilitati", "abilități", "ibilitate",
		"ivitate", "ivitati", "ivități",
		"icitate", "icitati", "icități",
		// Length 7-8
		"icatori", "icator", "atoare", "ătoare", "atori", "ător", "ători",
		"itoare", "itori", "itor",
		"ațiune", "ițiune",
		// Length 6
		"itatea", "itati", "ități",
		"ativa", "ative", "ativi", "ativă",
		"itiva", "itive", "itivi", "itivă",
		"iciva", "icive", "icivi", "icivă",
		"icala", "icale", "icali", "icalș",
		// Length 5
		"abil", "abila", "abile", "abili", "abilă",
		"ibil", "ibila", "ibile", "ibili", "ibilă",
		"oasa", "oasă", "oase", "oși",
		"anta", "ante", "anti", "antă",
		"iune", "iuni",
		// Length 4
		"ator", "atori",
		"ată", "ați", "ate",
		"ută", "uți", "ute",
		"ită", "iți", "ite",
		"ică", "ice", "ici", "icș",
		"iva", "ive", "ivi", "ivă",
		// Length 3
		"ata", "ati", "uta", "uti", "ita", "iti",
		// Length 2
		"at", "ut", "it", "ic", "iv", "os",
	}

	suffix := word.FirstSuffixIfIn(word.R2start, len(word.RS), suffixes...)
	if suffix != "" {
		word.RemoveLastNRunes(len([]rune(suffix)))
		return true
	}

	return false
}
