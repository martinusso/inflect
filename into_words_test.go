package inflect

import "testing"

func TestUnits(t *testing.T) {
	for value, expected := range units {
		got, err := IntoWords(float64(value))
		if err != nil {
			t.Errorf("It wasn't expected any error, got '%s'", err.Error())
		}
		if got != expected {
			t.Errorf("Expected '%s' got '%s'", expected, got)
		}
	}
}

func TestTeens(t *testing.T) {
	for value, expected := range teens {
		got, err := IntoWords(float64(value))
		if err != nil {
			t.Errorf("It wasn't expected any error, got '%s'", err.Error())
		}
		if got != expected {
			t.Errorf("Expected '%s' got '%s'", expected, got)
		}
	}
}

func TestTens(t *testing.T) {
	for value, expected := range tens {
		got, err := IntoWords(float64(value))
		if err != nil {
			t.Errorf("It wasn't expected any error, got '%s'", err.Error())
		}
		if got != expected {
			t.Errorf("Expected '%s' got '%s'", expected, got)
		}
	}
}

func TestHundreds(t *testing.T) {
	hundreds := []string{"",
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
		got, err := IntoWords(float64(value * 100))
		if err != nil {
			t.Errorf("It wasn't expected any error, got '%s'", err.Error())
		}
		if got != expected {
			t.Errorf("Expected '%s' got '%s'", expected, got)
		}
	}
}

func TestUpThousand(t *testing.T) {
	upThousands := map[float64]string{
		1984:      "one thousand, nine hundred and eighty-four",
		510072000: "five hundred and ten million, seventy-two thousand",
	}

	for key, value := range upThousands {
		got, err := IntoWords(key)
		if err != nil {
			t.Errorf("It wasn't expected any error, got '%s'", err.Error())
		}
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
		got, err := IntoWords(key)
		if err != nil {
			t.Errorf("It wasn't expected any error, got '%s'", err.Error())
		}
		if got != value {
			t.Errorf("Expected '%s' got '%s'", value, got)
		}
	}
}

func TestNegativeValue(t *testing.T) {
	expected := "Minus one hundred and forty-seven"
	got, err := IntoWords(-147)

	if err != nil {
		t.Errorf("It wasn't expected any error, got '%s'", err.Error())
	}
	if got != expected {
		t.Errorf("Expected '%s' got '%s'", expected, got)
	}
}
