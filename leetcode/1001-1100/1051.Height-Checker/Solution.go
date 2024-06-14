package Solution

import "sort"

func Solution(heights []int) int {
	dst := make([]int, len(heights))
	copy(dst, heights)
	sort.Ints(dst)
	ans := 0
	for i := 0; i < len(dst); i++ {
		if dst[i] != heights[i] {
			ans++
		}
	}
	return ans
}
