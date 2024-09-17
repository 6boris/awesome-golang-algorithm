package Solution

func Solution(nums []int) int {
	m := 0
	l := 0
	ans := 0
	for _, n := range nums {
		if n == m {
			l++
			ans = max(ans, l)
			continue
		}
		if n > m {
			m = n
			l = 1
			ans = 1
			continue
		}
		l = 0
	}
	return ans
}
