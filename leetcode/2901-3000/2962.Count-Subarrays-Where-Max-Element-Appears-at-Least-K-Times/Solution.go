package Solution

func Solution(nums []int, k int) int64 {
	ans := int64(0)
	m := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] > m {
			m = nums[i]
		}
	}
	indies := make([]int, 0)
	for i := 0; i < len(nums); i++ {
		if nums[i] == m {
			indies = append(indies, i)
		}
	}
	ll := len(nums)
	if ll < k {
		return ans
	}
	le := 0
	for s, e := 0, k-1; e < len(indies); s, e = s+1, e+1 {
		left := int64(indies[s] - le)
		right := int64(ll - 1 - indies[e])
		ans += left + right + left*right + 1
		le = indies[s] + 1
	}
	return ans
}
