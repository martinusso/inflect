package inflect

import "testing"

func TestParameterize(t *testing.T) {
	parametrizations := map[string]string{
		"J. R. R. Tolkien": "j-r-r-tolkien",
		"Eärendil":         "earendil",
		"Lúthien":          "luthien",
		"under_score":      "under-score",
		"Um1":              "um1",
	}
	for k, expected := range parametrizations {
		got := Parameterize(k)
		if got != expected {
			t.Errorf("Expected '%s' got '%s'", expected, got)
		}
	}
}

func TestParameterizeSep(t *testing.T) {
	parametrizations := map[string]string{
		"J. R. R. Tolkien": "j_r_r_tolkien",
		"Lúthien":          "luthien",
		"under_score":      "under_score",
	}
	for k, expected := range parametrizations {
		got := ParameterizeSep(k, "_")
		if got != expected {
			t.Errorf("Expected '%s' got '%s'", expected, got)
		}
	}
}
