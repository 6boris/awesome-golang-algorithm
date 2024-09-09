package Solution

import "sort"

func Solution(nums []int) {
	l := len(nums)
	tmp := make([]int, l)
	copy(tmp, nums)
	sort.Ints(tmp)
	// 1,5,1,1,6,4
	// 1, 2, 3, 4, 5, 6
	left, right := l/2, l-1
	if l&1 == 0 {
		left--
	}
	end := left
	// 1,5,1,1,6,4
	// 1, 1, 1, 2, 4, 5, 6
	// 1, 6, 1, 5, 1, 4, 2

	index := 0
	for right > end {
		nums[index] = tmp[left]
		nums[index+1] = tmp[right]
		right, left = right-1, left-1
		index += 2
	}
	if left >= 0 {
		nums[index] = tmp[left]
	}
}
