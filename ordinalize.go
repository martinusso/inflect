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
	switch number % 100 {
	case 11, 12, 13:
		return fmt.Sprintf("%dth", number)
	default:
		switch number % 10 {
		case 1:
			return fmt.Sprintf("%dst", number)
		case 2:
			return fmt.Sprintf("%dnd", number)
		case 3:
			return fmt.Sprintf("%drd", number)
		}
	}
	return fmt.Sprintf("%dth", number)
}
