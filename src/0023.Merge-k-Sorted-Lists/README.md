# [23. Merge k Sorted Lists]

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
>合并 k 个排序链表，返回合并后的排序链表

## 题解

### 思路1
>  蛮力解法：
- 1.遍历所有链表将值存到数组中
- 2.对数组进行排序
- 3.将排序好的数组转化为链表

> 复杂度：
- 时间复杂度
    - 将链表转化为数组需要O(N)的时间
    - 对数组排序需要O(N \log_2 N)的时间
    - 将数组转化为链表需要O(N)的时间
- 空间复杂度
    - 排序需要O(n)的空间
    - 数组转化为链表需要O(N)的空间

```go

```

### 思路2
> 思路2
```go

```

## 结语

如果你同我一样热爱数据结构、算法、LeetCode，可以关注我 GitHub 上的 LeetCode 题解：[awesome-golang-leetcode][me]

[title]: https://leetcode.com/problems/merge-k-sorted-lists/description/
[me]: https://github.com/kylesliu/awesome-golang-leetcode
