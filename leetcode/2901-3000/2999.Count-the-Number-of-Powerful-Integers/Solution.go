package Solution

import (
	"fmt"
	"strings"
)

func Solution(start int64, finish int64, limit int, s string) int64 {
	low := fmt.Sprintf("%d", start)
	high := fmt.Sprintf("%d", finish)
	n := len(high)
	low = strings.Repeat("0", n-len(low)) + low // align digits
	pre_len := n - len(s)                       // prefix length
	memo := make([]int64, n)
	for i := range memo {
		memo[i] = -1
	}

	var dfs func(int, bool, bool) int64
	dfs = func(i int, limit_low, limit_high bool) int64 {
		// recursive boundary
		if i == n {
			return 1
		}
		if !limit_low && !limit_high && memo[i] != -1 {
			return memo[i]
		}
		lo := 0
		if limit_low {
			lo = int(low[i] - '0')
		}
		hi := 9
		if limit_high {
			hi = int(high[i] - '0')
		}

		var res int64 = 0
		if i < pre_len {
			for digit := lo; digit <= min(hi, limit); digit++ {
				res += dfs(i+1, limit_low && digit == lo, limit_high && digit == hi)
			}
		} else {
			x := int(s[i-pre_len] - '0')
			if lo <= x && x <= min(hi, limit) {
				res = dfs(i+1, limit_low && x == lo, limit_high && x == hi)
			}
		}

		if !limit_low && !limit_high {
			memo[i] = res
		}
		return res
	}
	return dfs(0, true, true)
}
