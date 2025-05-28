package Solution

import "slices"

func dfs3372(adj map[int][]int, start, parent, limit int) int {
	if limit < 0 {
		return 0
	}

	cnt := 1 // self
	for _, rel := range adj[start] {
		if rel == parent {
			continue
		}
		cnt += dfs3372(adj, rel, start, limit-1)
	}
	return cnt

}

func calDistance(adj map[int][]int, n, limit int) []int {
	distance := make([]int, n)
	for i := 0; i < n; i++ {
		distance[i] = dfs3372(adj, i, -1, limit)
	}
	return distance
}

func Solution(edges1 [][]int, edges2 [][]int, k int) []int {
	n, m := 0, 0
	for _, e := range edges1 {
		n = max(n, e[0], e[1])
	}
	for _, e := range edges2 {
		m = max(m, e[0], e[1])
	}
	n, m = n+1, m+1
	adj1 := make(map[int][]int)
	adj2 := make(map[int][]int)

	for _, e := range edges1 {
		adj1[e[0]] = append(adj1[e[0]], e[1])
		adj1[e[1]] = append(adj1[e[1]], e[0])
	}
	for _, e := range edges2 {
		adj2[e[0]] = append(adj2[e[0]], e[1])
		adj2[e[1]] = append(adj2[e[1]], e[0])
	}

	dis1 := calDistance(adj1, n, k)
	dis2 := calDistance(adj2, m, k-1)

	cnt := slices.Max(dis2)
	for i := range n {
		dis1[i] += cnt
	}

	return dis1
}
