package Solution

type node787 struct {
	node int
	cost int
}

// 感觉dijkstra算法呢
func Solution(n int, flights [][]int, src int, dst int, k int) int {
	stops := 0
	dist := make([]int, n)
	for i := 0; i < n; i++ {
		dist[i] = -1
	}
	edges := make(map[int][]node787)
	for _, flight := range flights {
		from, to, cost := flight[0], flight[1], flight[2]
		if _, ok := edges[from]; !ok {
			edges[from] = make([]node787, 0)
		}
		edges[from] = append(edges[from], node787{to, cost})
	}
	queue := []node787{{src, 0}}
	for stops <= k && len(queue) > 0 {
		nextQ := make([]node787, 0)
		for _, item := range queue {
			for _, canReach := range edges[item.node] {
				nextCost := item.cost + canReach.cost
				if dist[canReach.node] == -1 || dist[canReach.node] > nextCost {
					dist[canReach.node] = nextCost
					nextQ = append(nextQ, node787{canReach.node, dist[canReach.node]})
				}

			}
		}
		queue = nextQ
		stops++
	}
	return dist[dst]

	// minCost := -1
	// //return dijkstra(edges, src, dst, n, steps)
	// visited := make([]bool, n)
	// visited[src] = true
	// var dfs func(int, int, int)
	// dfs = func(src, steps, cost int) {
	// 	if steps < 0 {
	// 		return
	// 	}
	// 	if src == dst {
	// 		if minCost == -1 || minCost > cost {
	// 			minCost = cost
	// 		}
	// 		return
	// 	}
	// 	for _, rel := range edges[src] {
	// 		if visited[rel.node] {
	// 			continue
	// 		}
	// 		visited[rel.node] = true
	// 		dfs(rel.node, steps-1, cost+rel.cost)
	// 		visited[rel.node] = false
	// 	}
	// }
	// dfs(src, steps, 0)
}
