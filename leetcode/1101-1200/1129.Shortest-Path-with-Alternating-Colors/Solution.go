package Solution

type pathColor struct {
	to, color int
}

type queueItem struct {
	node, steps, color int
}

func Solution(n int, redEdges [][]int, blueEdges [][]int) []int {
	ans := make([]int, n)
	for i := 1; i < n; i++ {
		ans[i] = -1
	}

	// bfs
	edges := make(map[int][]pathColor)
	for _, e := range redEdges {
		if _, ok := edges[e[0]]; !ok {
			edges[e[0]] = make([]pathColor, 0)
		}
		edges[e[0]] = append(edges[e[0]], pathColor{e[1], 0})
	}
	for _, e := range blueEdges {
		if _, ok := edges[e[0]]; !ok {
			edges[e[0]] = make([]pathColor, 0)
		}
		edges[e[0]] = append(edges[e[0]], pathColor{e[1], 1})
	}
	visited := make([][]bool, n)
	for i := 0; i < n; i++ {
		visited[i] = make([]bool, 2)
	}

	queue := []queueItem{{0, 0, -1}}
	visited[0][0], visited[0][1] = true, true
	for len(queue) > 0 {
		nextQ := make([]queueItem, 0)
		for _, item := range queue {
			for _, neigh := range edges[item.node] {
				if !visited[neigh.to][neigh.color] && neigh.color != item.color {
					visited[neigh.to][neigh.color] = true
					nextQ = append(nextQ, queueItem{neigh.to, item.steps + 1, neigh.color})
					if ans[neigh.to] == -1 {
						ans[neigh.to] = item.steps + 1
					}
				}
			}
		}
		queue = nextQ
	}
	return ans
}
