package Solution

func Solution(s string) int {
	stack := make([]byte, len(s))
	i := -1
	for _, b := range []byte(s) {
		if i == -1 || !(b == ')' && stack[i] == '(') {
			i++
			stack[i] = b
			continue
		}
		i--
	}
	return i + 1
}
