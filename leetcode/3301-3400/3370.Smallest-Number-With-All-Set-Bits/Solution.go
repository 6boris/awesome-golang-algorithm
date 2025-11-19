package Solution

func Solution(n int) int {
	base := 2
	for i := 0; i < 10; i++ {
		if base-1 >= n {
			return base - 1
		}
		base *= 2
	}
	return 0
}
