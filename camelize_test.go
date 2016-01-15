package inflect

import "testing"

func TestCamelize(t *testing.T) {
	m := map[string]string{
		"upper_camel_case":    "UpperCamelCase",
		"UPPER_CAMEL_CASE":    "UpperCamelCase",
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

func TestLowerCamelize(t *testing.T) {
	m := map[string]string{
		"LOWER_CAMEL_CASE":    "lowerCamelCase",
		"Pinky and the brain": "pinkyAndTheBrain",
	}
	for s, expected := range m {
		got := LowerCamelize(s)
		if got != expected {
			t.Errorf("Expected '%s' got '%s'", expected, got)
		}
	}
}
