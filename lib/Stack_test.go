package Solution

import (
	"fmt"
	"testing"
)

func TestNewStack(t *testing.T) {
	stk := NewStack()
	stk.Push(1)
	stk.Push(2)
	stk.Push(3)
	stk.Push(4)
	stk.Push(5)
	fmt.Println(stk.Pop())
	fmt.Println(stk.Pop())
	fmt.Println(stk.Pop())
	fmt.Println(stk.Pop())
	fmt.Println(stk.Pop())
	fmt.Println(stk.Empty())
}
