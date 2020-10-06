package Solution

func reverseLeftWords(s string, n int) string {
	ans := ""
	for i := n; i < len(s); i++ {
		ans += string(s[i])
	}
	for i := 0; i < n; i++ {
		ans += string(s[i])
	}
	return ans
}

func reverseLeftWords2(s string, n int) string {
	for i := 0; i < n; i++ {
		s += string(s[i])
	}
	return s[n:]
}
