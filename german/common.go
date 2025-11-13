package german

import (
	"github.com/kljensen/snowball/snowballword"
)

// German vowels are: a e i o u y ä ö ü
func isLowerVowel(r rune) bool {
	switch r {
	case 97, 101, 105, 111, 117, 121: // a e i o u y
		return true
	case 228, 246, 252: // ä ö ü
		return true
	}
	return false
}

// Find the starting point of regions R1 and R2
// R1 is the region after the first non-vowel following a vowel,
// or is the null region at the end of the word if there is no such non-vowel.
// R2 is the region after the first non-vowel following a vowel in R1
func findRegions(word *snowballword.SnowballWord) (r1start, r2start int) {
	r1start = len(word.RS)
	r2start = len(word.RS)

	// R1 must be at least 3 characters from the start
	minR1 := 3

	// Find R1
	for i := 0; i < len(word.RS)-1; i++ {
		if isLowerVowel(word.RS[i]) && !isLowerVowel(word.RS[i+1]) {
			r1start = i + 2
			if r1start < minR1 {
				r1start = minR1
			}
			break
		}
	}

	// Find R2 (start searching from R1)
	for i := r1start; i < len(word.RS)-1; i++ {
		if isLowerVowel(word.RS[i]) && !isLowerVowel(word.RS[i+1]) {
			r2start = i + 2
			break
		}
	}

	return
}

// Valid s-ending: the letter before s must be one of bdfghklmnrt
func isValidSEnding(word *snowballword.SnowballWord, idx int) bool {
	if idx <= 0 {
		return false
	}
	switch word.RS[idx-1] {
	case 98, 100, 102, 103, 104, 107, 108, 109, 110, 114, 116: // bdfghklmnrt
		return true
	}
	return false
}

// Valid st-ending: must be preceded by a valid st-ending character (bdfghklmnt)
// and must have at least 3 letters before it
func isValidStEnding(word *snowballword.SnowballWord, idx int) bool {
	if idx <= 2 {
		return false
	}
	switch word.RS[idx-1] {
	case 98, 100, 102, 103, 104, 107, 108, 109, 110, 116: // bdfghklmnt
		return true
	}
	return false
}

// Valid et-ending: must be one of dfgklmnrstUzä
func isValidEtEnding(word *snowballword.SnowballWord, idx int) bool {
	if idx <= 0 {
		return false
	}
	r := word.RS[idx-1]
	switch r {
	case 100, 102, 103, 107, 108, 109, 110, 114, 115, 116, 85, 122: // dfgklmnrstUz
		return true
	case 228: // ä
		return true
	}
	return false
}

// Return `true` if the input `word` is a German stop word.
func IsStopWord(word string) bool {
	switch word {
	case "aber", "alle", "allem", "allen", "aller", "alles", "als", "also", "am", "an",
		"ander", "andere", "anderem", "anderen", "anderer", "anderes", "anderm", "andern",
		"anderr", "anders", "auch", "auf", "aus", "bei", "bin", "bis", "bist", "da",
		"damit", "dann", "der", "den", "des", "dem", "die", "das", "daß", "dass", "derselbe",
		"derselben", "denselben", "desselben", "demselben", "dieselbe", "dieselben", "dasselbe",
		"dazu", "dein", "deine", "deinem", "deinen", "deiner", "deines", "denn", "derer",
		"dessen", "dich", "dies", "diese", "diesem", "diesen", "dieser", "dieses",
		"dir", "doch", "dort", "durch", "ein", "eine", "einem", "einen", "einer", "eines",
		"einig", "einige", "einigem", "einigen", "einiger", "einiges", "einmal", "er", "ihn",
		"ihm", "es", "etwas", "euer", "eure", "eurem", "euren", "eurer", "eures", "für",
		"gegen", "gewesen", "hab", "habe", "haben", "hat", "hatte", "hatten", "hier", "hin",
		"hinter", "ich", "mich", "mir", "ihr", "ihre", "ihrem", "ihren", "ihrer", "ihres",
		"euch", "im", "in", "indem", "ins", "ist", "jede", "jedem", "jeden", "jeder", "jedes",
		"jene", "jenem", "jenen", "jener", "jenes", "jetzt", "kann", "kein", "keine", "keinem",
		"keinen", "keiner", "keines", "können", "könnte", "machen", "man", "manche", "manchem",
		"manchen", "mancher", "manches", "mein", "meine", "meinem", "meinen", "meiner", "meines",
		"mit", "muss", "musste", "nach", "nicht", "nichts", "noch", "nun", "nur", "ob", "oder",
		"ohne", "sehr", "sein", "seine", "seinem", "seinen", "seiner", "seines", "selbst", "sich",
		"sie", "ihnen", "sind", "so", "solche", "solchem", "solchen", "solcher", "solches",
		"soll", "sollte", "sondern", "sonst", "über", "um", "und", "uns", "unse", "unsem",
		"unsen", "unser", "unses", "unter", "viel", "vom", "von", "vor", "während", "war",
		"waren", "warst", "was", "weg", "weil", "weiter", "welche", "welchem", "welchen",
		"welcher", "welches", "wenn", "wer", "werde", "werden", "wie", "wieder", "will", "wir",
		"wird", "wirst", "wo", "wollen", "wollte", "würde", "würden", "zu", "zum", "zur",
		"zwar", "zwischen":
		return true
	}
	return false
}
