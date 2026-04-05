func isValid(s string) bool {
    parenthesisMap := map[rune]rune{
		'(': ')',
		'{' : '}',
		'[' : ']',
	}

	var stack []rune
	for _, c := range s {
		if _, isOpening := parenthesisMap[c]; isOpening {
			stack = append(stack, c)
		} else {
			if len(stack) > 0 && parenthesisMap[stack[len(stack)-1]] == c {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		}
	}
	return len(stack) == 0
}
