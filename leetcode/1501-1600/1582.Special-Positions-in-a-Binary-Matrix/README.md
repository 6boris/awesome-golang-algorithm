# [1582.Special Positions in a Binary Matrix][title]

## Description
Given an `m x n` binary matrix `mat`, return the number of special positions in `mat`.

A position (`i, j`) is called **special** if `mat[i][j] == 1` and all other elements in row `i` and column `j` are `0` (rows and columns are **0-indexed**).

**Example 1:**  
![example1](./special1.jpg)

```
Input: mat = [[1,0,0],[0,0,1],[1,0,0]]
Output: 1
Explanation: (1, 2) is a special position because mat[1][2] == 1 and all other elements in row 1 and column 2 are 0.
```

**Example 2:**  
![example2](./special-grid.jpg)

```
Input: mat = [[1,0,0],[0,1,0],[0,0,1]]
Output: 3
Explanation: (0, 0), (1, 1) and (2, 2) are special positions.
```

## 结语

如果你同我一样热爱数据结构、算法、LeetCode，可以关注我 GitHub 上的 LeetCode 题解：[awesome-golang-algorithm][me]

[title]: https://leetcode.com/problems/special-positions-in-a-binary-matrix/
[me]: https://github.com/kylesliu/awesome-golang-algorithm
