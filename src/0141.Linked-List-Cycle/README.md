# [1. Add Sum][title]

## Description

Given two binary strings, return their sum (also a binary string).

The input strings are both **non-empty** and contains only characters `1` or `0`.

**Example 1:**

```
Input: a = "11", b = "1"
Output: "100"
```

**Example 2:**

```
Input: a = "1010", b = "1011"
Output: "10101"
```

**Tags:** Math, String

## 题意
>判断链表是否有环

## 题解

### 思路1
> 骗人的方法： 设置一个时间1s或者0.5s,在这个时间内遍历链表，判断是否可以遍历完


```go

```

### 思路2
> HashMap： 利用HashMap的结构特点将链表节点地址存到里面，循环判断 时间复杂度可以到O(n * 1)
```go

```

### 思路3
> 快慢指针: 用2个指针，一个每次走一步，一个每次走2步看他们是否会相遇
```go

```

## 结语

如果你同我一样热爱数据结构、算法、LeetCode，可以关注我 GitHub 上的 LeetCode 题解：[awesome-golang-leetcode][me]

[title]: https://leetcode.com/problems/two-sum/description/
[me]: https://github.com/kylesliu/awesome-golang-leetcode
