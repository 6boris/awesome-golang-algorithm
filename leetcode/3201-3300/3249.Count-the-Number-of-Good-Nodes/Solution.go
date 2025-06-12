package Solution

func Solution(edges [][]int) int {
	var ans int
	n := len(edges) + 1
	adj := make(map[int][]int)
	for _, e := range edges {
		adj[e[0]] = append(adj[e[0]], e[1])
		adj[e[1]] = append(adj[e[1]], e[0])
	}
	size := make([][]int, n)
	var dfs func(int, int) int
	dfs = func(root, parent int) int {
		cnt := 1
		for _, rel := range adj[root] {
			if rel == parent {
				continue
			}
			lz := dfs(rel, root)
			cnt += lz
			size[root] = append(size[root], lz)
		}
		return cnt
	}
	dfs(0, -1)
	for _, ls := range size {
		if len(ls) == 0 {
			ans++
			continue
		}

		cmp := ls[0]
		i := 1
		for ; i < len(ls); i++ {
			if cmp != ls[i] {
				break
			}
		}
		if i == len(ls) {
			ans++
		}
	}
	return ans
}
