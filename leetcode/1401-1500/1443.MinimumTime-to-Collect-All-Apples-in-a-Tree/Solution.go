package Solution

func Solution(n int, edges [][]int, hasApple []bool) int {
	graph := make(map[int][]int)

	for _, e := range edges {
		if _, ok := graph[e[0]]; !ok {
			graph[e[0]] = make([]int, 0)
		}
		if _, ok := graph[e[1]]; !ok {
			graph[e[1]] = make([]int, 0)
		}
		graph[e[0]] = append(graph[e[0]], e[1])
		graph[e[1]] = append(graph[e[1]], e[0])
	}

	var dfs func(int, int) int
	/*
		dfs = func(start, path int) {
			rels, ok := graph[start]
			if !ok {
				return
			}
			apples := 0
			for _, rel := range rels {
				if visited[rel] {
					continue
				}
				next := path + 1
				visited[rel] = true
				if hasApple[rel] {
					apples++
					if apples == 1 {
						totalEdges += path
					} else {
						totalEdges += 1
					}
					next = 1
				}
				dfs(rel, next)
				visited[rel] = false
			}
		}
	*/
	dfs = func(start, partent int) int {
		rels, ok := graph[start]
		if !ok {
			return 0
		}
		total := 0
		for _, rel := range rels {
			if rel == partent {
				continue
			}
			cost := dfs(rel, start)
			if cost > 0 || hasApple[rel] {
				total += cost + 2
			}
		}
		return total
	}
	return dfs(0, -1)
}
