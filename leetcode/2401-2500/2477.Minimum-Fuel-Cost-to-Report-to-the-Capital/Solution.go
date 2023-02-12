package Solution

func Solution(roads [][]int, seats int) int64 {
	edges := make(map[int][]int)
	for _, road := range roads {
		from, to := road[0], road[1]
		if _, ok := edges[from]; !ok {
			edges[from] = make([]int, 0)
		}
		if _, ok := edges[to]; !ok {
			edges[to] = make([]int, 0)
		}
		edges[from] = append(edges[from], to)
		edges[to] = append(edges[to], from)
	}

	var dfs func(int) (int, int, int)
	visited := make(map[int]struct{})
	visited[0] = struct{}{}
	dfs = func(node int) (int, int, int) {
		cars, empty, steps := 0, 0, 0
		skipAll := true
		for _, neigh := range edges[node] {
			if _, ok := visited[neigh]; ok {
				continue
			}
			skipAll = false
			visited[neigh] = struct{}{}
			c, m, s := dfs(neigh)
			cars += c
			empty += m
			steps += s
		}

		if node == 0 {
			return cars, empty, steps
		}
		if skipAll {
			return 1, seats - 1, 1
		}

		human := cars*seats - empty + 1
		cars = human / seats
		x := 0
		if r := human % seats; r != 0 {
			cars++
			x = seats - r
		}
		return cars, x, steps + cars
	}

	_, _, ans := dfs(0)
	return int64(ans)
}
