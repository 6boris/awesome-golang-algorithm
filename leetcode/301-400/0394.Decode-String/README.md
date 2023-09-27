# [394.Decode String][title]

## Description
Given an encoded string, return its decoded string.

The encoding rule is: `k[encoded_string]`, where the `encoded_string` inside the square brackets is being repeated exactly `k` times. Note that k is guaranteed to be a positive integer.

You may assume that the input string is always valid; there are no extra white spaces, square brackets are well-formed, etc. Furthermore, you may assume that the original data does not contain any digits and that digits are only for those repeat numbers, k. For example, there will not be input like `3a` or `2[4]`.

The test cases are generated so that the length of the output will never exceed `10^5`.

**Example 1:**

```
Input: s = "3[a]2[bc]"
Output: "aaabcbc"
```

**Example 2:**

```
Input: s = "3[a2[c]]"
Output: "accaccacc"
```

**Example 3:**

```
Input: s = "2[abc]3[cd]ef"
Output: "abcabccdcdcdef"
```

## 结语

如果你同我一样热爱数据结构、算法、LeetCode，可以关注我 GitHub 上的 LeetCode 题解：[awesome-golang-algorithm][me]

[title]: https://leetcode.com/problems/decode-string/
[me]: https://github.com/kylesliu/awesome-golang-algorithm
