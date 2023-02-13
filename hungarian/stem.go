package hungarian

import (
	"log"
	"strings"

	"github.com/kljensen/snowball/snowballword"
)

func printDebug(debug bool, w *snowballword.SnowballWord) {
	if debug {
		log.Println(w.DebugString())
	}
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
	if len(word) <= 2 || (stemStopwWords == false && isStopWord(word)) {
		return word
	}

	w := snowballword.New(word)

	// Stem the word.  Note, each of these
	// steps will alter `w` in place.
	//

	preprocess(w)
	step1(w)
	step2(w)

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
func step1(word *snowballword.SnowballWord) {
	n := len(word.RS)
	if n < 2 ||
		!(word.RS[n-1] == 'l' &&
			(word.RS[n-2] == 'a' || word.RS[n-2] == 'e')) {
		return
	}
	log.Println("R1", word.R1start, "n", n)
	// in R1
	if word.R1start > n-2 {
		return
	}
	// (In the case of consonant plus digraph, such as ccs, remove a c).
	if isDoubleConsonant(word.RS[n-5:n-2]) > 2 {
		word.RS[n-5], word.RS[n-4] = word.RS[n-4], word.RS[n-3]
		word.RemoveLastNRunes(3)
	} else if isDoubleConsonant(word.RS[n-4:n-2]) > 1 {
		// preceded by a double consonant
		word.RemoveLastNRunes(3)
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
	for _, suffix := range []string{
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
	} {
		rs := runesOf(suffix)
		if len(w.RS)-w.R1start >= len(rs) && w.HasSuffixRunes(rs) {
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
			return
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
	log.Println(w, w.R1start)
	for _, suffix := range []string{
		"ánként", "án",
		"én",
	} {
		rs := runesOf(suffix)
		if w.HasSuffixRunesIn(len(w.RS)-w.R1start, len(w.RS), rs) {
			repl := 'a'
			if suffix[0] == 'é' {
				repl = 'e'
			}
			w.RS[len(w.RS)-len(rs)] = repl
			w.RemoveLastNRunes(len(rs) - 1)
			return
		}
	}
}
