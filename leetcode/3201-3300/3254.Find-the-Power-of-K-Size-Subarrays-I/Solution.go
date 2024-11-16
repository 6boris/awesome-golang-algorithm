package Solution

func Solution(nums []int, k int) []int {
	l := len(nums)
	ml := make([]int, l)
	ml[0] = 1
	for i := 1; i < l; i++ {
		if nums[i]-nums[i-1] == 1 {
			ml[i] = ml[i-1] + 1
			continue
		}
		ml[i] = 1
	}
	ans := make([]int, l-k+1)
	idx := 0
	for i := k - 1; i < l; i, idx = i+1, idx+1 {
		ans[idx] = -1
		if ml[i] >= k {
			ans[idx] = nums[i]
		}
	}
	return ans
}
