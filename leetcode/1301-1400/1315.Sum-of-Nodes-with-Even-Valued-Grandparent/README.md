# [1315.Sum of Nodes with Even-Valued Grandparent][title]

## Description
Given the `root` of a binary tree, return the sum of values of nodes with an **even-valued grandparent**. If there are no nodes with an **even-valued grandparent**, return `0`.

A **grandparent** of a node is the parent of its parent if it exists.

**Example 1:**  

![example1](./even1-tree.jpg)

```
Input: root = [6,7,8,2,7,1,3,9,null,1,4,null,null,null,5]
Output: 18
Explanation: The red nodes are the nodes with even-value grandparent while the blue nodes are the even-value grandparents.
```

**Example 2:**  

![example2](./even2-tree.jpg)

```
Input: root = [1]
Output: 0
```

## 结语

如果你同我一样热爱数据结构、算法、LeetCode，可以关注我 GitHub 上的 LeetCode 题解：[awesome-golang-algorithm][me]

[title]: https://leetcode.com/problems/sum-of-nodes-with-even-valued-grandparent/
[me]: https://github.com/kylesliu/awesome-golang-algorithm
