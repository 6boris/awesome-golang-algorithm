package Solution

func Solution(adjacentPairs [][]int) []int {
	n := len(adjacentPairs) + 1
	adj := make(map[int][]int)
	for _, item := range adjacentPairs {
		a, b := item[0], item[1]
		if _, ok := adj[a]; !ok {
			adj[a] = make([]int, 0)
		}
		adj[a] = append(adj[a], b)
		if _, ok := adj[b]; !ok {
			adj[b] = make([]int, 0)
		}
		adj[b] = append(adj[b], a)
	}

	left := 0
	index := 0
	for v, next := range adj {
		if len(next) == 1 {
			left = v
			break
		}
	}

	ans := make([]int, n)
	ans[0] = left
	visited := make(map[int]struct{})
	visited[left] = struct{}{}
	start := left
	index = 1
	for index < n {
		for _, next := range adj[start] {
			if _, ok := visited[next]; ok {
				continue
			}
			start = next
			visited[next] = struct{}{}
			ans[index] = next
			index++
		}
	}
	return ans
}
