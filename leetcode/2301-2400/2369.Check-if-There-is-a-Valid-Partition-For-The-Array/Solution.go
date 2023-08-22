package Solution

func Solution(nums []int) bool {
	// c2, c1, c0 三种情况
	a, b, c := uint8(1), uint8(2), uint8(4)
	l := len(nums)

	dp := make([]uint8, l)
	dp[0] = 0
	if nums[1] == nums[0] {
		dp[1] = 1
	}
	if l == 2 {
		return dp[1] == 1
	}

	for idx := 2; idx < l; idx++ {
		if nums[idx] == nums[idx-1] {
			if dp[idx-2] > 0 {
				dp[idx] |= a
			}
		}
		if nums[idx] == nums[idx-1] && nums[idx-1] == nums[idx-2] {
			// try add b
			if idx == 2 || idx > 2 && dp[idx-3] > 0 {
				dp[idx] |= b
			}
		}
		if nums[idx-2]+2 == nums[idx] && nums[idx-1]+1 == nums[idx] {
			// try add c
			if idx == 2 || idx > 2 && dp[idx-3] > 0 {
				dp[idx] |= c
			}
		}
	}
	if l == 3 {
		return dp[l-1]&b == b || dp[l-1]&c == c
	}
	if dp[l-1]&a == a && dp[l-1-2] > 0 {
		return true
	}
	if l-4 >= 0 {
		if dp[l-1]&b == b && dp[l-4] > 0 {
			return true
		}
		if dp[l-1]&c == c && dp[l-4] > 0 {
			return true
		}
	}
	return false
}
