package Solution

func Solution(n int, edges [][]int) []int {
	if n == 1 {
		return []int{0}
	}
	adj := make([]map[int]struct{}, n)
	for i := 0; i < n; i++ {
		adj[i] = make(map[int]struct{})
	}
	for _, edge := range edges {
		a, b := edge[0], edge[1]
		adj[a][b] = struct{}{}
		adj[b][a] = struct{}{}
	}
	l := []int{}
	for i := 0; i < n; i++ {
		if len(adj[i]) == 1 {
			l = append(l, i)
		}
	}
	r := n
	for r > 2 {
		r -= len(l)
		nl := []int{}
		for _, lf := range l {
			var nei int
			for n := range adj[lf] {
				nei = n
			}
			delete(adj[nei], lf)
			if len(adj[nei]) == 1 {
				nl = append(nl, nei)
			}
		}
		l = nl
	}
	return l
}
