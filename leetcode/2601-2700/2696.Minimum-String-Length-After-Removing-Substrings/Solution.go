package Solution

func Solution(s string) int {
	stack := make([]byte, len(s))
	index := -1
	for _, b := range s {
		if index != -1 && ((b == 'B' && stack[index] == 'A') || (b == 'D' && stack[index] == 'C')) {
			index--
			continue
		}
		index++
		stack[index] = byte(b)
	}
	return index + 1
}
