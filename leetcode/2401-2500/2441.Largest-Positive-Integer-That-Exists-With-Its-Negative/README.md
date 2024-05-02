# [2441.Largest Positive Integer That Exists With Its Negative][title]

## Description
Given an integer array nums that **does not contain** any zeros, find **the largest positive** integer `k` such that `-k` also exists in the array.

Return the positive integer `k`. If there is no such integer, return `-1`.

**Example 1:**

```
Input: nums = [-1,2,-3,3]
Output: 3
Explanation: 3 is the only valid k we can find in the array.
```

**Exists 2:**

```
Input: nums = [-1,10,6,7,-7,1]
Output: 7
Explanation: Both 1 and 7 have their corresponding negative values in the array. 7 has a larger value.
```

**Exists 3:**

```
Input: nums = [-10,8,6,7,-2,-3]
Output: -1
Explanation: There is no a single valid k, we return -1.
```

## 结语

如果你同我一样热爱数据结构、算法、LeetCode，可以关注我 GitHub 上的 LeetCode 题解：[awesome-golang-algorithm][me]

[title]: https://leetcode.com/problems/largest-positive-integer-that-exists-with-its-negative/
[me]: https://github.com/kylesliu/awesome-golang-algorithm
