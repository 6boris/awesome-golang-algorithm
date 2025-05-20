# [3355.Zero Array Transformation I][title]

## Description
You are given an integer array `nums` of length n and a 2D array `queries`, where `queries[i] = [li, ri]`.

For each `queries[i]`:

- Select a subset of indices within the range `[li, ri]` in nums.
- Decrement the values at the selected indices by 1.

A **Zero Array** is an array where all elements are equal to 0.

Return `true` if it is possible to transform `nums` into a **Zero Array** after processing all the queries sequentially, otherwise return `false`.

**Example 1:**

```
Input: nums = [1,0,1], queries = [[0,2]]

Output: true

Explanation:

For i = 0:
Select the subset of indices as [0, 2] and decrement the values at these indices by 1.
The array will become [0, 0, 0], which is a Zero Array.
```

**Example 2:**

```
Input: nums = [4,3,2,1], queries = [[1,3],[0,2]]

Output: false

Explanation:

For i = 0:
Select the subset of indices as [1, 2, 3] and decrement the values at these indices by 1.
The array will become [4, 2, 1, 0].
For i = 1:
Select the subset of indices as [0, 1, 2] and decrement the values at these indices by 1.
The array will become [3, 1, 0, 0], which is not a Zero Array.
```

## 结语

如果你同我一样热爱数据结构、算法、LeetCode，可以关注我 GitHub 上的 LeetCode 题解：[awesome-golang-algorithm][me]

[title]: https://leetcode.com/problems/zero-array-transformation-i/
[me]: https://github.com/kylesliu/awesome-golang-algorithm
