package Solution

type ListNodeQueue []*ListNode

func (que ListNodeQueue) Len() int { return len(que) }

func (que ListNodeQueue) Less(i, j int) bool {
	//if que[i] == nil {
	//	return false
	//}
	//if que[j] == nil {
	//	return true
	//}
	return que[i].Val < que[j].Val
}

func (que ListNodeQueue) Swap(i, j int) {
	que[i], que[j] = que[j], que[i]
}

func (que *ListNodeQueue) Push(x interface{}) {
	item := x.(*ListNode)
	*que = append(*que, item)
}

func (que *ListNodeQueue) Pop() interface{} {
	old := *que
	n := len(old)
	item := old[n-1]
	*que = old[0 : n-1]
	return item
}
