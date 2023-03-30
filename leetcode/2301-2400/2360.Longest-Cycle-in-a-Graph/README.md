# [2360.Longest Cycle in a Graph][title]

## Description
You are given a **directed** graph of `n` nodes numbered from `0` to `n - 1`, where each node has **at most one** outgoing edge.

The graph is represented with a given **0-indexed** array `edges` of size `n`, indicating that there is a directed edge from node `i` to node `edges[i]`. If there is no outgoing edge from node `i`, then `edges[i] == -1`.

Return the length of the **longest** cycle in the graph. If no cycle exists, return `-1`.

A cycle is a path that starts and ends at the **same** node.

**Example 1:**  

![example1](./graph4drawio-5.png)

```
Input: edges = [3,3,4,2,3]
Output: 3
Explanation: The longest cycle in the graph is the cycle: 2 -> 4 -> 3 -> 2.
The length of this cycle is 3, so 3 is returned.
```

**Example 2:**  

![example2](./graph4drawio-1.png)


```
Input: edges = [2,-1,3,1]
Output: -1
Explanation: There are no cycles in this graph.
```

## 结语

如果你同我一样热爱数据结构、算法、LeetCode，可以关注我 GitHub 上的 LeetCode 题解：[awesome-golang-algorithm][me]

[title]: https://leetcode.com/problems/longest-cycle-in-a-graph/
[me]: https://github.com/kylesliu/awesome-golang-algorithm
