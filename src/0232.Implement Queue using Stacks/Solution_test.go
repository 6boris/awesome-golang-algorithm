package Solution

import (
	"testing"
)

func TestSolution(t *testing.T) {
	stack := &MyQueue{
		array: []int{12, 2, 3, 5, 6},
	}
	cloned := &MyQueue{}
	for ix := 1; ix < len(stack.array); ix++ {
		cloned.array = append(cloned.array, stack.array[ix])
	}
	t.Run("Peek Test", func(t *testing.T) {
		pk := stack.Peek()
		if pk != 12 {
			t.Fatalf("expected: %v, but got: %v", 12, pk)
		}
	})

	t.Run("Pop Test", func(t *testing.T) {
		pk := stack.Pop()
		// 值判断
		if pk != 12 {
			t.Fatalf("expected: %v, but got: %v", 12, pk)
		}
		// 校验队列
		for ix := 0; ix < len(stack.array); ix++ {
			if stack.array[ix] != cloned.array[ix] {
				t.Fatalf("The sequence is somewhat errors")
				break
			}
		}
	})
}
