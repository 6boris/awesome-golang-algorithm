# [2395.Find Subarrays With Equal Sum][title]

## Description
Given a **0-indexed** integer array `nums`, determine whether there exist `two` subarrays of length `2` with **equal** sum. Note that the two subarrays must begin at **different** indices.

Return `true` if these subarrays exist, and `false` otherwise.

A **subarray** is a contiguous non-empty sequence of elements within an array.

**Example 1:**

```
Input: nums = [4,2,4]
Output: true
Explanation: The subarrays with elements [4,2] and [2,4] have the same sum of 6.
```

**Example 2:**

```
Input: nums = [1,2,3,4,5]
Output: false
Explanation: No two subarrays of size 2 have the same sum.
```

**Example 3:**

```
Input: nums = [0,0,0]
Output: true
Explanation: The subarrays [nums[0],nums[1]] and [nums[1],nums[2]] have the same sum of 0. 
Note that even though the subarrays have the same content, the two subarrays are considered different because they are in different positions in the original array.
```

## 结语

如果你同我一样热爱数据结构、算法、LeetCode，可以关注我 GitHub 上的 LeetCode 题解：[awesome-golang-algorithm][me]

[title]: https://leetcode.com/problems/find-subarrays-with-equal-sum/
[me]: https://github.com/kylesliu/awesome-golang-algorithm
