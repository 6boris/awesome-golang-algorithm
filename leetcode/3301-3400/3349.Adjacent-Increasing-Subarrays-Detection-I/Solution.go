package Solution

func Solution(nums []int, k int) bool {
	indies := map[int]struct{}{}
	start, end := 0, 0
	curLen := 0
	pre := -1001
	for ; end < len(nums); end++ {
		if nums[end] <= pre {
			start, curLen = end, 1
		} else {
			curLen++
		}
		pre = nums[end]
		if curLen == k {
			indies[start] = struct{}{}
			start++
			curLen--
		}
	}
	for index := range indies {
		if _, ok := indies[index+k]; ok {
			return true
		}
	}
	return false
}
