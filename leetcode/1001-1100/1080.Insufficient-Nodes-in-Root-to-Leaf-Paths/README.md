# [1080.Insufficient Nodes in Root to Leaf Paths][title]

## Description
Given the `root` of a binary tree and an integer `limit`, delete all **insufficient nodes** in the tree simultaneously, and return the root of the resulting binary tree.

A node is **insufficient** if every root to **leaf** path intersecting this node has a sum strictly less than `limit`.

A **leaf** is a node with no children.

**Example 1:**  

![example1](./insufficient-11.png)

```
Input: root = [1,2,3,4,-99,-99,7,8,9,-99,-99,12,13,-99,14], limit = 1
Output: [1,2,3,4,null,null,7,8,9,null,14]
```

**Example 2:**  

![example2](./insufficient-11.png)

```
Input: root = [5,4,8,11,null,17,4,7,1,null,null,5,3], limit = 22
Output: [5,4,8,11,null,17,4,7,null,null,null,5]
```

**Example 3:**  

![example3](./screen-shot-2019-06-11-at-83301-pm.png)

```
Input: root = [1,2,-3,-5,null,4,null], limit = -1
Output: [1,null,-3,4]
```

## 结语

如果你同我一样热爱数据结构、算法、LeetCode，可以关注我 GitHub 上的 LeetCode 题解：[awesome-golang-algorithm][me]

[title]: https://leetcode.com/problems/insufficient-nodes-in-root-to-leaf-paths/
[me]: https://github.com/kylesliu/awesome-golang-algorithm
