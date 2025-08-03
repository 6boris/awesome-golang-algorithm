package Solution

func Solution(n int) int {
	if n&1 == 1 {
		return 2 * n
	}
	return n
}
