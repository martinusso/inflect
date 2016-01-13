package inflect

import "fmt"

// Ordinalize turns a number into an ordinal string
//
// Ex: `Ordinalize(1)     => "1st"`
//
// ```Ordinalize(1)     => "1st"```
//
// ```
// Ordinalize(1)     => "1st"
// ```
func Ordinalize(number int) string {
	ordinal := Ordinal(number)
	return fmt.Sprintf("%d%s", number, ordinal)
}
