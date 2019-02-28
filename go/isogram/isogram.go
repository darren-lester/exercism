package isogram

import "unicode"

// IsIsogram returns whether or not a string is an isogram
func IsIsogram(input string) bool {
	var seen = map[rune]bool{}

	for _, b := range input {
		if b == ' ' || b == '-' {
			continue
		}

		b = unicode.ToLower(b)
		if seen[b] {
			return false
		}
		seen[b] = true
	}

	return true
}
