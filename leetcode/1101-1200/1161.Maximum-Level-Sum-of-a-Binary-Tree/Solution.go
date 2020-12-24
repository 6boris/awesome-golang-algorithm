package Solution

import (
	"fmt"
)

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

// DFS 递归的中序遍历
func maxLevelSum2(root *TreeNode) int {
	m, ans := make(map[int]int), -1
	dfs(root, 1, m)
	fmt.Println(m)
	for _, v := range m {
		if v > ans {
			ans = v
		}
		print(v)
	}
	return ans
}

func dfs(root *TreeNode, level int, m map[int]int) {
	if root == nil {
		return
	}
	fmt.Println(level, root.Val)
	dfs(root.Left, level+1, m)
	m[level] += root.Val
	dfs(root.Right, level+1, m)
}
