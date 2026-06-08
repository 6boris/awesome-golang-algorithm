package Solution

func Solution(n, k int) int {
	i := 1
	for ; i <= n; i++ {
		if n%i == 0 {
			k--
			if k == 0 {
				return i
			}
		}
	}
	return -1
}
