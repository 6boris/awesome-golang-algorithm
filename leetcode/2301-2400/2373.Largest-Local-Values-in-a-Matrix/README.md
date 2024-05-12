# [2373.Largest Local Values in a Matrix][title]

## Description
You are given an `n x n` integer matrix `grid`.

Generate an integer matrix `maxLocal` of size `(n - 2) x (n - 2)` such that:

- `maxLocal[i][j]` is equal to the `largest` value of the `3 x 3` matrix in `grid` centered around row `i + 1` and column `j + 1`.
In other words, we want to find the largest value in every contiguous `3 x 3` matrix in `grid`.

Return the generated matrix.

**Example 1:**  

![1](./ex1.png)

```
Input: grid = [[9,9,8,1],[5,6,2,6],[8,2,6,4],[6,2,2,2]]
Output: [[9,9],[8,6]]
Explanation: The diagram above shows the original matrix and the generated matrix.
Notice that each value in the generated matrix corresponds to the largest value of a contiguous 3 x 3 matrix in grid.
```

**Example 2:**  

![2](./ex2new2.png)

```
Input: grid = [[1,1,1,1,1],[1,1,1,1,1],[1,1,2,1,1],[1,1,1,1,1],[1,1,1,1,1]]
Output: [[2,2,2],[2,2,2],[2,2,2]]
Explanation: Notice that the 2 is contained within every contiguous 3 x 3 matrix in grid.
```

## 结语

如果你同我一样热爱数据结构、算法、LeetCode，可以关注我 GitHub 上的 LeetCode 题解：[awesome-golang-algorithm][me]

[title]: https://leetcode.com/problems/largest-local-values-in-a-matrix/
[me]: https://github.com/kylesliu/awesome-golang-algorithm
