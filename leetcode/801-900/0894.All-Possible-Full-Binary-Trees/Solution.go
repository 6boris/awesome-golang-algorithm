package Solution

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func Solution(n int) []*TreeNode {
	cache := make([][]*TreeNode, 21)
	cache[1] = []*TreeNode{{Val: 0}}
	cache[3] = []*TreeNode{{Val: 0, Left: &TreeNode{Val: 0}, Right: &TreeNode{Val: 0}}}

	var buildTree func(left int) []*TreeNode
	buildTree = func(left int) []*TreeNode {
		result := make([]*TreeNode, 0)
		for leftNode := 1; leftNode < left; leftNode += 2 {
			for li := 0; li < len(cache[leftNode]); li++ {
				for ri := 0; ri < len(cache[left-leftNode-1]); ri++ {
					result = append(result, &TreeNode{Val: 0, Left: cache[leftNode][li], Right: cache[left-leftNode-1][ri]})
				}
			}
		}
		return result
	}

	for start := 5; start <= n; start += 2 {
		cache[start] = buildTree(start)
	}
	return cache[n]
}
