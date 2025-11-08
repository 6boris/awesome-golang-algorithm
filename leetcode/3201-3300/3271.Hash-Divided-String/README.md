# [3271.Hash Divided String][title]

## Description
You are given a string `s` of length `n` and an integer `k`, where `n` is a **multiple** of `k`. Your task is to hash the string `s` into a new string called `result`, which has a length of `n / k`.

First, divide `s` into `n / k` substrings, each with a length of `k`. Then, initialize `result` as an **empty** string.

For each **substring** in order from the beginning:

- The **hash value** of a character is the index of that character in the **English alphabet** (e.g., `'a' → 0, 'b' → 1, ..., 'z' → 25`).
- Calculate the sum of all the **hash values** of the characters in the substring.
- Find the remainder of this sum when divided by 26, which is called `hashedChar`.
- Identify the character in the English lowercase alphabet that corresponds to `hashedChar`.
- Append that character to the end of `result`.

Return `result`.

**Example 1:**

```
Input: s = "abcd", k = 2

Output: "bf"

Explanation:

First substring: "ab", 0 + 1 = 1, 1 % 26 = 1, result[0] = 'b'.

Second substring: "cd", 2 + 3 = 5, 5 % 26 = 5, result[1] = 'f'.
```

**Example 2:**

```
Input: s = "mxz", k = 3

Output: "i"

Explanation:

The only substring: "mxz", 12 + 23 + 25 = 60, 60 % 26 = 8, result[0] = 'i'.
```

## 结语

如果你同我一样热爱数据结构、算法、LeetCode，可以关注我 GitHub 上的 LeetCode 题解：[awesome-golang-algorithm][me]

[title]: https://leetcode.com/problems/hash-divided-string/
[me]: https://github.com/kylesliu/awesome-golang-algorithm
