# [1963.Minimum Number of Swaps to Make the String Balanced][title]

## Description
You are given a **0-indexed** string `s` of **even** length `n`. The string consists of **exactly** `n / 2` opening brackets `'['` and `n / 2` closing brackets `']'`.

A string is called **balanced** if and only if:

- It is the empty string, or
- It can be written as `AB`, where both `A` and `B` are **balanced** strings, or
- It can be written as `[C]`, where `C` is a **balanced** string.

You may swap the brackets at **any** two indices **any** number of times.

Return the **minimum** number of swaps to make `s` **balanced**.

**Example 1:**

```
Input: s = "][]["
Output: 1
Explanation: You can make the string balanced by swapping index 0 with index 3.
The resulting string is "[[]]".
```

**Example 2:**

```
Input: s = "]]][[["
Output: 2
Explanation: You can do the following to make the string balanced:
- Swap index 0 with index 4. s = "[]][][".
- Swap index 1 with index 5. s = "[[][]]".
The resulting string is "[[][]]".
```

**Example 3:**

```
Input: s = "[]"
Output: 0
Explanation: The string is already balanced.
```

## 结语

如果你同我一样热爱数据结构、算法、LeetCode，可以关注我 GitHub 上的 LeetCode 题解：[awesome-golang-algorithm][me]

[title]: https://leetcode.com/problems/minimum-number-of-swaps-to-make-the-string-balanced/
[me]: https://github.com/kylesliu/awesome-golang-algorithm
