package Solution

import "container/heap"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

const base = 1000

type nodePointer struct {
	Val int
	Row int
}

type nodePointerList []nodePointer

func (npl *nodePointerList) Len() int {
	return len(*npl)
}

func (npl *nodePointerList) Less(i, j int) bool {
	if (*npl)[i].Row == (*npl)[j].Row {
		return (*npl)[i].Val < (*npl)[j].Val
	}
	return (*npl)[i].Row < (*npl)[j].Row
}

func (npl *nodePointerList) Swap(i, j int) {
	(*npl)[i], (*npl)[j] = (*npl)[j], (*npl)[i]
}

func (npl *nodePointerList) Push(x interface{}) {
	*npl = append(*npl, x.(nodePointer))
}

func (npl *nodePointerList) Pop() interface{} {
	old := *npl
	n := len(old)
	x := old[n-1]
	*npl = old[:n-1]
	return x
}

func Solution(root *TreeNode) [][]int {
	ans := make([][]int, 0)

	if root == nil {
		return ans
	}
	columns := make([]*nodePointerList, base*2+1)
	for idx := 0; idx < base*2+1; idx++ {
		columns[idx] = &nodePointerList{}
	}
	var inOrder func(*TreeNode, int, int)
	inOrder = func(root *TreeNode, row int, col int) {
		if root.Left != nil {
			inOrder(root.Left, row+1, col-1)
		}

		heap.Push(columns[col+base], nodePointer{Val: root.Val, Row: row})
		if root.Right != nil {
			inOrder(root.Right, row+1, col+1)
		}
	}

	inOrder(root, 0, 0)
	for idx := 0; idx < base*2+1; idx++ {
		if columns[idx].Len() == 0 {
			continue
		}
		xs := make([]int, 0)
		for columns[idx].Len() > 0 {
			x := heap.Pop(columns[idx]).(nodePointer)
			xs = append(xs, x.Val)
		}
		ans = append(ans, xs)
	}
	return ans
}
