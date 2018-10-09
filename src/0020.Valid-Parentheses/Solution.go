package Solution

import (
	"container/list"
)

func isValid(s string) bool {
	stk := GetNormalStack()
	for i := 0; i < len(s); i++ {
		//	左括号时入栈
		if s[i] == '(' || s[i] == '[' || s[i] == '{' {
			stk.Push(s[i])
			//	为右括号是时弹出栈顶和当前字符对比
		} else if s[i] == ')' && stk.Pop() != uint8('(') {
			return false
		} else if s[i] == ']' && stk.Pop() != uint8('[') {
			return false
		} else if s[i] == '}' && stk.Pop() != uint8('{') {
			return false
		}
	}
	return stk.IsEmpty()
}

type normalStack struct {
	data list.List
}

//	入栈
func (s *normalStack) Push(data interface{}) {
	s.data.PushBack(data)
}

//	出栈
func (s *normalStack) Pop() interface{} {
	if s.IsEmpty() {
		return nil
	}
	data := s.data.Back().Value
	s.data.Remove(s.data.Back())
	return data
}

//	获取栈顶元素
func (s *normalStack) Top() interface{} {
	if s.IsEmpty() {
		return nil
	}

	return s.data.Back().Value
}

//	获取栈中元素的数量
func (s *normalStack) Len() int {
	return s.data.Len()
}

//	判断栈是否为空
func (s *normalStack) IsEmpty() bool {
	if s.data.Len() == 0 {
		return true
	}
	return false
}

func GetNormalStack() *normalStack {
	return &normalStack{}
}
