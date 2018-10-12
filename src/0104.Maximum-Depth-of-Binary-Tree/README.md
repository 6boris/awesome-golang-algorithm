# [104. Maximum Depth of Binary Tree][title]

## Description

Given a binary tree, find its maximum depth.

The maximum depth is the number of nodes along the longest path from the root node down to the farthest leaf node.

**Note:** A leaf is a node with no children.

**Example:**

Given binary tree `[3,9,20,null,null,15,7]`,

```
    3
   / \
  9  20
    /  \
   15   7
```

return its depth = 3.

**Tags:** Tree, Depth-first Search

## 题意
>寻找二叉树的最大深度

## 题解

### 思路1
>深搜即可，每深入一次节点加一即可，然后取左右子树的最大深度。

```go
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return Max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
```

### 思路2
> 思路2
```go

```

## 结语

如果你同我一样热爱数据结构、算法、LeetCode，可以关注我 GitHub 上的 LeetCode 题解：[awesome-golang-leetcode][me]

[title]: https://leetcode.com/problems/maximum-depth-of-binary-tree/description/
[me]: https://github.com/kylesliu/awesome-golang-leetcode
