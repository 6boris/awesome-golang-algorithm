package Solution

import (
	"container/heap"
)

type point struct {
	x, y, l int
}
type points []point

func (p *points) Len() int {
	return len(*p)
}

func (p *points) Swap(i, j int) {
	(*p)[i], (*p)[j] = (*p)[j], (*p)[i]
}

func (p *points) Less(i, j int) bool {
	return (*p)[i].l < (*p)[j].l
}

func (p *points) Push(x interface{}) {
	*p = append(*p, x.(point))
}

func (p *points) Pop() interface{} {
	old := *p
	n := len(old)
	x := old[n-1]
	*p = old[:n-1]
	return x
}

var dir1926 = [][]int{
	{1, 0},
	{-1, 0},
	{0, 1},
	{0, -1},
}

func Solution(maze [][]byte, entrance []int) int {
	init := &points{}
	heap.Init(init)
	m, n := len(maze), len(maze[0])
	visited := make([][]bool, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]bool, n)
	}

	root := point{x: entrance[0], y: entrance[1], l: 0}
	visited[entrance[0]][entrance[1]] = true
	heap.Push(init, root)
	for init.Len() > 0 {
		getOne := heap.Pop(init).(point)
		if (getOne.x == 0 || getOne.x == m-1 || getOne.y == 0 || getOne.y == n-1) &&
			(getOne.x != entrance[0] || getOne.y != entrance[1]) {
			return getOne.l
		}
		for _, dir := range dir1926 {
			nextX, nextY := getOne.x+dir[0], getOne.y+dir[1]
			if nextX >= 0 && nextX < m && nextY >= 0 && nextY < n && maze[nextX][nextY] == '.' && !visited[nextX][nextY] {
				heap.Push(init, point{x: nextX, y: nextY, l: getOne.l + 1})
				visited[nextX][nextY] = true
			}
		}
	}
	return -1
}
