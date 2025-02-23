package Solution

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func Solution(preorder []int, postorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	root := &TreeNode{
		Val: preorder[0],
	}
	if len(preorder) == 1 {
		return root
	}

	postIndex := 0
	next := preorder[1]
	for i := range postorder {
		if postorder[i] == next {
			postIndex = i
			break
		}
	}
	preIndex := 0
	m := map[int]int{}
	for i, n := range preorder {
		m[n] = i
	}
	for i := 0; i <= postIndex; i++ {
		preIndex = max(preIndex, m[postorder[i]])
	}
	root.Left = Solution(preorder[1:preIndex+1], postorder[:postIndex+1])
	root.Right = Solution(preorder[preIndex+1:], postorder[postIndex+1:len(postorder)-1])
	return root
}
