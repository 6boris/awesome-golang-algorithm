package Solution

func Solution(times [][]int, n int, k int) int {
	adj := make(map[int]map[int]int)
	for _, t := range times {
		from, to, weigth := t[0], t[1], t[2]
		if _, ok := adj[from]; !ok {
			adj[from] = make(map[int]int)
		}
		adj[from][to] = weigth
	}
	visited := make([]int, n+1)
	for i := 0; i <= n; i++ {
		visited[i] = -1
	}
	visited[k] = 0

	queue := []int{k}
	for len(queue) > 0 {
		nq := make([]int, 0)
		for _, q := range queue {
			for n, w := range adj[q] {
				costTime := visited[q] + w
				if visited[n] == -1 || visited[n] > costTime {
					visited[n] = costTime
					nq = append(nq, n)
				}
			}
		}
		queue = nq
	}
	ans := -1
	for i := 1; i <= n; i++ {
		if visited[i] == -1 {
			return -1
		}
		if ans == -1 || ans < visited[i] {
			ans = visited[i]
		}
	}
	return ans
}
