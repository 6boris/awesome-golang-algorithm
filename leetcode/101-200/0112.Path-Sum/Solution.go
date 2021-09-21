package Solution

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func hasPathSum_1(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	if root.Left == nil && root.Right == nil {
		return targetSum == root.Val
	}
	return hasPathSum_1(root.Left, targetSum-root.Val) || hasPathSum_1(root.Right, targetSum-root.Val)
}

type Point struct {
	Node *TreeNode
	Path int
}

func hasPathSum_2(root *TreeNode, targetSum int) bool {
	que := make([]*Point, 0)
	if root == nil {
		return false
	}
	que = append(que, &Point{root, root.Val})
	for len(que) > 0 {
		length := len(que)
		for i := 0; i < length; i++ {
			p := que[0]
			que = que[1:]
			if p.Node.Left == nil && p.Node.Right == nil && p.Path == targetSum {
				return true
			}
			if p.Node.Left != nil {
				que = append(que, &Point{p.Node.Left, p.Path + p.Node.Left.Val})
			}
			if p.Node.Right != nil {
				que = append(que, &Point{p.Node.Right, p.Path + p.Node.Right.Val})
			}
		}
	}
	return false
}
