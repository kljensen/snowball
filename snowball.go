package snowball

import (
	"fmt"
	"github.com/kljensen/snowball/english"
)

func Stem(word, language string, stemStopWords bool) (stemmed string, err error) {
	switch language {
	case "english":
		stemmed = english.Stem(word, stemStopWords)
		return
	}
	err = fmt.Errorf("Unknown language: %s", language)
	return

}
