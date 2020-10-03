package Solution

import "fmt"

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	res := [][]int{}
	que := Queue{}
	que.Push(root)

	for que.Len() != 0 {
		level_size := que.Len()
		current_level := []int{}

		for i := 0; i < level_size; i++ {
			node := que.Pop().(*TreeNode)
			current_level = append(current_level, node.Val)

			if node.Left != nil {
				que.Push(node.Left)
				fmt.Println(node.Left)

			}
			if node.Right != nil {
				que.Push(node.Right)
				fmt.Println(node.Right)
			}
		}
		res = append(res, current_level)
	}

	return res
}

func levelOrder2(root *TreeNode) [][]int {

	var ll [][]int

	if root != nil {
		ll = append(ll, []int{root.Val})
		recLevelOrder(root.Left, 1, &ll)
		recLevelOrder(root.Right, 1, &ll)
	}

	return ll
}

func recLevelOrder(head *TreeNode, depth int, llp *[][]int) {

	if head == nil {
		return
	}

	if depth >= len(*llp) {
		*llp = append(*llp, []int{})
	}

	(*llp)[depth] = append((*llp)[depth], head.Val)

	recLevelOrder(head.Left, depth+1, llp)
	recLevelOrder(head.Right, depth+1, llp)
}
