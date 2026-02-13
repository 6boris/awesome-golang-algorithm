package Solution

func Solution(complexity []int) int {
	n := len(complexity)
	for i := 1; i < n; i++ {
		if complexity[i] <= complexity[0] {
			return 0
		}
	}

	ans := 1
	mod := 1000000007
	for i := 2; i < n; i++ {
		ans = ans * i % mod
	}
	return ans
}
