package Solution

import "container/heap"

func Solution(nums []int) int {
	pq := &PriorityQueue{}
	heap.Init(pq)
	merged := make([]bool, len(nums))
	decreaseCount := 0
	count := 0
	head := &Node{value: int64(nums[0]), left: 0}
	current := head

	for i := 1; i < len(nums); i++ {
		newNode := &Node{value: int64(nums[i]), left: i}
		current.next = newNode
		newNode.prev = current

		heap.Push(pq, &Item{
			first:  current,
			second: newNode,
			cost:   current.value + newNode.value,
		})
		if nums[i-1] > nums[i] {
			decreaseCount++
		}

		current = newNode
	}

	for decreaseCount > 0 {
		item := heap.Pop(pq).(*Item)
		first := item.first
		second := item.second
		cost := item.cost

		if merged[first.left] || merged[second.left] || first.value+second.value != cost {
			continue
		}
		count++
		if first.value > second.value {
			decreaseCount--
		}

		prevNode := first.prev
		nextNode := second.next
		first.next = nextNode
		if nextNode != nil {
			nextNode.prev = first
		}

		if prevNode != nil {
			if prevNode.value > first.value && prevNode.value <= cost {
				decreaseCount--
			} else if prevNode.value <= first.value && prevNode.value > cost {
				decreaseCount++
			}
			heap.Push(pq, &Item{
				first:  prevNode,
				second: first,
				cost:   prevNode.value + cost,
			})
		}

		if nextNode != nil {
			if second.value > nextNode.value && cost <= nextNode.value {
				decreaseCount--
			} else if second.value <= nextNode.value && cost > nextNode.value {
				decreaseCount++
			}
			heap.Push(pq, &Item{
				first:  first,
				second: nextNode,
				cost:   cost + nextNode.value,
			})
		}

		first.value = cost
		merged[second.left] = true
	}

	return count
}

type Node struct {
	value int64
	left  int
	prev  *Node
	next  *Node
}

type Item struct {
	first  *Node
	second *Node
	cost   int64
	index  int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool {
	if pq[i].cost == pq[j].cost {
		return pq[i].first.left < pq[j].first.left
	}
	return pq[i].cost < pq[j].cost
}
func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}
func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.index = -1
	*pq = old[0 : n-1]
	return item
}
