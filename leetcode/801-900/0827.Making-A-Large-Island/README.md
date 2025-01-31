# [827.Making A Large Island][title]

## Description
You are given an `n x n` binary matrix `grid`. You are allowed to change **at most one** `0` to be `1`.

Return the size of the largest **island** in `grid` after applying this operation.

An **island** is a 4-directionally connected group of `1`s.

**Example 1:**

```
Input: grid = [[1,0],[0,1]]
Output: 3
Explanation: Change one 0 to 1 and connect two 1s, then we get an island with area = 3.
```

**Example 2:**

```
Input: grid = [[1,1],[1,0]]
Output: 4
Explanation: Change the 0 to 1 and make the island bigger, only one island with area = 4.
```

**Example 3:**

```
Input: grid = [[1,1],[1,1]]
Output: 4
Explanation: Can't change any 0 to 1, only one island with area = 4.
```

## 结语

如果你同我一样热爱数据结构、算法、LeetCode，可以关注我 GitHub 上的 LeetCode 题解：[awesome-golang-algorithm][me]

[title]: https://leetcode.com/problems/making-a-large-island/
[me]: https://github.com/kylesliu/awesome-golang-algorithm
