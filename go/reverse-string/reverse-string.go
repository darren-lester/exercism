package reverse

import "unicode/utf8"

func String(s string) string {
	if s == "" {
		return ""
	}

	reversed := make([]byte, 0)
	bytes := []byte(s)
	length := len(bytes)

	for length > 0 {
		lastRune, runeSize := utf8.DecodeLastRune(bytes)
		runeBuffer := make([]byte, runeSize)
		utf8.EncodeRune(runeBuffer, lastRune)
		reversed = append(reversed, runeBuffer...)
		bytes = bytes[:length-runeSize]
		length = len(bytes)
	}

	return string(reversed)
}
