package Solution

import "sort"

func Solution(n int) int {
	// 1, 2, 3, 4, 5 n = 5
	// 0, 1, 2, 3, 4
	i := sort.Search(n, func(i int) bool {
		stair := i + 1
		return stair*(stair+1) >= 2*n
	})
	return i + 1
}
