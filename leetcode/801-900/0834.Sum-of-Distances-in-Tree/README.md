# [834.Sum of Distances in Tree][title]

## Description
There is an undirected connected tree with n nodes labeled from `0` to `n - 1` and `n - 1` edges.

You are given the integer `n` and the array `edges` where `edges[i] = [ai, bi]` indicates that there is an edge between nodes a<sub>i</sub> and b<sub>i</sub> in the tree.

Return an array `answer` of length `n` where `answer[i]` is the sum of the distances between the i<sup>th</sup> node in the tree and all other nodes.

**Example 1:**  

![example1](./lc-sumdist1.jpg)

```
Input: n = 6, edges = [[0,1],[0,2],[2,3],[2,4],[2,5]]
Output: [8,12,6,10,10,10]
Explanation: The tree is shown above.
We can see that dist(0,1) + dist(0,2) + dist(0,3) + dist(0,4) + dist(0,5)
equals 1 + 1 + 2 + 2 + 2 = 8.
Hence, answer[0] = 8, and so on.
```

**Example 2:**  

![example2](./lc-sumdist2.jpg)

```
Input: n = 1, edges = []
Output: [0]
```

**Example 3:**  

![example3](./lc-sumdist3.jpg)

```
Input: n = 2, edges = [[1,0]]
Output: [1,1]
```

## 结语

如果你同我一样热爱数据结构、算法、LeetCode，可以关注我 GitHub 上的 LeetCode 题解：[awesome-golang-algorithm][me]

[title]: https://leetcode.com/problems/sum-of-distances-in-tree/
[me]: https://github.com/kylesliu/awesome-golang-algorithm
