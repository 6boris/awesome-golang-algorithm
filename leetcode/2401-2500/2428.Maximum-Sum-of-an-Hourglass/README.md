# [2428.Maximum Sum of an Hourglass][title]

## Description
You are given an `m x n` integer matrix `grid`.

We define an **hourglass** as a part of the matrix with the following form:

![1](./img.jpeg)

Return the **maximum** sum of the elements of an hourglass.

**Note** that an hourglass cannot be rotated and must be entirely contained within the matrix.


**Example 1:**  

![1](./1.jpeg)

```
Input: grid = [[6,2,1,3],[4,2,1,5],[9,2,8,7],[4,1,2,9]]
Output: 30
Explanation: The cells shown above represent the hourglass with the maximum sum: 6 + 2 + 1 + 2 + 9 + 2 + 8 = 30.
```

**Example 2:**  

![2](./2.jpeg)

```
Input: grid = [[1,2,3],[4,5,6],[7,8,9]]
Output: 35
Explanation: There is only one hourglass in the matrix, with the sum: 1 + 2 + 3 + 5 + 7 + 8 + 9 = 35.
```

## 结语

如果你同我一样热爱数据结构、算法、LeetCode，可以关注我 GitHub 上的 LeetCode 题解：[awesome-golang-algorithm][me]

[title]: https://leetcode.com/problems/maximum-sum-of-an-hourglass/
[me]: https://github.com/kylesliu/awesome-golang-algorithm
