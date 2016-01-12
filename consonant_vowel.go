package inflect

import "strings"

var (
	consonants = []string{"b", "c", "d", "f", "g", "h", "j", "k", "l", "m", "n",
		"p", "q", "r", "s", "t", "v", "x", "z", "w", "y"}
	vowels = []string{"a", "e", "i", "o", "u"}
)

// IsConsonant check if a character is a consonant
func IsConsonant(s string) bool {
	for _, c := range consonants {
		if c == strings.ToLower(s) {
			return true
		}
	}
	return false
}

// IsVowel check if a character is a vowel
func IsVowel(s string) bool {
	for _, v := range vowels {
		if v == strings.ToLower(s) {
			return true
		}
	}
	return false
}
