package snowball

import (
	"regexp"
	"testing"
)

func Test_Stem(t *testing.T) {
	testCases := []struct {
		in            string
		language      string
		stemStopWords bool
		out           string
		nilErr        bool
	}{
		{"aberration", "english", true, "aberr", true},
		{"abruptness", "english", true, "abrupt", true},
		{"absolute", "english", true, "absolut", true},
		{"abated", "english", true, "abat", true},
		{"acclivity", "english", true, "accliv", true},
		{"accumulations", "english", true, "accumul", true},
		{"agreement", "english", true, "agreement", true},
		{"breed", "english", true, "breed", true},
		{"ape", "english", true, "ape", true},
		{"skating", "english", true, "skate", true},
		{"fluently", "english", true, "fluentli", true},
		{"ied", "english", true, "ie", true},
		{"ies", "english", true, "ie", true},
		// Change stemStopWords
		{"above", "english", true, "abov", true},
		{"because", "english", false, "because", true},
		// Give invalid language
		{"because", "klingon", false, "", false},

		// Spanish tests, a few
		{"lejana", "spanish", true, "lejan", true},
		{"preocuparse", "spanish", true, "preocup", true},
		{"oposición", "spanish", true, "oposicion", true},
		{"prisionero", "spanish", true, "prisioner", true},
		{"ridiculización", "spanish", true, "ridiculiz", true},
		{"cotidianeidad", "spanish", true, "cotidian", true},
		{"portezuela", "spanish", true, "portezuel", true},
		{"enriquecerse", "spanish", true, "enriquec", true},
		{"campesinos", "spanish", true, "campesin", true},
		{"desalojó", "spanish", true, "desaloj", true},
		{"anticipadas", "spanish", true, "anticip", true},
		{"goyesca", "spanish", true, "goyesc", true},
		{"band", "spanish", true, "band", true},
	}
	for _, testCase := range testCases {
		out, err := Stem(testCase.in, testCase.language, testCase.stemStopWords)
		nilErr := true
		if err != nil {
			nilErr = false
		}
		if out != testCase.out || nilErr != testCase.nilErr {
			t.Errorf("Stem(\"%v\", \"%v\", %v) = \"%v, %v\", but expected %v, %v",
				testCase.in, testCase.language, testCase.stemStopWords,
				out, nilErr, testCase.out, testCase.nilErr,
			)
		}

	}
}

// Test if the VERSION constant is correctly formatted
//
func Test_Version(t *testing.T) {
	validVersionRegexp := regexp.MustCompile(`^v\d+\.\d+\.\d+$`)
	if validVersionRegexp.MatchString(VERSION) == false {
		t.Errorf("Invalid version specified: %v", VERSION)
	}
}
