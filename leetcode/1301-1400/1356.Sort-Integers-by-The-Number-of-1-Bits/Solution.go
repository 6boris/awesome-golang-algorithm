package Solution

import "sort"

func bitOfOne(n int) int {
	ans := 0
	for n > 0 {
		ans++
		n = n & (n - 1)
	}
	return ans
}
func Solution(arr []int) []int {
	sort.Slice(arr, func(i, j int) bool {
		a := bitOfOne(arr[i])
		b := bitOfOne(arr[j])
		if a == b {
			return arr[i] < arr[j]
		}
		return a < b
	})
	return arr
}
