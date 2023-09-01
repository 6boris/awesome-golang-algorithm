package Solution

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func dfs1530(root *TreeNode, dep, disLimit int, ans *int) map[int]int {
	if root == nil {
		return nil
	}
	h := make(map[int]int)
	if root.Left == nil && root.Right == nil {
		h[dep] = 1
		return h
	}
	left := dfs1530(root.Left, dep+1, disLimit, ans)
	right := dfs1530(root.Right, dep+1, disLimit, ans)
	for ld, lc := range left {
		h[ld] += lc
	}
	for ld, lc := range right {
		h[ld] += lc
	}
	for ld, lc := range left {
		for rd, rc := range right {
			if ld-dep+rd-dep <= disLimit {
				*ans += lc * rc
			}
		}
	}
	return h
}
func Solution(root *TreeNode, distance int) int {
	ans := 0
	dfs1530(root, 0, distance, &ans)
	return ans
}
