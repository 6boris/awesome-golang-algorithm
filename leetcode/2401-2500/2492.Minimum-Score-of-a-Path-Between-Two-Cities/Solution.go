package Solution

import "fmt"

func Solution(n int, roads [][]int) int {
	score := -1
	edges := make(map[int]map[int]int)
	for _, r := range roads {
		from, to, distance := r[0], r[1], r[2]
		if _, ok := edges[from]; !ok {
			edges[from] = make(map[int]int)
		}
		if _, ok := edges[to]; !ok {
			edges[to] = make(map[int]int)
		}
		edges[from][to] = distance
		edges[to][from] = distance
	}
	var dfs func(int)
	visited := make(map[string]struct{})
	dfs = func(city int) {

		for to, distnace := range edges[city] {
			key := fmt.Sprintf("%d-%d", city, to)
			if _, ok := visited[key]; ok {
				continue
			}
			revKey := fmt.Sprintf("%d-%d", to, city)
			visited[key] = struct{}{}
			visited[revKey] = struct{}{}

			if score == -1 || score > distnace {
				score = distnace
			}
			dfs(to)
		}
	}

	dfs(1)
	return score
}
