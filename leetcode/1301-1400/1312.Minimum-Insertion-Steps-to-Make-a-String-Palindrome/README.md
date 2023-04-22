# [1312.Minimum Insertion Steps to Make a String Palindrome][title]

## Description
Given a string `s`. In one step you can insert any character at any index of the string.

Return the minimum number of steps to make `s` palindrome.

A **Palindrome String** is one that reads the same backward as well as forward.

**Example 1:**

```
Input: s = "zzazz"
Output: 0
Explanation: The string "zzazz" is already palindrome we do not need any insertions.
```


**Example 2:**

```
Input: s = "mbadm"
Output: 2
Explanation: String can be "mbdadbm" or "mdbabdm".
```

**Example 3:**

```
Input: s = "leetcode"
Output: 5
Explanation: Inserting 5 characters the string becomes "leetcodocteel".
```

### 思路1

> 找出最长公共子序列，不同的地方，插入即可


## 结语

如果你同我一样热爱数据结构、算法、LeetCode，可以关注我 GitHub 上的 LeetCode 题解：[awesome-golang-algorithm][me]

[title]: https://leetcode.com/problems/minimum-insertion-steps-to-make-a-string-palindrome/
[me]: https://github.com/kylesliu/awesome-golang-algorithm
