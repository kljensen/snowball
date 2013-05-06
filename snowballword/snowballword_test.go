package snowballword

import "testing"

func Test_New(t *testing.T) {
	w := New("kyle")
	if w.String() != "kyle" {
		t.Errorf("Expected \"%v\" but got \"%v\"", "kyle", w.String())
	}
}

func Test_FirstPrefix(t *testing.T) {
	var testCases = []struct {
		input    string
		prefixes []string
		prefix   string
	}{
		{"firehose", []string{"x", "fi"}, "fi"},
		{"firehose", []string{"x", "fix", "fi"}, "fi"},
		{"firehose", []string{"x", "fi"}, "fi"},
		{"firehose", []string{"fire", "fi"}, "fire"},
		{"firehose", []string{"fixre", "xfi"}, ""},
		{"firehose", []string{"firehosex"}, ""},
	}
	for _, tc := range testCases {
		w := New(tc.input)
		prefix := w.FirstPrefix(tc.prefixes...)
		if prefix != tc.prefix {
			t.Errorf("Expected \"{%v}\" but got \"{%v}\"", tc.prefix, prefix)
		}
	}
}

func Test_FirstSuffix(t *testing.T) {
	var testCases = []struct {
		input    string
		suffixes []string
		suffix   string
	}{
		{"firehose", []string{"x", "fi"}, ""},
		{"firehose", []string{"x", "hose", "fi"}, "hose"},
		{"firehose", []string{"x", "se"}, "se"},
		{"firehose", []string{"fire", "xfirehose"}, ""},
	}
	for _, tc := range testCases {
		w := New(tc.input)
		suffix := w.FirstSuffix(tc.suffixes...)
		if suffix != tc.suffix {
			t.Errorf("Expected \"{%v}\" but got \"{%v}\"", tc.suffix, suffix)
		}
	}
}

func Test_ReplaceSuffix(t *testing.T) {
	var testCases = []struct {
		input          string
		r1start        int
		r2start        int
		suffix         string
		repl           string
		output         string
		outputR1String string
		outputR2String string
	}{
		{"accliviti", 2, 6, "iviti", "ive", "acclive", "clive", "e"},
		{"skating", 4, 6, "ing", "e", "skate", "e", ""},
	}
	for _, tc := range testCases {
		w := New(tc.input)
		w.R1start = tc.r1start
		w.R2start = tc.r2start
		w.ReplaceSuffix(tc.suffix, tc.repl, true)
		if w.String() != tc.output || w.R1String() != tc.outputR1String || w.R2String() != tc.outputR2String {
			t.Errorf("Expected \"{%v, %v, %v}\" but got \"{%v, %v, %v}\"", tc.output, tc.outputR1String, tc.outputR2String, w.String(), w.R1String(), w.R2String())
		}
	}
}
