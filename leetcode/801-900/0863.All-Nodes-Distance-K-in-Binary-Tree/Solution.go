package Solution

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func findFather(root *TreeNode, father map[int]*TreeNode) {
	if root == nil {
		return
	}
	if root.Left != nil {
		father[root.Left.Val] = root
		findFather(root.Left, father)
	}
	if root.Right != nil {
		father[root.Right.Val] = root
		findFather(root.Right, father)
	}
}

func Solution(root *TreeNode, target *TreeNode, k int) []int {
	ans := make([]int, 0)
	father := make(map[int]*TreeNode)
	findFather(root, father)
	visited := make(map[int]struct{})
	visited[target.Val] = struct{}{}
	queue := []*TreeNode{target}
	for len(queue) > 0 && k > 0 {
		nextQ := make([]*TreeNode, 0)
		for _, item := range queue {
			if item.Left != nil {
				if _, ok := visited[item.Left.Val]; !ok {
					visited[item.Left.Val] = struct{}{}
					nextQ = append(nextQ, item.Left)
				}
			}
			if item.Right != nil {
				if _, ok := visited[item.Right.Val]; !ok {
					visited[item.Right.Val] = struct{}{}
					nextQ = append(nextQ, item.Right)
				}
			}

			// father
			if fa, ok := father[item.Val]; ok {
				if _, ok1 := visited[fa.Val]; !ok1 {
					visited[fa.Val] = struct{}{}
					nextQ = append(nextQ, fa)
				}
			}
		}
		queue = nextQ
		k--
	}
	for _, item := range queue {
		ans = append(ans, item.Val)
	}
	return ans
}
