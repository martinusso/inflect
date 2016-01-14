package inflect

import "regexp"

var (
	transliterations = map[string]*regexp.Regexp{
		"A":  regexp.MustCompile(`À|Á|Â|Ã|Ä`),
		"AA": regexp.MustCompile(`Å`),
		"AE": regexp.MustCompile(`Æ`),
		"C":  regexp.MustCompile(`Ç`),
		"E":  regexp.MustCompile(`È|É|Ê|Ë`),
		"D":  regexp.MustCompile(`Ð`),
		"I":  regexp.MustCompile(`Ì|Í|Î|Ï`),
		"L":  regexp.MustCompile(`Ł`),
		"N":  regexp.MustCompile(`Ñ`),
		"O":  regexp.MustCompile(`Ò|Ó|Ô|Õ|Ö`),
		"OE": regexp.MustCompile(`Œ|Ø`),
		"Th": regexp.MustCompile(`Þ`),
		"U":  regexp.MustCompile(`Ù|Ú|Û|Ü`),
		"Y":  regexp.MustCompile(`Ý`),
		"a":  regexp.MustCompile(`à|á|â|ã|ä`),
		"aa": regexp.MustCompile(`å`),
		"ae": regexp.MustCompile(`æ`),
		"c":  regexp.MustCompile(`ç`),
		"e":  regexp.MustCompile(`è|é|ê|ë`),
		"i":  regexp.MustCompile(`ì|í|î|ï`),
		"n":  regexp.MustCompile(`ñ|ń`),
		"o":  regexp.MustCompile(`ò|ó|ô|õ|ö|ō`),
		"oe": regexp.MustCompile(`œ|ø`),
		"s":  regexp.MustCompile(`ś`),
		"ss": regexp.MustCompile(`ß`),
		"u":  regexp.MustCompile(`ù|ú|û|ü|ũ|ū|ŭ|ů|ű|ų`),
		"y":  regexp.MustCompile(`ý|ÿ`),
		"z":  regexp.MustCompile(`ż`),
		"d":  regexp.MustCompile(`ð`),
		"l":  regexp.MustCompile(`ł`),
		"th": regexp.MustCompile(`þ`),
	}
)

// Asciify replaces a set of accented characters with ascii equivalents
func Asciify(word string) string {
	for repl, regex := range transliterations {
		word = regex.ReplaceAllString(word, repl)
	}
	return word
}
