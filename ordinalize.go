package inflect

import (
	"errors"
	"fmt"
	"strconv"
)

// OrdinalizeStr turns a number (as string) into an ordinal string
func OrdinalizeStr(number string) string {
	n, err := strconv.Atoi(number)
	if err != nil {
		errors.New("")
	}
	return Ordinalize(n)
}

// Ordinalize turns a number into an ordinal string
func Ordinalize(number int) string {
	ordinal := Ordinal(number)
	return fmt.Sprintf("%d%s", number, ordinal)
}
