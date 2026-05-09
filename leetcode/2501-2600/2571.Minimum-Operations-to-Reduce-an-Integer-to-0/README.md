# [2571.Minimum Operations to Reduce an Integer to 0][title]

## Description
You are given a positive integer `n`, you can do the following operation **any** number of times:

- Add or subtract a **power** of `2` from `n`.

Return the **minimum** number of operations to make `n` equal to `0`.

A number `x` is power of `2` if `x == 2^i` where `i >= 0`.

**Example 1:**

```
Input: n = 39
Output: 3
Explanation: We can do the following operations:
- Add 20 = 1 to n, so now n = 40.
- Subtract 23 = 8 from n, so now n = 32.
- Subtract 25 = 32 from n, so now n = 0.
It can be shown that 3 is the minimum numb
```

**Example 2:**

```
Input: n = 54
Output: 3
Explanation: We can do the following operations:
- Add 21 = 2 to n, so now n = 56.
- Add 23 = 8 to n, so now n = 64.
- Subtract 26 = 64 from n, so now n = 0.
So the minimum number of operations is 3.
```

## 结语

如果你同我一样热爱数据结构、算法、LeetCode，可以关注我 GitHub 上的 LeetCode 题解：[awesome-golang-algorithm][me]

[title]: https://leetcode.com/problems/minimum-operations-to-reduce-an-integer-to-0
[me]: https://github.com/kylesliu/awesome-golang-algorithm
