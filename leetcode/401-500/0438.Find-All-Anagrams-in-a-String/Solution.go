package Solution

func Solution(s string, p string) []int {
	ls, lp := len(s), len(p)
	ans := make([]int, 0)
	pattern := make([]int, 26)
	for _, b := range p {
		pattern[b-'a']++
	}
	var equal func([]int, []int) bool
	equal = func(a, b []int) bool {
		for i := 0; i < 26; i++ {
			if a[i] != b[i] {
				return false
			}
		}
		return true
	}

	try := make([]int, 26)
	for i := 0; i < lp; i++ {
		try[s[i]-'a']++
	}
	if equal(try, pattern) {
		ans = append(ans, 0)
	}
	for next := lp; next < ls; next++ {
		del := next - lp
		try[s[del]-'a']--
		try[s[next]-'a']++
		if equal(try, pattern) {
			ans = append(ans, del+1)
		}
	}
	return ans
}
