package polish

import (
	"github.com/kljensen/snowball/snowballword"
)

// Polish vowels: a ą e ę i o ó u y
func isVowel(r rune) bool {
	switch r {
	case 'a', 'e', 'i', 'o', 'u', 'y':
		return true
	case 0x0105, 0x0119, 0x00F3: // ą ę ó
		return true
	}
	return false
}

// Find R1 region for Polish
func markRegions(word *snowballword.SnowballWord) {
	// R1: region after first non-vowel following a vowel
	r1start := len(word.RS)

	// gopast vowel
	foundVowel := false
	for i := 0; i < len(word.RS); i++ {
		if isVowel(word.RS[i]) {
			foundVowel = true
			// gopast non-vowel
			for j := i + 1; j < len(word.RS); j++ {
				if !isVowel(word.RS[j]) {
					r1start = j + 1
					break
				}
			}
			break
		}
	}

	if !foundVowel {
		r1start = len(word.RS)
	}

	word.R1start = r1start
}

// Return true if the input word is a Polish stop word.
func IsStopWord(word string) bool {
	// Common Polish stop words
	switch word {
	case "a", "aby", "ach", "acz", "aczkolwiek", "aj", "albo", "ale",
		"ależ", "ani", "aż",
		"bardziej", "bardzo", "bądź", "będzie", "bez", "bo", "bowiem",
		"by", "byli", "bym", "bynajmniej", "być", "był", "była", "było",
		"były", "będą", "będę",
		"cali", "cała", "cały", "ci", "cię", "ciebie", "co", "cokolwiek",
		"coś", "czasami", "czasem", "czemu", "czy", "czyli", "często",
		"daleko", "dla", "dlaczego", "dlatego", "do", "dobrze", "dokąd",
		"dość", "dużo", "dwa", "dwaj", "dwie", "dwoje",
		"dzisiaj", "dziś",
		"gdyby", "gdzie", "gdziekolwiek", "gdzieś", "gdy", "gdyż",
		"go", "godz",
		"i", "ich", "ile", "im", "inna", "inne", "inny", "innych", "iż",
		"ja", "ją", "jak", "jakaś", "jakby", "jaki", "jakiś", "jakkolwiek",
		"jako", "jakoś", "je", "jeden", "jedna", "jedno", "jednym", "jedynie",
		"jego", "jej", "jemu", "jest", "jestem", "jeszcze", "jeśli", "jeżeli",
		"już",
		"każdy", "kiedy", "kierunku", "kilka", "kilku", "kimś", "kto", "ktokolwiek",
		"ktoś", "która", "które", "którego", "której", "który", "których", "którym",
		"którzy", "ku",
		"lat", "lecz", "lub",
		"ma", "mają", "mam", "mi", "miał", "mimo", "między", "mnie", "mną",
		"mogą", "moi", "moim", "moja", "moje", "może", "możliwe", "można", "mój",
		"mu", "musi", "my",
		"na", "nad", "nam", "nami", "nas", "nasi", "nasz", "nasza", "nasze",
		"naszego", "naszych", "natomiast", "natychmiast", "nawet", "nią", "nic",
		"nich", "nie", "niego", "niej", "niemu", "nigdy", "nim", "nimi", "niż",
		"no", "nowe", "np", "nr",
		"o", "o.o.", "ob", "obok", "od", "okay", "on", "ona", "one", "oni",
		"ono", "oraz", "oto", "owszem",
		"pan", "pana", "pani", "pl", "po", "pod", "podczas", "pomimo", "ponad",
		"ponieważ", "powinien", "powinna", "powinni", "powinno", "poza", "prawie",
		"przecież", "przed", "przede", "przedtem", "przez", "przy", "raz", "razie",
		"roku", "również",
		"sam", "sama", "są", "się", "skąd", "sobie", "sobą", "sposób", "swoje",
		"ta", "tak", "taka", "taki", "takie", "także", "tam", "te", "tego", "tej",
		"temu", "ten", "teraz", "też", "to", "toba", "tobą", "tobie", "totobą",
		"trzeba", "tu", "tutaj", "twoi", "twoim", "twoja", "twoje",
		"twym", "twój", "ty", "tych", "tylko", "tym", "tys", "tzw", "tę",
		"u", "w", "wam", "wami", "was", "wasi", "wasz", "wasza", "wasze", "we",
		"według", "wiele", "wielu", "więc", "więcej", "wszyscy", "wszystkich",
		"wszystkie", "wszystkim", "wszystko", "wtedy", "wy", "właśnie", "wśród",
		"z", "za", "zapewne", "zawsze", "ze", "zeznowu", "znowu", "znów", "został",
		"żaden", "żadna", "żadne", "żadnych", "że", "żeby":
		return true
	}
	return false
}
