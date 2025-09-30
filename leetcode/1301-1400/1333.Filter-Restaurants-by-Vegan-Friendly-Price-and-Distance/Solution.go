package Solution

import "sort"

func Solution(restaurants [][]int, veganFriendly int, maxPrice int, maxDistance int) []int {
	left := make([]int, 0)
	for i := 0; i < len(restaurants); i++ {
		if veganFriendly == 1 && restaurants[i][2] == 0 {
			continue
		}
		if restaurants[i][3] > maxPrice {
			continue
		}
		if restaurants[i][4] > maxDistance {
			continue
		}
		left = append(left, i)
	}
	sort.Slice(left, func(i, j int) bool {
		a, b := restaurants[left[i]], restaurants[left[j]]
		if a[1] == b[1] {
			return a[0] > b[0]
		}
		return a[1] > b[1]
	})
	ret := make([]int, len(left))
	for i, idx := range left {
		ret[i] = restaurants[idx][0]
	}
	return ret
}
