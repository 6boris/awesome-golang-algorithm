# [3407.Substring Matching Pattern][title]

## Description
You are given a string `s` and a pattern string `p`, where `p` contains **exactly one** `'*'` character.

The `'*'` in p can be replaced with any sequence of zero or more characters.

Return `true` if p can be made a substring of `s`, and `false` otherwise.

**Example 1:**

```
Input: s = "leetcode", p = "ee*e"

Output: true

Explanation:

By replacing the '*' with "tcod", the substring "eetcode" matches the pattern.
```

**Example 2:**

```
Input: s = "car", p = "c*v"

Output: false

Explanation:

There is no substring matching the pattern.
```

**Example 3:**

```
Input: s = "luck", p = "u*"

Output: true

Explanation:

The substrings "u", "uc", and "uck" match the pattern.
```

## 结语

如果你同我一样热爱数据结构、算法、LeetCode，可以关注我 GitHub 上的 LeetCode 题解：[awesome-golang-algorithm][me]

[title]: https://leetcode.com/problems/substring-matching-pattern/
[me]: https://github.com/kylesliu/awesome-golang-algorithm
