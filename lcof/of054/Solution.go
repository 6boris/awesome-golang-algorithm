package Solution

func Solution(x bool) bool {
	return x
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func kthLargest(root *TreeNode, k int) int {
	var nums []int
	nums = dfs(&nums, root)
	return nums[k-1]
}

//逆中序遍历（右中左 -- 递减序列）
func dfs(nums *[]int, r *TreeNode) []int {
	if r.Right != nil {
		dfs(nums, r.Right)
	}
	if r != nil {
		*nums = append(*nums, r.Val)
	}
	if r.Left != nil {
		dfs(nums, r.Left)
	}
	return *nums
}
