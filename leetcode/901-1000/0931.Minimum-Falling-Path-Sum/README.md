# [931.Minimum Falling Path Sum][title]

## Description
Given an `n x n` array of integers `matrix`, return the **minimum sum** of any **falling path** through `matrix`.

A **falling path** starts at any element in the first row and chooses the element in the next row that is either directly below or diagonally left/right. Specifically, the next element from position (`row, col`) will be (`row + 1, col - 1`), (`row + 1, col`), or (`row + 1, col + 1`).

**Example 1:**  

![example1](./failing1-grid.jpg)

```
Input: matrix = [[2,1,3],[6,5,4],[7,8,9]]
Output: 13
Explanation: There are two falling paths with a minimum sum as shown.
```

**Example 2:**  

![example2](./failing2-grid.jpg)

```
Input: matrix = [[-19,57],[-40,-5]]
Output: -59
Explanation: The falling path with a minimum sum is shown.
```

## 结语

如果你同我一样热爱数据结构、算法、LeetCode，可以关注我 GitHub 上的 LeetCode 题解：[awesome-golang-algorithm][me]

[title]: https://leetcode.com/problems/minimum-falling-path-sum/
[me]: https://github.com/kylesliu/awesome-golang-algorithm
