# [2563.Count the Number of Fair Pairs][title]

## Description
Given a **0-indexed** integer array `nums` of size `n` and two integers `lower` and `upper`, return the number of fair pairs.

A pair `(i, j)` is **fair** if:

- `0 <= i < j < n`, and
- `lower <= nums[i] + nums[j] <= upper`


**Example 1:**

```
Input: nums = [0,1,7,4,4,5], lower = 3, upper = 6
Output: 6
Explanation: There are 6 fair pairs: (0,3), (0,4), (0,5), (1,3), (1,4), and (1,5).
```

**Example 2:**

```
Input: nums = [1,7,9,2,5], lower = 11, upper = 11
Output: 1
Explanation: There is a single fair pair: (2,3).
```

## 结语

如果你同我一样热爱数据结构、算法、LeetCode，可以关注我 GitHub 上的 LeetCode 题解：[awesome-golang-algorithm][me]

[title]: https://leetcode.com/problems/count-the-number-of-fair-pairs
[me]: https://github.com/kylesliu/awesome-golang-algorithm
