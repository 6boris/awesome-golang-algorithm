package Solution

const mod1795 = 1000000007

func Solution(s string) int {
	ans := 0
	count := 1
	for idx := 1; idx < len(s); idx++ {
		ans = (ans + count) % mod1795
		if s[idx] == s[idx-1] {
			count++
		} else {
			count = 1
		}
	}
	ans = (ans + count) % mod1795
	return ans
}
