package Solution

const mod940 = 1000000007

func Solution(s string) int {
	dp := [26]int{}
	total := 0
	for i := 0; i < len(s); i++ {
		c := s[i] - 'a'
		// 出现s[i]后新增的不同子序列为total+1
		nc := (total + 1) % mod940
		// 如果s[i]出现过，那么就说明以s[i]结尾的贝多算了
		total = (total + nc - dp[c] + mod940) % mod940
		dp[c] = nc
	}
	return total
}
