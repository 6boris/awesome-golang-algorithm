package Solution

func Solution(gain []int) int {
	start := 0
	ans := 0
	for _, h := range gain {
		start += h
		if start > ans {
			ans = start
		}
	}
	return ans
}
