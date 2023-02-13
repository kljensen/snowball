package hungarian

import (
	"reflect"
	"testing"

	"github.com/kljensen/snowball/snowballword"
)

func TestStemSentence(t *testing.T) {
	var pairs [][2]string
	var got []string
	for k, want := range map[string][]string{
		`Tisztelettel az alábbi bankszámlára szeretném kérni az utalás. Raiffeisen
Bank:999999999999999999999999.Tisztelettel:Horváth Péter

Az alábbi email a KöBE hálózatán kívüli forrásból érkezett, kérjük, legyen óvatos a beágyazott linkekkel és csatolmányokkal!
`: []string{
			"tisztel", "az", "alább", "bankszáml", "szeretne", "kérn", "az", "utalás",
			"raiffeis", "ba", "999999999999999999999999", "tisztel", "horváth", "péter",
			"az", "alább", "email", "a", "kö", "hálózat", "kívül", "forrás", "érkezet",
			"kér", "legyen", "óvatos", "a", "beágyazot", "link", "és", "csatolmány",
		},
	} {
		pairs = StemSentence(pairs[:0], k)
		got = got[:0]
		for _, p := range pairs {
			got = append(got, p[1])
		}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("%q: got %q, wanted %q", k, got, want)
		}
	}
}
func TestStem(t *testing.T) {
	for k, want := range map[string]string{
		"fiaiéi": "fi",
		"megkelkáposztásíthatatlanságoskodásaitokért": "megkelkáposztásíthatatlanságoskodás",
	} {
		if got := Stem(k, false); got != want {
			t.Errorf("%q: got %q, wanted %q", k, got, want)
		}
	}
}

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
		"bányánként": "bánya",
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
func TestStep4(t *testing.T) {
	for k, want := range map[string]string{
		"házastul":   "ház",
		"képestül":   "kép",
		"akóstul":    "akó",
		"ruhástul":   "ruha",
		"vízeséstül": "vízese",
	} {
		w := snowballword.New(k)
		preprocess(w)
		step4(w)
		if got := string(w.RS); got != want {
			t.Errorf("%q: got %q, wanted %q", k, got, want)
		}
	}
}
func TestStep5(t *testing.T) {
	for k, want := range map[string]string{
		"fiaié":  "fiaié",
		"blatté": "blat",
	} {
		w := snowballword.New(k)
		preprocess(w)
		step5(w)
		if got := string(w.RS); got != want {
			t.Errorf("%q: got %q, wanted %q", k, got, want)
		}
	}
}
func TestStep6(t *testing.T) {
	for k, want := range map[string]string{
		"fiatoké": "fiat",
		"fiáéi":   "fia",
	} {
		w := snowballword.New(k)
		preprocess(w)
		step6(w)
		if got := string(w.RS); got != want {
			t.Errorf("%q: got %q, wanted %q", k, got, want)
		}
	}
}
func TestStep7(t *testing.T) {
	for k, want := range map[string]string{
		"mamájuk": "mama",
		"fenéjük": "fene",
		"bánatod": "bánat",
	} {
		w := snowballword.New(k)
		preprocess(w)
		step7(w)
		if got := string(w.RS); got != want {
			t.Errorf("%q: got %q, wanted %q", k, got, want)
		}
	}
}
func TestStep8(t *testing.T) {
	for k, want := range map[string]string{
		"mamáid":   "mama",
		"fenéitek": "fene",
		"bánatai":  "bánat",
	} {
		w := snowballword.New(k)
		preprocess(w)
		step8(w)
		if got := string(w.RS); got != want {
			t.Errorf("%q: got %q, wanted %q", k, got, want)
		}
	}
}
func TestStep9(t *testing.T) {
	for k, want := range map[string]string{
		"mamák":   "mama",
		"fenék":   "fene",
		"bánatok": "bánat",
	} {
		w := snowballword.New(k)
		preprocess(w)
		step9(w)
		if got := string(w.RS); got != want {
			t.Errorf("%q: got %q, wanted %q", k, got, want)
		}
	}
}
