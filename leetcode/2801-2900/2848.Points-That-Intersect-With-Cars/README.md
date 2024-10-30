# [2848.Points That Intersect With Cars][title]

## Description
You are given a **0-indexed** 2D integer array `nums` representing the coordinates of the cars parking on a number line. For any index `i`, `nums[i] = [starti, endi]` where `starti` is the starting point of the `ith` car and `endi` is the ending point of the `ith` car.

Return the number of integer points on the line that are covered with **any part** of a car.

**Example 1:**

```
Input: nums = [[3,6],[1,5],[4,7]]
Output: 7
Explanation: All the points from 1 to 7 intersect at least one car, therefore the answer would be 7.
```

**Example 2:**

```
Input: nums = [[1,3],[5,8]]
Output: 7
Explanation: Points intersecting at least one car are 1, 2, 3, 5, 6, 7, 8. There are a total of 7 points, therefore the answer would be 7.
```

## 结语

如果你同我一样热爱数据结构、算法、LeetCode，可以关注我 GitHub 上的 LeetCode 题解：[awesome-golang-algorithm][me]

[title]: https://leetcode.com/problems/points-that-intersect-with-cars/
[me]: https://github.com/kylesliu/awesome-golang-algorithm
