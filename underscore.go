package inflect

import "unicode"

// Underscore makes an underscored/lowercase string
func Underscore(s string) string {
	sep := "_"
	var safe string
	for i, r := range s {
		if unicode.IsUpper(r) && i > 0 {
			safe += sep
		}
		safe += string(r)
	}
	return ParameterizeSep(safe, sep)
}
