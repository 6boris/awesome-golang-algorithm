# [501.Find Mode in Binary Search Tree][title]

## Description
Given the `root` of a binary search tree (BST) with duplicates, return all the [mode(s)](https://en.wikipedia.org/wiki/Mode_(statistics)) (i.e., the most frequently occurred element) in it.

If the tree has more than one mode, return them in **any order**.

Assume a BST is defined as follows:

- The left subtree of a node contains only nodes with keys **less than or equal to** the node's key.
- The right subtree of a node contains only nodes with keys **greater than or equal to** the node's key.
Both the left and right subtrees must also be binary search trees.

**Example 1:**  

![example1](./mode-tree.jpeg)

```
Input: root = [1,null,2,2]
Output: [2]
```

**Example 2:**

```
Input: root = [0]
Output: [0]
```

## 结语

如果你同我一样热爱数据结构、算法、LeetCode，可以关注我 GitHub 上的 LeetCode 题解：[awesome-golang-algorithm][me]

[title]: https://leetcode.com/problems/find-mode-in-binary-search-tree/
[me]: https://github.com/kylesliu/awesome-golang-algorithm
