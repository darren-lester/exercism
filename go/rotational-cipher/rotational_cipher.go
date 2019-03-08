package rotationalcipher

import "unicode"

// RotationalCipher encodes an input string with a Caesar cipher with a given shift size
func RotationalCipher(input string, shiftKey int) string {
	length := len(input)
	output := make([]byte, length)
	offset := int('a')

	for i, r := range input {
		if !unicode.IsLetter(r) {
			output[i] = input[i]
			continue
		}

		lowered := unicode.ToLower(r)
		rotated := rune((int(lowered)-offset+shiftKey)%26 + offset)

		if unicode.IsUpper(r) {
			rotated = unicode.ToUpper(rotated)
		}
		output[i] = byte(rotated)
	}

	return string(output)
}
