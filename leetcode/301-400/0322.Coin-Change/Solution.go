package Solution

import "math"

func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := 0; i < amount+1; i++ {
		dp[i] = amount + 1
	}
	dp[0] = 0

	for i := 1; i <= amount; i++ {
		for j := 0; j < len(coins); j++ {
			if coins[j] <= i {
				dp[i] = min(dp[i], dp[i-coins[j]]+1)

			}
		}
	}
	if dp[amount] > amount {
		return -1
	}
	return dp[amount]
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

//动态规划（不压缩空间）
func coinChange2(coins []int, amount int) int {
	dp := make([][]int, len(coins)+1) //定义状态
	for i := 0; i <= len(coins); i++ {
		dp[i] = make([]int, amount+1)
	}
	//初始条件
	for j := 0; j <= amount; j++ {
		dp[0][j] = amount + 1
	}
	dp[0][0] = 0
	for i := 1; i <= len(coins); i++ {
		for j := 0; j <= amount; j++ {
			//状态转移方程
			if j >= coins[i-1] {
				dp[i][j] = int(math.Min(float64(dp[i-1][j]), float64(dp[i][j-coins[i-1]]+1)))
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}
	if dp[len(coins)][amount] > amount {
		return -1
	} else {
		return dp[len(coins)][amount]
	}
}

//动态规划（压缩空间）
func coinChange3(coins []int, amount int) int {
	dp := make([]int, amount+1)
	//初始值
	for i := 0; i <= amount; i++ {
		dp[i] = amount + 1
	}
	dp[0] = 0

	for i := 1; i <= len(coins); i++ {
		for j := 0; j <= amount; j++ {
			if j >= coins[i-1] {
				dp[j] = int(math.Min(float64(dp[j]), float64(dp[j-coins[i-1]]+1)))
			}
		}
	}
	if dp[amount] > amount {
		return -1
	} else {
		return dp[amount]
	}
}
