package inflect

import "testing"

func TestCamelize(t *testing.T) {
	m := map[string]string{
		"upper_camel_case":    "UpperCamelCase",
		"pinky and the brain": "PinkyAndTheBrain",
		"j. r. r. tolkien":    "JRRTolkien",
		"eärendil":            "Earendil",
	}
	for s, expected := range m {
		got := Camelize(s)
		if got != expected {
			t.Errorf("Expected '%s' got '%s'", expected, got)
		}
	}
}
