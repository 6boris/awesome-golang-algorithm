package Solution

import "sort"

func combinationSum2(nums []int, target int) (result [][]int) {
	sort.Ints(nums)
	combinationSum2Helper(nums, nil, target, 0, 0, &result)
	return result
}

func combinationSum2Helper(nums, combo []int, target, sum, startIndex int, result *[][]int) {
	if sum == target {
		*result = append(*result, append([]int{}, combo...))
		return
	}
	for i := startIndex; i < len(nums) && (sum+nums[i]) <= target; i++ {
		if i != startIndex && nums[i] == nums[i-1] {
			continue
		}
		combinationSum2Helper(nums, append(combo, nums[i]), target, sum+nums[i], i+1, result)
	}
}
