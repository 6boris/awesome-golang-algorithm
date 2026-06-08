package Solution

const MOD = 1000000007

func Solution(zero, one, limit int) int {
	memo := make([][][]int, zero+1)
	for i := range memo {
		memo[i] = make([][]int, one+1)
		for j := range memo[i] {
			memo[i][j] = []int{-1, -1}
		}
	}

	var dp func(int, int, int) int
	dp = func(zero, one, lastBit int) int {
		if zero == 0 {
			if lastBit == 0 || one > limit {
				return 0
			} else {
				return 1
			}
		} else if one == 0 {
			if lastBit == 1 || zero > limit {
				return 0
			} else {
				return 1
			}
		}

		if memo[zero][one][lastBit] == -1 {
			res := 0
			if lastBit == 0 {
				res = (dp(zero-1, one, 0) + dp(zero-1, one, 1)) % MOD
				if zero > limit {
					res = (res - dp(zero-limit-1, one, 1) + MOD) % MOD
				}
			} else {
				res = (dp(zero, one-1, 0) + dp(zero, one-1, 1)) % MOD
				if one > limit {
					res = (res - dp(zero, one-limit-1, 0) + MOD) % MOD
				}
			}
			memo[zero][one][lastBit] = res % MOD
		}
		return memo[zero][one][lastBit]
	}

	return (dp(zero, one, 0) + dp(zero, one, 1)) % MOD
}
