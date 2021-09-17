package Solution

import (
	"container/list"
)

type Stack struct {
	data list.List
}

//	入栈
func (s *Stack) Push(data interface{}) {
	s.data.PushBack(data)
}

func (s *Stack) Pop() interface{} {
	data := s.data.Back().Value
	s.data.Remove(s.data.Back())
	return data
}

func (s *Stack) Top() interface{} {
	return s.data.Back().Value
}

func (s *Stack) Len() int {
	return s.data.Len()
}

func (s *Stack) IsEmpty() bool {
	if s.data.Len() == 0 {
		return true
	}
	return false
}

func NewStack() {
}
