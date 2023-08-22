package Solution

import "math"

func Solution(mat [][]int) [][]int {
	rows := len(mat)
	cols := len(mat[0])
	zeroPoints := make([][2]int, 0)
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if mat[r][c] == 0 {
				zeroPoints = append(zeroPoints, [2]int{r, c})
				continue
			}
			mat[r][c] = math.MaxInt
		}
	}
	var dirs = [][]int{
		{1, 0}, {0, 1}, {-1, 0}, {0, -1},
	}
	for len(zeroPoints) > 0 {
		next := make([][2]int, 0)
		for _, item := range zeroPoints {
			for _, dir := range dirs {
				nx, ny := item[0]+dir[0], item[1]+dir[1]
				if nx < 0 || nx >= rows || ny < 0 || ny >= cols || mat[nx][ny] <= mat[item[0]][item[1]]+1 {
					continue
				}
				mat[nx][ny] = mat[item[0]][item[1]] + 1
				next = append(next, [2]int{nx, ny})
			}
		}
		zeroPoints = next
	}
	return mat

}
