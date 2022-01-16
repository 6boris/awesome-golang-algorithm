package Solution

func Solution(nums []int) int {
	length := len(nums)
	dp := make([][]int, length+1)
	for idx := 0; idx < length+1; idx++ {
		dp[idx] = make([]int, length)
	}

	for _len := 1; _len <= length; _len++ {
		for i := 0; i+_len <= length; i++ {
			j := i + _len - 1
			for k := i; k <= j; k++ {
				p, n := 1, 1
				if i-1 >= 0 {
					p = nums[i-1]
				}
				if j+1 < length {
					n = nums[j+1]
				}
				left, right := 0, dp[k+1][j]
				if k-1 >= 0 {
					left = dp[i][k-1]
				}
				if r := left + right + nums[k]*p*n; r > dp[i][j] {
					dp[i][j] = r
				}
			}
		}
	}

	return dp[0][length-1]
}
