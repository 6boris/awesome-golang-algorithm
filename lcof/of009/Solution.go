package Solution

import "container/list"

type CQueue struct {
	stack1, stack2 *list.List
}

func Constructor() CQueue {
	return CQueue{
		stack1: list.New(),
		stack2: list.New(),
	}
}

func (q *CQueue) AppendTail(value int) {
	q.stack1.PushBack(value)
}

func (q *CQueue) DeleteHead() int {
	// 如果第二个栈为空
	if q.stack2.Len() == 0 {
		for q.stack1.Len() > 0 {
			q.stack2.PushBack(q.stack1.Remove(q.stack1.Back()))
		}
	}
	if q.stack2.Len() != 0 {
		e := q.stack2.Back()
		q.stack2.Remove(e)
		return e.Value.(int)
	}
	return -1
}
