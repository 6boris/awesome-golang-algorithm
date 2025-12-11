package Solution

import "sort"

func Solution(n int, buildings [][]int) int {
	var ret int
	l := len(buildings)
	byY := make([][]int, l)
	byX := make([][]int, l)
	copy(byY, buildings)
	copy(byX, buildings)
	sort.Slice(byY, func(i, j int) bool {
		return byY[i][1] < byY[j][1]
	})
	sort.Slice(byX, func(i, j int) bool {
		return byX[i][0] < byX[j][0]
	})

	yLoc := make(map[int][]int)
	xLoc := make(map[int][]int)
	for _, build := range byY {
		xLoc[build[0]] = append(xLoc[build[0]], build[1])
	}
	for _, build := range byX {
		yLoc[build[1]] = append(yLoc[build[1]], build[0])
	}

	for _, build := range buildings {
		x, y := build[0], build[1]
		index := sort.Search(len(xLoc[x]), func(i int) bool {
			return xLoc[x][i] >= y
		})
		if index > 0 && index < len(xLoc[x])-1 {
			j := sort.Search(len(yLoc[y]), func(ii int) bool {
				return yLoc[y][ii] >= x
			})
			if j > 0 && j < len(yLoc[y])-1 {
				ret++
			}
		}

	}

	return ret
}
