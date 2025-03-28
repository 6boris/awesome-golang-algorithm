package Solution

import "sort"

func Solution(grid [][]int, queries []int) []int {
	var dirs = [][2]int{
		{0, 1}, {1, 0}, {0, -1}, {-1, 0},
	}
	rows, cols := len(grid), len(grid[0])
	visited := make([][]bool, rows)
	for i := range rows {
		visited[i] = make([]bool, cols)
	}
	indies := make([]int, len(queries))
	for i := range queries {
		indies[i] = i
	}
	sort.Slice(indies, func(i, j int) bool {
		ai, bi := indies[i], indies[j]
		return queries[ai] < queries[bi]
	})

	ans := make([]int, len(queries))
	queue := [][2]int{{0, 0}}
	visited[0][0] = true
	cal := make(map[int]int)
	cnt := 0
	for _, index := range indies {
		target := queries[index]
		if v, ok := cal[target]; ok {
			ans[index] = v
			continue
		}
		for {
			nq := make([][2]int, 0)
			all := true
			for _, cur := range queue {
				if grid[cur[0]][cur[1]] >= target {
					nq = append(nq, [2]int{cur[0], cur[1]})
					continue
				}
				all = false
				cnt++
				for _, d := range dirs {
					nx, ny := cur[0]+d[0], cur[1]+d[1]
					if nx < 0 || nx >= rows || ny < 0 || ny >= cols || visited[nx][ny] {
						continue
					}
					nq = append(nq, [2]int{nx, ny})
					visited[nx][ny] = true
				}
			}
			queue = nq
			if all {
				break
			}
		}
		cal[target] = cnt
		ans[index] = cnt
	}
	return ans
}
