package inflect

import (
	"math"
	"strings"
)

const (
	unsupportedValueError = "Unsupported value to be written out in words."
	andSeparator          = "and"
	point                 = "point"
	minus                 = "Minus"
	zero                  = "zero"
	hundred               = "hundred"

	comma = ","
)

var (
	units = []string{"", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

	teens = map[int]string{
		10: "ten",
		11: "eleven",
		12: "twelve",
		13: "thirteen",
		14: "fourteen",
		15: "fifteen",
		16: "sixteen",
		17: "seventeen",
		18: "eighteen",
		19: "nineteen"}

	tens = map[int]string{
		10: "ten",
		20: "twenty",
		30: "thirty",
		40: "forty",
		50: "fifty",
		60: "sixty",
		70: "seventy",
		80: "eighty",
		90: "ninety"}

	thousands = []string{"", "thousand", "million", "billion", "trillion",
		"quadrillion", "quintillion", "sextillion", "septillion", "octillion", "nonillion",
		"decillion", "undecillion", "duodecillion", "tredecillion", "quattuordecillion",
		"sexdecillion", "septendecillion", "octodecillion", "novemdecillion", "vigintillion"}
)

// IntoWords convert numbers (float64) to words
func IntoWords(number float64) (string, error) {
	if number == 0 {
		return zero, nil
	}

	words := []string{}
	// Minus
	if number < 0 {
		words = append(words, minus)
	}
	number = math.Abs(number)

	integer, fractional := math.Modf(number)
	if integer != 0 || fractional == 0 {
		numberIntoWords := writeOutNumbersInWords(integer)
		words = append(words, numberIntoWords...)
	}
	if fractional > 0 {
		words = append(words, point)
		fractional := round(math.Abs(fractional) * 100)
		fracIntoWords := writeOutNumbersInWords(fractional)
		words = append(words, fracIntoWords...)
	}

	return sanitize(words), nil
}

func writeOutNumbersInWords(n float64) []string {
	switch {
	case n < 10:
		return getUnderTen(n)
	case n < 20:
		return getUnderTwenty(n)
	case n < 100:
		return getUnderHundred(n)
	case n < 1000:
		return getUnderThousand(n)
	default:
		return getUpThousand(n)
	}
}

func getUnderTen(n float64) []string {
	return []string{units[int(n)]}
}

func getUnderTwenty(n float64) []string {
	return []string{teens[int(n)]}
}

func getUnderHundred(f float64) []string {
	i := int(f/10) * 10
	value := tens[i]

	mod := math.Mod(f, 10)
	if mod != 0 {
		remaining := writeOutNumbersInWords(mod)
		value += "-" + remaining[0]
	}

	return []string{value}
}

func getUnderThousand(f float64) []string {
	unit := units[int(f/100)]
	words := []string{unit, hundred}

	mod := math.Mod(f, 100)
	if mod != 0 {
		remaining := writeOutNumbersInWords(mod)
		words = append(words, andSeparator)
		words = append(words, remaining...)
	}
	return words
}

func getUpThousand(n float64) []string {
	words := []string{}
	for i := 0; n >= 1; i++ {
		var r float64
		n, r = math.Modf(n / 1000)

		w := writeOutNumbersInWords(r * 1000)
		if (len(w) == 0) || (len(w) == 1 && w[0] == "") {
			continue
		}

		if i > 0 {
			w = append(w, thousands[i])
			w = append(w, comma)
		}
		words = append(w, words...)
	}
	if words[len(words)-1] == comma {
		words = words[:len(words)-1]
	}
	return words
}

func sanitize(words []string) string {
	s := strings.Join(words, " ")
	s = strings.Replace(s, " , ", ", ", -1)
	s = strings.Trim(s, " ")
	return s
}

func round(val float64) float64 {
	const (
		roundOn = .5
		places  = 2
	)

	var round float64
	pow := math.Pow(10, float64(places))
	digit := pow * val
	_, div := math.Modf(digit)
	if div >= roundOn {
		round = math.Ceil(digit)
	} else {
		round = math.Floor(digit)
	}
	return round / pow
}
