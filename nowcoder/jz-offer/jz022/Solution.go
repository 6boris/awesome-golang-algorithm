package Solution

func Solution(x bool) bool {
	return x
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	res := []int{}
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
