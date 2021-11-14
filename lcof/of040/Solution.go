package Solution

import (
	"sort"
)

func getLeastNumbers(arr []int, k int) []int {
	ans := make([]int, k)
	sort.Ints(arr)
	for i := 0; i < k; i++ {
		ans[i] = arr[i]
	}
	return ans
}
