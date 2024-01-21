# [522.Longest Uncommon Subsequence II][title]

## Description
Given an array of strings `strs`, return the length of the **longest uncommon subsequence** between them. If the longest uncommon subsequence does not exist, return `-1`.

An **uncommon subsequence** between an array of strings is a string that is a **subsequence of one string but not the others**.

A **subsequence** of a string `s` is a string that can be obtained after deleting any number of characters from `s`.

- For example, `"abc"` is a subsequence of `"aebdc"` because you can delete the underlined characters in `"aebdc"` to get `"abc"`. Other subsequences of `"aebdc"` include `"aebdc"`, `"aeb"`, and `""` (empty string).

**Example 1:**

```
Input: strs = ["aba","cdc","eae"]
Output: 3
```

**Example 2:**

```
Input: strs = ["aaa","aaa","aa"]
Output: -1
```

## 结语

如果你同我一样热爱数据结构、算法、LeetCode，可以关注我 GitHub 上的 LeetCode 题解：[awesome-golang-algorithm][me]

[title]: https://leetcode.com/problems/longest-uncommon-subsequence-ii/
[me]: https://github.com/kylesliu/awesome-golang-algorithm
