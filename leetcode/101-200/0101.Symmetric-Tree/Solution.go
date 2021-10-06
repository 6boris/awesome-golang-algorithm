package Solution

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric_1(root *TreeNode) bool {
	var dfs func(*TreeNode, *TreeNode) bool
	dfs = func(p *TreeNode, q *TreeNode) bool {
		if p == nil && q == nil {
			return true
		}
		if p == nil || q == nil {
			return false
		}
		return p.Val == q.Val &&
			dfs(p.Left, q.Right) &&
			dfs(p.Right, q.Left)
	}
	return dfs(root, root)
}
func isSymmetric_2(root *TreeNode) bool {
	if root == nil || (root.Left == nil && root.Right == nil) {
		return true
	}
	que := make([]*TreeNode, 0)
	que = append(que, root.Left, root.Right)
	for len(que) > 0 {
		left, right := que[0], que[1]
		que = que[2:]
		if left == nil && right == nil {
			continue
		}
		if left == nil || right == nil || left.Val != right.Val {
			return false
		}
		que = append(que, left.Left, right.Right, right.Left, left.Right)
	}
	return true
}
