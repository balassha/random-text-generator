package trigram

import "strings"

// Normalize makes the string lower case and removes any special characters and
// numbers.
func Normalize(s string) string {
	s = strings.ToLower(s)
	s = strings.Map(func(r rune) rune {
		if r < 'a' || r > 'z' {
			return -1
		}
		return r
	}, s)

	return s
}
