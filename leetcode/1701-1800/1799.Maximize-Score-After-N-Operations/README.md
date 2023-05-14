# [1799.Maximize Score After N Operations][title]

## Description

You are given `nums`, an array of positive integers of size `2 * n`. You must perform `n` operations on this array.

In the i<sup>th</sup> operation **(1-indexed)**, you will:

- Choose two elements, `x` and `y`.
- Receive a score of `i * gcd(x, y)`.
- Remove `x` and `y` from `nums`.

Return the maximum score you can receive after performing `n` operations.

The function `gcd(x, y)` is the greatest common divisor of `x` and `y`.

**Example 1:**

```
Input: nums = [1,2]
Output: 1
Explanation: The optimal choice of operations is:
(1 * gcd(1, 2)) = 1
```

**Example 2:**

```
Input: nums = [3,4,6,8]
Output: 11
Explanation: The optimal choice of operations is:
(1 * gcd(3, 6)) + (2 * gcd(4, 8)) = 3 + 8 = 11
```

**Example 3:**

```
Input: nums = [1,2,3,4,5,6]
Output: 14
Explanation: The optimal choice of operations is:
(1 * gcd(1, 5)) + (2 * gcd(2, 4)) + (3 * gcd(3, 6)) = 1 + 4 + 9 = 14
```

## 结语

如果你同我一样热爱数据结构、算法、LeetCode，可以关注我 GitHub 上的 LeetCode 题解：[awesome-golang-algorithm][me]

[title]: https://leetcode.com/problems/maximize-score-after-n-operations/
[me]: https://github.com/kylesliu/awesome-golang-algorithm
