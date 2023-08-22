package Solution

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func Solution(root *TreeNode) *TreeNode {
	var dfs func(*TreeNode, int)
	maxDep := 0
	targetLeafNodes := 0
	nodeDep := make(map[int]int)
	dfs = func(now *TreeNode, dep int) {
		if now == nil {
			return
		}
		nodeDep[now.Val] = dep
		if now.Left == nil && now.Right == nil {
			if dep > maxDep {
				maxDep = dep
				targetLeafNodes = 1
			} else if dep == maxDep {
				targetLeafNodes++
			}
			return
		}
		dfs(now.Left, dep+1)
		dfs(now.Right, dep+1)
	}
	dfs(root, 1)

	var ans *TreeNode
	var dfs1 func(*TreeNode) int
	dfs1 = func(now *TreeNode) int {
		if now == nil {
			return 0
		}
		left := dfs1(now.Left)
		right := dfs1(now.Right)

		addOne := 0
		if nodeDep[now.Val] == maxDep {
			addOne = 1
		}
		r := left + right + addOne
		if r == targetLeafNodes {
			if ans == nil || nodeDep[ans.Val] < nodeDep[now.Val] {
				ans = now
			}
		}
		return r
	}
	dfs1(root)

	return ans
}
