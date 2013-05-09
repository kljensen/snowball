package snowball

import "testing"

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
