package Solution

func Solution(s string) string {
	result := make([]byte, 0)

	for i := 0; i < len(s); i++ {
		char := s[i]

		switch {
		case char >= 'a' && char <= 'z':
			result = append(result, char)

		case char == '*':
			if len(result) > 0 {
				result = result[:len(result)-1]
			}

		case char == '#':
			result = append(result, result...)

		case char == '%':
			for left, right := 0, len(result)-1; left < right; left, right = left+1, right-1 {
				result[left], result[right] = result[right], result[left]
			}
		}
	}

	return string(result)
}
