package Solution

import "sort"

func Solution(nums []int) []int {
	var ret []int
	sort.Ints(nums)
	index, cur := 1, nums[0]+1
	for ; index != len(nums); cur++ {
		if cur == nums[index] {
			index++
			continue
		}
		ret = append(ret, cur)
	}
	return ret
}
