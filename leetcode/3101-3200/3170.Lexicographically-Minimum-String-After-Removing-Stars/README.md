# [3170.Lexicographically Minimum String After Removing Stars][title]

## Description
You are given a string `s`. It may contain any number of `'*'` characters. Your task is to remove all `'*'` characters.

While there is a `'*'`, do the following operation:

- Delete the leftmost `'*'` and the **smallest** non-`'*'` character to its left. If there are several smallest characters, you can delete any of them.

Return the **lexicographically smallest** resulting string after removing all `'*'` characters.

**Example 1:**

```
Input: s = "aaba*"

Output: "aab"

Explanation:

We should delete one of the 'a' characters with '*'. If we choose s[3], s becomes the lexicographically smallest.
```

**Example 2:**

```
Input: s = "abc"

Output: "abc"

Explanation:

There is no '*' in the string.
```

## 结语

如果你同我一样热爱数据结构、算法、LeetCode，可以关注我 GitHub 上的 LeetCode 题解：[awesome-golang-algorithm][me]

[title]: https://leetcode.com/problems/lexicographically-minimum-string-after-removing-stars/
[me]: https://github.com/kylesliu/awesome-golang-algorithm
