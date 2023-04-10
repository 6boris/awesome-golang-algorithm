package Solution

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func buildTree1110(root *TreeNode, deleteValues map[int]struct{}) (*TreeNode, []*TreeNode) {
	if root == nil {
		return nil, nil
	}
	ans := []*TreeNode{}
	if _, ok := deleteValues[root.Val]; ok {
		if root.Left != nil {
			ans = append(ans, root.Left)
		}
		if root.Right != nil {
			ans = append(ans, root.Right)
		}
		return nil, ans
	}

	node := &TreeNode{Val: root.Val}
	lNode, c1 := buildTree1110(root.Left, deleteValues)
	rNode, c2 := buildTree1110(root.Right, deleteValues)
	node.Left = lNode
	node.Right = rNode
	if len(c1) > 0 {
		ans = append(ans, c1...)
	}
	if len(c2) > 0 {
		ans = append(ans, c2...)
	}
	return node, ans
}

func Solution(root *TreeNode, to_delete []int) []*TreeNode {
	if len(to_delete) == 0 {
		return []*TreeNode{root}
	}

	deleteValues := make(map[int]struct{})
	for _, v := range to_delete {
		deleteValues[v] = struct{}{}
	}
	ans := make([]*TreeNode, 0)
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		nextQ := make([]*TreeNode, 0)
		for _, tree := range queue {
			t, otherTree := buildTree1110(tree, deleteValues)
			if t != nil {
				ans = append(ans, t)
			}
			nextQ = append(nextQ, otherTree...)
		}
		queue = nextQ
	}
	return ans
}
