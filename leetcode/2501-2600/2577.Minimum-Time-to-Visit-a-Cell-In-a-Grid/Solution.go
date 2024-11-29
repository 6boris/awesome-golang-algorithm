package Solution

import "container/heap"

var dirs2577 = [][2]int{
	{0, 1}, {1, 0}, {0, -1}, {-1, 0},
}

type heap2577 [][3]int

func (h *heap2577) Len() int {
	return len(*h)
}

func (h *heap2577) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *heap2577) Less(i, j int) bool {
	return (*h)[i][2] < (*h)[j][2]
}

func (h *heap2577) Push(x any) {
	*h = append(*h, x.([3]int))
}

func (h *heap2577) Pop() any {
	old := *h
	l := len(old)
	x := old[l-1]
	*h = old[:l-1]
	return x
}

func Solution(grid [][]int) int {
	if grid[0][1] > 1 && grid[1][0] > 1 {
		// 根本走不了
		return -1
	}
	m, n := len(grid), len(grid[0])
	visited := make(map[[2]int]int)
	visited[[2]int{0, 0}] = 0
	q := heap2577{{0, 0, 0}}

	for q.Len() > 0 {
		cur := heap.Pop(&q).([3]int)
		if cur[0] == m-1 && cur[1] == n-1 {
			return cur[2]
		}
		for _, dir := range dirs2577 {
			nx, ny := cur[0]+dir[0], cur[1]+dir[1]
			if nx >= 0 && nx < m && ny >= 0 && ny < n {
				cost := cur[2] + 1 // 耗时比目前小，直接+1过去
				if diff := grid[nx][ny] - cur[2]; diff >= 1 {
					cost = grid[nx][ny]
					if diff&1 == 0 {
						cost++
					}
				}
				if v, ok := visited[[2]int{nx, ny}]; !ok || v > cost {
					heap.Push(&q, [3]int{nx, ny, cost})
					visited[[2]int{nx, ny}] = cost
				}
			}
		}
	}

	return -1
}
