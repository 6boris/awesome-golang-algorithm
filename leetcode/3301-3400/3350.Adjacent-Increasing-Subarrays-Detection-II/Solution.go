package Solution

import "sort"

func ok(nums []int, k int) bool {
	// 存储所有ok的字数组的下标
	// 然后找index的间隔是否存在不想等就可以了
	indies := []int{}
	start, end := 0, 0
	curLen := 0
	pre := -1001
	// 6, 13, -17, -20, 2
	for ; end < len(nums); end++ {
		if nums[end] <= pre {
			start, curLen = end, 1
		} else {
			curLen++
		}
		pre = nums[end]
		if curLen == k {
			indies = append(indies, start)
			start++
			curLen--
		}
	}
	keys := make(map[int]struct{})
	for _, index := range indies {
		keys[index] = struct{}{}
	}
	for i := 0; i < len(indies)-1; i++ {
		if _, ok := keys[indies[i]+k]; ok {
			return true
		}
	}
	return false
}
func Solution(nums []int) int {
	index := sort.Search(len(nums)/2, func(i int) bool {
		return !ok(nums, i+1)
	})
	return index
}
