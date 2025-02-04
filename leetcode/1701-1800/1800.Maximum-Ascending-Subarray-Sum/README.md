# [1800.Maximum Ascending Subarray Sum][title]

## Description
Given an array of positive integers `nums`, return the maximum possible sum of an **ascending** subarray in `nums`.

A subarray is defined as a contiguous sequence of numbers in an array.

A subarray `[numsl, numsl+1, ..., numsr-1, numsr]` is **ascending** if for all `i` where `l <= i < r`, `numsi  < numsi+1`. Note that a subarray of size 1 is **ascending**.

**Example 1:**

```
Input: nums = [10,20,30,5,10,50]
Output: 65
Explanation: [5,10,50] is the ascending subarray with the maximum sum of 65.
```

**Example 2:**

```
Input: nums = [10,20,30,40,50]
Output: 150
Explanation: [10,20,30,40,50] is the ascending subarray with the maximum sum of 150.
```

**Example 3:**

```
Input: nums = [12,17,15,13,10,11,12]
Output: 33
Explanation: [10,11,12] is the ascending subarray with the maximum sum of 33.
```

## 结语

如果你同我一样热爱数据结构、算法、LeetCode，可以关注我 GitHub 上的 LeetCode 题解：[awesome-golang-algorithm][me]

[title]: https://leetcode.com/problems/maximum-ascending-subarray-sum/
[me]: https://github.com/kylesliu/awesome-golang-algorithm
