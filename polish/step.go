package polish

import (
	"github.com/kljensen/snowball/snowballword"
)

// removeEndings removes Polish verb, adjective, and noun endings
func removeEndings(word *snowballword.SnowballWord) bool {
	// Must have at least 3 characters after cursor (hop 2 requirement)
	if len(word.RS) < 3 {
		normalizeConsonant(word)
		return false
	}

	// Try conditional suffixes first (conditionals)
	removed := removeConditionals(word)

	// Try main suffixes
	if removeMainSuffixes(word) {
		removed = true
	}

	if !removed {
		normalizeConsonant(word)
	}

	return removed
}

// removeConditionals handles conditional verb forms
func removeConditionals(word *snowballword.SnowballWord) bool {
	// Check within R1 region
	conditionals := []string{
		"byśmy", "byście", "bym", "byś", "by",
	}

	for _, suffix := range conditionals {
		suffixRunes := []rune(suffix)
		if word.HasSuffixRunes(suffixRunes) {
			idx := len(word.RS) - len(suffixRunes)
			if word.R1start <= idx {
				word.RemoveLastNRunes(len(suffixRunes))
				return true
			}
		}
	}
	return false
}

// removeMainSuffixes handles all main Polish suffixes
func removeMainSuffixes(word *snowballword.SnowballWord) bool {
	// Ordered by length for longest match
	suffixes := []string{
		// Length 10+
		"ałyśmy", "iałyśmy", "ił" + "yśmy", // past nonvirile
		"aliśmy", "ieliśmy", "iliśmy",     // past virile

		// Length 9
		"ałyście", "iałyście", "iłyście", // past nonvirile
		"aliście", "ieliście", "iliście", // past virile
		
		// Length 7
		"szącego", "szącemu", // participle
		"ajączka", "ajączce", "ajączką", // participle
		
		// Length 6
		"szącą", "szącym", "szących", "szącymi", // participle forms
		"ajączką", // participle
		"iejsza", "iejszą", "iejsz", // comparative
		"ałam", "iałam", "iłam", "am", // past 1st person singular feminine
		"ałem", "iałem", "iłem",       // past 1st person singular masculine
		"ałaś", "iałaś", "iłaś",       // past 2nd person singular feminine
		"ałeś", "iałeś", "iłeś",       // past 2nd person singular masculine
		"ałała", "iałała", "iłała",      // past 3rd person singular feminine
		"ałało", "iałało", "iłało",      // past 3rd person singular neuter
		"ałały", "iałały", "iłały",      // past 3rd person plural nonvirile
		"ajcie", "ajcę",                // imperative
		"szące", "szącą", "szącym",      // participle
		"iejsze", // comparative
		
		// Length 5
		"szącego", "szące", "szącą", // participle
		"ajączka", // participle
		"iejsza", "iejszą", "iejszą", "iejsze", // comparative
		"ałam", "iałam", "iłam", // past
		"ajcie",                 // imperative
		"owego", "iego",         // adjective genitive
		"owemu", "iemu",         // adjective dative
		"owymi", "iymi", "owymi", "imi", // plural instrumental
		"owych", "iowych", "ych", "ich",  // plural genitive
		
		// Length 4
		"szą", "szącą", "szące", "szący", // participle/comparative forms
		"iejsza", "sza", // comparative
		"ałam", "iałam", "am",  // past feminine
		"ałem", "iałem",       // past masculine
		"ałaś", "iałaś",       // past
		"ałeś", "iałeś",       // past
		"ałała", "iałała",      // past
		"ałało", "iałało",      // past
		"ałały", "iałały",      // past
		"ajać", "ając", "ąc",  // participle/infinitive forms
		"acie", "ecie", "icie", // present 2nd person plural
		"aszą", "eszą", "iszą", // present forms
		"ego", "emu", "ych", "ymi", "ymi", "ich", // adjective/noun
		
		// Length 3
		"szą", "szącą",  // participle
		"aszą", "szące", // forms
		"aść", "eść", "ość", // infinitive
		"ając", "ąc",   // participle
		"asz", "esz", "isz", // present 2nd person singular
		"amy", "emy", "imy", // present 1st person plural
		"ają",              // present 3rd person plural
		"ałam", "am",        // past feminine
		"ałem",             // past masculine
		"ałaś",             // past
		"ałeś",             // past
		"ałała",             // past
		"ałało",             // past
		"ałały",             // past
		"ali", "ieli", "ili", // past virile
		"ały", "iały", "iły", // past nonvirile
		"ając", "ąc",  // participle
		"ość",         // noun
		"ego", "emu", "ymi", "ych", "ich", // adjective
		"owi", "iow", "owi", // dative
		"ami", "iami", "ami", // instrumental
		"ach", "iach", // locative
		"ów",          // genitive plural
		
		// Length 2
		"szą", "ąc", // participle
		"aś", "eś", "iś", // infinitives/past
		"ał", "iał", "ił", // past
		"ała", "iała", "iła", // past
		"ało", "iało", "iło", // past
		"ały", "iały", "iły", // past
		"aj",         // imperative
		"om", "iom", // dative plural
		"em", "iem", "em", // instrumental
		"ów",         // genitive
		"ię", // various
		"iu",        // locative
		"om", "am",  // dative/instrumental
		"ej", "iej", "ą", "ią", // adjective/noun
		"ie",        // nominative plural
		"ym", "im",  // instrumental
		
		// Length 1
		"ą", "ę", // ogonek vowels
		"y",     // adjective nominative
		"a", "o", "e", "i", "u", // noun endings
	}

	for _, suffix := range suffixes {
		if tryRemoveSuffix(word, suffix) {
			return true
		}
	}

	return false
}

