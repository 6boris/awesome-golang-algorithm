package Solution

// n^2, TLE
func Solution(arr []int) int {
	ans := make(map[int]struct{})
	for idx := 0; idx < len(arr); idx++ {
		ans[arr[idx]] = struct{}{}
		now := arr[idx]
		for x := idx - 1; x >= 0; x-- {
			now |= arr[x]
			ans[now] = struct{}{}
		}
	}
	return len(ans)
}

// 上一个算法，就是多做了很多无用功，比如某些结果计算不会产生新的数字
// 所以将数字紧凑一下 将到达i时候能产生的数字放到map中，下一次直接前面的即可，最后把所有整合
func Solution1(arr []int) int {
	dp := make([]map[int]struct{}, len(arr))
	dp[0] = map[int]struct{}{
		arr[0]: {},
	}

	ans := map[int]struct{}{
		arr[0]: {},
	}
	for idx := 1; idx < len(arr); idx++ {
		dp[idx] = map[int]struct{}{
			arr[idx]: {},
		}
		for x := range dp[idx-1] {
			dp[idx][arr[idx]|x] = struct{}{}
		}
		for k := range dp[idx] {
			ans[k] = struct{}{}
		}
	}
	return len(ans)
}
