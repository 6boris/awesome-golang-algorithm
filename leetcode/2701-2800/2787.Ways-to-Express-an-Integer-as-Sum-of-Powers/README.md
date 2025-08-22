# [2787.Ways to Express an Integer as Sum of Powers][title]

## Description
Given two **positive** integers `n` and `n`.

Return the number of ways `n` can be expressed as the sum of the `xth` power of **unique** positive integers, in other words, the number of sets of unique integers `[n1, n2, ..., nk]` where `n = n1x + n2x + ... + nkx`.

Since the result can be very large, return it modulo `10^9 + 7`.

For example, if `n = 160` and `x = 3`, one way to express `n` is `n = 2^3 + 3^3 + 5^3`.

**Example 1:**

```
Input: n = 10, x = 2
Output: 1
Explanation: We can express n as the following: n = 32 + 12 = 10.
It can be shown that it is the only way to express 10 as the sum of the 2nd power of unique integers.
```

**EXample 2:**

```
Input: n = 4, x = 1
Output: 2
Explanation: We can express n in the following ways:
- n = 41 = 4.
- n = 31 + 11 = 4.
```

## 结语

如果你同我一样热爱数据结构、算法、LeetCode，可以关注我 GitHub 上的 LeetCode 题解：[awesome-golang-algorithm][me]

[title]: https://leetcode.com/problems/ways-to-express-an-integer-as-sum-of-powers
[me]: https://github.com/kylesliu/awesome-golang-algorithm
