package Solution

func Solution(n int) []int {
	l := 2*n - 1
	res := make([]int, l)
	used := make([]bool, n+1)
	var dfs func(int) bool
	dfs = func(start int) bool {
		if start == l {
			return true
		}
		if res[start] != 0 {
			return dfs(start + 1)
		}

		for s := n; s >= 1; s-- {
			if used[s] {
				continue
			}
			res[start] = s
			used[s] = true
			if s == 1 {
				if dfs(start + 1) {
					return true
				}
			} else {
				if end := start + s; end < l && res[end] == 0 {
					res[end] = s
					if dfs(start + 1) {
						return true
					}
					res[end] = 0
				}
			}
			res[start] = 0
			used[s] = false
		}
		return false
	}

	dfs(0)
	return res
}
