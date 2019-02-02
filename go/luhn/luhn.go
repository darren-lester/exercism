package luhn

import (
	"regexp"
	"strconv"
	"strings"
)

func Valid(input string) bool {
	if !isValidFormat(input) {
		return false
	}

	sanitisedInput := sanitiseInput(input)

	length := len(sanitisedInput)
	if length < 2 {
		return false
	}

	luhnSum := calculateLuhnSum(sanitisedInput)

	return luhnSum%10 == 0
}

func isValidFormat(input string) bool {
	re := regexp.MustCompile(`^[ \d]*$`)
	return re.Match([]byte(input))
}

func sanitiseInput(input string) string {
	return strings.Replace(input, " ", "", -1)
}

func calculateLuhnSum(input string) int {
	luhnSum := 0
	for i := len(input) - 1; i >= 1; i -= 2 {
		v1 := toInt(input[i])
		v2 := toInt(input[i-1])
		v2 *= 2
		if v2 > 9 {
			v2 -= 9
		}
		luhnSum += v1 + v2
	}

	return luhnSum
}

func toInt(b byte) int {
	i, _ := strconv.Atoi(string(b))
	return i
}
