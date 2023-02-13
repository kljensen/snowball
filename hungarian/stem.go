package hungarian

import (
	"log"
	"strings"
	"unicode"

	"github.com/kljensen/snowball/snowballword"
)

func printDebug(debug bool, w *snowballword.SnowballWord) {
	if debug {
		log.Println(w.DebugString())
	}
}

func StemSentence(pairs [][2]string, s string) [][2]string {
	for _, word := range strings.FieldsFunc(s, func(r rune) bool {
		return unicode.IsPunct(r) || unicode.IsSpace(r)
	}) {
		pairs = append(pairs, [2]string{word, Stem(word, false)})
	}
	return pairs
}

// Stem an Hungarian word.  This is the only exported
// function in this package.
//
//	This stemming algorithm removes the inflectional suffixes of nouns. Nouns are inflected for case, person/possession and number.
//
// Letters in Hungarian include the following accented forms,
//
//	á   é   í   ó   ö   ő   ú   ü   ű
//
// The following letters are vowels:
//
//	a   á   e   é   i   í   o   ó   ö   ő   u   ú   ü   ű
//
// The following letters are digraphs:
//
//	cs   dz   dzs   gy   ly   ny   ty   zs
//
// A double consonant is defined as:
//
//	bb   cc   ccs   dd   ff   gg   ggy   jj   kk   ll   lly   mm   nn   nny   pp   rr   ss   ssz   tt   tty   vv   zz   zzs
func Stem(word string, stemStopwWords bool) string {

	word = strings.ToLower(strings.TrimSpace(word))

	// Return small words and stop words
	if len(word) <= 2 || (!stemStopwWords && isStopWord(word)) {
		return word
	}

	w := snowballword.New(word)

	// Stem the word.  Note, each of these
	// steps will alter `w` in place.
	//

	preprocess(w)
	step1(w)
	step2(w)
	step3(w)
	step4(w)
	step5(w)
	step6(w)
	step7(w)
	step8(w)
	step9(w)

	return w.String()

}

func preprocess(w *snowballword.SnowballWord) {
	w.R1start = findRegions(w)
}

// step1 Remove instrumental case
//
// Search for one of the following suffixes and perform the action indicated.
//
//	al   el
//
// delete if in R1 and preceded by a double consonant,
// and remove one of the double consonants.
// (In the case of consonant plus digraph, such as ccs, remove a c).
func step1(w *snowballword.SnowballWord) {
	n := len(w.RS)
	if n < 2 ||
		!(w.RS[n-1] == 'l' &&
			(w.RS[n-2] == 'a' || w.RS[n-2] == 'e')) {
		return
	}
	// in R1
	if w.R1start > n-2 {
		return
	}
	// (In the case of consonant plus digraph, such as ccs, remove a c).
	if isDoubleConsonant(w.RS[n-5:n-2]) > 2 {
		w.RS[n-5], w.RS[n-4] = w.RS[n-4], w.RS[n-3]
		w.RemoveLastNRunes(3)
	} else if isDoubleConsonant(w.RS[n-4:n-2]) > 1 {
		// preceded by a double consonant
		w.RemoveLastNRunes(3)
	}
}

//	Step 2: Remove frequent cases
//
// Search for the longest among the following suffixes and perform the action indicated.
//
//	ban   ben   ba   be   ra   re   nak   nek   val   vel   tól   től   ról   ről   ból   ből   hoz   hez   höz   nál   nél   ig   at   et   ot   öt   ért   képp   képpen   kor   ul   ül   vá   vé   onként   enként   anként   ként   en   on   an   ön   n   t
//
// delete if in R1
//
// if the remaining word ends á replace by a
// if the remaining word ends é replace by e
func step2(w *snowballword.SnowballWord) {
	if suffix := firstSuffixInR1(w, []string{
		"onként", "enként", "anként",
		"képpen",
		"ként",
		"képp",
		"kor",
		"ban", "ben", "nak", "nek", "val", "vel", "tól", "től", "ról", "ről", "ból", "ből", "hoz", "hez", "höz", "nál", "nél",
		"ért",
		"ba", "be", "ra", "re", "ig", "at", "et", "ot", "öt",
		"ul", "ül", "vá", "vé",
		"en", "on", "an", "ön",
		"n", "t",
	}); suffix != "" {
		rs := runesOf(suffix)
		// delete if in R1
		w.RemoveLastNRunes(len(rs))
		if len(w.RS) == 0 {
			return
		}
		switch w.RS[len(w.RS)-1] {
		case 'á':
			// if the remaining word ends á replace by a
			w.RS[len(w.RS)-1] = 'a'
		case 'é':
			// if the remaining word ends é replace by e
			w.RS[len(w.RS)-1] = 'e'
		}
	}
}

