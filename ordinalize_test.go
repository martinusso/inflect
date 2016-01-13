package inflect

import (
	"strconv"
	"testing"
)

var (
	ordinals = map[int]string{
		1:      "1st",
		2:      "2nd",
		3:      "3rd",
		4:      "4th",
		5:      "5th",
		6:      "6th",
		7:      "7th",
		8:      "8th",
		9:      "9th",
		10:     "10th",
		11:     "11th",
		12:     "12th",
		13:     "13th",
		14:     "14th",
		15:     "15th",
		16:     "16th",
		17:     "17th",
		18:     "18th",
		19:     "19th",
		20:     "20th",
		21:     "21st",
		22:     "22nd",
		23:     "23rd",
		24:     "24th",
		25:     "25th",
		26:     "26th",
		27:     "27th",
		28:     "28th",
		29:     "29th",
		30:     "30th",
		31:     "31st",
		32:     "32nd",
		33:     "33rd",
		34:     "34th",
		35:     "35th",
		36:     "36th",
		37:     "37th",
		38:     "38th",
		39:     "39th",
		100:    "100th",
		101:    "101st",
		102:    "102nd",
		103:    "103rd",
		104:    "104th",
		200:    "200th",
		201:    "201st",
		202:    "202nd",
		203:    "203rd",
		204:    "204th",
		1000:   "1000th",
		1001:   "1001st",
		1002:   "1002nd",
		1003:   "1003rd",
		1004:   "1004th",
		10000:  "10000th",
		10001:  "10001st",
		10002:  "10002nd",
		10003:  "10003rd",
		10004:  "10004th",
		100000: "100000th",
		100001: "100001st",
		100002: "100002nd",
		100003: "100003rd",
		100004: "100004th",
	}
)

func TestOrdinalize(t *testing.T) {
	for k, expected := range ordinals {
		got := Ordinalize(k)
		if got != expected {
			t.Errorf("Expected '%s' got '%s'", expected, got)
		}
	}
}

func TestOrdinalizeStr(t *testing.T) {
	for k, expected := range ordinals {
		got := OrdinalizeStr(strconv.Itoa(k))
		if got != expected {
			t.Errorf("Expected '%s' got '%s'", expected, got)
		}
	}
}
