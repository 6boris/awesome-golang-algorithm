# [2583.Kth Largest Sum in a Binary Tree][title]

## Description
You are given the `root` of a binary tree and a positive integer `k`.

The **level sum** in the tree is the sum of the values of the nodes that are on the **same** level.

Return the k<sup>th</sup> **largest** level sum in the tree (not necessarily distinct). If there are fewer than `k` levels in the tree, return `-1`.

**Note** that two nodes are on the same level if they have the same distance from the root.

**Example 1:**  

![example1](./binaryytreeedrawio-2.png)

```
Input: root = [5,8,9,2,1,3,7,4,6], k = 2
Output: 13
Explanation: The level sums are the following:
- Level 1: 5.
- Level 2: 8 + 9 = 17.
- Level 3: 2 + 1 + 3 + 7 = 13.
- Level 4: 4 + 6 = 10.
The 2nd largest level sum is 13.
```

**Example 2:**

```
Input: root = [1,2,null,3], k = 1
Output: 3
Explanation: The largest level sum is 3.
```

## 结语

如果你同我一样热爱数据结构、算法、LeetCode，可以关注我 GitHub 上的 LeetCode 题解：[awesome-golang-algorithm][me]

[title]: https://leetcode.com/problems/kth-largest-sum-in-a-binary-tree
[me]: https://github.com/kylesliu/awesome-golang-algorithm
