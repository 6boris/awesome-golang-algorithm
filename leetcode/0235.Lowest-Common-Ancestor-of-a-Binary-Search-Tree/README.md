# [236. Lowest Common Ancestor of a Binary Tree][title]

## Description

Given a binary search tree (BST), find the lowest common ancestor (LCA) of two given nodes in the BST.

According to the definition of LCA on Wikipedia: “The lowest common ancestor is defined between two nodes p and q as the lowest node in T that has both p and q as descendants (where we allow a node to be a descendant of itself).”

Given binary search tree:  `root = [6,2,8,0,4,7,9,null,null,3,5]`

```
       _______6______
       /              \
    ___2__          ___8__
   /      \        /      \
   0      _4       7       9
         /  \
         3   5
```

**Example 1:**

Input: root = [6,2,8,0,4,7,9,null,null,3,5], p = 2, q = 8

Output: 6

Explanation: The LCA of nodes 2 and 8 is 6.

**Example 2:**

Input: root = [6,2,8,0,4,7,9,null,null,3,5], p = 2, q = 4

Output: 2

Explanation: The LCA of nodes 2 and 4 is 2, since a node can be a descendant of itself
             according to the LCA definition.

## NOTES

- All of the nodes' values will be unique.
- `p` and `q` are different and both values will exist in the BST.

**Tags:** Binary Search Tree (BST)

## 题意
> 在二叉搜索树中查询两个节点最近的祖先元素

### 思路1
定义节点，直接左右递归, 详见代码。

代码实现可以查考`0236`实现， 但是考虑到`BST`的有序性, 判断逻辑可以更加精简。
```go
func Lowest(root *TreeNode, p *TreeNode, q *TreeNode) *TreeNode {
	if p.val < root.val && q.val < root.val {
		return Lowest(root.left, p, q)
	}
	if p.val > root.val && q.val > root.val {
		return Lowest(root.right, p, q)
	}
	return root
}
```

## 结语

如果你同我一样热爱数据结构、算法、LeetCode，可以关注我 GitHub 上的 LeetCode 题解：[awesome-golang-letcode][me]

[title]: https://leetcode.com/problems/valid-anagram/description/
[me]: https://github.com/kylesliu/awesome-golang-algorithm
