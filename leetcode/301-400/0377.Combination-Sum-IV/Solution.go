package Solution

func combinationSum4(nums []int, target int) int {
	dp := make([]int, target+1)
	dp[0] = 1
	for i := 1; i <= target; i++ {
		for _, num := range nums {
			if i-num >= 0 {
				dp[i] += dp[i-num]
			}
		}
	}
	return dp[target]
}

func Solution2(nums []int, target int) int {
	var lookup func(int) int
	// dfs + cache
	cache := make(map[int]int)
	lookup = func(sum int) int {
		if sum < 0 {
			return 0
		}
		if sum == 0 {
			return 1
		}
		if v, ok := cache[sum]; ok {
			return v
		}

		ans := 0
		for _, i := range nums {
			next := sum - i
			ans += lookup(next)
		}
		cache[sum] = ans
		return ans
	}
	lookup(target)
	return cache[target]
}
