package Solution

func isValid(s string) bool {
	stack := make([]rune, len(s))
	top := 0

	for _, c := range s {
		switch c {
		case '(':
			stack[top] = ')'
			top += 1
			break
		case '{':
			stack[top] = '}'
			top += 1
			break
		case '[':
			stack[top] = ']'
			top += 1
			break
		default:
			if top == 0 || stack[top-1] != c {
				return false
			}
			top -= 1
			break
		}
	}

	return top == 0
}
