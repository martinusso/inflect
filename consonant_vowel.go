package inflect

import "regexp"

// IsConsonant check if a character is a consonant
func IsConsonant(s string) bool {
	matched, _ := regexp.MatchString("^[^aeiuo]+$", s)
	return matched
}

// IsVowel check if a character is a vowel
func IsVowel(s string) bool {
	matched, _ := regexp.MatchString("^[aeiuo]+$", s)
	return matched
}
