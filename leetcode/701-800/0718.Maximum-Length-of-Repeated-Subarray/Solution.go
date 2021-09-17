package Solution

func Solution(nums1 []int, nums2 []int) int {
	n, m := len(nums1), len(nums2)
	dp := make([]int, m+1)
	res := 0
	for i := 1; i < n+1; i++ {
		for j := m; j > 0; j-- {
			if nums1[i-1] == nums2[j-1] {
				dp[j] = 1 + dp[j-1]
			} else {
				dp[j] = 0
			}
			res = max(res, dp[j])
		}
	}
	return res
}

func max(v1, v2 int) int {
	if v1 > v2 {
		return v1
	}
	return v2
}
