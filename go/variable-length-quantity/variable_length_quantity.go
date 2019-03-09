package variablelengthquantity

import "errors"

func DecodeVarint(input []byte) ([]uint32, error) {
	decoded := []uint32{}
	var nextVal uint32 = 0
	cont := true

	for _, b := range input {
		nextVal = nextVal<<7 + uint32(b&0x7F)

		if b>>7 == 1 {
			cont = true
		} else {
			decoded = append(decoded, nextVal)
			nextVal = 0
			cont = false
		}
	}

	if cont {
		return decoded, errors.New("Incomplete input")
	}

	return decoded, nil
}

func EncodeVarint(input []uint32) []byte {
	encoded := []byte{}

	for _, val := range input {
		chunk := val & 0x7F
		encodedVal := []byte{byte(0x00 | chunk)}

		for val = val >> 7; val > 0; val = val >> 7 {
			chunk := val & 0x7F
			encodedVal = append([]byte{byte(0x80 | chunk)}, encodedVal...)
		}

		encoded = append(encoded, encodedVal...)
	}

	return encoded
}
