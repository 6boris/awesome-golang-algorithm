package Solution

import "sort"

func Solution(n int, restrictions [][]int) int {
	restrictions = append(restrictions, []int{1, 0})

	sort.Slice(restrictions, func(i, j int) bool {
		return restrictions[i][0] < restrictions[j][0]
	})

	m := len(restrictions)

	for i := 1; i < m; i++ {
		dist := restrictions[i][0] - restrictions[i-1][0]
		if restrictions[i][1] > restrictions[i-1][1]+dist {
			restrictions[i][1] = restrictions[i-1][1] + dist
		}
	}

	for i := m - 2; i >= 0; i-- {
		dist := restrictions[i+1][0] - restrictions[i][0]
		if restrictions[i][1] > restrictions[i+1][1]+dist {
			restrictions[i][1] = restrictions[i+1][1] + dist
		}
	}

	ret := 0

	for i := 0; i < m-1; i++ {
		id1, h1 := restrictions[i][0], restrictions[i][1]
		id2, h2 := restrictions[i+1][0], restrictions[i+1][1]

		peak := (h1 + h2 + (id2 - id1)) / 2
		ret = max(ret, peak)
	}

	lastId := restrictions[m-1][0]
	lastH := restrictions[m-1][1]
	ret = max(ret, lastH+(n-lastId))

	return ret
}
