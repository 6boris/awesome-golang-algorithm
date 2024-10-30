package Solution

func Solution(s string) int {
	used := make(map[string]struct{})
	var dfs func(int) int
	dfs = func(index int) int {
		if index >= len(s) {
			return 0
		}
		m := 0
		for next := index + 1; next <= len(s); next++ {
			cur := s[index:next]
			if _, ok := used[cur]; ok {
				continue
			}
			used[cur] = struct{}{}
			m = max(m, dfs(next)+1)
			delete(used, cur)
		}
		return m
	}
	return dfs(0)
}
