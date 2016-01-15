package inflect

import "testing"

func TestConsonant(t *testing.T) {
	consonants := []string{"b", "c", "d", "f", "g", "h", "j", "k", "l", "m", "n",
		"p", "q", "r", "s", "t", "v", "x", "z", "w", "y"}

	for _, c := range consonants {
		if !IsConsonant(c) {
			t.Errorf("Expected '%s' is a consonant got false", c)
		}
	}
	if IsConsonant("abc") {
		t.Errorf("'abc' should not be considered consonant")
	}
}

func TestVowel(t *testing.T) {
	vowels := []string{"a", "e", "i", "o", "u"}

	for _, v := range vowels {
		if !IsVowel(v) {
			t.Errorf("Expected '%s' is a vowel got false", v)
		}
	}
	if IsVowel("abc") {
		t.Errorf("'abc' should not be considered vowel")
	}
}
