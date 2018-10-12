# [101. Symmetric Tree][title]

## Description

Given a binary tree, check whether it is a mirror of itself (ie, symmetric around its center).

For example, this binary tree `[1,2,2,3,4,4,3]` is symmetric:

```
    1
   / \
  2   2
 / \ / \
3  4 4  3
```

But the following `[1,2,2,null,3,null,3]` is not:

```
    1
   / \
  2   2
   \   \
   3    3
```

**Note:**

Bonus points if you could solve it both recursively and iteratively.

**Tags:** Tree, Depth-first Search, Breadth-first Search

## 题意
>题意是判断一棵二叉树是否左右对称

## 题解

### 思路1
> 用深度搜索，比较根结点的左右两棵子树是否对称，如果左右子树的值相同，那么再分别对左子树的左节点和右子树的右节点，左子树的右节点和右子树的左节点做比较即可。

```go
func isSymmetric(root *TreeNode) bool {
	return root == nil || isSymmetricHelper(root.Left, root.Right)
}

func isSymmetricHelper(left *TreeNode, right *TreeNode) bool {

	if left == nil || right == nil {
		return left == right
	}

	//	检查节点的值
	if left.Val != right.Val {
		return false
	}
	return isSymmetricHelper(left.Left, right.Right) &&
		isSymmetricHelper(left.Right, right.Left)
}
```

### 思路2
> 思路2
```go

```

## 结语

如果你同我一样热爱数据结构、算法、LeetCode，可以关注我 GitHub 上的 LeetCode 题解：[awesome-golang-leetcode][me]

[title]: https://leetcode.com/problems/symmetric-tree/description/
[me]: https://github.com/kylesliu/awesome-golang-leetcode
