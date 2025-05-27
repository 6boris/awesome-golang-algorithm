package Solution

const mod1931 = 1e9 + 7

var colors = []byte{'R', 'G', 'B'}

func columns(m int, cur string, preColor byte) []string {
	ans := []string{}
	if m == 0 {
		ans = append(ans, cur)
		return ans
	}

	for _, color := range colors {
		if color != preColor {
			ans = append(ans, columns(m-1, cur+string(color), color)...)
		}
	}
	return ans
}
func can(a, b string) bool {
	for i := 0; i < len(a); i++ {
		if a[i] == b[i] {
			return false
		}
	}
	return true
}

func Solution(m int, n int) int {
	groups := columns(m, "", ' ')
	cache := make(map[string]map[int]int)
	var dfs func(int, string) int
	dfs = func(left int, cur string) int {
		if left == 0 {
			return 1
		}
		if v, ok := cache[cur]; ok {
			if c, ok1 := v[left]; ok1 {
				return c
			}
		}
		ans := 0
		for i := 0; i < len(groups); i++ {
			if !can(cur, groups[i]) {
				continue
			}
			ans = (ans + dfs(left-1, groups[i])) % mod1931
		}

		v, ok := cache[cur]
		if !ok {
			v = map[int]int{}
		}
		v[left] = ans
		cache[cur] = v
		return ans
	}
	// 挑选n个
	cur := ""
	for i := 0; i < m; i++ {
		cur += " "
	}
	return dfs(n, cur)
}
