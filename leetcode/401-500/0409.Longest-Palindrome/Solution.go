package Solution

func longestPalindrome(s string) int {
	m, ans := make(map[rune]int, 128), 0
	for _, v := range s {
		m[v]++
	}
	for _, v := range m {
		ans += v / 2 * 2
		if v%2 == 1 && ans%2 == 0 {
			ans++
		}
	}
	return ans
}
