package Solution

func Solution(nums []int, k int) int {
	m := 0
	in := make(map[int]struct{})
	for i := range nums {
		if nums[i]%k == 0 {
			in[nums[i]] = struct{}{}
			m = max(m, nums[i])
		}
	}
	for i := k; i <= m; i += k {
		if _, ok := in[i]; !ok {
			return i
		}
	}
	return m + k
}
