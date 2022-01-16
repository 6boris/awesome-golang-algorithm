package Solution

func Solution(s string) [][]string {
	res := make([][]string, 0)
	buildPartition(s, 0, []string{}, &res)
	return res
}

func buildPartition(s string, start int, path []string, res *[][]string) {
	if start >= len(s) {
		dst := make([]string, len(path))
		copy(dst, path)
		*res = append(*res, dst)
		return
	}

	for end := start + 1; end <= len(s); end++ {
		str := s[start:end]
		if isPalindrome(str) {
			buildPartition(s, end, append(path, str), res)
		}
	}
}

func isPalindrome(s string) bool {
	for _s, e := 0, len(s)-1; _s < e; _s, e = _s+1, e-1 {
		if s[_s] != s[e] {
			return false
		}
	}
	return true
}
