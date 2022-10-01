package Solution

func Solution(n int, k int) int {

	ans := 0
	for base := n; base != 0; base /= k {
		ans += base % k
	}
	return ans
}
