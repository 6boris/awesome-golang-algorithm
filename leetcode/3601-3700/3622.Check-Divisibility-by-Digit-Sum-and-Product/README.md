# [3622.Check Divisibility by Digit Sum and Product][title]

## Description
You are given a positive integer `n`. Determine whether `n` is divisible by the **sum** of the following two values:

- The **digit sum** of `n` (the sum of its digits).
- The **digit product** of `n` (the product of its digits).

Return `true` if n is divisible by this sum; otherwise, return `false`.

**Example 1:**

```
Input: n = 99

Output: true

Explanation:

Since 99 is divisible by the sum (9 + 9 = 18) plus product (9 * 9 = 81) of its digits (total 99), the output is true.
```

**Example 2:**

```
Input: n = 23

Output: false

Explanation:

Since 23 is not divisible by the sum (2 + 3 = 5) plus product (2 * 3 = 6) of its digits (total 11), the output is false.
```

## 结语

如果你同我一样热爱数据结构、算法、LeetCode，可以关注我 GitHub 上的 LeetCode 题解：[awesome-golang-algorithm][me]

[title]: https://leetcode.com/problems/check-divisibility-by-digit-sum-and-product/
[me]: https://github.com/kylesliu/awesome-golang-algorithm
