package romanian

import (
	"github.com/kljensen/snowball/snowballword"
)

// Romanian vowels: a e i o u â î ă
func isVowel(r rune) bool {
	switch r {
	case 'a', 'e', 'i', 'o', 'u':
		return true
	case 0x00E2, 0x00EE, 0x0103: // â î ă
		return true
	}
	return false
}

// Find R1, R2, RV regions for Romanian (similar to Italian/Portuguese)
func findRegions(word *snowballword.SnowballWord) (r1start, r2start, rvstart int) {
	// R1: region after first non-vowel following a vowel
	r1start = len(word.RS)
	for i := 0; i < len(word.RS)-1; i++ {
		if isVowel(word.RS[i]) && !isVowel(word.RS[i+1]) {
			r1start = i + 2
			break
		}
	}

	// R2: region after first non-vowel following a vowel in R1
	r2start = len(word.RS)
	for i := r1start; i < len(word.RS)-1; i++ {
		if isVowel(word.RS[i]) && !isVowel(word.RS[i+1]) {
			r2start = i + 2
			break
		}
	}

	// RV: like Italian/Portuguese
	rvstart = len(word.RS)
	if len(word.RS) >= 3 {
		switch {
		case !isVowel(word.RS[1]):
			// Second letter is consonant
			for i := 2; i < len(word.RS); i++ {
				if isVowel(word.RS[i]) {
					rvstart = i + 1
					break
				}
			}
		case isVowel(word.RS[0]) && isVowel(word.RS[1]):
			// First two are vowels
			for i := 2; i < len(word.RS); i++ {
				if !isVowel(word.RS[i]) {
					rvstart = i + 1
					break
				}
			}
		default:
			// Consonant-vowel case
			rvstart = 3
		}
	}

	return
}

// Return true if the input word is a Romanian stop word.
func IsStopWord(word string) bool {
	switch word {
	case "a", "abia", "acea", "aceasta", "această", "aceea", "aceeași", "acei", "aceia",
		"acel", "acela", "acelasi", "același", "aceste", "acestea", "acestia",
		"aceștia", "acestuia", "ai", "aia", "aibă", "aici", "al", "ala", "ăla", "ale",
		"alea", "ălea", "altceva", "altcineva", "am", "ar", "are", "așa", "asemenea",
		"asta", "ăsta", "astăzi", "astea", "ăstea", "ăștia", "asupra", "aș", "așadar",
		"atât", "atâta", "atâtea", "atâția", "atunci", "au",
		"bine",
		"ca", "că", "căci", "cand", "când", "care", "cărei", "căror", "cărui", "cât",
		"câte", "câtva", "ce", "cea", "ceea", "cei", "ceia", "cel", "ceva", "chiar", "cînd",
		"cine", "cineva", "cît", "cîte", "cîtva", "contra", "cu", "cum", "cumva",
		"curând", "curînd",
		"da", "dă", "dacă", "dar", "dată", "datorită", "dau", "de", "deci", "deja",
		"deoarece", "departe", "deși", "din", "dinaintea", "dintr", "dintr-", "dintre",
		"doar", "doi", "doilea", "două", "drept", "după", "dupa",
		"ea", "ei", "el", "ele", "eram", "este", "eu", "exact",
		"față", "fără", "face", "fi", "fie", "fiecare", "fii", "fim", "fiți",
		"fiu", "fost", "frumos",
		"graţie",
		"halbă",
		"i", "ia", "iar", "ieri", "ești", "îi", "ăi", "îl", "ăl", "îmi", "împotriva",
		"în", "înainte", "înaintea", "încât", "încît", "încotro", "între", "întrucât",
		"întrucît", "sunt", "s-ar",
		"încă", "ști",
		"la", "lângă", "le", "li", "lîngă", "lor", "lui",
		"ma", "mă", "mai", "mare", "mea", "mei", "mele", "mereu", "meu", "mi", "mie",
		"mîine", "mine", "mod", "mult",
		"ne", "nevoie", "ni", "nici", "nicăieri", "nimeni", "nimeri", "nimic", "nişte",
		"noastră", "noastre", "noi", "noroc", "nostru", "nouă", "noștri", "nu",
		"o", "opt", "or", "ori", "oricând", "oricare", "oricât", "orice", "oricînd",
		"oricine", "oricît", "oricum", "oriunde",
		"până", "pănă", "pe", "pentru", "peste", "pic", "pînă", "poate", "pot", "prea",
		"prima", "primul", "prin", "printr-", "putea", "puțin", "puțina", "puțină",
		"sa", "să", "săi", "sale", "sau", "său", "se", "și", "sînt", "sînteți",
		"spate", "spre", "sub", "suntem", "sunteți",
		"șapte", "șase", "șaizeci",
		"ta", "tăi", "tale", "tău", "te", "ți", "ție", "timp", "tine", "toată", "toate",
		"tot", "totuși", "trei", "treia", "treilea", "tu", "tuturor",
		"un", "una", "unde", "undeva", "unei", "uneia", "unele", "uneori", "unii", "unor",
		"unora", "unu", "unui", "unuia", "unul", "vă", "vi", "voastră", "voastre", "voi",
		"vostru", "vouă", "voștri",
		"zece", "zero", "zi", "zice":
		return true
	}
	return false
}
