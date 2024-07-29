package Solution

func Solution(n int, edges [][]int, restricted []int) int {
	restrictedMap := make(map[int]struct{})
	for _, n := range restricted {
		restrictedMap[n] = struct{}{}
	}
	adj := make([][]int, n)
	for _, e := range edges {
		adj[e[0]] = append(adj[e[0]], e[1])
		adj[e[1]] = append(adj[e[1]], e[0])
	}
	ans := 0
	q := []int{0}
	v := make([]bool, n)
	v[0] = true
	for len(q) > 0 {
		nq := make([]int, 0)
		ans += len(q)
		for _, cur := range q {
			for _, next := range adj[cur] {
				if v[next] {
					continue
				}
				if _, ok := restrictedMap[next]; ok {
					continue
				}
				nq = append(nq, next)
				v[next] = true
			}
		}
		q = nq
	}
	return ans
}