// step3: Remove special cases:
//
// Search for the longest among the following suffixes and perform the action indicated.
//
//	án   ánként
//
// replace by a if in R1
//
//	én
//
// replace by e if in R1
func step3(w *snowballword.SnowballWord) {
	if suffix := firstSuffixInR1(w, []string{
		"ánként", "án",
		"én",
	}); suffix != "" {
		rs := runesOf(suffix)
		repl := 'a'
		if rs[0] == 'é' {
			repl = 'e'
		}
		w.RS[len(w.RS)-len(rs)] = repl
		w.RemoveLastNRunes(len(rs) - 1)
	}
}

// step4: Remove other cases:
//
// Search for the longest among the following suffixes and perform the action indicated
//
//	astul   estül   stul   stül
//
// delete if in R1
//
//	ástul
//
// replace with a if in R1
//
//	éstül
//
// replace with e if in R1
func step4(w *snowballword.SnowballWord) {
	if suffix := firstSuffixInR1(w, []string{"ástul"}); suffix != "" {
		w.RemoveLastNRunes(4)
		w.RS[len(w.RS)-1] = 'a'
		return
	}
	if suffix := firstSuffixInR1(w, []string{"éstül"}); suffix != "" {
		w.RemoveLastNRunes(4)
		w.RS[len(w.RS)-1] = 'e'
		return
	}
	// astul   estül   stul   stül
	if suffix := firstSuffixInR1(w, []string{"astul", "estül", "stul", "stül"}); suffix != "" {
		w.RemoveLastNRunes(len(runesOf(suffix)))
		return
	}
}

// step5: Remove factive case
//
// Search for one of the following suffixes and perform the action indicated.
//
//	á   é
//
// delete if in R1 and preceded by a double consonant,
// and remove one of the double consonants (as in step 1).
func step5(w *snowballword.SnowballWord) {
	n := len(w.RS)
	if n < 3 || w.R1start >= n || !(w.RS[n-1] == 'á' || w.RS[n-1] == 'é') {
		return
	}
	// (In the case of consonant plus digraph, such as ccs, remove a c).
	if isDoubleConsonant(w.RS[n-4:n-1]) > 2 {
		w.RS[n-4], w.RS[n-3] = w.RS[n-3], w.RS[n-1]
		w.RemoveLastNRunes(2)
	} else if isDoubleConsonant(w.RS[n-3:n-1]) > 1 {
		// preceded by a double consonant
		w.RemoveLastNRunes(2)
	}
}

// step6: Remove owned
// Search for the longest among the following suffixes and perform the action indicated.
//
//	oké   öké   aké   eké   ké   éi   é
//
// delete if in R1
//
//	áké   áéi
//
// replace with a if in R1
//
//	éké   ééi   éé
//
// replace with e if in R1
func step6(w *snowballword.SnowballWord) {
	if suffix := firstSuffixInR1(w, []string{
		"áké", "áéi",
		"éké", "ééi", "éé",
		"oké", "öké", "aké", "eké", "ké", "éi", "é",
	}); suffix != "" {
		switch suffix {

		case "áké", "áéi":
			w.RemoveLastNRunes(2)
			w.RS[len(w.RS)-1] = 'a'

		case "éké", "ééi", "éé":
			w.RemoveLastNRunes(len(runesOf(suffix)) - 1)
			w.RS[len(w.RS)-1] = 'e'

		default:
			w.RemoveLastNRunes(len(runesOf(suffix)))
		}
	}
}

