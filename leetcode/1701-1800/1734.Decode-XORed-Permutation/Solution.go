package Solution

func xor(n int) int {
	switch n % 4 {
	case 1:
		return 1
	case 0:
		return n
	case 2:
		return n - 1
	}
	return 0
}
func Solution(encoded []int) []int {
	l := len(encoded)
	toN := xor(l + 1)
	pick := 0
	for step := 0; step < l; step += 2 {
		pick ^= encoded[step]
	}

	lastElem := toN ^ pick
	ans := make([]int, l+1)
	ans[l] = lastElem
	for idx := l - 1; idx >= 0; idx-- {
		ans[idx] = ans[idx+1] ^ encoded[idx]
	}
	return ans
}
