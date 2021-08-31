package Solution

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func Solution(root *TreeNode, k int) bool {
	nodes := make([]int, 0)
	treeArray(root, &nodes)
	return binarySearch(nodes, k)
}

func treeArray(root *TreeNode, nodes *[]int) {
	if root == nil {
		return
	}
	if root.Left != nil {
		treeArray(root.Left, nodes)
	}

	*nodes = append(*nodes, root.Val)

	if root.Right != nil {
		treeArray(root.Right, nodes)
	}
}
func binarySearch(nodes []int, target int) bool {
	start, end := 0, len(nodes)-1
	for start < end {
		num := nodes[start] + nodes[end]
		if num == target {
			return true
		}
		if num < target {
			start++
			continue
		}
		end--
	}

	return false
}
