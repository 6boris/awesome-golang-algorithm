package Solution

import "container/heap"

type heap3341 [][3]int

func (h *heap3341) Len() int {
	return len(*h)
}

func (h *heap3341) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *heap3341) Less(i, j int) bool {
	return (*h)[i][2] < (*h)[j][2]
}

func (h *heap3341) Push(x any) {
	*h = append(*h, x.([3]int))
}

func (h *heap3341) Pop() any {
	old := *h
	l := len(old)
	x := old[l-1]
	*h = old[:l-1]
	return x
}

var dirs3341 = [][2]int{
	{0, 1}, {0, -1}, {-1, 0}, {1, 0},
}

func Solution(moveTime [][]int) int {
	m, n := len(moveTime), len(moveTime[0])
	h := heap3341{{0, 0, 0}}
	visited := map[[2]int]struct{}{
		[2]int{0, 0}: {},
	}
	for h.Len() > 0 {
		top := heap.Pop(&h).([3]int)
		if top[0] == m-1 && top[1] == n-1 {
			return top[2]
		}
		for _, dir := range dirs3341 {
			nx, ny := top[0]+dir[0], top[1]+dir[1]
			if nx >= 0 && nx < m && ny >= 0 && ny < n {
				if _, ok := visited[[2]int{nx, ny}]; !ok {
					visited[[2]int{nx, ny}] = struct{}{}
					cost := moveTime[nx][ny]
					if cost <= top[2] {
						cost = top[2] + 1
					} else {
						cost++
					}
					heap.Push(&h, [3]int{nx, ny, cost})
				}
			}
		}
	}
	return -1
}
