# [1267.Count Servers that Communicate][title]

## Description
You are given a map of a server center, represented as a `m * n` integer matrix `grid`, where 1 means that on that cell there is a server and 0 means that it is no server. Two servers are said to communicate if they are on the same row or on the same column.

Return the number of servers that communicate with any other server.

**Example 1:**  

![example1](./untitled-diagram-6.jpeg)

```
Input: grid = [[1,0],[0,1]]
Output: 0
Explanation: No servers can communicate with others.
```

**Example 2:**  

![example2](./untitled-diagram-4.jpeg)

```
Input: grid = [[1,0],[1,1]]
Output: 3
Explanation: All three servers can communicate with at least one other server.
```

**Example 3:**  

![example3](./untitled-diagram-1-3.jpeg)

```
Input: grid = [[1,1,0,0],[0,0,1,0],[0,0,1,0],[0,0,0,1]]
Output: 4
Explanation: The two servers in the first row can communicate with each other. The two servers in the third column can communicate with each other. The server at right bottom corner can't communicate with any other server.
```

## 结语

如果你同我一样热爱数据结构、算法、LeetCode，可以关注我 GitHub 上的 LeetCode 题解：[awesome-golang-algorithm][me]

[title]: https://leetcode.com/problems/count-servers-that-communicate/
[me]: https://github.com/kylesliu/awesome-golang-algorithm
