package Solution

func longestPalindrome(s string) int {
	m, ans, has_odd := make(map[rune]int, 128), 0, false
	for _, v := range s {
		m[v]++
	}
	for _, v := range m {
		ans += v
		if v%2 == 1 {
			ans--
			has_odd = true
		}
	}
	if has_odd {
		ans++
	}

	return ans
}
