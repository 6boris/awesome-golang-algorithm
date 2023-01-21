package Solution

func Solution(n int, edges [][]int, source int, destination int) bool {
	graph := make(map[int]map[int]struct{})
	for _, e := range edges {
		if _, ok := graph[e[0]]; !ok {
			graph[e[0]] = make(map[int]struct{})
		}
		if _, ok := graph[e[1]]; !ok {
			graph[e[1]] = make(map[int]struct{})
		}

		graph[e[0]][e[1]] = struct{}{}
		graph[e[1]][e[0]] = struct{}{}
	}

	visited := map[int]struct{}{}
	visited[source] = struct{}{}
	found := false
	var dfs func(int, *bool)
	dfs = func(source int, found *bool) {
		if *found {
			return
		}
		for rel := range graph[source] {
			if _, ok := visited[rel]; ok {
				continue
			}
			if rel == destination {
				*found = true
				return
			}
			visited[rel] = struct{}{}
			dfs(rel, found)
			delete(visited, rel)
		}
	}
	dfs(source, &found)
	return found
}

func Solution1(n int, edges [][]int, source int, destination int) bool {
	if source == destination {
		return true
	}
	graph := make(map[int]map[int]struct{})
	for _, e := range edges {
		if _, ok := graph[e[0]]; !ok {
			graph[e[0]] = make(map[int]struct{})
		}
		if _, ok := graph[e[1]]; !ok {
			graph[e[1]] = make(map[int]struct{})
		}

		graph[e[0]][e[1]] = struct{}{}
		graph[e[1]][e[0]] = struct{}{}
	}

	queue := []int{source}

	visited := map[int]struct{}{}
	visited[source] = struct{}{}

	for len(queue) > 0 {
		nq := make([]int, 0)
		for _, item := range queue {
			for rel := range graph[item] {
				if rel == destination {
					return true
				}

				if _, ok := visited[rel]; !ok {
					visited[rel] = struct{}{}
					nq = append(nq, rel)
				}
			}
		}
		queue = nq
	}
	return false
}
