package Solution

func Solution(n int, left []int, right []int) int {
	ans := 0
	for _, l := range left {
		if l > ans {
			ans = l
		}
	}
	for _, r := range right {
		if l := n - r; l > ans {
			ans = l
		}
	}
	return ans
}
