package Solution

import "sort"

func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	var res [][]int
	for i := 0; i < len(nums)-3; i++ {
		if i != 0 && nums[i] == nums[i-1] {
			continue
		}
		threeSum(&res, nums[i], nums[i+1:], target-nums[i])
	}
	return res
}

func threeSum(res *[][]int, first int, nums []int, target int) {
	nlen := len(nums)
	for i := 0; i < nlen-2; i++ {
		if i != 0 && nums[i] == nums[i-1] {
			continue
		}

		left, right := i+1, nlen-1
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			if sum == target {
				*res = append(*res, []int{first, nums[i], nums[left], nums[right]})
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				for left < right && nums[right] == nums[right-1] {
					right--
				}
				left++
				right--
			} else if sum < target {
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				left++
			} else {
				for left < right && nums[right] == nums[right-1] {
					right--
				}
				right--
			}
		}
	}
}
