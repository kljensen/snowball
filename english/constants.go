/*
	Constants used in the English Snowball stemmer.
*/
package english

const vowels string = "aeiouy"

var liEnding = "cdeghkmnrt"
var step1aSuffixes = [...]string{"sses", "ied", "ies", "us", "ss", "s"}
var step1bSuffixes = [...]string{"eedly", "ingly", "edly", "eed", "ing", "ed"}
var step2Suffixes = [...]string{
	"ization", "ational", "fulness", "ousness",
	"iveness", "tional", "biliti", "lessli",
	"entli", "ation", "alism", "aliti", "ousli",
	"iviti", "fulli", "enci", "anci", "abli",
	"izer", "ator", "alli", "bli", "ogi", "li",
}
var step3Suffixes = [...]string{
	"ational", "tional", "alize", "icate", "iciti",
	"ative", "ical", "ness", "ful",
}
var step4Suffixes = [...]string{
	"ement", "ance", "ence", "able", "ible", "ment",
	"ant", "ent", "ism", "ate", "iti", "ous",
	"ive", "ize", "ion", "al", "er", "ic",
}
var step5Suffixes = [...]string{"e", "l"}
