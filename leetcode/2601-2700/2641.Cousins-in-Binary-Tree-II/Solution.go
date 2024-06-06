package Solution

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func Solution(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	sum := 0
	queue := [][2]*TreeNode{{nil, root}}
	root.Val = 0
	for len(queue) > 0 {
		nq := make([][2]*TreeNode, 0)
		sum = 0
		del := make(map[*TreeNode]int)
		for _, item := range queue {
			if item[1].Left != nil {
				nq = append(nq, [2]*TreeNode{item[1], item[1].Left})
				sum += item[1].Left.Val
				del[item[1]] += item[1].Left.Val
			}
			if item[1].Right != nil {
				nq = append(nq, [2]*TreeNode{item[1], item[1].Right})
				sum += item[1].Right.Val
				del[item[1]] += item[1].Right.Val
			}
		}
		for _, item := range nq {
			item[1].Val = sum - del[item[0]]
		}
		queue = nq
	}
	return root
}
