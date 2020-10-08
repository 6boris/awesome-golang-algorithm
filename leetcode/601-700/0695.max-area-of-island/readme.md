# [695. Longest Continuous Increasing Subsequence][title]

## Description

Given a non-empty 2D array grid of 0's and 1's, an island is a group of 1's (representing land) connected 4-directionally (horizontal or vertical.) You may assume all four edges of the grid are surrounded by water.

Find the maximum area of an island in the given 2D array. (If there is no island, the maximum area is 0.)

**Example 1:**
```
[[0,0,1,0,0,0,0,1,0,0,0,0,0],
 [0,0,0,0,0,0,0,1,1,1,0,0,0],
 [0,1,1,0,1,0,0,0,0,0,0,0,0],
 [0,1,0,0,1,1,0,0,1,0,1,0,0],
 [0,1,0,0,1,1,0,0,1,1,1,0,0],
 [0,0,0,0,0,0,0,0,0,0,1,0,0],
 [0,0,0,0,0,0,0,1,1,1,0,0,0],
 [0,0,0,0,0,0,0,1,1,0,0,0,0]]
```
Given the above grid, return 6. Note the answer is not 11, because the island must be connected 4-directionally.

**Example 2:**
```
[[0,0,0,0,0,0,0,0]]
```
Given the above grid, return 0.

## Note
 The length of each dimension in the given grid does not exceed 50.

## 题意
> 求岛屿的最大面积。

## 题解

### 思路 1
> 使用 DFS 深度优先搜索。

```go
func maxAreaOfIsland(grid [][]int) int {
    ans := 0
    height, width := len(grid), len(grid[0])

    for i := 0; i < height; i++ {
        for j := 0; j < width; j++ {
            tmp := 0;
            if 1 == grid[i][j] {
                tmp = dfs(grid, i, j)
                if tmp > ans {
                    ans = tmp
                }
            }
        }
    }
    return ans
}

func dfs(grid [][]int, i, j int) int {
    res := 0
    if i >= 0 && i < len(grid) && j >= 0 && j < len(grid[0]) && 1 == grid[i][j] {
        grid[i][j] = 0
        up := dfs(grid, i-1, j)
        down := dfs(grid, i+1, j)
        left := dfs(grid, i, j-1)
        right := dfs(grid, i, j+1)
        res = up + down + left + right + 1
    }
    return res
}
```

## 结语

如果你同我一样热爱数据结构、算法、LeetCode，可以关注我 GitHub 上的 LeetCode 题解：[awesome-golang-leetcode][me]

[title]: https://leetcode.com/problems/max-area-of-island/description/  
[me]: https://github.com/kylesliu/awesome-golang-algorithm
