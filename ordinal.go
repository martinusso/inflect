package inflect

import (
	"fmt"
	"math"
)

const (
	st = "st"
	nd = "nd"
	rd = "rd"
	th = "th"
)

// Ordinal returns the ordinal suffix that should be added to a number to denote the position in an ordered sequence such as 1st, 2nd, 3rd, 4th...
func Ordinal(number int) string {
	switch abs(number) % 100 {
	case 11, 12, 13:
		return th
	default:
		switch abs(number) % 10 {
		case 1:
			return st
		case 2:
			return nd
		case 3:
			return rd
		}
	}
	return th
}

func abs(number int) int {
	return int(math.Abs(float64(number)))
}

// Ordinalize turns a number into an ordinal string
func Ordinalize(number int) string {
	ordinal := Ordinal(number)
	return fmt.Sprintf("%d%s", number, ordinal)
}
