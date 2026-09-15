# [2472.Maximum Number of Non-overlapping Palindrome Substrings][title]

## Description
You are given a string `s` and a **positive** integer `k`.

Select a set of **non-overlapping** substrings from the string `s` that satisfy the following conditions:

- The **length** of each substring is **at least** `k`.
- Each substring is a **palindrome**.

Return the **maximum** number of substrings in an optimal selection.

A **substring** is a contiguous sequence of characters within a string.

**Example 1:**

```
Input: s = "abaccdbbd", k = 3
Output: 2
Explanation: We can select the substrings underlined in s = "abaccdbbd". Both "aba" and "dbbd" are palindromes and have a length of at least k = 3.
It can be shown that we cannot find a selection with more than two valid substrings.
```

**Example 2:**

```
Input: s = "adbcda", k = 2
Output: 0
Explanation: There is no palindrome substring of length at least 2 in the string.
```

## 结语

如果你同我一样热爱数据结构、算法、LeetCode，可以关注我 GitHub 上的 LeetCode 题解：[awesome-golang-algorithm][me]

[title]: https://leetcode.com/problems/maximum-number-of-non-overlapping-palindrome-substrings/
[me]: https://github.com/kylesliu/awesome-golang-algorithm
