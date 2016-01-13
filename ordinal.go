package inflect

const (
	st = "st"
	nd = "nd"
	rd = "rd"
	th = "th"
)

// Ordinal returns the suffix that should be added to a number to denote the position in an ordered sequence such as 1st, 2nd, 3rd, 4th...
func Ordinal(number int) string {
	switch number % 100 {
	case 11, 12, 13:
		return th
	default:
		switch number % 10 {
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
