package Solution

func Solution(num string, k int) string {
	n := len(num)
	if n == k {
		return "0"
	}
	stack := make([]byte, 0)
	for idx := 0; idx < n; idx++ {
		for k > 0 && len(stack) > 0 && stack[len(stack)-1] > num[idx] {
			stack = stack[:len(stack)-1]
			k--
		}
		stack = append(stack, num[idx])
	}
	stack = stack[:len(stack)-k]
	i := 0
	for ; i < len(stack) && stack[i] == '0'; i++ {
	}
	if i == len(stack) {
		return "0"
	}
	return string(stack[i:])
}
