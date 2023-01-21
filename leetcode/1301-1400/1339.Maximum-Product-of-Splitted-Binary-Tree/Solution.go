package Solution

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

const mod1339 = 1000000007

func Solution(root *TreeNode) int {
	sum := treeSum(root)
	ans := -1
	splitTree(root, sum, &ans)
	return ans % mod1339
}

func treeSum(root *TreeNode) int {
	if root == nil {
		return 0
	}

	left := treeSum(root.Left)
	right := treeSum(root.Right)
	root.Val += left + right
	return root.Val
}

func splitTree(root *TreeNode, sum int, ans *int) {
	if root == nil {
		return
	}
	if root.Left != nil {
		splitTree(root.Left, sum, ans)
	}

	if root.Right != nil {
		splitTree(root.Right, sum, ans)
	}
	diff := sum - root.Val
	s := diff * root.Val
	if *ans == -1 || s > *ans {
		*ans = s
	}
}
