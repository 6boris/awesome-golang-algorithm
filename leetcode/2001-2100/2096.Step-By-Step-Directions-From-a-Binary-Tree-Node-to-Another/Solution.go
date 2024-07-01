package Solution

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func Solution(root *TreeNode, startValue int, destValue int) string {
	var dfs func(*TreeNode, *[]int, *[]byte, int) bool
	dfs = func(tree *TreeNode, nums *[]int, dirs *[]byte, value int) bool {
		if tree == nil {
			return false
		}
		if tree.Val == value {
			return true
		}
		if tree.Left != nil {
			*nums = append(*nums, tree.Left.Val)
			*dirs = append(*dirs, 'L')
			if dfs(tree.Left, nums, dirs, value) {
				return true
			}
			*nums = (*nums)[:len(*nums)-1]
			*dirs = (*dirs)[:len(*dirs)-1]

		}
		if tree.Right != nil {
			*nums = append(*nums, tree.Right.Val)
			*dirs = append(*dirs, 'R')
			if dfs(tree.Right, nums, dirs, value) {
				return true
			}
			*nums = (*nums)[:len(*nums)-1]
			*dirs = (*dirs)[:len(*dirs)-1]
		}
		return false
	}
	nums1, nums2 := []int{}, []int{}
	dirs1, dirs2 := []byte{}, []byte{}
	_ = dfs(root, &nums1, &dirs1, startValue)
	_ = dfs(root, &nums2, &dirs2, destValue)

	i := 0
	for ; i < len(nums1) && i < len(nums2) && nums1[i] == nums2[i]; i++ {
	}
	for i1 := i; i1 < len(dirs1); i1++ {
		dirs1[i1] = 'U'
	}
	return string(dirs1[i:]) + string(dirs2[i:])
}
