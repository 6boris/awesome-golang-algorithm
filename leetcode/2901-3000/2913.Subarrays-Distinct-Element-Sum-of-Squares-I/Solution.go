package Solution

func Solution(nums []int) int {
	n := len(nums)
	var ret int
	for i := 0; i < n; i++ {
		var seen [101]bool
		distinct := 0
		for j := i; j < n; j++ {
			if !seen[nums[j]] {
				seen[nums[j]] = true
				distinct++
			}
			ret += distinct * distinct
		}
	}
	return ret
}
