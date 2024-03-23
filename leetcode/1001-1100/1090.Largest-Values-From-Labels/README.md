# [1090.Largest Values From Labels][title]

## Description
There is a set of `n` items. You are given two integer arrays `values` and `values` where the value and the label of the i<sup>th</sup> element are `values[i]` and `labels[i]` respectively. You are also given two integers `numWanted` and `useLimit`.

Choose a subset `s` of the `n` elements such that:

- The size of the subset `s` is **less than or equal to** `numWanted`.
- There are **at most** `useLimit` items with the same label in `s`.

The **score** of a subset is the sum of the values in the subset.

Return the maximum **score** of a subset `s`.

**Example 1:**

```
Input: values = [5,4,3,2,1], labels = [1,1,2,2,3], numWanted = 3, useLimit = 1
Output: 9
Explanation: The subset chosen is the first, third, and fifth items.
```

**Example 2:**

```
Input: values = [5,4,3,2,1], labels = [1,3,3,3,2], numWanted = 3, useLimit = 2
Output: 12
Explanation: The subset chosen is the first, second, and third items.
```

**Example 3:**

```
Input: values = [9,8,8,7,6], labels = [0,0,0,1,1], numWanted = 3, useLimit = 1
Output: 16
Explanation: The subset chosen is the first and fourth items.
```

## 结语

如果你同我一样热爱数据结构、算法、LeetCode，可以关注我 GitHub 上的 LeetCode 题解：[awesome-golang-algorithm][me]

[title]: https://leetcode.com/problems/largest-values-from-labels/
[me]: https://github.com/kylesliu/awesome-golang-algorithm
