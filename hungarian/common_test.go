package hungarian

import (
	"testing"

	"github.com/kljensen/snowball/snowballword"
)

func TestFindRegions(t *testing.T) {
	for k, want := range map[string]int{
		"tóban":   2, //          consonant-vowel
		"ablakan": 2, //       vowel-consonant
		"acsony":  3, //         vowel-digraph
		"cvs":     3, //          null R1 region
	} {
		got := findRegions(snowballword.New(k))
		if got != want {
			t.Errorf("%q: got %d, wanted %d", k, got, want)
		}
	}
}
