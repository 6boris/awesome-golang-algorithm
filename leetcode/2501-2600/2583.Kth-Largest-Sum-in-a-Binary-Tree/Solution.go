package Solution

import "container/heap"

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}
type list2583 []int64

func (l *list2583) Len() int {
	return len(*l)
}

func (l *list2583) Swap(i, j int) {
	(*l)[i], (*l)[j] = (*l)[j], (*l)[i]
}

func (l *list2583) Less(i, j int) bool {
	return (*l)[i] < (*l)[j]
}

func (l *list2583) Push(x interface{}) {
	*l = append(*l, x.(int64))
}

func (l *list2583) Pop() interface{} {
	old := *l
	ll := len(old)
	x := old[ll-1]
	*l = old[:ll-1]
	return x
}

func Solution(root *TreeNode, k int) int64 {
	list := list2583{}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		next := make([]*TreeNode, 0)
		sum := int64(0)
		for _, item := range queue {
			sum += int64(item.Val)
			if item.Left != nil {
				next = append(next, item.Left)
			}
			if item.Right != nil {
				next = append(next, item.Right)
			}
		}

		queue = next
		if len(list) < k {
			heap.Push(&list, sum)
		} else {
			top := heap.Pop(&list).(int64)
			if top < sum {
				top = sum
			}
			heap.Push(&list, top)
		}
	}
	if len(list) != k {
		return -1
	}
	return heap.Pop(&list).(int64)
}
