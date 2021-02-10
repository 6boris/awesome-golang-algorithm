package Solution

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//	递归求解
func buildTree(pre []int, in []int) *TreeNode {
	if len(pre) == 0 || len(in) == 0 {
		return nil
	}
	//	找到跟节点位置
	mid := search(in, pre[0])
	return &TreeNode{
		Val:   pre[0],
		Left:  buildTree(pre[1:mid+1], in[:mid+1]),
		Right: buildTree(pre[mid+1:], in[mid+1:]),
	}
}
func search(nodes []int, val int) int {
	for p, v := range nodes {
		if v == val {
			return p
		}
	}
	return -1
}
