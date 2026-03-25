# [3546.Equal Sum Grid Partition I][title]

## Description
You are given an `m x n` matrix `grid` of positive integers. Your task is to determine if it is possible to make **either one horizontal or one vertical cut** on the grid such that:

- Each of the two resulting sections formed by the cut is **non-empty**.
- The sum of the elements in both sections is **equal**.

Return `true` if such a partition exists; otherwise return `false`.

**Example 1:**  

![1](./1.jpeg)

```
Input: grid = [[1,4],[2,3]]

Output: true

Explanation:

A horizontal cut between row 0 and row 1 results in two non-empty sections, each with a sum of 5. Thus, the answer is true.
```

**Example 2:**

```
Input: grid = [[1,3],[2,4]]

Output: false

Explanation:

No horizontal or vertical cut results in two non-empty sections with equal sums. Thus, the answer is false.
```

## 结语

如果你同我一样热爱数据结构、算法、LeetCode，可以关注我 GitHub 上的 LeetCode 题解：[awesome-golang-algorithm][me]

[title]: https://leetcode.com/problems/equal-sum-grid-partition-i/
[me]: https://github.com/kylesliu/awesome-golang-algorithm
