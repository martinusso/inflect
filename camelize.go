package inflect

import (
	"log"
	"regexp"
	"strings"
)

// Camelize converts strings to UpperCamelCase
func Camelize(s string) string {
	safe := Transliterate(s)
	safe = strings.ToLower(safe)
	reg, err := regexp.Compile("[^A-Za-z0-9]+")
	if err != nil {
		log.Fatal(err)
	}
	safe = reg.ReplaceAllString(safe, " ")
	return strings.Replace(strings.Title(safe), " ", "", -1)
}

// LowerCamelize converts strings to lowerCamelCase
func LowerCamelize(s string) string {
	camelized := Camelize(s)
	return strings.ToLower(string(camelized[0])) + camelized[1:]
}
