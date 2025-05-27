package Solution

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func Solution(root *TreeNode) int {
	cache := make(map[*TreeNode][2]int)
	var dfs func(*TreeNode, bool) int
	dfs = func(tree *TreeNode, selectedParent bool) int {
		if tree == nil {
			return 0
		}
		if selectedParent {
			v, ok := cache[tree]
			if ok {
				if v[0] != -1 {
					return v[0]
				}
			}
			left := dfs(tree.Left, false)
			right := dfs(tree.Right, false)
			if !ok {
				v = [2]int{left + right, -1}
			} else {
				v[0] = left + right
			}
			cache[tree] = v
			return v[0]
		}
		v, ok := cache[tree]
		if !ok {
			v = [2]int{-1, -1}
		}
		if v[0] == -1 {
			left := dfs(tree.Left, false)
			right := dfs(tree.Right, false)
			v[0] = left + right
		}
		if v[1] == -1 {
			left1 := dfs(tree.Left, true)
			right1 := dfs(tree.Right, true)
			v[1] = left1 + right1 + tree.Val
		}
		cache[tree] = v
		return max(v[0], v[1])
	}
	return dfs(root, false)
}
