# [2769.Find the Maximum Achievable Number][title]

## Description
Given two integers, `num` and `t`. A **number** `x` is **achievable** if it can become equal to `num` after applying the following operation **at most** `t` times:

- Increase or decrease `x` by `1`, and simultaneously increase or decrease `num` by `1`.

Return the **maximum** possible value of `x`.

**Example 1:**

```
Input: num = 4, t = 1

Output: 6

Explanation:

Apply the following operation once to make the maximum achievable number equal to num:

Decrease the maximum achievable number by 1, and increase num by 1.
```

**Example 2:**

```
Input: num = 3, t = 2

Output: 7

Explanation:

Apply the following operation twice to make the maximum achievable number equal to num:

Decrease the maximum achievable number by 1, and increase num by 1.
```

## 结语

如果你同我一样热爱数据结构、算法、LeetCode，可以关注我 GitHub 上的 LeetCode 题解：[awesome-golang-algorithm][me]

[title]: https://leetcode.com/problems/find-the-maximum-achievable-number
[me]: https://github.com/kylesliu/awesome-golang-algorithm
