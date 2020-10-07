package Solution

import (
	"fmt"
	"testing"
)

// TestSolution Example for solution test cases
func TestSolution(t *testing.T) {
	obj := Constructor()
	obj.Push(2)
	obj.Push(3)
	obj.Push(1)
	obj.Pop()
	param_3 := obj.Top()
	param_4 := obj.Min()
	fmt.Println(param_3, param_4)
}
