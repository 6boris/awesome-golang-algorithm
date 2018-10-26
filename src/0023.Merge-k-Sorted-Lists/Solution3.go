package Solution

import "container/heap"

func mergeKLists3(lists []*ListNode) *ListNode {
	pq := make(ListNodeQueue, len(lists))
	for i := range lists {
		pq[i] = lists[i]
	}
	heap.Init(&pq)
	if pq.Len() < 1 {
		return nil
	}

	result := heap.Pop(&pq).(*ListNode)
	if result == nil {
		return result
	}
	if result.Next != nil {
		heap.Push(&pq, result.Next)
	}
	tail := result
	for pq.Len() > 0 {
		tail.Next = heap.Pop(&pq).(*ListNode)
		tail = tail.Next
		if tail == nil {
			break
		}
		if tail.Next != nil {
			heap.Push(&pq, tail.Next)
		}
	}

	return result
}
