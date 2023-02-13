package hungarian

import (
	"testing"

	"github.com/kljensen/snowball/snowballword"
)

func TestStep1(t *testing.T) {
	for k, want := range map[string]string{
		"taccsal": "tacs",
		"téttel":  "tét",
		"paddal":  "pad",
		"padló":   "padló",
	} {
		w := snowballword.New(k)
		preprocess(w)
		step1(w)
		if got := string(w.RS); got != want {
			t.Errorf("%q: got %q, wanted %q", k, got, want)
		}
	}
}

func TestStep2(t *testing.T) {
	for k, want := range map[string]string{
		"padonként": "pad",
		"tétről":    "tét",
		"palából":   "pala",
	} {
		w := snowballword.New(k)
		preprocess(w)
		step2(w)
		if got := string(w.RS); got != want {
			t.Errorf("%q: got %q, wanted %q", k, got, want)
		}
	}
}
func TestStep3(t *testing.T) {
	for k, want := range map[string]string{
		"banánként":  "bana",
		"vágyanként": "vágya",
		"lepkén":     "lepke",
	} {
		w := snowballword.New(k)
		preprocess(w)
		step3(w)
		if got := string(w.RS); got != want {
			t.Errorf("%q: got %q, wanted %q", k, got, want)
		}
	}
}