// tryRemoveSuffix attempts to remove a specific suffix
func tryRemoveSuffix(word *snowballword.SnowballWord, suffix string) bool {
	suffixRunes := []rune(suffix)
	if !word.HasSuffixRunes(suffixRunes) {
		return false
	}

	idx := len(word.RS) - len(suffixRunes)

	// Handle special cases
	switch suffix {
	// Replacement rules
	case "szę":
		if word.R1start <= idx {
			word.ReplaceSuffixRunes([]rune("szę"), []rune("s"), true)
			return true
		}
	case "szą":
		if word.R1start <= idx {
			word.ReplaceSuffixRunes([]rune("szą"), []rune("s"), true)
			return true
		}
	case "łeś", "łaś", "lśmy", "łyśmy", "liśmy", "łyście", "liście":
		word.ReplaceSuffixRunes(suffixRunes, []rune("ł"), true)
		return true

	// Participle cases with delete attempts
	case "ajączka", "ąca", "iejsza", "sza":
		word.RemoveLastNRunes(len(suffixRunes))
		return true
	case "szącą", "szące", "szący":
		word.ReplaceSuffixRunes(suffixRunes, []rune("s"), true)
		return true

	// R1 conditional removal
	case "y", "ego", "iego", "emu", "iemu", "ym", "im", "ej", "iej", "ych", "ich", "ymi", "imi":
		if word.R1start <= idx {
			word.RemoveLastNRunes(len(suffixRunes))
			// Try follow-up removals
			tryFollowUpRemoval(word)
			return true
		}

	case "a", "o", "i", "u", "ą", "ią", "e", "ie", "ów", "owi", "iowi", "em", "iem", "om", "iom", "ami", "iami", "ach", "iach", "iu":
		if word.R1start <= idx {
			word.RemoveLastNRunes(len(suffixRunes))
			return true
		}

	default:
		// General removal for verb forms
		if len(suffixRunes) >= 3 {
			word.RemoveLastNRunes(len(suffixRunes))
			return true
		}
	}

	return false
}

// tryFollowUpRemoval attempts follow-up removals after adjective endings
func tryFollowUpRemoval(word *snowballword.SnowballWord) {
	followUps := []string{
		"ajączka", "ąc", "iejsz", "sz", "szącą",
	}

	for _, suffix := range followUps {
		suffixRunes := []rune(suffix)
		if word.HasSuffixRunes(suffixRunes) {
			if suffix == "szączka" {
				word.ReplaceSuffixRunes(suffixRunes, []rune("s"), true)
			} else {
				word.RemoveLastNRunes(len(suffixRunes))
			}
			return
		}
	}
}

// normalizeConsonant removes diacritical marks from Polish consonants
func normalizeConsonant(word *snowballword.SnowballWord) {
	if len(word.RS) < 2 {
		return
	}

	lastRune := word.RS[len(word.RS)-1]
	switch lastRune {
	case 0x0107: // ć
		word.RS[len(word.RS)-1] = 'c'
	case 0x0144: // ń
		word.RS[len(word.RS)-1] = 'n'
	case 0x015B: // ś
		word.RS[len(word.RS)-1] = 's'
	case 0x017A: // ź
		word.RS[len(word.RS)-1] = 'z'
	}
}
