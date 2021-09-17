package Solution

var (
	idxMap map[int]int
	p      int
)

func Solution(inorder []int, postorder []int) *TreeNode {
	p = len(postorder) - 1
	idxMap = make(map[int]int)
	for i, v := range inorder {
		idxMap[v] = i
	}
	return build(0, len(inorder)-1, postorder)
}

func build(left, right int, postorder []int) *TreeNode {
	if left > right {
		return nil
	}
	root := &TreeNode{Val: postorder[p]}
	p--
	idx := idxMap[root.Val]
	root.Right = build(idx+1, right, postorder)
	root.Left = build(left, idx-1, postorder)
	return root
}
