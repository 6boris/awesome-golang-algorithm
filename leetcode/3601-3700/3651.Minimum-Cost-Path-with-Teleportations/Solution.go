package Solution

import (
	"math"
	"sort"
)

func Solution(grid [][]int, k int) int {
	m, n := len(grid), len(grid[0])
	type point struct{ x, y int }
	points := make([]point, 0, m*n)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			points = append(points, point{i, j})
		}
	}
	sort.Slice(points, func(i, j int) bool {
		return grid[points[i].x][points[i].y] < grid[points[j].x][points[j].y]
	})
	costs := make([][]int, m)
	for i := range costs {
		costs[i] = make([]int, n)
		for j := range costs[i] {
			costs[i][j] = math.MaxInt
		}
	}
	for t := 0; t <= k; t++ {
		minCost := math.MaxInt
		for i, j := 0, 0; i < len(points); i++ {
			minCost = min(minCost, costs[points[i].x][points[i].y])
			if i+1 < len(points) && grid[points[i].x][points[i].y] == grid[points[i+1].x][points[i+1].y] {
				continue
			}
			for r := j; r <= i; r++ {
				costs[points[r].x][points[r].y] = minCost
			}
			j = i + 1
		}
		for i := m - 1; i >= 0; i-- {
			for j := n - 1; j >= 0; j-- {
				if i == m-1 && j == n-1 {
					costs[i][j] = 0
					continue
				}
				if i != m-1 {
					costs[i][j] = min(costs[i][j], costs[i+1][j]+grid[i+1][j])
				}
				if j != n-1 {
					costs[i][j] = min(costs[i][j], costs[i][j+1]+grid[i][j+1])
				}
			}
		}
	}
	return costs[0][0]
}
