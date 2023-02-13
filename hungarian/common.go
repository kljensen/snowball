package hungarian

import (
	"sync"

	"github.com/kljensen/snowball/snowballword"
)

var (
	runesMapMu sync.Mutex
	runesMap   = make(map[string][]rune)
)

func runesOf(s string) []rune {
	runesMapMu.Lock()
	rs := runesMap[s]
	if rs == nil {
		rs = []rune(s)
		runesMap[s] = rs
	}
	runesMapMu.Unlock()
	return rs
}

// findRegions returns start of R1.
//
// If the word begins with a vowel, R1 is defined as the region after the first consonant or digraph in the word.
// If the word begins with a consonant, it is defined as the region after the first vowel in the word.
// If the word does not contain both a vowel and consonant, R1 is the null region at the end of the word.
func findRegions(word *snowballword.SnowballWord) (r1start int) {
	if len(word.RS) < 2 {
		return 0
	}

	// If the word begins with a vowel, R1 is defined as the region
	// after the first consonant or digraph in the word.
	if isVowel(word.RS[0]) {
		for i := 1; i < len(word.RS); i++ {
			if isVowel(word.RS[i]) {
				continue
			}
			if j := isDigraph(word.RS[i:]); j > 0 {
				return i + j
			}
			// consonant
			return i + 1
		}
		return len(word.RS)
	}

	// If the word begins with a consonant, it is defined as the region
	// after the first vowel in the word.
	for i := 1; i < len(word.RS); i++ {
		if isVowel(word.RS[i]) {
			return i + 1
		}
	}
	return len(word.RS)
}

func isVowel(r rune) bool {
	switch r {
	case 'a', 'á', 'e', 'é', 'i', 'í', 'o', 'ó', 'ö', 'ő', 'u', 'ú', 'ü', 'ű':
		return true
	}
	return false
}
func isDigraph(rs []rune) int {
	if len(rs) < 2 {
		return 0
	}
	switch rs[0] {
	case 'c', 'z': // cs, zs
		if rs[1] == 's' {
			return 2
		}
	case 'd':
		if rs[1] == 'z' {
			if len(rs) > 2 && rs[2] == 's' { // dzs
				return 3
			}
			return 2 // dz
		}
	case 'g', 'l', 'n', 't':
		if rs[1] == 'y' {
			return 2
		}
	}
	return 0
}

func isConsonant(r rune) bool {
	switch r {
	case 'b', 'c', 'd', 'f', 'g', 'j', 'k', 'l', 'm', 'n', 'p', 'r', 's', 't', 'v', 'z':
		return true
	}
	return false
}
func isDoubleConsonant(rs []rune) int {
	if len(rs) < 2 || !isConsonant(rs[0]) || rs[0] != rs[1] {
		return 0
	}
	if len(rs) > 2 {
		switch rs[0] {
		case 'c', 'z':
			if rs[2] == 's' {
				return 3
			}
		case 's':
			if rs[2] == 'z' {
				return 3
			}
		case 'g', 'l', 'n', 't':
			if rs[2] == 'y' {
				return 3
			}
		}
	}
	return 2
}

// isStopWord returns true it the word is a stop word.
//
// # Hungarian stop word list prepared by Anna Tordai
//
// https://snowballstem.org/algorithms/hungarian/stop.txt
func isStopWord(word string) bool {
	switch word {
	case "a",
		"ahogy",
		"ahol",
		"aki",
		"akik",
		"akkor",
		"alatt",
		"által",
		"általában",
		"amely",
		"amelyek",
		"amelyekben",
		"amelyeket",
		"amelyet",
		"amelynek",
		"ami",
		"amit",
		"amolyan",
		"amíg",
		"amikor",
		"át",
		"abban",
		"ahhoz",
		"annak",
		"arra",
		"arról",
		"az",
		"azok",
		"azon",
		"azt",
		"azzal",
		"azért",
		"aztán",
		"azután",
		"azonban",
		"bár",
		"be",
		"belül",
		"benne",
		"cikk",
		"cikkek",
		"cikkeket",
		"csak",
		"de",
		"e",
		"eddig",
		"egész",
		"egy",
		"egyes",
		"egyetlen",
		"egyéb",
		"egyik",
		"egyre",
		"ekkor",
		"el",
		"elég",
		"ellen",
		"elő",
		"először",
		"előtt",
		"első",
		"én",
		"éppen",
		"ebben",
		"ehhez",
		"emilyen",
		"ennek",
		"erre",
		"ez",
		"ezt",
		"ezek",
		"ezen",
		"ezzel",
		"ezért",
		"és",
		"fel",
		"felé",
		"hanem",
		"hiszen",
		"hogy",
		"hogyan",
		"igen",
		"így",
		"illetve",
		"ill.",
		"ill",
		"ilyen",
		"ilyenkor",
		"ison",
		"ismét",
		"itt",
		"jó",
		"jól",
		"jobban",
		"kell",
		"kellett",
		"keresztül",
		"keressünk",
		"ki",
		"kívül",
		"között",
		"közül",
		"legalább",
		"lehet",
		"lehetett",
		"legyen",
		"lenne",
		"lenni",
		"lesz",
		"lett",
		"maga",
		"magát",
		"majd",
		"már",
		"más",
		"másik",
		"meg",
		"még",
		"mellett",
		"mert",
		"mely",
		"melyek",
		"mi",
		"mit",
		"míg",
		"miért",
		"milyen",
		"mikor",
		"minden",
		"mindent",
		"mindenki",
		"mindig",
		"mint",
		"mintha",
		"mivel",
		"most",
		"nagy",
		"nagyobb",
		"nagyon",
		"ne",
		"néha",
		"nekem",
		"neki",
		"nem",
		"néhány",
		"nélkül",
		"nincs",
		"olyan",
		"ott",
		"össze",
		"ő",
		"ők",
		"őket",
		"pedig",
		"persze",
		"rá",
		"s",
		"saját",
		"sem",
		"semmi",
		"sok",
		"sokat",
		"sokkal",
		"számára",
		"szemben",
		"szerint",
		"szinte",
		"talán",
		"tehát",
		"teljes",
		"tovább",
		"továbbá",
		"több",
		"úgy",
		"ugyanis",
		"új",
		"újabb",
		"újra",
		"után",
		"utána",
		"utolsó",
		"vagy",
		"vagyis",
		"valaki",
		"valami",
		"valamint",
		"való",
		"vagyok",
		"van",
		"vannak",
		"volt",
		"voltam",
		"voltak",
		"voltunk",
		"vissza",
		"vele",
		"viszont",
		"volna":
		return true
	}
	return false
}
