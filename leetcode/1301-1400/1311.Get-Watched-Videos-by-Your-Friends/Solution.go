package Solution

import "sort"

func Solution(watchedVideos [][]string, friends [][]int, id int, level int) []string {
	now := 0
	var ans []string
	queue := []int{id}
	visited := [100]bool{}
	visited[id] = true
	for ; now < level; now++ {
		nq := make([]int, 0)
		videos := map[string]int{}
		next := make([]string, 0)
		for _, u := range queue {
			for _, f := range friends[u] {
				if visited[f] {
					continue
				}
				visited[f] = true
				nq = append(nq, f)
				for _, w := range watchedVideos[f] {
					videos[w]++
				}
			}
		}
		for v := range videos {
			next = append(next, v)
		}
		sort.Slice(next, func(i, j int) bool {
			a, b := videos[next[i]], videos[next[j]]
			if a == b {
				return next[i] < next[j]
			}
			return a < b
		})
		ans = next
		queue = nq
	}
	return ans

}
