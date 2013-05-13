package snowball

import (
	"fmt"
	"github.com/kljensen/snowball/english"
	"github.com/kljensen/snowball/spanish"
)

func Stem(word, language string, stemStopWords bool) (stemmed string, err error) {
	switch language {
	case "english":
		stemmed = english.Stem(word, stemStopWords)
	case "spanish":
		stemmed = spanish.Stem(word, stemStopWords)
	default:
		err = fmt.Errorf("Unknown language: %s", language)
	}
	return

}
