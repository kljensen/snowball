package snowball

import (
	"fmt"

	"github.com/kljensen/snowball/danish"
	"github.com/kljensen/snowball/dutch"
	"github.com/kljensen/snowball/english"
	"github.com/kljensen/snowball/finnish"
	"github.com/kljensen/snowball/french"
	"github.com/kljensen/snowball/german"
	"github.com/kljensen/snowball/hungarian"
	"github.com/kljensen/snowball/italian"
	"github.com/kljensen/snowball/norwegian"
	"github.com/kljensen/snowball/polish"
	"github.com/kljensen/snowball/portuguese"
	"github.com/kljensen/snowball/romanian"
	"github.com/kljensen/snowball/russian"
	"github.com/kljensen/snowball/spanish"
	"github.com/kljensen/snowball/swedish"
	"github.com/kljensen/snowball/turkish"
)

const (
	VERSION string = "v0.7.0"
)

// Stem a word in the specified language.
func Stem(word, language string, stemStopWords bool) (stemmed string, err error) {

	var f func(string, bool) string
	switch language {
	case "danish":
		f = danish.Stem
	case "dutch":
		f = dutch.Stem
	case "english":
		f = english.Stem
	case "finnish":
		f = finnish.Stem
	case "french":
		f = french.Stem
	case "german":
		f = german.Stem
	case "hungarian":
		f = hungarian.Stem
	case "italian":
		f = italian.Stem
	case "norwegian":
		f = norwegian.Stem
	case "polish":
		f = polish.Stem
	case "portuguese":
		f = portuguese.Stem
	case "romanian":
		f = romanian.Stem
	case "russian":
		f = russian.Stem
	case "spanish":
		f = spanish.Stem
	case "swedish":
		f = swedish.Stem
	case "turkish":
		f = turkish.Stem
	default:
		err = fmt.Errorf("Unknown language: %s", language)
		return
	}
	stemmed = f(word, stemStopWords)
	return

}
