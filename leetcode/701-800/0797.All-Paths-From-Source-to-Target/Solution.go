package Solution

func Solution(graph [][]int) [][]int {
	ans := make([][]int, 0)
	var dfs func(int, []int)
	dfs = func(node int, path []int) {
		if node == len(graph)-1 {
			dst := make([]int, len(path))
			copy(dst, path)
			ans = append(ans, dst)
			return
		}

		for _, relNode := range graph[node] {
			dfs(relNode, append(path, relNode))
		}
	}
	dfs(0, []int{0})
	return ans
}
