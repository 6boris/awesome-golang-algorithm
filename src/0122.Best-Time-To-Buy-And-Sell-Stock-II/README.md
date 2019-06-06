# [122. Best Time to Buy and Sell Stock II][title]

## Description

Given two binary strings, return their sum (also a binary string).

The input strings are both **non-empty** and contains only characters `1` or `0`.

**Example 1:**

```
Input: a = "11", b = "1"
Output: "100"
```

**Example 2:**

```
Input: a = "1010", b = "1011"
Output: "10101"
```

**Tags:** Math, String

## 题意
>假设你有一个数组，其中第i个元素表示第i天某个股票的价格。
 设计一种算法以找到最大利润，可以完成任意多次交易，但必须先购买股票再出售股票，不能同时多次交易。
 
## 题解

### 思路1
> (动态规划) O(n)

- 设计状态`f[i]`表示第`i`天，当前不持有股票的最大收益；`g[i]`表示第`i`天，当前持有股票的最大收益。
- 状态转移为`f[i] = max(f[i - 1], g[i - 1] + prices[i])`，表示构成第i天不持有股票有两种方式，一种是前一天也不持有，另一种是前一天持有且这一天售出；二者取最大值。
- `g[i] = max(g[i - 1], f[i - 1] - prices[i])`，表示构成第i天持有股票有两种方式，一种是前一天持有，另一种是前一天不持有，但这一天刚刚买入。
- 最终答案为`f[n - 1]`，即最后一天不持有股票的最大收益。

```go
func maxProfit2(prices []int) int {
	// 初始化DP
	n := len(prices)
	dp := make([][]int, 0)
	for i := 0; i <= n; i++ {
		dp = append(dp, make([]int, 2))
	}
	dp[0][0], dp[0][1] = 0, math.MinInt32
	profit_max := 0
	for i := 0; i < n; i++ {
		dp[i+1][0] = max(dp[i][1]+prices[i], dp[i][0])
		dp[i+1][1] = max(dp[i][0]-prices[i], dp[i][1])

		profit_max = max(profit_max, dp[i+1][0])
		profit_max = max(profit_max, dp[i+1][1])
	}
	return profit_max
}
```

### 思路2
> 贪心:遍历一次数组，低进高出，把正的价格差相加起来就是最终利润。

遍历一次数组，低进高出，把正的价格差相加起来就是最终利润。
- 递增，如`[1,2,3]`，那么1买3卖 与 每天都买入卖出 等价
- 递减，如`[3,2,1]`，赚钱是赚不了的
- 先高再低，如`[1,3,2]`，那么只能在1买3卖捞一笔
- 先低再高，如`[2,1,3]`，那么同样只能在1买3卖捞一笔

```go
func maxProfit(prices []int) int {
	profit := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			profit += prices[i] - prices[i-1]
		}
	}
	return profit
}
```

## 结语

如果你同我一样热爱数据结构、算法、LeetCode，可以关注我 GitHub 上的 LeetCode 题解：[awesome-golang-leetcode][me]

[title]: https://leetcode.com/problems/best-time-to-buy-and-sell-stock-ii/
[me]: https://github.com/kylesliu/awesome-golang-leetcode
