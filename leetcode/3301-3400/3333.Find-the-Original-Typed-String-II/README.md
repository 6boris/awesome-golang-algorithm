# [3333.Find the Original Typed String II][title]

## Description
Alice is attempting to type a specific string on her computer. However, she tends to be clumsy and **may** press a key for too long, resulting in a character being typed **multiple** times.

You are given a string `word`, which represents the **final** output displayed on Alice's screen. You are also given a **positive** integer `k`.

Return the total number of possible original strings that Alice might have intended to type, if she was trying to type a string of size **at least** `k`.

Since the answer may be very large, return it **modulo** `10^9 + 7`.

**Example 1:**

```
Input: word = "aabbccdd", k = 7

Output: 5

Explanation:

The possible strings are: "aabbccdd", "aabbccd", "aabbcdd", "aabccdd", and "abbccdd".
```

**Example 2:**

```
Input: word = "aabbccdd", k = 8

Output: 1

Explanation:

The only possible string is "aabbccdd".
```

**Example 3:**

```
Input: word = "aaabbb", k = 3

Output: 8
```

## 结语

如果你同我一样热爱数据结构、算法、LeetCode，可以关注我 GitHub 上的 LeetCode 题解：[awesome-golang-algorithm][me]

[title]: https://leetcode.com/problems/find-the-original-typed-string-ii/
[me]: https://github.com/kylesliu/awesome-golang-algorithm
