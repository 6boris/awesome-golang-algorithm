package Solution

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxLevelSum(root *TreeNode) int {
	tree := make(map[*TreeNode]int)
	tree[root] = 0
	var treeList []*TreeNode
	treeList = append(treeList, root)
	sum := root.Val
	l := 0
	var sumList []int
	level := tree[root]
	for len(treeList) > 0 {
		r := treeList[0]
		treeList = treeList[1:]

		if tree[r] > level {
			sumList = append(sumList, sum)
			sum = 0
			level = tree[r]
		}
		sum += r.Val

		if r.Left != nil {
			treeList = append(treeList, r.Left)
			tree[r.Left] = level + 1
		}

		if r.Right != nil {
			treeList = append(treeList, r.Right)
			tree[r.Right] = level + 1
		}

	}

	sumList = append(sumList, sum)
	max := sumList[0]
	for i, j := range sumList {
		if j > max {
			max = j
			l = i
		}
	}
	return l + 1
}
