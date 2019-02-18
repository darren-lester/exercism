package brackets

var openBrackets = map[rune]bool{
	'(': true,
	'[': true,
	'{': true,
}

var closedToOpenBracket = map[rune]rune{
	')': '(',
	']': '[',
	'}': '{',
}

// Bracket returns whether or not the parentheses in a string are balanced
func Bracket(s string) bool {
	brackets := Stack{}

	for _, r := range s {
		_, isOpenBracket := openBrackets[r]

		if isOpenBracket {
			brackets.push(r)
			continue
		}

		openBracket, isCloseBracket := closedToOpenBracket[r]

		if isCloseBracket {
			if len(brackets) == 0 {
				return false
			}

			if brackets.pop() != openBracket {
				return false
			}
		}
	}

	return len(brackets) == 0
}

// Stack is a LIFO list
type Stack []interface{}

func (s *Stack) push(item interface{}) {
	*s = append(*s, item)
}

func (s *Stack) pop() interface{} {
	lastIndex := len(*s) - 1
	top := (*s)[lastIndex]
	*s = (*s)[:lastIndex]
	return top
}
