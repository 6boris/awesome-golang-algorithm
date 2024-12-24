# [3243.Shortest Distance After Road Addition Queries I][title]

## Description
You are given an integer `n` and a 2D integer array `queries`.

There are `n` cities numbered from `0` to `n - 1`. Initially, there is a **unidirectional** road from city `i` to city `i + 1` for all `0 <= i < n - 1`.

`queries[i] = [ui, vi]` represents the addition of a new **unidirectional** road from city `ui` to city `vi`. After each query, you need to find the **length** of the **shortest path** from city `0` to city `n - 1`.

Return an array `answer` where for each `i` in the range `[0, queries.length - 1]`, `answer[i]` is the length of the shortest path from city `0` to city `n - 1` after processing the **first** `i + 1` queries.

**Example 1:**  

![1](./1.jpg)
![2](./2.jpg)
![3](./3.jpg)

```
Input: n = 5, queries = [[2,4],[0,2],[0,4]]
Output: [3,2,1]
```

**Example 2:**  

![4](./4.jpg)
![5](./5.jpg)

```
Input: n = 4, queries = [[0,3],[0,2]]

Output: [1,1]
```

## 结语

如果你同我一样热爱数据结构、算法、LeetCode，可以关注我 GitHub 上的 LeetCode 题解：[awesome-golang-algorithm][me]

[title]: https://leetcode.com/problems/shortest-distance-after-road-addition-queries-i/
[me]: https://github.com/kylesliu/awesome-golang-algorithm
