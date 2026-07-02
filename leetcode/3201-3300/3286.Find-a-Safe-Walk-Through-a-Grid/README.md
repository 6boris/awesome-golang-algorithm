# [3286.Find a Safe Walk Through a Grid][title]

## Description
You are given an `m x n` binary matrix grid and an integer `health`.

You start on the upper-left corner (`0, 0`) and would like to get to the lower-right corner (`m - 1, n - 1`).

You can move up, down, left, or right from one cell to another adjacent cell as long as your health remains **positive**.

Cells `(i, j)` with `grid[i][j] = 1` are considered **unsafe** and reduce your health by 1.

Return `true` if you can reach the final cell with a health value of 1 or more, and `false` otherwise.

**Example 1:**  

![1](./1.png)

```
Input: grid = [[0,1,0,0,0],[0,1,0,1,0],[0,0,0,1,0]], health = 1

Output: true

Explanation:

The final cell can be reached safely by walking along the gray cells below.
```

**Example 2:**  

![2](./2.png)

```
Input: grid = [[0,1,1,0,0,0],[1,0,1,0,0,0],[0,1,1,1,0,1],[0,0,1,0,1,0]], health = 3

Output: false

Explanation:

A minimum of 4 health points is needed to reach the final cell safely.
```

**Example 3:**  

![3](./3.png)

```
Input: grid = [[1,1,1],[1,0,1],[1,1,1]], health = 5

Output: true

Explanation:

The final cell can be reached safely by walking along the gray cells below.



Any path that does not go through the cell (1, 1) is unsafe since your health will drop to 0 when reaching the final cell.
```

## 结语

如果你同我一样热爱数据结构、算法、LeetCode，可以关注我 GitHub 上的 LeetCode 题解：[awesome-golang-algorithm][me]

[title]: https://leetcode.com/problems/find-a-safe-walk-through-a-grid
[me]: https://github.com/kylesliu/awesome-golang-algorithm
