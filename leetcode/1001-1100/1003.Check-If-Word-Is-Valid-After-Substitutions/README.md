# [1003.Check If Word Is Valid After Substitutions][title]

## Description
Given a string `s`, determine if it is **valid**.

A string `s` is **valid** if, starting with an empty string `t = ""`, you can **transform** `s` **into** `s` after performing the following operation **any number of times**:

- Insert string `"abc"` into any position in `t`. More formally, `t` becomes t<sub>left</sub> + `"abc"` + t<sub>right</sub>, where t == t<sub>left</sub> + t<sub>right</sub>. Note that t<sub>left</sub> and t<sub>right</sub> may be **empty**.

Return `true` if `s` is a **valid** string, otherwise, return `false`.

**Example 1:**

```
Input: s = "aabcbc"
Output: true
Explanation:
"" -> "abc" -> "aabcbc"
Thus, "aabcbc" is valid.
```

**Example 2:**

```
Input: s = "abcabcababcc"
Output: true
Explanation:
"" -> "abc" -> "abcabc" -> "abcabcabc" -> "abcabcababcc"
Thus, "abcabcababcc" is valid.
```

**Example 3:**

```
Input: s = "abccba"
Output: false
Explanation: It is impossible to get "abccba" using the operation.
```

## 结语

如果你同我一样热爱数据结构、算法、LeetCode，可以关注我 GitHub 上的 LeetCode 题解：[awesome-golang-algorithm][me]

[title]: https://leetcode.com/problems/check-if-word-is-valid-after-substitutions/
[me]: https://github.com/kylesliu/awesome-golang-algorithm
