package inflect

import "fmt"

// Ordinalize turns a number into an ordinal string
func Ordinalize(number int) string {
	ordinal := Ordinal(number)
	return fmt.Sprintf("%d%s", number, ordinal)
}
