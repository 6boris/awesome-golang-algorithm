package Solution

import "sort"

func Solution(nums []int) [][]int {
	res := [][]int{}
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	visited := make(map[int]bool)
	helper(nums, []int{}, visited, &res)
	return res
}

func helper(nums, res []int, visited map[int]bool, final *[][]int) {
	if len(nums) == len(res) {
		temp := make([]int, len(res))
		copy(temp, res)
		*final = append(*final, temp)
		return
	}
	for i := 0; i < len(nums); i++ {
		if visited[i] {
			continue
		}
		if i > 0 && nums[i] == nums[i-1] && !visited[i-1] {
			continue
		}
		res = append(res, nums[i])
		visited[i] = true
		helper(nums, res, visited, final)
		visited[i] = false
		res = res[:len(res)-1]
	}
}
