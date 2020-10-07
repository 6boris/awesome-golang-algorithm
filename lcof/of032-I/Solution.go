package Solution

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	queue, res := make([]*TreeNode, 0), []int{}
	queue = append(queue, root)
	for len(queue) != 0 {
		cRoot := queue[0]
		res = append(res, cRoot.Val)
		if cRoot.Left != nil {
			queue = append(queue, cRoot.Left)
		}
		if cRoot.Right != nil {
			queue = append(queue, cRoot.Right)
		}
		queue = queue[1:]
	}
	return res
}
