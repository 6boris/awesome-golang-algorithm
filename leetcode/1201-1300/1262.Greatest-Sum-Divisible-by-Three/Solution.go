package Solution

func Solution(nums []int) int {
	dp := []int{-1, -1, -1}
	dp[0] = 0

	for _, num := range nums {
		t := make([]int, 3)
		copy(t, dp)
		for _, i := range t {
			if i >= 0 {
				newR := (i + num) % 3
				if i+num > dp[newR] {
					dp[newR] = i + num
				}
			}
		}
	}
	return dp[0]
}
