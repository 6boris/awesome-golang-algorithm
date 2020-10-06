package Solution

import "fmt"

type MinSlack struct {
	stack    Stack
	stackMin Stack
}

func (s *MinSlack) push(v int) {
	s.stack.Push(v)
	if s.stackMin.IsEmpty() {
		s.stackMin.Push(v)
	} else {
		//	入 stack 元素是最小的
		if minV, ok := s.stackMin.Top(); ok == nil && minV.(int) >= v {
			fmt.Println("push", v)
			s.stackMin.Push(v)
		} else {
			val, _ := s.stackMin.Top()
			fmt.Println("p", val)
			s.stackMin.Push(val)
		}
	}
}
func (s *MinSlack) pop() {
	_, _ = s.stack.Pop()
	_, _ = s.stackMin.Pop()
}
func (s *MinSlack) top() (interface{}, error) {
	return s.stack.Top()
}

func (s *MinSlack) min() (interface{}, error) {
	return s.stackMin.Top()
}
