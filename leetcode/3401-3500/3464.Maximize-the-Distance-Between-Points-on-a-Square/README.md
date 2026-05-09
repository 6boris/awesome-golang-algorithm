# [3464.Maximize the Distance Between Points on a Square][title]

## Description
You are given an integer `side`, representing the edge length of a square with corners at `(0, 0)`, `(0, side)`, `(side, 0)`, and `(side, side)` on a Cartesian plane.

You are also given a **positive** integer `k` and a 2D integer array `points`, where `points[i] = [xi, yi]` represents the coordinate of a point lying on the **boundary** of the square.

You need to select `k` elements among `points` such that the minimum Manhattan distance between any two points is **maximized**.

Return the **maximized** possible minimum Manhattan distance between the selected `k` points.

The Manhattan Distance between two cells `(xi, yi)` and `(xj, yj)` is `|xi - xj| + |yi - yj|`.

**Example 1:**  

![1](./1.png)

```
Input: side = 2, points = [[0,2],[2,0],[2,2],[0,0]], k = 4

Output: 2

Explanation:

Select all four points.
```

**Example 2:**  

![2](./2.png)

```
Input: side = 2, points = [[0,0],[1,2],[2,0],[2,2],[2,1]], k = 4

Output: 1

Explanation:

Select the points (0, 0), (2, 0), (2, 2), and (2, 1).
```

**Example 3:**  

![3](./3.png)

```
Input: side = 2, points = [[0,0],[0,1],[0,2],[1,2],[2,0],[2,2],[2,1]], k = 5

Output: 1

Explanation:

Select the points (0, 0), (0, 1), (0, 2), (1, 2), and (2, 2).
```

## 结语

如果你同我一样热爱数据结构、算法、LeetCode，可以关注我 GitHub 上的 LeetCode 题解：[awesome-golang-algorithm][me]

[title]: https://leetcode.com/problems/maximize-the-distance-between-points-on-a-square/
[me]: https://github.com/kylesliu/awesome-golang-algorithm
