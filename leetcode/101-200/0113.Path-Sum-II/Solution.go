package Solution

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum_1(root *TreeNode, targetSum int) [][]int {
	ans := make([][]int, 0)
	var dfs func(*TreeNode, int, []int)
	dfs = func(node *TreeNode, target int, path []int) {
		if node == nil {
			return
		}
		target -= node.Val
		path = append(path, node.Val)
		if node.Left == nil && node.Right == nil && target == 0 {
			ans = append(ans, append([]int(nil), path...))
			return
		}
		dfs(node.Left, target, path)
		dfs(node.Right, target, path)
		if len(path) != 0 {
			path = path[:len(path)-1]
		}
	}
	dfs(root, targetSum, []int{})
	return ans
}

type pair struct {
	node *TreeNode
	left int
}

func pathSum_2(root *TreeNode, targetSum int) (ans [][]int) {
	if root == nil {
		return
	}
	parent := map[*TreeNode]*TreeNode{}
	getPath := func(node *TreeNode) (path []int) {
		for ; node != nil; node = parent[node] {
			path = append(path, node.Val)
		}
		for i, j := 0, len(path)-1; i < j; i++ {
			path[i], path[j] = path[j], path[i]
			j--
		}
		return
	}
	queue := []pair{{root, targetSum}}
	for len(queue) > 0 {
		p := queue[0]
		queue = queue[1:]
		node := p.node
		left := p.left - node.Val
		if node.Left == nil && node.Right == nil {
			if left == 0 {
				ans = append(ans, getPath(node))
			}
		} else {
			if node.Left != nil {
				parent[node.Left] = node
				queue = append(queue, pair{node.Left, left})
			}
			if node.Right != nil {
				parent[node.Right] = node
				queue = append(queue, pair{node.Right, left})
			}
		}
	}
	return
}
