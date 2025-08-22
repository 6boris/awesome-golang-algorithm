# [679.24 Game][title]

## Description
You are given an integer array `cards` of length `4`. You have four cards, each containing a number in the range `[1, 9]`. You should arrange the numbers on these cards in a mathematical expression using the operators `['+', '-', '*', '/']` and the parentheses `'('` and `')'` to get the value 24.

You are restricted with the following rules:

- The division operator `'/'` represents real division, not integer division.

    - For example, `4 / (1 - 2 / 3) = 4 / (1 / 3) = 12`.

- Every operation done is between two numbers. In particular, we cannot use `'-'` as a unary operator.

    - For example, if `cards = [1, 1, 1, 1]`, the expression `"-1 - 1 - 1 - 1"` is **not allowed**.

- You cannot concatenate numbers together

    - For example, if `cards = [1, 2, 1, 2]`, the expression `"12 + 12"` is not valid.

Return `true` if you can get such expression that evaluates to `24`, and `false` otherwise.

**Example 1:**

```
Input: cards = [4,1,8,7]
Output: true
Explanation: (8-4) * (7-1) = 24
```

**Example 2:**

```
Input: cards = [1,2,1,2]
Output: false
```

## 结语

如果你同我一样热爱数据结构、算法、LeetCode，可以关注我 GitHub 上的 LeetCode 题解：[awesome-golang-algorithm][me]

[title]: https://leetcode.com/problems/24-game/
[me]: https://github.com/kylesliu/awesome-golang-algorithm
