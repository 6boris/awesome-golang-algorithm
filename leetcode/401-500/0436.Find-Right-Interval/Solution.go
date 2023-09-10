package Solution

import (
	"fmt"
	"sort"
)

func Solution(intervals [][]int) []int {
	l := len(intervals)
	ans := make([]int, 0)
	inCopy := make([][]int, l)
	ai := make(map[string]int)
	bi := make(map[string]int)

	for i := 0; i < l; i++ {
		inCopy[i] = make([]int, 2)
		copy(inCopy[i], intervals[i])
		k := fmt.Sprintf("%d-%d", intervals[i][0], intervals[i][1])
		ai[k] = i
	}

	sort.Slice(inCopy, func(i, j int) bool {
		if inCopy[i][0] == inCopy[j][0] {
			return inCopy[i][1] < inCopy[j][1]
		}
		return inCopy[i][0] < inCopy[j][0]
	})
	for i := 0; i < l; i++ {
		k := fmt.Sprintf("%d-%d", inCopy[i][0], inCopy[i][1])
		bi[k] = i
	}

	var bsearch func(int, int, int) int
	bsearch = func(left, right, target int) int {
		ll, r := left, right
		for ll < r {
			mid := ll + (r-ll)/2
			if inCopy[mid][0] < target {
				ll = mid + 1
				continue
			}
			r = mid
		}
		if ll == right || r == left {
			return -1
		}
		k := fmt.Sprintf("%d-%d", inCopy[r][0], inCopy[r][1])
		return ai[k]
	}

	for idx, in := range intervals {
		k := fmt.Sprintf("%d-%d", in[0], in[1])
		if in[0] == in[1] {
			ans = append(ans, idx)
			continue
		}
		left := bi[k]
		ans = append(ans, bsearch(left, l, in[1]))
	}
	return ans
}
