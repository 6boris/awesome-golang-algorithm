package Solution

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func Solution(root *TreeNode, voyage []int) []int {
	ok := true
	array := make([]int, 0)
	var dfs func(*TreeNode, *int)
	dfs = func(tree *TreeNode, index *int) {
		if !ok {
			return
		}
		if *index >= len(voyage) {
			return
		}
		if tree == nil {
			return
		}
		if tree.Val != voyage[*index] {
			ok = false
			array = []int{-1}
			return
		}
		*index++
		left, right := tree.Left, tree.Right
		// 如果当前左子树与index的值相等，就不需要交换
		// 如果交换，就是左侧遍历右子树，右侧遍历左子树
		// 左右边的值不等于currentIndex，需要交换，添加节点
		if left != nil && *index < len(voyage) && left.Val != voyage[*index] {
			array = append(array, tree.Val)
			dfs(right, index)
			dfs(left, index)
		} else {
			dfs(left, index)
			dfs(right, index)
		}
	}
	index := 0
	dfs(root, &index)
	return array
}
