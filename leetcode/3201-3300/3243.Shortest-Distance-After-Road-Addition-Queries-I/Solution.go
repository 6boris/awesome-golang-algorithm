package Solution

func Solution(n int, queries [][]int) []int {
	ans := make([]int, len(queries))
	adj := make(map[int]map[int]struct{})
	for i := range n {
		adj[i] = make(map[int]struct{})
	}
	//  初始化路径
	for i := range n - 1 {
		adj[i][i+1] = struct{}{}
	}
	var bfs func() int
	bfs = func() int {
		queue := []int{0}
		step := 0
		used := map[int]struct{}{
			0: {},
		}
		for len(queue) > 0 {
			nq := make([]int, 0)
			for _, q := range queue {
				if q == n-1 {
					return step
				}
				for rel := range adj[q] {
					if _, ok := used[rel]; !ok {
						nq = append(nq, rel)
						used[rel] = struct{}{}
					}
				}
			}
			queue = nq
			step++
		}
		return -1
	}
	for i, q := range queries {
		adj[q[0]][q[1]] = struct{}{}
		ans[i] = bfs()
	}
	return ans
}
