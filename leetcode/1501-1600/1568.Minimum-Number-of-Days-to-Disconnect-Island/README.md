# [1568.Minimum Number of Days to Disconnect Island][title]

## Description
You are given an `m x n` binary `grid` grid where `1` represents land and `0` represents water. An **island** is a maximal **4-directionally** (horizontal or vertical) connected group of `1`'s.

The grid is said to be **connected** if we have **exactly one island**, otherwise is said **disconnected**.

In one day, we are allowed to change any single land cell `(1)` into a water cell `(0)`.

Return the minimum number of days to disconnect the grid.

**Example 1:**  

![1](./1.jpg)

```
Input: grid = [[0,1,1,0],[0,1,1,0],[0,0,0,0]]

Output: 2
Explanation: We need at least 2 days to get a disconnected grid.
Change land grid[1][1] and grid[0][2] to water and get 2 disconnected island.
```

**Example 2:**  

![2](./2.jpg)

```
Input: grid = [[1,1]]
Output: 2
Explanation: Grid of full water is also disconnected ([[1,1]] -> [[0,0]]), 0 islands.
```

## 结语

如果你同我一样热爱数据结构、算法、LeetCode，可以关注我 GitHub 上的 LeetCode 题解：[awesome-golang-algorithm][me]

[title]: https://leetcode.com/problems/minimum-number-of-days-to-disconnect-island/
[me]: https://github.com/kylesliu/awesome-golang-algorithm
