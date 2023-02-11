# [1129.Shortest Path with Alternating Colors][title]

## Description
You are given an integer `n`, the number of nodes in a directed graph where the nodes are labeled from `n` to `n - 1`. Each edge is red or blue in this graph, and there could be self-edges and parallel edges.

You are given two arrays `redEdges` and `blueEdges` where:

- `redEdges[i] = [ai, bi]` indicates that there is a directed red edge from node a<sub>i</sub> to node b<sub>i</sub> in the graph, and
- `blueEdges[j] = [uj, vj]` indicates that there is a directed blue edge from node u<sub>j</sub> to node v<sub>j</sub> in the graph.

Return an array `answer` of length n, where each `answer[x]` is the length of the shortest path from node 0 to node x such that the edge colors alternate along the path, or `-1` if such a path does not exist.

**Example 1:**

```
Input: n = 3, redEdges = [[0,1],[1,2]], blueEdges = []
Output: [0,1,-1]
```

**Example 2:**

```
Input: n = 3, redEdges = [[0,1]], blueEdges = [[2,1]]
Output: [0,1,-1]
```

## 结语

如果你同我一样热爱数据结构、算法、LeetCode，可以关注我 GitHub 上的 LeetCode 题解：[awesome-golang-algorithm][me]

[title]: https://leetcode.com/problems/shortest-path-with-alternating-colors/
[me]: https://github.com/kylesliu/awesome-golang-algorithm
