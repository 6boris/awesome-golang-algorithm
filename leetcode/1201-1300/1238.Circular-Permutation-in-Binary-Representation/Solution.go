package Solution

func g(n int) int {
	return n ^ (n >> 1)
}

func revg(x int) int {
	n := 0
	for ; x > 0; x >>= 1 {
		n ^= x
	}
	return n
}

// k 位的格雷码
func Solution(n int, start int) []int {
	alloc := 1
	for i := 0; i < n; i++ {
		alloc *= 2
	}
	ans := make([]int, alloc)
	ans[0] = start

	startIndex := revg(start)
	index := 1
	for next := (startIndex + 1) % alloc; next != startIndex; next = (next + 1) % alloc {
		ans[index] = g(next)
		index++
	}
	return ans
}
