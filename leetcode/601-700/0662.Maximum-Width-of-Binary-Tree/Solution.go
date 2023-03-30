package Solution

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}
type treeNodeWithIndex struct {
	node  *TreeNode
	index int
}

func Solution(root *TreeNode) int {
	if root == nil {
		return 0
	}
	queue := []*treeNodeWithIndex{{node: root, index: 0}}
	ans := 0
	for len(queue) > 0 {
		nextQ := make([]*treeNodeWithIndex, 0)

		left, right := queue[0], queue[len(queue)-1]
		if diff := right.index - left.index + 1; diff > ans {
			ans = diff
		}
		for _, item := range queue {
			i := item.index * 2
			if item.node.Left != nil {
				nextQ = append(nextQ, &treeNodeWithIndex{node: item.node.Left, index: i})
			}
			if item.node.Right != nil {
				nextQ = append(nextQ, &treeNodeWithIndex{node: item.node.Right, index: i + 1})
			}
		}

		queue = nextQ
	}
	return ans
}
