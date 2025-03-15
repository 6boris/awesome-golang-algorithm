package Solution

import "slices"

func Solution(nums []int, k int) int {
	l := len(nums)
	left, right := 0, slices.Max(nums)

	capFn := func(capability int) bool {
		count := 0
		i := 0

		for i < l {
			if nums[i] <= capability {
				count++
				i++
				if i < l {
					i++
				}
			} else {
				i++
			}
		}
		return count >= k
	}

	for left < right {
		mid := (left + right) / 2
		if capFn(mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}

	return left
}
