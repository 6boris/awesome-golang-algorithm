package Solution

func Solution(nums []int, pattern []int) int {
	var ans int
	m := len(pattern)
	n := len(nums)
	k := 0
	for i := 0; i < n-m; i++ {
		k = 0
		for ; k < m; k++ {
			if pattern[k] == 1 && nums[i+k+1] > nums[i+k] {
				continue
			}
			if pattern[k] == 0 && nums[i+k+1] == nums[i+k] {
				continue
			}
			if pattern[k] == -1 && nums[i+k+1] < nums[i+k] {
				continue
			}
			break
		}
		if k == m {
			ans++
		}
	}
	return ans
}
