# [728.Self Dividing Numbers][title]

## Description
A **self-dividing number** is a number that is divisible by every digit it contains.

- For example, `128` is a **self-dividing number** because `128 % 1 == 0`, `128 % 2 == 0`, and `128 % 8 == 0`.

A **self-dividing number** is not allowed to contain the digit zero.

Given two integers `left` and `right`, return a list of all the **self-dividing numbers** in the range `[left, right]` (both **inclusive**).

**Example 1:**

```
Input: left = 1, right = 22
Output: [1,2,3,4,5,6,7,8,9,11,12,15,22]
```

**Example 2:**

```
Input: left = 47, right = 85
Output: [48,55,66,77]
```

## 结语

如果你同我一样热爱数据结构、算法、LeetCode，可以关注我 GitHub 上的 LeetCode 题解：[awesome-golang-algorithm][me]

[title]: https://leetcode.com/problems/self-dividing-numbers/
[me]: https://github.com/kylesliu/awesome-golang-algorithm
