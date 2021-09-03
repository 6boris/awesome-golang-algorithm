package Solution

func isValid_1(s string) bool {
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

func isValid_2(s string) bool {
	n := len(s)
	if n%2 == 1 {
		return false
	}
	pairs := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}
	stack := []byte{}
	for i := 0; i < n; i++ {
		if pairs[s[i]] > 0 {
			if len(stack) == 0 || stack[len(stack)-1] != pairs[s[i]] {
				return false
			}
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, s[i])
		}
	}
	return len(stack) == 0
}
