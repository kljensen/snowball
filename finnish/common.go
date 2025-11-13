package finnish

import (
	"github.com/kljensen/snowball/snowballword"
)

// Finnish vowels: a e i o u y ä ö
func isV1(r rune) bool {
	switch r {
	case 'a', 'e', 'i', 'o', 'u', 'y':
		return true
	case 0x00E4, 0x00F6: // ä ö
		return true
	}
	return false
}


// Finnish consonants
func isConsonant(r rune) bool {
	return !isV1(r)
}

// Check if AEI vowel
func isAEI(r rune) bool {
	switch r {
	case 'a', 'e', 'i', 0x00E4: // a e i ä
		return true
	}
	return false
}

// Mark regions p1 and p2
func markRegions(word *snowballword.SnowballWord) {
	// p1: after first V1, then first non-V1
	p1 := len(word.RS)
	for i := 0; i < len(word.RS)-1; i++ {
		if isV1(word.RS[i]) {
			for j := i + 1; j < len(word.RS); j++ {
				if !isV1(word.RS[j]) {
					p1 = j + 1
					break
				}
			}
			break
		}
	}

	// p2: after first V1 in p1 region, then first non-V1
	p2 := len(word.RS)
	for i := p1; i < len(word.RS)-1; i++ {
		if isV1(word.RS[i]) {
			for j := i + 1; j < len(word.RS); j++ {
				if !isV1(word.RS[j]) {
					p2 = j + 1
					break
				}
			}
			break
		}
	}

	word.R1start = p1
	word.R2start = p2
}

// Check if word ends with particle_end (V1 or 'nt')
func hasParticleEnd(word *snowballword.SnowballWord) bool {
	if len(word.RS) == 0 {
		return false
	}
	lastRune := word.RS[len(word.RS)-1]
	if isV1(lastRune) {
		return true
	}
	if len(word.RS) >= 2 && string(word.RS[len(word.RS)-2:]) == "nt" {
		return true
	}
	return false
}

// Return true if the input word is a Finnish stop word.
func IsStopWord(word string) bool {
	switch word {
	case "ei", "en", "et", "emme", "ette", "eivät",
		"ja", "että", "on", "ovat", "oli", "ollut", "olla",
		"kanssa", "mukaan", "niin", "siitä", "sitä", "tämä", "tulee",
		"hän", "itse", "kaikki", "ne", "nyt", "näin", "vain", "vai",
		"vielä", "vuoden", "vuonna", "voi", "yli":
		return true
	}
	return false
}
