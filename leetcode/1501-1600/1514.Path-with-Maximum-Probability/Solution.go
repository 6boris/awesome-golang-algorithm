package Solution

// 哈哈哈哈哈， challengeProgrammingDatastructure/13-3-single-source-shortest-path-1.go dijkstra模板
func Solution(n int, edges [][]int, succProb []float64, start int, end int) float64 {
	adj := make(map[int]map[int]float64)
	for idx, edge := range edges {
		f, t := edge[0], edge[1]
		if _, ok := adj[f]; !ok {
			adj[f] = make(map[int]float64)
		}
		if _, ok := adj[t]; !ok {
			adj[t] = make(map[int]float64)
		}
		adj[f][t] = succProb[idx]
		adj[t][f] = succProb[idx]
	}
	edgesProb := make([]float64, n)
	for i := 0; i < n; i++ {
		edgesProb[i] = -1
	}
	edgesProb[start] = 1
	visited := make([]bool, n)
	for i := 0; i < n-1; i++ {
		selectNode := -1
		for from, prob := range edgesProb {
			if prob == -1 {
				continue
			}
			if !visited[from] && (selectNode == -1 || edgesProb[from] > edgesProb[selectNode]) {
				selectNode = from
			}
		}

		if selectNode == -1 {
			break
		}
		visited[selectNode] = true
		for next, prob := range adj[selectNode] {
			if prob == -1 {
				continue
			}
			if edgesProb[next] == -1.0 || prob*edgesProb[selectNode] > edgesProb[next] {
				edgesProb[next] = prob * edgesProb[selectNode]
			}
		}
	}
	if edgesProb[end] == -1 {
		return 0
	}
	return edgesProb[end]
}
