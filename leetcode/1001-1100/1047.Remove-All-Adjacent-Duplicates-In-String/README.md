# [1047.Remove All Adjacent Duplicates In String][title]

## Description
You are given a string `s` consisting of lowercase English letters. A **duplicate removal** consists of choosing two **adjacent** and **equal** letters and removing them.

We repeatedly make **duplicate removals** on `s` until we no longer can.

Return the final string after all such duplicate removals have been made. It can be proven that the answer is **unique**.



**Example 1:**

```
Input: s = "abbaca"
Output: "ca"
Explanation: 
For example, in "abbaca" we could remove "bb" since the letters are adjacent and equal, and this is the only possible move.  The result of this move is that the string is "aaca", of which only "aa" is possible, so the final string is "ca".
```


**Example 2:**

```
Input: s = "azxxzy"
Output: "ay"
```

## 结语

如果你同我一样热爱数据结构、算法、LeetCode，可以关注我 GitHub 上的 LeetCode 题解：[awesome-golang-algorithm][me]

[title]: https://leetcode.com/problems/remove-all-adjacent-duplicates-in-string/
[me]: https://github.com/kylesliu/awesome-golang-algorithm
