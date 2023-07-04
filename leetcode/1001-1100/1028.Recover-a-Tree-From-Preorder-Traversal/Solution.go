package Solution

import "strconv"

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func Solution(traversal string) *TreeNode {
	curDep, backIndex := 0, 0
	return recoverTree(traversal, 0, 0, &curDep, &backIndex)
}

func recoverTree(traversal string, index, nodeDep int, curDep, backIndex *int) *TreeNode {
	if index >= len(traversal) {
		return nil
	}
	node := &TreeNode{}
	next := index + 1
	for ; next < len(traversal) && traversal[next] >= '0' && traversal[next] <= '9'; next++ {
	}
	v, _ := strconv.Atoi(traversal[index:next])
	node.Val = v
	dep := 0
	for ; next < len(traversal) && traversal[next] == '-'; next++ {
		dep++
	}

	*curDep = dep
	*backIndex = next
	if dep > nodeDep {
		node.Left = recoverTree(traversal, next, dep, curDep, backIndex)
	}
	if nodeDep < *curDep {
		node.Right = recoverTree(traversal, *backIndex, *curDep, curDep, backIndex)
	}
	return node
}
