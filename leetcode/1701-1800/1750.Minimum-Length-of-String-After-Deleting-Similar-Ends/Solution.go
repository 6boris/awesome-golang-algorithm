package Solution

func Solution(s string) int {
	ans := len(s)
	l, r := 0, len(s)-1
	lc, rc := s[0], s[r]
	if lc != rc {
		return ans
	}
	for l < r && lc == rc {
		for ; l < r && s[l] == lc; l++ {
		}
		for ; r >= l && s[r] == rc; r-- {
		}
		ans = min(ans, r-l+1)
		lc, rc = s[l], s[r]
	}
	return ans
}
