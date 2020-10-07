package Solution

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	res := make([][]int, 0, 1<<5)
	dfs(root, 0, &res)
	return res
}

func dfs(root *TreeNode, depth int, res *[][]int) {
	if root == nil {
		return
	}
	if len(*res) <= depth {
		// 初始化当前层数的切片
		*res = append(*res, []int{root.Val})
		goto DFS
	}
	switch depth % 2 {
	case 0:
		// 当前层数为奇数就追加在队尾
		(*res)[depth] = append((*res)[depth], root.Val)
	case 1:
		// 当前层数为偶数就放在队首
		(*res)[depth] = append([]int{root.Val}, (*res)[depth]...)
	}

DFS:
	dfs(root.Left, depth+1, res)
	dfs(root.Right, depth+1, res)
}
