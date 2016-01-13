package inflect

import "math"

const (
	st = "st"
	nd = "nd"
	rd = "rd"
	th = "th"
)

/*
Ordinal returns the suffix that should be added to a number to denote the position in an ordered sequence such as 1st, 2nd, 3rd, 4th...

	Ordinal(1)     => "st"
	Ordinal(2)     => "nd"
	Ordinal(1002)  => "nd"
	Ordinal(1003)  => "rd"
	Ordinal(-11)   => "th"
	Ordinal(-1021) => "st"
*/
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
