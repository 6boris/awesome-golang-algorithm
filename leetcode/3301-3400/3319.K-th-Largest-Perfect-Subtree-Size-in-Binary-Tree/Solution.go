package Solution

import "sort"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func Solution(root *TreeNode, k int) int {
	var dfs func(*TreeNode) int
	count := make(map[int]int)
	dfs = func(cur *TreeNode) int {
		if cur == nil {
			return -1
		}
		if cur.Left == nil && cur.Right == nil {
			count[1]++
			// leaf
			return 1
		}
		left := dfs(cur.Left)
		right := dfs(cur.Right)

		if left != -1 && left == right {
			count[left+right+1]++
			return left + right + 1
		}
		return -1
	}
	dfs(root)
	keys := make([]int, 0)
	for i := range count {
		keys = append(keys, i)
	}

	sort.Slice(keys, func(i, j int) bool {
		return keys[i] > keys[j]
	})
	var cnt int
	for _, n := range keys {
		cnt = count[n]
		if cnt >= k {
			return n
		}
		k -= cnt
	}
	return -1
}
