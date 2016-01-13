package inflect

import "fmt"

// Ordinalize turns a number into an ordinal string
//
//		Ordinalize(1)     => "1st"
//		Ordinalize(2)     => "2nd"
//		Ordinalize(1002)  => "1002nd"
//		Ordinalize(1003)  => "1003rd"
//		Ordinalize(-11)   => "-11th"
//		Ordinalize(-1021) => "-1021st"
func Ordinalize(number int) string {
	ordinal := Ordinal(number)
	return fmt.Sprintf("%d%s", number, ordinal)
}
