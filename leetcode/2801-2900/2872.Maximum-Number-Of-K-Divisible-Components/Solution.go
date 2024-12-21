package Solution

func Solution(n int, edges [][]int, values []int, k int) int {
	for i := range values {
		values[i] %= k
	}
	adj := make([][]int, n)
	for _, e := range edges {
		a, b := e[0], e[1]
		adj[a] = append(adj[a], b)
		adj[b] = append(adj[b], a)
	}

	count := 0
	var dfs func(int, int) int
	dfs = func(parent, cur int) int {
		s := 0
		for _, nei := range adj[cur] {
			if nei == parent {
				continue
			}
			s = (s + dfs(cur, nei)) % k
		}
		s = (s + values[cur]) % k
		if s == 0 {
			count++
		}
		return s
	}
	dfs(-1, 0)
	return count
}
