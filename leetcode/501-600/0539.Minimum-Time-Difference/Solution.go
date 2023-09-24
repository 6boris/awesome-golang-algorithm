package Solution

import (
	"fmt"
	"sort"
)

func Solution(timePoints []string) int {
	hm := make([][]int, len(timePoints))
	var a, b int
	for idx, tp := range timePoints {
		fmt.Sscanf(tp, "%d:%d", &a, &b)
		if a == 0 {
			a = 24
		}
		hm[idx] = []int{a, b}
	}
	// "12:12","00:13"
	sort.Slice(hm, func(i, j int) bool {
		if hm[i][0] == hm[j][0] {
			return hm[i][1] > hm[j][1]
		}
		return hm[i][0] > hm[j][0]
	})
	ans := -1
	l := len(hm)

	for idx := 0; idx < l; idx++ {
		next := (idx + 1) % l
		diffh := hm[idx][0] - hm[next][0]
		diffm := hm[idx][1] - hm[next][1]
		r := diffh*60 + diffm
		if r < 0 {
			r = -r
		}
		if x := 1440 - r; x < r {
			r = x
		}
		if ans == -1 || r < ans {
			ans = r
		}
	}

	return ans
}
