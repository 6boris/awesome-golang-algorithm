package Solution

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isCompleteTree_1(root *TreeNode) bool {
	que, end := []*TreeNode{root}, false
	for len(que) > 0 {
		node := que[0]
		que = que[1:]
		if node == nil {
			end = true
		} else {
			if end == true {
				return false
			}
			que = append(que, node.Left, node.Right)
		}
	}
	return true
}
