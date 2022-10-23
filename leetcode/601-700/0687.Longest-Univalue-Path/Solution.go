package Solution

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func Solution(root *TreeNode) int {
	ans := 0
	longestUnivaluePath687(root, &ans)
	if ans > 0 {
		ans--
	}
	return ans
}

func longestUnivaluePath687(root *TreeNode, ans *int) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}
	left := longestUnivaluePath687(root.Left, ans)
	right := longestUnivaluePath687(root.Right, ans)
	tryReturn := 0
	current := 1
	if left != 0 && root.Val == root.Left.Val {
		current += left
		tryReturn = left
	}

	if right != 0 && root.Val == root.Right.Val {
		current += right
		if tryReturn < right {
			tryReturn = right
		}
	}
	if current > *ans {
		*ans = current
	}
	return tryReturn + 1
}
