package Solution

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func Solution(root *TreeNode) int {
	ans := 0
	path := make([]bool, 9)
	pseudoPalindromic1457(root, path, &ans)
	return ans
}

// 00000000 00000000 00000000 00000000
func pseudoPalindromic1457(root *TreeNode, path []bool, ans *int) {
	if root == nil {
		return
	}
	path[root.Val-1] = !path[root.Val-1]

	pseudoPalindromic1457(root.Left, path, ans)
	pseudoPalindromic1457(root.Right, path, ans)

	if root.Left == nil && root.Right == nil {
		odd := false
		idx := 0
		for ; idx < 9; idx++ {
			if !path[idx] {
				continue
			}
			if odd {
				break
			}
			odd = true
		}
		if idx == 9 {
			*ans++
		}
	}

	path[root.Val-1] = !path[root.Val-1]

}
