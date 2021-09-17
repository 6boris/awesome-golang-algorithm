package Solution

import "sort"

//	排序 + 遍历
func findRepeatNumber(nums []int) int {
	sort.Ints(nums)
	for idx := range nums {
		if idx > 0 {
			if nums[idx] == nums[idx-1] {
				return nums[idx]
			}
		}
	}
	return 0
}

//	map哈希表
func findRepeatNumber2(nums []int) int {
	mapTmp := make(map[int]bool)
	for _, val := range nums {
		if mapTmp[val] {
			return val
		} else {
			mapTmp[val] = true
		}
	}
	return 0
}

//	原地置换
func findRepeatNumber3(nums []int) int {
	for idx, val := range nums {
		if idx == val {
			continue
		}
		if nums[val] == val {
			return val
		}
		nums[val], nums[idx] = nums[idx], nums[val]
	}
	return 0
}
