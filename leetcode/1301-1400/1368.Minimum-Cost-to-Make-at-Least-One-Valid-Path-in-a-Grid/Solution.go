package Solution

import "container/heap"

// x, y, cost, dir
type heap1368 [][4]int

func (h *heap1368) Len() int {
	return len(*h)
}
func (h *heap1368) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *heap1368) Less(i, j int) bool {
	return (*h)[i][2] < (*h)[j][2]
}
func (h *heap1368) Push(x any) {
	*h = append(*h, x.([4]int))
}

func (h *heap1368) Pop() any {
	old := *h
	l := len(old)
	x := old[l-1]
	*h = old[:l-1]
	return x
}

var dir1368 = [4][2]int{
	{0, 1}, {0, -1}, {1, 0}, {-1, 0},
}

func Solution(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	if m == 1 && n == 1 {
		return 0
	}
	inf := m*n + 1
	dp := make([][]int, m)
	for i := range m {
		dp[i] = make([]int, n)
		for j := range n {
			dp[i][j] = inf
		}
	}
	// 就是优先队列
	dp[0][0] = 0
	h := &heap1368{}
	if grid[0][0]&1 == 0 {
		dp[0][0] = 1
		heap.Push(h, [4]int{0, 0, 1, 1})
		heap.Push(h, [4]int{0, 0, 1, 3})
	} else {
		heap.Push(h, [4]int{0, 0, 0, grid[0][0]})
	}
	for h.Len() > 0 {
		cur := heap.Pop(h).([4]int)
		if cur[0] == m-1 && cur[1] == n-1 {
			return cur[2]
		}
		for i, d := range dir1368 {
			nx, ny := cur[0]+d[0], cur[1]+d[1]
			if nx < m && nx >= 0 && ny < n && ny >= 0 {
				cost := cur[2]
				if i != cur[3]-1 {
					cost++
				}
				if dp[nx][ny] > cost {
					dp[nx][ny] = cost
					heap.Push(h, [4]int{nx, ny, cost, grid[nx][ny]})
				}
			}
		}
	}

	return -1
}
