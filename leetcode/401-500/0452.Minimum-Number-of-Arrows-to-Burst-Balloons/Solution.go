package Solution

import (
	"sort"
)

func findMinArrowShots(points [][]int) int {
	if len(points) == 0 {
		return 0
	}
	sort.Sort(ByLen(points))

	res := 0
	end := -1<<31 - 1
	for _, interval := range points {
		if interval[0] > end {
			res += 1
			end = interval[1]
		}
	}
	return res
}

type ByLen [][]int

func (a ByLen) Len() int           { return len(a) }
func (a ByLen) Less(i, j int) bool { return a[i][1] < a[j][1] }
func (a ByLen) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
