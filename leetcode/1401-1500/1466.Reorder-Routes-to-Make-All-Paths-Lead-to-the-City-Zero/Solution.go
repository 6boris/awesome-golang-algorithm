package Solution

func Solution(n int, connections [][]int) int {
	to := make(map[int][]int)
	from := make(map[int][]int)
	for _, conn := range connections {
		if _, ok := to[conn[1]]; !ok {
			to[conn[1]] = make([]int, 0)
		}
		if _, ok := from[conn[0]]; !ok {
			from[conn[0]] = make([]int, 0)
		}
		to[conn[1]] = append(to[conn[1]], conn[0])
		from[conn[0]] = append(from[conn[0]], conn[1])
	}
	ans := 0
	var dfs func(int, int)
	dfs = func(start, parent int) {
		for _, rel := range to[start] {
			if rel != parent {
				dfs(rel, start)
			}
		}
		for _, rel := range from[start] {
			if rel != parent {
				ans++
				dfs(rel, start)
			}
		}
	}
	dfs(0, -1)
	return ans
}
