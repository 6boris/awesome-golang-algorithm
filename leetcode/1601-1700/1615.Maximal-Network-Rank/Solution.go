package Solution

func Solution(n int, roads [][]int) int {
	adj := make([][]bool, n)
	for i := 0; i < n; i++ {
		adj[i] = make([]bool, n)
	}
	count := make([]int, n)
	for _, road := range roads {
		from, to := road[0], road[1]
		adj[from][to] = true
		adj[to][from] = true
		count[from]++
		count[to]++
	}
	ans := 0

	for c1 := 0; c1 < n-1; c1++ {
		for c2 := c1 + 1; c2 < n; c2++ {
			r := count[c1] + count[c2]
			if adj[c1][c2] {
				r--
			}
			if r > ans {
				ans = r
			}
		}
	}

	return ans
}
