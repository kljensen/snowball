package turkish

import (
	"github.com/kljensen/snowball/snowballword"
)

// Step 1: Remove nominal verb suffixes
func nominalVerbSuffixes(word *snowballword.SnowballWord) bool {
	suffixes := []string{
		"mışsınız", "mıştık", "mışlar", "mıştı", "mışım", "mışsın",
		"muşsunuz", "muştuk", "muşlar", "muştu", "muşum", "muşsun",
		"müşsünüz", "müştük", "müşlar", "müştü", "müşüm", "müşsün",
		"mışsınız", "mıştık", "mışız", "mışsın", "mıştı", "mışım",
		"yorum", "iyorum", "uyorum", "üyorum",
		"yoruz", "iyoruz", "uyoruz", "üyoruz",
		"dım", "dim", "dum", "düm", "tım", "tim", "tum", "tüm",
		"dık", "dik", "duk", "dük", "tık", "tik", "tuk", "tük",
		"sa", "se",
	}

	for _, suffix := range suffixes {
		suffixRunes := []rune(suffix)
		if word.HasSuffixRunes(suffixRunes) {
			idx := len(word.RS) - len(suffixRunes)
			if word.R1start <= idx {
				word.RemoveLastNRunes(len(suffixRunes))
				return true
			}
		}
	}
	return false
}

// Step 2: Remove noun suffixes
func nounSuffixes(word *snowballword.SnowballWord) bool {
	suffixes := []string{
		"ların", "lerin", "lardan", "lerden", "larda", "lerde",
		"ları", "leri", "lar", "ler",
		"dan", "den", "tan", "ten",
		"nın", "nin", "nun", "nün",
		"nı", "ni", "nu", "nü",
		"na", "ne", "a", "e",
		"da", "de", "ta", "te",
		"ı", "i", "u", "ü",
	}

	for _, suffix := range suffixes {
		suffixRunes := []rune(suffix)
		if word.HasSuffixRunes(suffixRunes) {
			idx := len(word.RS) - len(suffixRunes)
			if word.R1start <= idx {
				word.RemoveLastNRunes(len(suffixRunes))
				return true
			}
		}
	}
	return false
}

// Step 3: Remove derivational suffixes
func derivationalSuffixes(word *snowballword.SnowballWord) bool {
	suffixes := []string{
		"lık", "lik", "luk", "lük",
		"cık", "cik", "cuk", "cük",
		"çık", "çik", "çuk", "çük",
	}

	for _, suffix := range suffixes {
		suffixRunes := []rune(suffix)
		if word.HasSuffixRunes(suffixRunes) {
			idx := len(word.RS) - len(suffixRunes)
			if word.R1start <= idx {
				word.RemoveLastNRunes(len(suffixRunes))
				return true
			}
		}
	}
	return false
}
