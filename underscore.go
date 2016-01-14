package inflect

import "strings"

// Underscore makes an underscored/lowercase string
func Underscore(s string) string {
	var safe string
	for i, r := range s {
		s := string(r)
		if s == strings.ToUpper(s) && i > 0 {
			safe += "_"
		}
		safe += s
	}
	return ParameterizeSep(safe, "_")
}
