# [123. Best Time to Buy and Sell Stock III][title]

## Description

Say you have an array for which the ith element is the price of a given stock on day i.

Design an algorithm to find the maximum profit. You may complete at most two transactions.

Note: You may not engage in multiple transactions at the same time (i.e., you must sell the stock before you buy again).

**Example 1:**

```
Input: [3,3,5,0,0,3,1,4]
Output: 6
Explanation: Buy on day 4 (price = 0) and sell on day 6 (price = 3), profit = 3-0 = 3.
             Then buy on day 7 (price = 1) and sell on day 8 (price = 4), profit = 4-1 = 3.
```

**Example 2:**

```
Input: [1,2,3,4,5]
Output: 4
Explanation: Buy on day 1 (price = 1) and sell on day 5 (price = 5), profit = 5-1 = 4.
             Note that you cannot buy on day 1, buy on day 2 and sell them later, as you are
             engaging multiple transactions at the same time. You must sell before buying again.
```
**Example 2:**

```
Input: [7,6,4,3,1]
Output: 0
Explanation: In this case, no transaction is done, i.e. max profit = 0.
```

**Tags:** Math, String

## 题意
>买卖股票，持有一股，买卖2此

## 题解

### 思路1
> 

```go
func maxProfit(prices []int) int {
	t1_b, t1_s, t2_b, t2_s := math.MinInt32, 0, math.MinInt32, 0

	for _, v := range prices {
		t1_b = max(t1_b, 0-v)
		t1_s = max(t1_s, t1_b+v)
		t2_b = max(t2_b, t1_s-v)
		t2_s = max(t2_s, t2_b+v)
	}
	return t2_s
}
func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

```



## 结语

如果你同我一样热爱数据结构、算法、LeetCode，可以关注我 GitHub 上的 LeetCode 题解：[awesome-golang-leetcode][me]

[title]: https://leetcode.com/problems/best-time-to-buy-and-sell-stock-iii/
[me]: https://github.com/kylesliu/awesome-golang-leetcode
