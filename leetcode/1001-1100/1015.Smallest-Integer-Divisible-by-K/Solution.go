package Solution

func Solution(k int) int {
	base := 0
	// 1, 11, 111,
	for i := 1; i <= k; i++ {
		base = (base*10 + 1) % k
		if base == 0 {
			return i
		}
	}
	return -1
}
