package Solution

func Solution(isWater [][]int) [][]int {
	var dir = [][2]int{
		{0, 1}, {0, -1}, {1, 0}, {-1, 0},
	}
	queue := [][2]int{}
	m, n := len(isWater), len(isWater[0])
	res := make([][]int, m)
	for i := range m {
		res[i] = make([]int, n)
		for j := range n {
			res[i][j] = -1
		}
	}
	for i := range m {
		for j := range n {
			if isWater[i][j] == 1 {
				queue = append(queue, [2]int{i, j})
				res[i][j] = 0
			}
		}
	}
	h := 1
	for len(queue) > 0 {
		nq := make([][2]int, 0)
		for _, cur := range queue {
			for _, d := range dir {
				nx, ny := cur[0]+d[0], cur[1]+d[1]
				if nx >= 0 && nx < m && ny >= 0 && ny < n && res[nx][ny] == -1 {
					res[nx][ny] = h
					nq = append(nq, [2]int{nx, ny})
				}
			}
		}
		queue = nq
		h++
	}

	return res
}
