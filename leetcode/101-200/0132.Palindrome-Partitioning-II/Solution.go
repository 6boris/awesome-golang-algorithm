package Solution

func Solution(s string) int {
	cache := map[string]int{}
	var (
		dfs          func(string) int
		isPalindrome func(string) bool
	)
	isPalindrome = func(s string) bool {
		l, r := 0, len(s)-1
		for ; l < r; l, r = l+1, r-1 {
			if s[l] != s[r] {
				return false
			}
		}
		return true
	}
	dfs = func(cur string) int {
		l := len(cur)
		if l == 0 || l == 1 {
			return 0
		}
		if v, ok := cache[cur]; ok {
			return v
		}
		if isPalindrome(cur) {
			cache[cur] = 0
			return 0
		}
		m := -1
		for end := 1; end < len(cur); end++ {
			if isPalindrome(cur[:end]) {
				r := dfs(cur[end:]) + 1
				if m == -1 || m > r {
					m = r
				}
			}
		}
		cache[cur] = m
		return m
	}
	return dfs(s)
}
