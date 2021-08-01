# [915.Partition Array into Disjoint Intervals][title]


## Description
Given an array `nums`, partition it into two (contiguous) subarrays `left` and `right` so that:

 - Every element in `left` is less than or equal to every element in `right`.
 - `left` and `right` are non-empty.
 - `left` has the smallest possible size.

Return the `length` of `left` after such a partitioning. It is guaranteed that such a partitioning exists.

**Example 1:**

```
Input: nums = [5,0,3,8,6]
Output: 3
Explanation: left = [5,0,3], right = [8,6]
```

__Example 2:__

```
Input: nums = [1,1,1,0,6,12]
Output: 4
Explanation: left = [1,1,1,0], right = [6,12]
```

## 结语

如果你同我一样热爱数据结构、算法、LeetCode，可以关注我 GitHub 上的 LeetCode 题解：[awesome-golang-leetcode][me]

[title]: https://leetcode.com/problems/partition-array-into-disjoint-intervals/
[me]: https://github.com/kylesliu/awesome-golang-algorithm
