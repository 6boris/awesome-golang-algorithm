package Solution

func Solution(s string) bool {
	stack := make([]byte, len(s))
	i := -1
	for idx := 0; idx < len(s); idx++ {
		if s[idx] == 'c' && i >= 1 && stack[i] == 'b' && stack[i-1] == 'a' {
			i -= 2
			continue
		}
		i++
		stack[i] = s[idx]
	}
	return i == -1
}
