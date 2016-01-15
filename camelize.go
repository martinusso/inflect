package inflect

import (
	"log"
	"regexp"
	"strings"
)

// Camelize converts strings to UpperCamelCase
func Camelize(s string) string {
	safe := Transliterate(s)
	reg, err := regexp.Compile("[^A-Za-z0-9]+")
	if err != nil {
		log.Fatal(err)
	}
	safe = reg.ReplaceAllString(safe, " ")
	return strings.Replace(strings.Title(safe), " ", "", -1)
}
