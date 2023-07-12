package Solution

func Solution(graph [][]int) []int {
	ans := make([]int, 0)
	cache := make(map[int]struct{})
	var dfs func(int, []bool) bool
	// 根据终端节点反推？
	dfs = func(node int, visited []bool) bool {
		if _, ok := cache[node]; ok {
			// 已经确定某个节点可以到
			return true
		}
		if visited[node] {
			// 遇到访问过的，那肯定是出现环了
			return false
		}

		visited[node] = true
		for _, next := range graph[node] {
			if !dfs(next, visited) {
				return false
			}
		}
		cache[node] = struct{}{}
		return true
	}
	l := len(graph)
	visited := make([]bool, l)
	for i := 0; i < l; i++ {
		if len(graph[i]) == 0 {
			cache[i] = struct{}{}
		}
	}
	for i := 0; i < l; i++ {
		if _, ok := cache[i]; ok {
			ans = append(ans, i)
			continue
		}
		for i := 0; i < l; i++ {
			visited[i] = false
		}
		if dfs(i, visited) {
			ans = append(ans, i)
		}
	}
	return ans
}
