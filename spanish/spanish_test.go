package spanish

import (
	"github.com/kljensen/snowball/snowballword"
	"testing"
)

// Test stopWords for things we know should be true
// or false.
//
func Test_stopWords(t *testing.T) {
	testCases := []struct {
		word   string
		result bool
	}{
		{"el", true},
		{"queso", false},
	}
	for _, testCase := range testCases {
		result := isStopWord(testCase.word)
		if result != testCase.result {
			t.Errorf("Expect isStopWord(\"%v\") = %v, but got %v",
				testCase.word, testCase.result, result,
			)
		}
	}
}

// Test isLowerVowel for things we know should be true
// or false.
//
func Test_isLowerVowel(t *testing.T) {
	testCases := []struct {
		chars  string
		result bool
	}{
		// These are all vowels.
		{"aeiouáéíóúü", true},
		// None of these are vowels.
		{"cbfqhkl", false},
	}
	for _, testCase := range testCases {
		for _, r := range testCase.chars {
			result := isLowerVowel(r)
			if result != testCase.result {
				t.Errorf("Expect isLowerVowel(\"%v\") = %v, but got %v",
					r, testCase.result, result,
				)
			}

		}
	}
}

// Test isLowerVowel for things we know should be true
// or false.
//
func Test_findRegions(t *testing.T) {
	testCases := []struct {
		word    string
		r1start int
		r2start int
		rvstart int
	}{
		{"macho", 3, 5, 4},
		{"olivia", 2, 4, 3},
		{"trabajo", 4, 6, 3},
		{"áureo", 3, 5, 3},
		{"piñaolayas", 3, 6, 3},
		{"terminales", 3, 6, 3},
		{"durmió", 3, 6, 3},
		{"cobija", 3, 5, 3},
		{"anderson", 2, 5, 4},
		{"cervezas", 3, 6, 3},
		{"climáticas", 4, 6, 3},
		{"expide", 2, 5, 4},
		{"cenizas", 3, 5, 3},
		{"maximiliano", 3, 5, 3},
		{"específicos", 2, 5, 4},
		{"menor", 3, 5, 3},
		{"generis", 3, 5, 3},
		{"casero", 3, 5, 3},
		{"pululan", 3, 5, 3},
		{"suscitado", 3, 6, 3},
		{"pesadez", 3, 5, 3},
		{"interno", 2, 5, 4},
		{"agredido", 2, 5, 4},
		{"desprendía", 3, 7, 3},
		{"vistazo", 3, 6, 3},
		{"frecuentan", 4, 7, 3},
		{"noviembre", 3, 6, 3},
		{"sintética", 3, 6, 3},
		{"newagismo", 3, 5, 3},
		{"eliseo", 2, 4, 3},
		{"desbordado", 3, 6, 3},
		{"dispongo", 3, 6, 3},
		{"dilatar", 3, 5, 3},
		{"xochitl", 3, 6, 3},
		{"proporcionaba", 4, 6, 3},
		{"pue", 3, 3, 3},
		{"alpargatado", 2, 5, 4},
		{"exigida", 2, 4, 3},
		{"céntricas", 3, 7, 3},
		{"prende", 4, 6, 3},
		{"estructural", 2, 6, 5},
		{"ilegalmente", 2, 4, 3},
		{"freeport", 5, 7, 3},
		{"sonrisas", 3, 6, 3},
		{"cobró", 3, 5, 3},
		{"dioses", 4, 6, 3},
		{"consistieron", 3, 6, 3},
		{"policiales", 3, 5, 3},
		{"conciliador", 3, 6, 3},
		{"fierro", 4, 6, 3},
		{"aparadores", 2, 4, 3},
		{"coreados", 3, 6, 3},
		{"posición", 3, 5, 3},
		{"adversidades", 2, 5, 4},
		{"comprometido", 3, 7, 3},
		{"aventuras", 2, 4, 3},
		{"santiso", 3, 6, 3},
		{"talentos", 3, 5, 3},
		{"apreciar", 2, 5, 4},
		{"sprints", 5, 7, 4},
		{"zarco", 3, 5, 3},
		{"concretos", 3, 7, 3},
		{"gavica", 3, 5, 3},
		{"suavemente", 4, 6, 3},
		{"españolitos", 2, 5, 4},
		{"grabará", 4, 6, 3},
		{"entregados", 2, 6, 5},
		{"gustaría", 3, 6, 3},
		{"nickin", 3, 6, 3},
		{"sogem", 3, 5, 3},
		{"prohíbe", 4, 6, 3},
		{"espinoso", 2, 5, 4},
		{"atraviesan", 2, 5, 4},
		{"bancomext", 3, 6, 3},
		{"paraguay", 3, 5, 3},
		{"amamos", 2, 4, 3},
		{"consigna", 3, 6, 3},
		{"funcionarios", 3, 7, 3},
		{"marquis", 3, 7, 3},
		{"desactivaron", 3, 5, 3},
		{"concentrados", 3, 6, 3},
		{"democratizante", 3, 5, 3},
		{"afianzadora", 2, 5, 3},
		{"homicidio", 3, 5, 3},
		{"promovidos", 4, 6, 3},
		{"maquiladora", 3, 6, 3},
		{"bike", 3, 4, 3},
		{"recuerdos", 3, 6, 3},
		{"géneros", 3, 5, 3},
		{"rechaza", 3, 6, 3},
		{"sentarían", 3, 6, 3},
		{"quererlo", 4, 6, 3},
		{"sofisticado", 3, 5, 3},
		{"miriam", 3, 6, 3},
		{"echara", 2, 5, 4},
		{"mico", 3, 4, 3},
		{"enferma", 2, 5, 4},
		{"reforzamiento", 3, 5, 3},
		{"circunscrito", 3, 6, 3},
		{"indiana", 2, 6, 4},
		{"metrópoli", 3, 6, 3},
		{"libreta", 3, 6, 3},
		{"gonzalez", 3, 6, 3},
		{"antidemocrática", 2, 5, 4},
	}
	for _, testCase := range testCases {
		w := snowballword.New(testCase.word)
		r1start, r2start, rvstart := findRegions(w)
		if r1start != testCase.r1start || r2start != testCase.r2start || rvstart != testCase.rvstart {
			t.Errorf("Expect findRegions(\"%v\") = %v, %v, %v, but got %v, %v, %v",
				testCase.word, testCase.r1start, testCase.r2start, testCase.rvstart,
				r1start, r2start, rvstart,
			)
		}

	}
}
