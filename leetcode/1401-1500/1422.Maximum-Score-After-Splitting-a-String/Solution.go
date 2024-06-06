package Solution

func Solution(s string) int {
	l := len(s)
	sum := make([]int, l)
	sum[0] = int(s[0] - '0')

	for idx := 1; idx < l; idx++ {
		sum[idx] = int(s[idx]-'0') + sum[idx-1]
	}

	ans := 0
	for idx := 0; idx < l-1; idx++ {
		right := sum[l-1] - sum[idx]
		left := idx + 1 - sum[idx]
		if r := left + right; r > ans {
			ans = r
		}
	}
	return ans
}
