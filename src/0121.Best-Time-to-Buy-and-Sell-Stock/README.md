# [121. Best Time to Buy and Sell Stock][title]

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
 如果您只允许完成至多一笔交易（即买入一只股票并卖出一只股票），则设计一种算法以找到最大利润。
 必须先购买股票再出售股票。

## 题解

### 思路1
> 贪心

- 遍历一次即可得出结果

```go
func maxProfit2(prices []int) int {
	profit, low := 0, prices[0]
	for i := 0; i < len(prices); i++ {
		profit = max(profit, prices[i]-low)
		low = min(low, prices[i])
	}
	return profit
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
```

### 思路2
> 思路2
```go

```

## 结语

如果你同我一样热爱数据结构、算法、LeetCode，可以关注我 GitHub 上的 LeetCode 题解：[awesome-golang-leetcode][me]

[title]: https://leetcode.com/problems/best-time-to-buy-and-sell-stock/
[me]: https://github.com/kylesliu/awesome-golang-leetcode
