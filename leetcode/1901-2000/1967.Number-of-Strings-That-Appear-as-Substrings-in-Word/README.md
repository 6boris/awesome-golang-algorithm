# [1967.Number of Strings That Appear as Substrings in Word][title]

## Description
Given an array of strings `patterns` and a string `word`, return the **number** of strings in `patterns` that exist as a **substring** in `word`.

A **substring** is a contiguous sequence of characters within a string.

**Example 1:**

```
Input: patterns = ["a","abc","bc","d"], word = "abc"
Output: 3
Explanation:
- "a" appears as a substring in "abc".
- "abc" appears as a substring in "abc".
- "bc" appears as a substring in "abc".
- "d" does not appear as a substring in "abc".
3 of the strings in patterns appear as a substring in word.
```

**Example 2:**

```
Input: patterns = ["a","b","c"], word = "aaaaabbbbb"
Output: 2
Explanation:
- "a" appears as a substring in "aaaaabbbbb".
- "b" appears as a substring in "aaaaabbbbb".
- "c" does not appear as a substring in "aaaaabbbbb".
2 of the strings in patterns appear as a substring in word.
```

**Example 3:**

```
Input: patterns = ["a","a","a"], word = "ab"
Output: 3
Explanation: Each of the patterns appears as a substring in word "ab".
```

## 结语

如果你同我一样热爱数据结构、算法、LeetCode，可以关注我 GitHub 上的 LeetCode 题解：[awesome-golang-algorithm][me]

[title]: https://leetcode.com/problems/number-of-strings-that-appear-as-substrings-in-word/
[me]: https://github.com/kylesliu/awesome-golang-algorithm
