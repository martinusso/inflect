package inflect

import "testing"

func TestOrdinal(t *testing.T) {
	ordinals := map[int]string{
		-1: "st", 0: "th",
		1: "st", 2: "nd", 3: "rd", 4: "th", 5: "th", 6: "th", 7: "th", 8: "th", 9: "th",
		10: "th", 11: "th", 12: "th", 13: "th", 14: "th", 15: "th", 16: "th", 17: "th", 18: "th", 19: "th",
		20: "th", 21: "st", 22: "nd", 23: "rd", 24: "th", 25: "th", 26: "th", 27: "th", 28: "th", 29: "th",
		30: "th", 31: "st", 32: "nd", 33: "rd", 34: "th", 35: "th", 36: "th", 37: "th", 38: "th", 39: "th",
		100: "th", 101: "st", 102: "nd", 103: "rd", 104: "th",
		200: "th", 201: "st", 202: "nd", 203: "rd", 204: "th",
		1000: "th", 1001: "st", 1002: "nd", 1003: "rd", 1004: "th",
		10000: "th", 10001: "st", 10002: "nd", 10003: "rd", 10004: "th",
		100000: "th", 100001: "st", 100002: "nd", 100003: "rd", 100004: "th",
	}
	for k, expected := range ordinals {
		got := Ordinal(k)
		if got != expected {
			t.Errorf("Expected '%s' got '%s'", expected, got)
		}
	}
}
