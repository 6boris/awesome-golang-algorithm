package Solution

func Solution(s string) int {
	ans := 1
	cur := 1
	for i := 1; i < len(s); i++ {
		if s[i] == s[i-1]+1 {
			cur++
		} else {
			cur = 1
		}
		ans = max(ans, cur)
	}
	return ans
}
