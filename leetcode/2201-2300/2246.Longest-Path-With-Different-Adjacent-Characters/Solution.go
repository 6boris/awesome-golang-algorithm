package Solution

import "sort"

func Solution(parent []int, s string) int {
	nodes := len(s)
	tree := make(map[int][]int)

	for i := 1; i < nodes; i++ {
		if _, ok := tree[i]; !ok {
			tree[i] = make([]int, 0)
		}
		if _, ok := tree[parent[i]]; !ok {
			tree[parent[i]] = make([]int, 0)
		}

		tree[i] = append(tree[i], parent[i])
		tree[parent[i]] = append(tree[parent[i]], i)
	}

	max := 1
	var dfs func(int, int) int
	dfs = func(node, parent int) int {
		total := make([]int, 0)
		for _, rel := range tree[node] {
			if rel == parent {
				continue
			}
			next := dfs(rel, node)
			if s[rel] != s[node] {
				total = append(total, next)
			}
		}
		if len(total) == 0 {
			return 1
		}
		ans := 0
		sort.Slice(total, func(i, j int) bool {
			return total[i] > total[j]
		})
		for i := 0; i < len(total) && i < 2; i++ {
			ans += total[i]
		}
		if ans+1 > max {
			max = ans + 1
		}
		return total[0] + 1
	}
	dfs(0, -1)

	return max
}
