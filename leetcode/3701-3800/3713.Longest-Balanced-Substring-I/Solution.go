package Solution

func Solution(s string) int {
	ret := 1
	n := len(s)
	count := make([][26]int, n)
	pre := [26]int{}
	for i := range s {
		pre[s[i]-'a']++
		count[i] = pre
	}
	check := func(end, start int) bool {
		test := count[end]
		if start != -1 {
			for i := range 26 {
				test[i] -= count[start][i]
			}
		}
		cur := 0
		for i := range 26 {
			if test[i] == 0 {
				continue
			}
			if cur == 0 {
				cur = test[i]
			}
			if cur != test[i] {
				return false
			}
		}
		return true
	}
	for end := 1; end < n; end++ {
		if check(end, -1) {
			ret = max(ret, end+1)
			continue
		}
		for i := 0; i < end; i++ {
			if check(end, i) {
				ret = max(ret, end-i)
				break
			}
		}
	}
	return ret
}
