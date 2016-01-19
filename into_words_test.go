package inflect

import "testing"

func TestZero(t *testing.T) {
	got := IntoWords(0)
	if got != zero {
		t.Errorf("Expected '%s' got '%s'", zero, got)
	}
}

func TestUnits(t *testing.T) {
	for i, expected := range units {
		if i == 0 {
			continue
		}
		got := IntoWords(float64(i))
		if got != expected {
			t.Errorf("Expected '%s' got '%s'", expected, got)
		}
	}
}

func TestTeens(t *testing.T) {
	for value, expected := range teens {
		got := IntoWords(float64(value))
		if got != expected {
			t.Errorf("Expected '%s' got '%s'", expected, got)
		}
	}
}

func TestTens(t *testing.T) {
	for value, expected := range tens {
		got := IntoWords(float64(value))
		if got != expected {
			t.Errorf("Expected '%s' got '%s'", expected, got)
		}
	}
}

func TestHundreds(t *testing.T) {
	hundreds := []string{zero,
		"one hundred",
		"two hundred",
		"three hundred",
		"four hundred",
		"five hundred",
		"six hundred",
		"seven hundred",
		"eight hundred",
		"nine hundred"}

	for value, expected := range hundreds {
		got := IntoWords(float64(value * 100))
		if got != expected {
			t.Errorf("Expected '%s' got '%s'", expected, got)
		}
	}
}

func TestUpThousand(t *testing.T) {
	upThousands := map[float64]string{
		1000:      "one thousand",
		1984:      "one thousand, nine hundred and eighty-four",
		510072000: "five hundred and ten million, seventy-two thousand",
		775398007: "seven hundred and seventy-five million, three hundred and ninety-eight thousand, seven",
	}

	for key, value := range upThousands {
		got := IntoWords(key)
		if got != value {
			t.Errorf("Expected '%s' got '%s'", value, got)
		}
	}
}

func TestFraction(t *testing.T) {
	upThousands := map[float64]string{
		1.99:    "one point ninety-nine",
		12756.2: "twelve thousand, seven hundred and fifty-six point twenty",
		0.42:    "point forty-two",
	}

	for key, value := range upThousands {
		got := IntoWords(key)
		if got != value {
			t.Errorf("Expected '%s' got '%s'", value, got)
		}
	}
}

func TestNegativeValue(t *testing.T) {
	expected := "Minus one hundred and forty-seven"
	got := IntoWords(-147)
	if got != expected {
		t.Errorf("Expected '%s' got '%s'", expected, got)
	}
}
