package Solution

import "sort"

func Solution(arr []int) int {
	sort.Ints(arr)
	arr[0] = 1
	for idx := 1; idx < len(arr); idx++ {
		diff := arr[idx] - arr[idx-1]
		if diff > 1 {
			arr[idx] = arr[idx-1] + 1
			continue
		}
	}
	return arr[len(arr)-1]
}
