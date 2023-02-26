# [669.Trim a Binary Search Tree][title]

## Description
Given the `root` of a binary search tree and the lowest and highest boundaries as `low` and `high`, trim the tree so that all its elements lies in `[low, high]`. Trimming the tree should **not** change the relative structure of the elements that will remain in the tree (i.e., any node's descendant should remain a descendant). It can be proven that there is a **unique answer**.

Return the root of the trimmed binary search tree. Note that the root may change depending on the given bounds.

**Example 1:**  

![example1](./trim1.jpg)

```
Input: root = [1,0,2], low = 1, high = 2
Output: [1,null,2]
```

**Example 2:**  

![example2](./trim2.jpg)

```
Input: root = [3,0,4,null,2,null,null,1], low = 1, high = 3
Output: [3,2,null,1]
```

## 结语

如果你同我一样热爱数据结构、算法、LeetCode，可以关注我 GitHub 上的 LeetCode 题解：[awesome-golang-algorithm][me]

[title]: https://leetcode.com/problems/trim-a-binary-search-tree/
[me]: https://github.com/kylesliu/awesome-golang-algorithm
