# [1372.Longest ZigZag Path in a Binary Tree][title]

## Description
You are given the `root` of a binary tree.

A ZigZag path for a binary tree is defined as follow:

- Choose **any** node in the binary tree and a direction (right or left).
- If the current direction is right, move to the right child of the current node; otherwise, move to the left child.
- Change the direction from right to left or from left to right.
- Repeat the second and third steps until you can't move in the tree.
Zigzag length is defined as the number of nodes visited - 1. (A single node has a length of 0).

Return the longest **ZigZag** path contained in that tree.

**Example 1:**  

![example1](./sample_1_1702.png)

```
Input: root = [1,null,1,1,1,null,null,1,1,null,1,null,null,null,1,null,1]
Output: 3
Explanation: Longest ZigZag path in blue nodes (right -> left -> right).
```

**Example 2:**  

![example2](./sample_2_1702.png)

```
Input: root = [1,1,1,null,1,null,null,1,1,null,1]
Output: 4
Explanation: Longest ZigZag path in blue nodes (left -> right -> left -> right).
```


**Example 3:**

```
Input: root = [1]
Output: 0
```

## 结语

如果你同我一样热爱数据结构、算法、LeetCode，可以关注我 GitHub 上的 LeetCode 题解：[awesome-golang-algorithm][me]

[title]: https://leetcode.com/problems/longest-zigzag-path-in-a-binary-tree
[me]: https://github.com/kylesliu/awesome-golang-algorithm
