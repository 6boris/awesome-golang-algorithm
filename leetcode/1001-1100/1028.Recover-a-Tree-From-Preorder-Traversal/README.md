# [1028.Recover a Tree From Preorder Traversal][title]

## Description
We run a preorder depth-first search (DFS) on the `root` of a binary tree.

At each node in this traversal, we output D dashes (where `D` is the depth of this node), then we output the value of this node.  If the depth of a node is `D`, the depth of its immediate child is `D + 1`.  The depth of the `root` node is `0`.

If a node has only one child, that child is guaranteed to be **the left child**.

Given the output `traversal` of this traversal, recover the tree and return its `root`.

**Example 1:**  

![example1](./recover-a-tree-from-preorder-traversal.png)

```
Input: traversal = "1-2--3--4-5--6--7"
Output: [1,2,5,3,4,6,7]
```

**Example 2:**  

![example2](./screen-shot-2019-04-10-at-114101-pm.png)

```
Input: traversal = "1-2--3---4-5--6---7"
Output: [1,2,5,3,null,6,null,4,null,7]
```

**Example 3:**  

![example3](./screen-shot-2019-04-10-at-114955-pm.png)

```
Input: traversal = "1-401--349---90--88"
Output: [1,401,null,349,88,90]
```

## 结语

如果你同我一样热爱数据结构、算法、LeetCode，可以关注我 GitHub 上的 LeetCode 题解：[awesome-golang-algorithm][me]

[title]: https://leetcode.com/problems/recover-a-tree-from-preorder-traversal/
[me]: https://github.com/kylesliu/awesome-golang-algorithm
