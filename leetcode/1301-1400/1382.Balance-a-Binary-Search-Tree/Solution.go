package Solution

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func Solution(root *TreeNode) *TreeNode {
	var targetRoot *TreeNode
	depth := make(map[*TreeNode]int)
	var inOrder func(node *TreeNode)
	inOrder = func(root *TreeNode) {
		if root.Left != nil {
			inOrder(root.Left)
		}
		targetRoot = rebuildBST(targetRoot, root.Val, depth)
		if root.Right != nil {
			inOrder(root.Right)
		}
	}
	inOrder(root)
	return targetRoot
}

func max1382(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func leftRotate(node *TreeNode, depth map[*TreeNode]int) *TreeNode {
	right := node.Right
	rightLeft := right.Left
	right.Left = node
	node.Right = rightLeft

	depth[node] = max1382(depth[node.Left], depth[node.Right]) + 1
	depth[right] = max1382(depth[right.Left], depth[right.Right]) + 1
	return right
}

func rightRotate(node *TreeNode, depth map[*TreeNode]int) *TreeNode {
	left := node.Left
	leftRight := left.Right
	node.Left = leftRight
	left.Right = node

	depth[node] = max1382(depth[node.Left], depth[node.Right]) + 1
	depth[left] = max1382(depth[left.Left], depth[left.Right]) + 1
	return left
}

// 树的高度都是left-right
func rebuildBST(root *TreeNode, val int, depth map[*TreeNode]int) *TreeNode {
	newNode := &TreeNode{
		Val: val,
	}
	depth[newNode] = 1

	if root == nil {
		return newNode
	}
	if root.Val < val {
		root.Right = rebuildBST(root.Right, val, depth)
	} else if root.Val > val {
		root.Left = rebuildBST(root.Left, val, depth)
	} else {
		return root
	}

	hl := depth[root.Left]
	hr := depth[root.Right]
	rootHeight := hl + 1
	if hr+1 > rootHeight {
		rootHeight = hr + 1
	}
	depth[root] = rootHeight

	diff := hl - hr
	if diff > 1 && val < root.Left.Val {
		return rightRotate(root, depth)
	}
	if diff > 1 && val > root.Left.Val {
		root.Left = leftRotate(root.Left, depth)
		return rightRotate(root, depth)
	}
	if diff < -1 && val > root.Right.Val {
		return leftRotate(root, depth)
	}
	if diff < -1 && val < root.Right.Val {
		root.Right = rightRotate(root.Right, depth)
		return leftRotate(root, depth)
	}
	return root
}
