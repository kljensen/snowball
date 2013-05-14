package snowball

import (
	"fmt"
	"github.com/kljensen/snowball/english"
	"github.com/kljensen/snowball/french"
	"github.com/kljensen/snowball/spanish"
)

const (
	VERSION string = "v0.2.0"
)

// Stem a word in the specified language.
//
func Stem(word, language string, stemStopWords bool) (stemmed string, err error) {

	var f func(string, bool) string
	switch language {
	case "english":
		f = english.Stem
	case "spanish":
		f = spanish.Stem
	case "french":
		f = french.Stem
	default:
		err = fmt.Errorf("Unknown language: %s", language)
		return
	}
	stemmed = f(word, stemStopWords)
	return

}
