# [378.Kth Smallest Element in a Sorted Matrix][title]

## Description
Given an `n x n` `matrix` where each of the rows and columns is sorted in ascending order, return the k<sup>th</sup> smallest element in the matrix.

Note that it is the k<sup>th</sup> smallest element **in the sorted order**, not the k<sup>th</sup> **distinct** element.

You must find a solution with a memory complexity better than `O(n^2)`.

**Example 1:**

```
Input: matrix = [[1,5,9],[10,11,13],[12,13,15]], k = 8
Output: 13
Explanation: The elements in the matrix are [1,5,9,10,11,12,13,13,15], and the 8th smallest number is 13
```

**Example 2:**

```
Input: matrix = [[-5]], k = 1
Output: -5
```

## 结语

如果你同我一样热爱数据结构、算法、LeetCode，可以关注我 GitHub 上的 LeetCode 题解：[awesome-golang-algorithm][me]

[title]: https://leetcode.com/problems/kth-smallest-element-in-a-sorted-matrix/
[me]: https://github.com/kylesliu/awesome-golang-algorithm
