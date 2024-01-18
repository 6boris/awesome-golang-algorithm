package Solution

import "fmt"

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}
type qItem struct {
	node *TreeNode
	cur  int
}

func treeHight(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := treeHight(root.Left)
	right := treeHight(root.Right)
	if left > right {
		right = left
	}
	return right + 1
}

func Solution(root *TreeNode) [][]string {
	th := treeHight(root)
	width := (1 << th) - 1
	ans := make([][]string, th)

	for i := 0; i < th; i++ {
		ans[i] = make([]string, width)
	}

	th--
	cur := width / 2

	q := []qItem{{node: root, cur: cur}}
	index := 0
	for len(q) > 0 {
		nq := make([]qItem, 0)
		shift := th - index - 1
		if shift < 0 {
			for _, item := range q {
				ans[index][item.cur] = fmt.Sprintf("%d", item.node.Val)
			}
			break
		}

		x := 1 << shift
		for _, item := range q {
			ans[index][item.cur] = fmt.Sprintf("%d", item.node.Val)
			if item.node.Left != nil {
				nq = append(nq, qItem{node: item.node.Left, cur: item.cur - x})
			}
			if item.node.Right != nil {
				nq = append(nq, qItem{node: item.node.Right, cur: item.cur + x})
			}
		}
		q = nq
		index++
	}
	return ans
}
