# [993.Cousins in Binary Tree][title]

## Description
Given the `root` of a binary tree with unique values and the values of two different nodes of the tree `x` and `y`, return `true` if the nodes corresponding to the values `x` and `y` in the tree are **cousins**, or `false` otherwise.

Two nodes of a binary tree are **cousins** if they have the same depth with different parents.

Note that in a binary tree, the root node is at the depth `0`, and children of each depth `k` node are at the depth `k + 1`.

**Example 1:**  

![1](./1.png)

```
Input: root = [1,2,3,4], x = 4, y = 3
Output: false
```

**Example 2:**  

![2](./2.png)

```
Input: root = [1,2,3,null,4,null,5], x = 5, y = 4
Output: true
```

**Example 3:**  

![3](./3.png)

```
Input: root = [1,2,3,null,4], x = 2, y = 3
Output: false
```

## 结语

如果你同我一样热爱数据结构、算法、LeetCode，可以关注我 GitHub 上的 LeetCode 题解：[awesome-golang-algorithm][me]

[title]: https://leetcode.com/problems/cousins-in-binary-tree/
[me]: https://github.com/kylesliu/awesome-golang-algorithm
