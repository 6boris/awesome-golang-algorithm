package Solution

import "sort"

func arrayPairSum(nums []int) int {
	sort.Ints(nums)
	ans := 0
	for i := 0; i < len(nums); i += 2 {
		ans += nums[i]
	}
	return ans
}

func Sort(nums []int) []int {
	if len(nums) == 0 {
		return []int{}
	}
	max := nums[0]
	var left_res, right_res, res []int
	for i := 1; i < len(nums); i++ {
		if nums[i] < max {
			left_res = append(left_res, nums[i])
		} else {
			right_res = append(right_res, nums[i])
		}
	}
	left_res = Sort(left_res)
	right_res = Sort(right_res)
	res = append(res, left_res...)
	res = append(res, max)
	res = append(res, right_res...)
	return res
}
