package inflect

import (
	"log"
	"regexp"
	"strings"
)

// Parameterize replaces special characters in a pretty string
func Parameterize(s string) string {
	sep := "-"
	return ParameterizeSep(s, sep)
}

// ParameterizeSep replaces special characters, according to the separator, in a string
func ParameterizeSep(s, sep string) string {
	safe := Asciify(s)
	safe = strings.ToLower(safe)
	reg, err := regexp.Compile("[^A-Za-z0-9]+")
	if err != nil {
		log.Fatal(err)
	}
	safe = reg.ReplaceAllString(safe, sep)
	return safe
}
