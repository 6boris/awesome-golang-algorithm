package Solution

import "container/heap"

type heap2290 [][3]int

func (h *heap2290) Len() int {
	return len(*h)
}

func (h *heap2290) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *heap2290) Less(i, j int) bool {
	return (*h)[i][2] < (*h)[j][2]
}

func (h *heap2290) Push(x any) {
	*h = append(*h, x.([3]int))
}

func (h *heap2290) Pop() any {
	old := *h
	l := len(old)
	x := old[l-1]
	*h = old[:l-1]
	return x
}

var dirs2290 = [][2]int{
	{0, 1}, {1, 0}, {-1, 0}, {0, -1},
}

func Solution(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	h := heap2290{{0, 0, 0}}
	v := map[[2]int]int{
		[2]int{0, 0}: 0,
	}
	for h.Len() > 0 {
		cur := heap.Pop(&h).([3]int)
		if cur[0] == m-1 && cur[1] == n-1 {
			return cur[2]
		}
		for _, dir := range dirs2290 {
			nx, ny := cur[0]+dir[0], cur[1]+dir[1]
			if !(nx >= 0 && nx < m && ny >= 0 && ny < n) {
				continue
			}
			need := cur[2] + grid[nx][ny]
			key := [2]int{nx, ny}
			//key = [3]int{nx, ny, need}
			if vv, ok := v[key]; !ok || vv > need {
				heap.Push(&h, [3]int{nx, ny, need})
				v[key] = need
			}
		}
	}
	return -1
}
