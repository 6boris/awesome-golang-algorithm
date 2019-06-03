# [322. Coin Change][title]

## Description

You are given coins of different denominations and a total amount of money amount. Write a function to compute the fewest number of coins that you need to make up that amount. If that amount of money cannot be made up by any combination of the coins, return -1.

**Example 1:**

```
Input: coins = [1, 2, 5], amount = 11
Output: 3 
Explanation: 11 = 5 + 5 + 1
```

**Example 2:**

```
Input: coins = [2], amount = 3
Output: -1
```

**Tags:** Math, String

## 题意
> 给定 nn 种不同硬币的面值，以及需要凑出的总面值 totaltotal。请写一个函数，求最少需要多少硬币，可以凑出 totaltotal 的钱。
  如果不存在任何一种拼凑方案，则返回-1。
  


## 题解

### 动态规划
> 状态表示：
- dp[i]表示凑出 i 价值的钱，最少需要多少个硬币。
- 第i种硬币 dp[i] = min(dp[i], min(dp[i - coins[i]] ) + 1)

```go
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
```

### 思路2
> 思路2
```go

```

## 结语

如果你同我一样热爱数据结构、算法、LeetCode，可以关注我 GitHub 上的 LeetCode 题解：[awesome-golang-leetcode][me]

[title]: https://leetcode.com/problems/two-sum/description/
[me]: https://github.com/kylesliu/awesome-golang-leetcode
