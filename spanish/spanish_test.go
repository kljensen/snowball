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

type stepFunc func(*snowballword.SnowballWord) bool
type stepTest struct {
	WordIn     string
	R1start    int
	R2start    int
	RVstart    int
	Changed    bool
	WordOut    string
	R1startOut int
	R2startOut int
	RVstartOut int
}

func runStepTest(t *testing.T, f stepFunc, tcs []stepTest) {
	for _, testCase := range tcs {
		w := snowballword.New(testCase.WordIn)
		w.R1start = testCase.R1start
		w.R2start = testCase.R2start
		w.RVstart = testCase.RVstart
		retval := f(w)
		if retval != testCase.Changed || w.String() != testCase.WordOut || w.R1start != testCase.R1startOut || w.R2start != testCase.R2startOut || w.RVstart != testCase.RVstartOut {
			t.Errorf("Expected %v -> \"{%v, %v, %v, %v}\", but got \"{%v, %v, %v, %v}\"", testCase.WordIn, testCase.WordOut, testCase.R1startOut, testCase.R2startOut, testCase.RVstartOut, w.String(), w.R1start, w.R2start, w.RVstart)
		}
	}
}

// Test step0, the removal of pronoun suffixes.
//
func Test_step0(t *testing.T) {
	testCases := []stepTest{
		{"liberarlo", 3, 5, 3, true, "liberar", 3, 5, 3},
		{"ejecutarse", 2, 4, 3, true, "ejecutar", 2, 4, 3},
		{"convirtiéndolas", 3, 6, 3, true, "convirtiendo", 3, 6, 3},
		{"perfeccionarlo", 3, 6, 3, true, "perfeccionar", 3, 6, 3},
		{"formarlo", 3, 6, 3, true, "formar", 3, 6, 3},
		{"negociarlo", 3, 5, 3, true, "negociar", 3, 5, 3},
		{"dirigirla", 3, 5, 3, true, "dirigir", 3, 5, 3},
		{"malograrlas", 3, 5, 3, true, "malograr", 3, 5, 3},
		{"atacarlos", 2, 4, 3, true, "atacar", 2, 4, 3},
		{"originarla", 2, 4, 3, true, "originar", 2, 4, 3},
		{"ponerlos", 3, 5, 3, true, "poner", 3, 5, 3},
		{"ubicándolo", 2, 4, 3, true, "ubicando", 2, 4, 3},
		{"dejarme", 3, 5, 3, true, "dejar", 3, 5, 3},
		{"regalarnos", 3, 5, 3, true, "regalar", 3, 5, 3},
		{"resolverlas", 3, 5, 3, true, "resolver", 3, 5, 3},
		{"esperarse", 2, 5, 4, true, "esperar", 2, 5, 4},
		{"cuidarlo", 4, 6, 3, true, "cuidar", 4, 6, 3},
		{"empezarlos", 2, 5, 4, true, "empezar", 2, 5, 4},
		{"gastarla", 3, 6, 3, true, "gastar", 3, 6, 3},
		{"levantarme", 3, 5, 3, true, "levantar", 3, 5, 3},
		{"ausentarse", 3, 5, 3, true, "ausentar", 3, 5, 3},
		{"colocándose", 3, 5, 3, true, "colocando", 3, 5, 3},
		{"suponerse", 3, 5, 3, true, "suponer", 3, 5, 3},
		{"someterlos", 3, 5, 3, true, "someter", 3, 5, 3},
		{"criticarlos", 4, 6, 3, true, "criticar", 4, 6, 3},
		{"consolidarlo", 3, 6, 3, true, "consolidar", 3, 6, 3},
		{"globalizarse", 4, 6, 3, true, "globalizar", 4, 6, 3},
		{"corregirla", 3, 6, 3, true, "corregir", 3, 6, 3},
		{"aplicarle", 2, 5, 4, true, "aplicar", 2, 5, 4},
		{"casarse", 3, 5, 3, true, "casar", 3, 5, 3},
		{"costándole", 3, 6, 3, true, "costando", 3, 6, 3},
		{"rescindirlo", 3, 6, 3, true, "rescindir", 3, 6, 3},
		{"quitándole", 4, 6, 3, true, "quitando", 4, 6, 3},
		{"conservarse", 3, 6, 3, true, "conservar", 3, 6, 3},
		{"venderlo", 3, 6, 3, true, "vender", 3, 6, 3},
		{"garantizarse", 3, 5, 3, true, "garantizar", 3, 5, 3},
		{"disfrutarse", 3, 7, 3, true, "disfrutar", 3, 7, 3},
		{"comunicarse", 3, 5, 3, true, "comunicar", 3, 5, 3},
		{"propiciarse", 4, 6, 3, true, "propiciar", 4, 6, 3},
		{"otorgarnos", 2, 4, 3, true, "otorgar", 2, 4, 3},
		{"contorsionarse", 3, 6, 3, true, "contorsionar", 3, 6, 3},
		{"motivarlas", 3, 5, 3, true, "motivar", 3, 5, 3},
		{"congelarse", 3, 6, 3, true, "congelar", 3, 6, 3},
		{"generandoles", 3, 5, 3, true, "generando", 3, 5, 3},
		{"evitarlo", 2, 4, 3, true, "evitar", 2, 4, 3},
		{"atenderlos", 2, 4, 3, true, "atender", 2, 4, 3},
		{"apoyándola", 2, 4, 3, true, "apoyando", 2, 4, 3},
		{"pasarse", 3, 5, 3, true, "pasar", 3, 5, 3},
		{"escucharlos", 2, 5, 4, true, "escuchar", 2, 5, 4},
		{"intervenirse", 2, 5, 4, true, "intervenir", 2, 5, 4},
		{"contratarle", 3, 7, 3, true, "contratar", 3, 7, 3},
		{"retirándose", 3, 5, 3, true, "retirando", 3, 5, 3},
		{"quitarles", 4, 6, 3, true, "quitar", 4, 6, 3},
		{"reforzarlas", 3, 5, 3, true, "reforzar", 3, 5, 3},
		{"obtenerla", 2, 5, 4, true, "obtener", 2, 5, 4},
		{"considerarlo", 3, 6, 3, true, "considerar", 3, 6, 3},
		{"regresarse", 3, 6, 3, true, "regresar", 3, 6, 3},
		{"ponerse", 3, 5, 3, true, "poner", 3, 5, 3},
		{"llevándose", 4, 6, 3, true, "llevando", 4, 6, 3},
		{"ocuparse", 2, 4, 3, true, "ocupar", 2, 4, 3},
		{"aprovecharse", 2, 5, 4, true, "aprovechar", 2, 5, 4},
		{"corregirlo", 3, 6, 3, true, "corregir", 3, 6, 3},
		{"probarle", 4, 6, 3, true, "probar", 4, 6, 3},
		{"comernos", 3, 5, 3, true, "comer", 3, 5, 3},
		{"iniciarme", 2, 4, 3, true, "iniciar", 2, 4, 3},
		{"concentrarse", 3, 6, 3, true, "concentrar", 3, 6, 3},
		{"llevarse", 4, 6, 3, true, "llevar", 4, 6, 3},
		{"difundirlo", 3, 5, 3, true, "difundir", 3, 5, 3},
		{"basándose", 3, 5, 3, true, "basando", 3, 5, 3},
		{"destinarlos", 3, 6, 3, true, "destinar", 3, 6, 3},
		{"reubicarse", 4, 6, 3, true, "reubicar", 4, 6, 3},
		{"manteniéndose", 3, 6, 3, true, "manteniendo", 3, 6, 3},
		{"colocarla", 3, 5, 3, true, "colocar", 3, 5, 3},
		{"pasarles", 3, 5, 3, true, "pasar", 3, 5, 3},
		{"depositarse", 3, 5, 3, true, "depositar", 3, 5, 3},
		{"tragarse", 4, 6, 3, true, "tragar", 4, 6, 3},
		{"eliminarla", 2, 4, 3, true, "eliminar", 2, 4, 3},
		{"eliminarse", 2, 4, 3, true, "eliminar", 2, 4, 3},
		{"apegarnos", 2, 4, 3, true, "apegar", 2, 4, 3},
		{"asociarse", 2, 4, 3, true, "asociar", 2, 4, 3},
		{"cambiarlos", 3, 7, 3, true, "cambiar", 3, 7, 3},
		{"envolviéndose", 2, 5, 4, true, "envolviendo", 2, 5, 4},
		{"lograrse", 3, 6, 3, true, "lograr", 3, 6, 3},
		{"mostrarse", 3, 7, 3, true, "mostrar", 3, 7, 3},
		{"pasarle", 3, 5, 3, true, "pasar", 3, 5, 3},
		{"enfrentándose", 2, 6, 5, true, "enfrentando", 2, 6, 5},
		{"permitirse", 3, 6, 3, true, "permitir", 3, 6, 3},
		{"sanearlas", 3, 6, 3, true, "sanear", 3, 6, 3},
		{"refugiarse", 3, 5, 3, true, "refugiar", 3, 5, 3},
		{"relacionarse", 3, 5, 3, true, "relacionar", 3, 5, 3},
		{"sacarlo", 3, 5, 3, true, "sacar", 3, 5, 3},
		{"organizarse", 2, 5, 4, true, "organizar", 2, 5, 4},
		{"familiarizarse", 3, 5, 3, true, "familiarizar", 3, 5, 3},
		{"decidirse", 3, 5, 3, true, "decidir", 3, 5, 3},
		{"tomarle", 3, 5, 3, true, "tomar", 3, 5, 3},
		{"volverlas", 3, 6, 3, true, "volver", 3, 6, 3},
		{"efectuarse", 2, 4, 3, true, "efectuar", 2, 4, 3},
		{"elegirse", 2, 4, 3, true, "elegir", 2, 4, 3},
		{"establecerse", 2, 5, 4, true, "establecer", 2, 5, 4},
		{"ponerles", 3, 5, 3, true, "poner", 3, 5, 3},
	}
	runStepTest(t, step0, testCases)
}
