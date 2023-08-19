# [1584.Min Cost to Connect All Points][title]

## Description
You are given an array `points` representing integer coordinates of some points on a 2D-plane, where `points[i] = [xi, yi]`.

The cost of connecting two points `[xi, yi]` and `[xj, yj]` is the **manhattan distance** between them: `|xi - xj| + |yi - yj|`, where `|val|` denotes the absolute value of `val`.

Return the minimum cost to make all points connected. All points are connected if there is **exactly one** simple path between any two points.

**Example 1:**  

![example1](./d.png)

```
Input: points = [[0,0],[2,2],[3,10],[5,2],[7,0]]
Output: 20
Explanation:
We can connect the points as shown above to get the minimum cost of 20.
Notice that there is a unique path between every pair of points.
```

![explanation](./c.png)

**Example 2:**

```
Input: points = [[3,12],[-2,5],[-4,1]]
Output: 18
```

## 结语

如果你同我一样热爱数据结构、算法、LeetCode，可以关注我 GitHub 上的 LeetCode 题解：[awesome-golang-algorithm][me]

[title]: https://leetcode.com/problems/min-cost-to-connect-all-points/
[me]: https://github.com/kylesliu/awesome-golang-algorithm
