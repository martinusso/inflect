package inflect

import "testing"

func TestConsonant(t *testing.T) {
	for _, c := range consonants {
		if !IsConsonant(c) {
			t.Errorf("Expected '%s' is a consonant got false", c)
		}
	}
}

func TestVowel(t *testing.T) {
	for _, v := range vowels {
		if !IsVowel(v) {
			t.Errorf("Expected '%s' is a vowel got false", v)
		}
	}
}
