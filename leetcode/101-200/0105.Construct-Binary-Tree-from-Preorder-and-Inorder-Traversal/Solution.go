package Solution

var (
	idxMap      map[int]int
	preorderIdx int
)

func Solution(preorder []int, inorder []int) *TreeNode {
	idxMap = make(map[int]int)
	preorderIdx = 0
	for i, v := range inorder {
		idxMap[v] = i
	}
	return build(0, len(inorder)-1, preorder)
}

func build(left, right int, preorder []int) *TreeNode {
	if left > right {
		return nil
	}
	root := &TreeNode{Val: preorder[preorderIdx]}
	preorderIdx++
	idx := idxMap[root.Val]
	root.Left = build(left, idx-1, preorder)
	root.Right = build(idx+1, right, preorder)
	return root
}
