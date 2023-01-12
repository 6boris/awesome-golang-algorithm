package Solution

func Solution(n int, edges [][]int, labels string) []int {
	ans := make([]int, n)
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
	var dfs func(int, int) []int
	dfs = func(node, parent int) []int {
		rels := graph[node]
		total := make([]int, 26)
		total[labels[node]-'a']++
		for _, next := range rels {
			if next == parent {
				continue
			}
			next := dfs(next, node)
			for i, n := range next {
				if n != 0 {
					total[i] += n
				}
			}
		}
		ans[node] = total[labels[node]-'a']
		return total
	}
	dfs(0, -1)
	return ans
}
