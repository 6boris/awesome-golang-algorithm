package Solution

import (
	"testing"
)

func TestSolution(t *testing.T) {
	stack := &MyStack{}
	for ix := 0; ix < 5; ix++ {
		stack.Push(ix)
	}
	pv := stack.Pop()

	if pv != 4 {
		t.Fatalf("Expected pv to be %v, but got %v", 0, pv)
	}
	stack.Push(6)
	stack.Pop()
	pv = stack.Pop()
	if pv != 2 {
		t.Fatalf("Pop Test: Expected pv to be %v, but got %v", 2, pv)
	}
	pv = stack.Top()
	if pv != 1 {
		t.Fatalf("Top Test: Expected pv to be %v,  but got %v", 1, pv)
	}
}

func TestSolution1(t *testing.T) {
	stack := Constructor225()
	for ix := 0; ix < 5; ix++ {
		stack.Push(ix)
	}
	pv := stack.Pop()

	if pv != 4 {
		t.Fatalf("Expected pv to be %v, but got %v", 0, pv)
	}
	stack.Push(6)
	stack.Pop()
	pv = stack.Pop()
	if pv != 3 {
		t.Fatalf("Pop Test: Expected pv to be %v, but got %v", 2, pv)
	}
	pv = stack.Top()
	if pv != 2 {
		t.Fatalf("Top Test: Expected pv to be %v,  but got %v", 1, pv)
	}
}
