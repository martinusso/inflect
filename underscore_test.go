package inflect

import "testing"

func TestUnderscore(t *testing.T) {
	m := map[string]string{
		"UnderScore":          "under_score",
		"Pinky and the Brain": "pinky_and_the_brain",
		"J. R. R. Tolkien":    "j_r_r_tolkien",
	}
	for s, expected := range m {
		got := Underscore(s)
		if got != expected {
			t.Errorf("Expected '%s' got '%s'", expected, got)
		}
	}
}
