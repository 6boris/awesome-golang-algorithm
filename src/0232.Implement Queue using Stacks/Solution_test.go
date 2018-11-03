package Solution

import (
	"testing"
)

func TestSolution(t *testing.T) {
	queue := &MyQueue{}
	for ix := 0; ix < 5; ix++ {
		queue.Push(ix)
	}
	pv := queue.Peek()
	if pv != 0 {
		t.Fatalf("Peek: expected to be %v, but got %v", 0, pv)
	}
	pv = queue.Pop()
	if pv != 0 {
		t.Fatalf("Pop: expected to be %v, but got %v", 0, pv)
	}
	queue.Push(5)
	pv = queue.Pop()
	if pv != 1 {
		t.Fatalf("Pop: expected to be %v, but got %v", 1, pv)
	}
	queue.Pop()
	queue.Pop()
	queue.Pop()
	// queue.Pop()
	if !queue.Empty() {
		t.Fatalf("Empty: expected to be %v, but got %v", true, false)
	}
}
