# [1209.Remove All Adjacent Duplicates in String II][title]

## Description
You are given a string `s` and an integer `k`, a `k` **duplicate removal** consists of choosing `k` adjacent and equal letters from `s` and removing them, causing the left and the right side of the deleted substring to concatenate together.

We repeatedly make `k` **duplicate removals** on `s` until we no longer can.

Return the final string after all such duplicate removals have been made. It is guaranteed that the answer is **unique**.

**Example 1:**

```
Input: s = "abcd", k = 2
Output: "abcd"
Explanation: There's nothing to delete.
```

**Example 2:**

```
Input: s = "deeedbbcccbdaa", k = 3
Output: "aa"
Explanation: 
First delete "eee" and "ccc", get "ddbbbdaa"
Then delete "bbb", get "dddaa"
Finally delete "ddd", get "aa"
```

**Example 3:**

```
Input: s = "pbbcggttciiippooaais", k = 2
Output: "ps"
```

## 结语

如果你同我一样热爱数据结构、算法、LeetCode，可以关注我 GitHub 上的 LeetCode 题解：[awesome-golang-algorithm][me]

[title]: https://leetcode.com/problems/remove-all-adjacent-duplicates-in-string-ii/
[me]: https://github.com/kylesliu/awesome-golang-algorithm
