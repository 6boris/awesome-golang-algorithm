package Solution

import (
	"container/heap"
	"fmt"
	"testing"
)

func TestIntHeap_Less(t *testing.T) {
	h := &IntMaxHeap{}
	heap.Init(h)
	heap.Push(h, 7)
	heap.Push(h, 3)
	heap.Push(h, 2)
	heap.Push(h, 1)
	heap.Push(h, 5)
	heap.Push(h, 5)
	heap.Push(h, 6)
	heap.Push(h, 7)
	fmt.Printf("minimum: %d\n", (*h))


	for h.Len() > 0 {
		fmt.Printf("%d ", heap.Pop(h))
	}
	//fmt.Printf("minimum: %d\n", (*h))
	//fmt.Printf("minimum: %d\n", (*h)[2])
}
