# [3737.Count Subarrays With Majority Element I][title]

## Description
You are given an integer array `nums` and an integer `target`.

Return the number of subarrays of `nums` in which `target` is the **majority element**.

The **majority element** of a subarray is the element that appears **strictly more than half** of the times in that subarray.

**Example 1:**

```
Input: nums = [1,2,2,3], target = 2

Output: 5

Explanation:

Valid subarrays with target = 2 as the majority element:

nums[1..1] = [2]
nums[2..2] = [2]
nums[1..2] = [2,2]
nums[0..2] = [1,2,2]
nums[1..3] = [2,2,3]
So there are 5 such subarrays.
```

**Example 2:**

```
Input: nums = [1,1,1,1], target = 1

Output: 10

Explanation:

10 subarrays have 1 as the majority element.
```

**Example 3:**

```
Input: nums = [1,2,3], target = 4

Output: 0

Explanation:

target = 4 does not appear in nums at all. Therefore, there cannot be any subarray where 4 is the majority element. Hence the answer is 0.
```

## 结语

如果你同我一样热爱数据结构、算法、LeetCode，可以关注我 GitHub 上的 LeetCode 题解：[awesome-golang-algorithm][me]

[title]: https://leetcode.com/problems/count-subarrays-with-majority-element-i/
[me]: https://github.com/kylesliu/awesome-golang-algorithm
