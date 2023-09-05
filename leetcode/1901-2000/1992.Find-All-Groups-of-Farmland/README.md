# [1992.Find All Groups of Farmland][title]

## Description
You are given a **0-indexed** `m x n` binary matrix `land` where a `0` represents a hectare of forested land and a `1` represents a hectare of farmland.

To keep the land organized, there are designated rectangular areas of hectares that consist **entirely** of farmland. These rectangular areas are called **groups**. No two groups are adjacent, meaning farmland in one group is **not** four-directionally adjacent to another farmland in a different group.

`land` can be represented by a coordinate system where the top left corner of `land` is `(0, 0)` and the bottom right corner of `land` is `(m-1, n-1)`. Find the coordinates of the top left and bottom right corner of each **group** of farmland. A **group** of farmland with a top left corner at `(r1, c1)` and a bottom right corner at `(r2, c2)` is represented by the 4-length array `[r1, c1, r2, c2]`.

Return a 2D array containing the 4-length arrays described above for each **group** of farmland in `land`. If there are no groups of farmland, return an empty array. You may return the answer in **any order**.

**Example 1:**  

![example1](./screenshot-2021-07-27-at-12-23-15-copy-of-diagram-drawio-diagrams-net.png)

```
Input: land = [[1,0,0],[0,1,1],[0,1,1]]
Output: [[0,0,0,0],[1,1,2,2]]
Explanation:
The first group has a top left corner at land[0][0] and a bottom right corner at land[0][0].
The second group has a top left corner at land[1][1] and a bottom right corner at land[2][2].
```

**Example 2:**  

![example2](./screenshot-2021-07-27-at-12-30-26-copy-of-diagram-drawio-diagrams-net.png)

```
Input: land = [[1,1],[1,1]]
Output: [[0,0,1,1]]
Explanation:
The first group has a top left corner at land[0][0] and a bottom right corner at land[1][1].
```

**Example 3:**  

![example3](./screenshot-2021-07-27-at-12-32-24-copy-of-diagram-drawio-diagrams-net.png)

```
Input: land = [[0]]
Output: []
Explanation:
There are no groups of farmland.
```

## 结语

如果你同我一样热爱数据结构、算法、LeetCode，可以关注我 GitHub 上的 LeetCode 题解：[awesome-golang-algorithm][me]

[title]: https://leetcode.com/problems/find-all-groups-of-farmland/
[me]: https://github.com/kylesliu/awesome-golang-algorithm
