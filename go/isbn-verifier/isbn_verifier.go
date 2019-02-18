package isbn

import (
	"regexp"
	"strconv"
)

// IsValidISBN returns whether or not a string is a valid ISBN
func IsValidISBN(input string) bool {
	re := regexp.MustCompile("[0-9X]")
	matches := re.FindAll([]byte(input), -1)

	if len(matches) != 10 {
		return false
	}

	digitRunes := matches[:9]
	digits := make([]int, len(digitRunes))
	for i, r := range digitRunes {
		d, _ := strconv.Atoi(string(r))
		digits[i] = d
	}

	checkRune := matches[9][0]
	var check int
	if checkRune == 'X' {
		check = 10
	} else {
		d, _ := strconv.Atoi(string(checkRune))
		check = d
	}

	return (digits[0]*10+digits[1]*9+digits[2]*8+digits[3]*7+digits[4]*6+digits[5]*5+digits[6]*4+digits[7]*3+digits[8]*2+check)%11 == 0
}
