package Solution

type Queue []*ListNode

func (q Queue) Len() int {
	return len(q)
}

func (q Queue) Less(i, j int) bool {
		if q[i] == nil {
			return false
		}
		if q[j] == nil {
			return true
		}
	return q[i].Val < q[j].Val
}

func (q Queue) Swap(i, j int) {
	q[i], q[j] = q[j], q[i]
}

func (q *Queue) Push(x interface{}) {
	item := x.(*ListNode)
	*q = append(*q, item)
}

func (q *Queue) Pop() interface{} {
	old := *q
	n := len(old)
	item := old[n-1]
	*q = old[0 : n-1]
	return item
}
