# [907.Sum of Subarray Minimums][title]

## Description
Given an array of integers arr, find the sum of `min(b)`, where `b` ranges over every (contiguous) subarray of `arr`. Since the answer may be large, return the answer **modulo** 10<sup>9</sup> + 7.

**Example 1:**

```
Input: arr = [3,1,2,4]
Output: 17
Explanation: 
Subarrays are [3], [1], [2], [4], [3,1], [1,2], [2,4], [3,1,2], [1,2,4], [3,1,2,4]. 
Minimums are 3, 1, 2, 4, 1, 1, 2, 1, 1, 1.
Sum is 17.
```

**Example 2:**

```
Input: arr = [11,81,94,43,3]
Output: 444
```

## 结语

如果你同我一样热爱数据结构、算法、LeetCode，可以关注我 GitHub 上的 LeetCode 题解：[awesome-golang-algorithm][me]

[title]: https://leetcode.com/problems/sum-of-subarray-minimums/
[me]: https://github.com/kylesliu/awesome-golang-algorithm
