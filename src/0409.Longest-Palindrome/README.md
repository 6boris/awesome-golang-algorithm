# [409. Longest Palindrome][title]

## Description

Given a string which consists of lowercase or uppercase letters, find the length of the longest palindromes that can be built with those letters.

This is case sensitive, for example "Aa" is not considered a palindrome here.

Note:
Assume the length of given string will not exceed 1,010.

**Example 1:**

```
Input:
"abccccdd"

Output:
7

Explanation:
One longest palindrome that can be built is "dccaccd", whose length is 7.
```

**Tags:** Math, String

## 题意
> 求2数之和

## 题解

### 思路1
> 用哈希表统计每个字符出现的次数
  判断奇偶性：
  若字符出现次数为偶数，则肯定能够组成回文串，计入累加器
  若为奇数，添加个数-1并计入累加器，并且标记存在中心字符

```go

```

### 思路2
> 思路2
```go

```

## 结语

如果你同我一样热爱数据结构、算法、LeetCode，可以关注我 GitHub 上的 LeetCode 题解：[awesome-golang-leetcode][me]

[title]: https://leetcode.com/problems/longest-palindrome/
[me]: https://github.com/kylesliu/awesome-golang-leetcode
