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

func TestOrdinalize(t *testing.T) {
	ordinals := map[int]string{
		-1: "-1st", 0: "0th",
		1: "1st", 2: "2nd", 3: "3rd", 4: "4th", 5: "5th", 6: "6th", 7: "7th", 8: "8th", 9: "9th",
		10: "10th", 11: "11th", 12: "12th", 13: "13th", 14: "14th", 15: "15th", 16: "16th", 17: "17th", 18: "18th", 19: "19th",
		20: "20th", 21: "21st", 22: "22nd", 23: "23rd", 24: "24th", 25: "25th", 26: "26th", 27: "27th", 28: "28th", 29: "29th",
		30: "30th", 31: "31st", 32: "32nd", 33: "33rd", 34: "34th", 35: "35th", 36: "36th", 37: "37th", 38: "38th", 39: "39th",
		100: "100th", 101: "101st", 102: "102nd", 103: "103rd", 104: "104th",
		200: "200th", 201: "201st", 202: "202nd", 203: "203rd", 204: "204th",
		1000: "1000th", 1001: "1001st", 1002: "1002nd", 1003: "1003rd", 1004: "1004th",
		10000: "10000th", 10001: "10001st", 10002: "10002nd", 10003: "10003rd", 10004: "10004th",
		100000: "100000th", 100001: "100001st", 100002: "100002nd", 100003: "100003rd", 100004: "100004th",
	}
	for k, expected := range ordinals {
		got := Ordinalize(k)
		if got != expected {
			t.Errorf("Expected '%s' got '%s'", expected, got)
		}
	}
}
