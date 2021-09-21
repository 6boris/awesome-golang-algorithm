package Solution

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth_1(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := maxDepth_1(root.Left)
	right := maxDepth_1(root.Right)

	return max(left, right) + 1
}

func maxDepth_2(root *TreeNode) int {
	maxDep := 0
	var dfs func(*TreeNode, int)
	dfs = func(node *TreeNode, v int) {
		if node == nil {
			return
		}
		maxDep = max(maxDep, v)
		dfs(node.Left, v+1)
		dfs(node.Right, v+1)
	}
	if root != nil {
		dfs(root, 1)
	}
	return maxDep
}

func maxDepth_3(root *TreeNode) int {
	ans, que := 0, make([]*TreeNode, 0)
	if root == nil {
		return 0
	}
	que = append(que, root)
	for len(que) > 0 {
		length := len(que)
		for i := 0; i < length; i++ {
			node := que[0]
			que = que[1:]
			if node.Left != nil {
				que = append(que, node.Left)
			}
			if node.Right != nil {
				que = append(que, node.Right)
			}
		}
		ans++
	}
	return ans
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
