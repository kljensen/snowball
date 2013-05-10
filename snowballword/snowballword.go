/*
	This package defines a SnowballWord struct that is used
	to encapsulate most of the "state" variables we must track
	when stemming a word.  The SnowballWord struct also has
	a few methods common to stemming in a variety of languages.
*/
package snowballword

// SnowballWord represents a word that is going to be stemmed.
// 
type SnowballWord struct {

	// A slice of runes
	RS []rune

	// The index in RS where the R1 region begins
	R1start int

	// The index in RS where the R2 region begins
	R2start int

	// The index in RS where the RV region begins
	RVstart int
}

// Create a new SnowballWord struct
func New(in string) (word *SnowballWord) {
	word = &SnowballWord{RS: []rune(in)}
	word.R1start = len(word.RS)
	word.R2start = len(word.RS)
	word.RVstart = len(word.RS)
	return
}

// Replace a suffix and adjust R1start and R2start as needed.
// If `force` is false, check to make sure the suffix exists first.
//
func (w *SnowballWord) ReplaceSuffix(suffix, replacement string, force bool) bool {

	var (
		doReplacement bool
		suffixRunes   []rune
	)
	if force {
		doReplacement = true
		suffixRunes = []rune(suffix)
	} else {
		var foundSuffix string
		foundSuffix, suffixRunes = w.FirstSuffix(suffix)
		if foundSuffix == suffix {
			doReplacement = true
		}
	}
	if doReplacement == false {
		return false
	}
	w.ReplaceSuffixRunes(suffixRunes, []rune(replacement), true)
	return true
}

// Replace a suffix and adjust R1start and R2start as needed.
// If `force` is false, check to make sure the suffix exists first.
//
func (w *SnowballWord) ReplaceSuffixRunes(suffixRunes []rune, replacementRunes []rune, force bool) bool {

	if force || w.HasSuffixRunes(suffixRunes) {
		lenWithoutSuffix := len(w.RS) - len(suffixRunes)
		w.RS = append(w.RS[:lenWithoutSuffix], replacementRunes...)

		// If R, R2, & RV are now beyond the length
		// of the word, they are set to the length
		// of the word.  Otherwise, they are left
		// as they were.
		w.resetR1R2()
		return true
	}
	return false
}

// Resets R1start and R2start to ensure they 
// are within bounds of the current rune slice.
func (w *SnowballWord) resetR1R2() {
	rsLen := len(w.RS)
	if w.R1start > rsLen {
		w.R1start = rsLen
	}
	if w.R2start > rsLen {
		w.R2start = rsLen
	}
	if w.RVstart > rsLen {
		w.RVstart = rsLen
	}
}

// Return a slice of w.RS, allowing the start
// and stop to be out of bounds.
//
func (w *SnowballWord) slice(start, stop int) []rune {
	startMin := 0
	if start < startMin {
		start = startMin
	}
	max := len(w.RS) - 1
	if start > max {
		start = max
	}
	if stop > max {
		stop = max
	}
	return w.RS[start:stop]
}

// Return the R1 region as a slice of runes
func (w *SnowballWord) R1() []rune {
	return w.RS[w.R1start:]
}

// Return the R1 region as a string
func (w *SnowballWord) R1String() string {
	return string(w.R1())
}

// Return the R2 region as a slice of runes
func (w *SnowballWord) R2() []rune {
	return w.RS[w.R2start:]
}

// Return the R2 region as a string
func (w *SnowballWord) R2String() string {
	return string(w.R2())
}

// Return the RV region as a slice of runes
func (w *SnowballWord) RV() []rune {
	return w.RS[w.RVstart:]
}

// Return the RV region as a string
func (w *SnowballWord) RVString() string {
	return string(w.RV())
}

// Return the SnowballWord as a string
func (w *SnowballWord) String() string {
	return string(w.RS)
}

// Return the first prefix found or the empty string.
func (w *SnowballWord) FirstPrefix(prefixes ...string) (foundPrefix string, foundPrefixRunes []rune) {
	found := false
	rsLen := len(w.RS)

	for _, prefix := range prefixes {
		prefixRunes := []rune(prefix)
		if len(prefixRunes) > rsLen {
			continue
		}

		found = true
		for i, r := range prefixRunes {
			if i > rsLen-1 || (w.RS)[i] != r {
				found = false
				break
			}
		}
		if found {
			foundPrefix = prefix
			foundPrefixRunes = prefixRunes
			break
		}
	}
	return
}

// Return true if the first `endPos` runes of `w` end in `suffixRunes`
//
func (w *SnowballWord) HasSuffixRunesAt(endPos int, suffixRunes []rune) bool {
	rsLen := len(w.RS)
	numMatching := 0
	suffixLen := len(suffixRunes)
	for i := 0; i < endPos && i < rsLen && i < suffixLen; i++ {
		if w.RS[endPos-i-1] != suffixRunes[suffixLen-i-1] {
			break
		} else {
			numMatching += 1
		}
	}
	if numMatching == suffixLen {
		return true
	}
	return false
}

// Return true if `w` ends with `suffixRunes`
//
func (w *SnowballWord) HasSuffixRunes(suffixRunes []rune) bool {
	return w.HasSuffixRunesAt(len(w.RS), suffixRunes)
}

// Return the first suffix found or the empty string.
func (w *SnowballWord) FirstSuffixAt(endPos int, sufficies ...string) (suffix string, suffixRunes []rune) {
	for _, suffix := range sufficies {
		suffixRunes := []rune(suffix)
		if w.HasSuffixRunesAt(endPos, suffixRunes) {
			return suffix, suffixRunes
		}
	}

	// Empty out suffixRunes
	suffixRunes = suffixRunes[:0]
	return "", suffixRunes
}

// Return the first suffix found or the empty string.
func (w *SnowballWord) FirstSuffix(sufficies ...string) (suffix string, suffixRunes []rune) {
	return w.FirstSuffixAt(len(w.RS), sufficies...)
}