// step7: Remove singular owner suffixes
//
// Search for the longest among the following suffixes and perform the action indicated.
//
//	ünk   unk   nk   juk   jük   uk   ük   em   om   am   m   od   ed   ad   öd   d   ja   je   a   e o
//
// delete if in R1
//
//	ánk ájuk ám ád á
//
// replace with a if in R1
//
//	énk éjük ém éd é
//
// replace with e if in R1
func step7(w *snowballword.SnowballWord) {
	if suffix := firstSuffixInR1(w, []string{
		"ájuk", "éjük",
		"énk",
		"ünk", "unk",
		"juk", "jük",
		"ánk",
		"nk",
		"uk", "ük", "em", "om", "am",
		"od", "ed", "ad", "öd", "ja", "je",
		"ám", "ád", "ém", "éd",
		"m", "d",
		"a", "e", "o",
		"á", "é",
	}); suffix != "" {
		n := len(runesOf(suffix))
		switch suffix {
		case "ánk", "ájuk", "ám", "ád", "á":
			w.RemoveLastNRunes(n - 1)
			w.RS[len(w.RS)-1] = 'a'
		case "énk", "éjük", "ém", "éd", "é":
			w.RemoveLastNRunes(n - 1)
			w.RS[len(w.RS)-1] = 'e'
		default:
			w.RemoveLastNRunes(n)
		}
	}
}

// step8: Remove plural owner suffixes
// Search for the longest among the following suffixes and perform the action indicated.
//
//	jaim   jeim   aim   eim   im   jaid   jeid   aid   eid   id   jai   jei   ai   ei   i   jaink   jeink   eink   aink   ink   jaitok   jeitek   aitok   eitek   itek   jeik   jaik   aik   eik   ik
//
// delete if in R1
//
//	áim   áid   ái   áink   áitok   áik
//
// replace with a if in R1
//
//	éim   éid     éi   éink   éitek   éik
//
// replace with e if in R1
func step8(w *snowballword.SnowballWord) {
	if suffix := firstSuffixInR1(w, []string{
		"jaitok", "jeitek",
		"jaink", "jeink", "aitok", "eitek", "áitok", "éitek",
		"áink", "éink", "itek", "jeik", "jaik",
		"eink", "aink", "jaim", "jeim", "jaid", "jeid",
		"áim", "áid", "áik", "éim", "éid", "éik",
		"ink", "aik", "eik", "jai", "jei",
		"aim", "eim", "aid", "eid",
		"ái", "éi", "ik", "id", "ai", "ei",
		"im",
		"i",
	}); suffix != "" {
		n := len(runesOf(suffix))
		switch suffix {
		case "áim", "áid", "ái", "áink", "áitok", "áik":
			w.RemoveLastNRunes(n - 1)
			w.RS[len(w.RS)-1] = 'a'
		case "éim", "éid", "éi", "éink", "éitek", "éik":
			w.RemoveLastNRunes(n - 1)
			w.RS[len(w.RS)-1] = 'e'
		default:
			w.RemoveLastNRunes(n)
		}
	}
}

// step9: Remove plural suffixes
//
// Search for the longest among the following suffixes and perform the action indicated.
//
//	ák
//
// replace with a if in R1
// replace with e if in R1
//
//	ök   ok   ek   ak   k
//
// delete if in R1
func step9(w *snowballword.SnowballWord) {
	if suffix := firstSuffixInR1(w, []string{
		"ák", "ék",
		"ök", "ok", "ek", "ak", "k",
	}); suffix != "" {
		switch suffix {
		case "ák":
			w.RemoveLastNRunes(1)
			w.RS[len(w.RS)-1] = 'a'
		case "ék":
			w.RemoveLastNRunes(1)
			w.RS[len(w.RS)-1] = 'e'
		default:
			w.RemoveLastNRunes(len(runesOf(suffix)))
		}
	}
}

func firstSuffixInR1(w *snowballword.SnowballWord, suffixes []string) string {
	for _, suffix := range suffixes {
		rs := runesOf(suffix)
		if len(w.RS)-w.R1start >= len(rs) && w.HasSuffixRunes(rs) {
			return suffix
		}
	}
	return ""
}
