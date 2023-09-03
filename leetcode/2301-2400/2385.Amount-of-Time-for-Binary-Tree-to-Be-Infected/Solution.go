package Solution

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func Solution(root *TreeNode, start int) int {
	if root.Left == nil && root.Right == nil {
		return 0
	}
	father := make(map[*TreeNode]*TreeNode)
	var startNode *TreeNode
	var dfs func(*TreeNode)
	dfs = func(t *TreeNode) {
		if t == nil {
			return
		}
		if t.Val == start {
			startNode = t
		}
		if t.Left != nil {
			father[t.Left] = t
			dfs(t.Left)
		}
		if t.Right != nil {
			father[t.Right] = t
			dfs(t.Right)
		}
	}

	dfs(root)
	q := []*TreeNode{startNode}
	visited := make(map[*TreeNode]struct{})
	visited[startNode] = struct{}{}
	ans := 0
	for len(q) > 0 {
		next := make([]*TreeNode, 0)
		for _, n := range q {
			for _, nn := range []*TreeNode{n.Left, n.Right, father[n]} {
				if nn == nil {
					continue
				}
				if _, ok := visited[nn]; ok {
					continue
				}
				visited[nn] = struct{}{}
				next = append(next, nn)
			}
		}
		q = next
		ans++
	}
	return ans - 1
}
