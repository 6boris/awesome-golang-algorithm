# [236. Lowest Common Ancestor of a Binary Tree][title]

## Description

Given a binary tree, find the lowest common ancestor (LCA) of two given nodes in the tree.

According to the definition of [LCA on Wikipedia](https://en.wikipedia.org/wiki/Lowest_common_ancestor): “The lowest common ancestor is defined between two nodes p and q as the lowest node in T that has both p and q as descendants (where we allow a node to be a descendant of itself).”

Given the following binary tree:  root = [3,5,1,6,2,0,8,null,null,7,4]

```
_______3______
       /              \
    ___5__          ___1__
   /      \        /      \
   6      _2       0       8
         /  \
         7   4
```

**Example 1:**

Input: root = [3,5,1,6,2,0,8,null,null,7,4], p = 5, q = 1

Output: 3

Explanation: The LCA of of nodes 5 and 1 is 3.

**Example 2:**

Input: root = [3,5,1,6,2,0,8,null,null,7,4], p = 5, q = 4

Output: 5

Explanation: The LCA of nodes 5 and 4 is 5, since a node can be a descendant of itself
             according to the LCA definition.

## NOTES
- All of the nodes' values will be unique.
- `p` and `q` are different and both values will exist in the binary tree.

**Tags:** Binary Tree

## 题意
> 查找二叉树某两个节点最近的祖先元素

### 思路1
定义节点，直接左右递归, 详见代码
```go
type TreeNode struct {
	val   int
	left  *TreeNode
	right *TreeNode
}
```

## 结语

如果你同我一样热爱数据结构、算法、LeetCode，可以关注我 GitHub 上的 LeetCode 题解：[awesome-golang-letcode][me]

[title]: https://leetcode.com/problems/valid-anagram/description/
[me]: https://github.com/kylesliu/awesome-golang-leetcode
