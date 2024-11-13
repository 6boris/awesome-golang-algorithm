package Solution

import "sort"

func Solution(nums []int, lower, upper int) int64 {
	sort.Ints(nums)
	var ans int64 = 0
	l := len(nums)
	for i := 0; i < len(nums); i++ {
		// 1, 2, 3, 4
		tmpL := l - i - 1
		leftIndex := sort.Search(tmpL, func(ii int) bool {
			return nums[i]+nums[ii+1+i] >= lower
		})
		if leftIndex == tmpL {
			continue
		}
		rightIndex := sort.Search(tmpL, func(ii int) bool {
			return nums[i]+nums[ii+1+i] > upper
		})
		if rightIndex == tmpL && nums[i]+nums[l-1] > upper {
			continue
		}
		ans += int64(rightIndex - leftIndex)
	}
	return ans
}
