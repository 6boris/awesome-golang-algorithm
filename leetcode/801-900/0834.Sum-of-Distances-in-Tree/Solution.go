package Solution

// 借鉴官方算法
func Solution(n int, edges [][]int) []int {
	adj := make(map[int]map[int]struct{})
	for u := 0; u < n; u++ {
		adj[u] = make(map[int]struct{})
	}
	for _, e := range edges {
		adj[e[0]][e[1]] = struct{}{}
		adj[e[1]][e[0]] = struct{}{}
	}

	ans := make([]int, n)
	count := make([]int, n)
	for i := 0; i < n; i++ {
		count[i] = 1
	}
	var dfs func(int, int)
	var dfs1 func(int, int)
	dfs = func(node, parent int) {
		for child := range adj[node] {
			if child != parent {
				dfs(child, node)
				count[node] += count[child]
				ans[node] += ans[child] + count[child]
			}
		}
	}
	dfs1 = func(node, parent int) {
		for child := range adj[node] {
			if child != parent {
				ans[child] = ans[node] - count[child] + n - count[child]
				dfs1(child, node)
			}
		}
	}
	dfs(0, -1)
	dfs1(0, -1)
	return ans
}
