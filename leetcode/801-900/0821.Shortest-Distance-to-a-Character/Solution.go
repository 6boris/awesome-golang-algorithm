package Solution

import "sort"

func Solution(s string, c byte) []int {
	ret := make([]int, len(s))
	indies := make([]int, 0)
	for i := range s {
		if s[i] == c {
			indies = append(indies, i)
		}
	}
	n := len(indies)
	for i := range s {
		idx := sort.Search(n, func(ii int) bool {
			return indies[ii] >= i
		})
		if idx == n {
			ret[i] = i - indies[n-1]
			continue
		}
		distance := indies[idx] - i
		if idx > 0 {
			distance = min(distance, i-indies[idx-1])
		}
		ret[i] = distance
	}
	return ret
}
