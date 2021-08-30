package Solution

import "container/list"

func movingCount1(m int, n int, k int) int {
	queue := list.List{}
	visited := make([][]bool, m)
	for i := 0; i < len(visited); i++ {
		visited[i] = make([]bool, n)
	}
	queue.PushBack([]int{0, 0, 0, 0})
	res := 0
	for queue.Len() > 0 {
		back := queue.Back()
		queue.Remove(back)
		bfs := back.Value.([]int)
		i := bfs[0]
		j := bfs[1]
		si := bfs[2]
		sj := bfs[3]
		if i >= m || j >= n || si+sj > k || visited[i][j] {
			continue
		}

		res++
		visited[i][j] = true

		sj1 := sj + 1
		si1 := si + 1
		if (j+1)%10 == 0 {
			sj1 = sj - 8
		}
		if (i+1)%10 == 0 {
			si1 = si - 8
		}

		queue.PushBack([]int{i + 1, j, si1, sj})
		queue.PushBack([]int{i, j + 1, si, sj1})
	}
	return res
}

var n1, m1, k1 int
var visited [][]bool

func movingCount2(m int, n int, k int) int {
	m1 = m
	n1 = n
	k1 = k
	visited = make([][]bool, m)
	for i := 0; i < len(visited); i++ {
		visited[i] = make([]bool, n)
	}

	return dfs(0, 0, 0, 0)
}

func dfs(i int, j int, si int, sj int) int {
	if i >= m1 || j >= n1 || si+sj > k1 || visited[i][j] {
		return 0
	}
	visited[i][j] = true

	sj1 := sj + 1
	si1 := si + 1
	if (j+1)%10 == 0 {
		sj1 = sj - 8
	}
	if (i+1)%10 == 0 {
		si1 = si - 8
	}

	return 1 + dfs(i, j+1, si, sj1) + dfs(i+1, j, si1, sj)
}
