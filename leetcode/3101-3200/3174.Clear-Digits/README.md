# [3174.Clear Digits][title]

## Description
You are given a string `s`.

Your task is to remove **all** digits by doing this operation repeatedly:

- Delete the first digit and the **closest non-digit** character to its left.

Return the resulting string after removing all digits.

**Example 1:**

```
Input: s = "abc"

Output: "abc"

Explanation:

There is no digit in the string.
```

**Example 2:**

```
Input: s = "cb34"

Output: ""

Explanation:

First, we apply the operation on s[2], and s becomes "c4".

Then we apply the operation on s[1], and s becomes "".
```

## 结语

如果你同我一样热爱数据结构、算法、LeetCode，可以关注我 GitHub 上的 LeetCode 题解：[awesome-golang-algorithm][me]

[title]: https://leetcode.com/problems/clear-digits/
[me]: https://github.com/kylesliu/awesome-golang-algorithm
