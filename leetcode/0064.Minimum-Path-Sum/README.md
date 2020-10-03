# [64. Minimum Path Sum (DP) ][title]

## Description

Given a m x n grid filled with non-negative numbers, find a path from top left to bottom right which minimizes the sum of all numbers along its path.

> Note: You can only move either down or right at any point in time.

**Example 1:**

```
Input:
[
  [1,3,1],
  [1,5,1],
  [4,2,1]
]
Output: 7
Explanation: Because the path 1→3→1→1→1 minimizes the sum.
```


**Tags:** Math, String

## 题意
>给定一个包含非负整数的 m x n 网格，请找出一条从左上角到右下角的路径，使得路径上的数字总和为最小。


## 题解
### 思路1
> 用DP的思想，想把矩阵的第一行和第一列算出来，然后在算其他的
```go
//  1.计算第一列
[1 0 0]
[2 0 0]
[0 0 0]

[1 0 0]
[2 0 0]
[6 0 0]

//  2.计算第一行
[1 4 0]
[2 0 0]
[6 0 0]

[1 4 5]
[2 0 0]
[6 0 0]
//  3.计算其他
[1 4 5]
[2 7 0]
[6 0 0]

[1 4 5]
[2 7 6]
[6 0 0]

[1 4 5]
[2 7 6]
[6 8 0]

[1 4 5]
[2 7 6]
[6 8 7]

[1 4 5]
[2 7 6]
[6 8 7]

```

```go
func minPathSum(grid [][]int) int {
	row := len(grid)
	col := len(grid[0])

	//	Init the DP
	dp := make([][]int, row)
	for i := 0; i < row; i++ {
		dp[i] = make([]int, col)
	}
	dp[0][0] = grid[0][0]

	//	Calculate first col
	for i := 1; i < row; i++ {
		dp[i][0] = dp[i-1][0] + grid[i][0]
	}

	//	Calculate first row
	for i := 1; i < col; i++ {
		dp[0][i] = dp[0][i-1] + grid[0][i]
	}

	//	Calculate other number
	for i := 1; i < row; i++ {
		for j := 1; j < col; j++ {
			dp[i][j] = Min(dp[i-1][j], dp[i][j-1]) + grid[i][j]
		}
	}
	return dp[row-1][col-1]
}

func Min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

```

### 思路2
> 思路2
```go

```

## 结语

如果你同我一样热爱数据结构、算法、LeetCode，可以关注我 GitHub 上的 LeetCode 题解：[awesome-golang-leetcode][me]

[title]: https://leetcode.com/problems/minimum-path-sum/
[me]: https://github.com/kylesliu/awesome-golang-algorithm
