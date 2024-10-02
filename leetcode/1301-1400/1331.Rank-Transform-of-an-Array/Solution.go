package Solution

import "sort"

type index [][2]int

func Solution(arr []int) []int {
	l := len(arr)
	ans := make([]int, l)
	if l == 0 {
		return ans
	}
	slices := make(index, l)
	for i := range l {
		slices[i] = [2]int{i, arr[i]}
	}
	sort.Slice(slices, func(i, j int) bool {
		return slices[i][1] < slices[j][1]
	})

	rank := 1
	ans[slices[0][0]] = rank
	for i := 1; i < l; i++ {
		if slices[i][1] != slices[i-1][1] {
			rank++
		}
		ans[slices[i][0]] = rank
	}
	return ans
}
