package Solution

type MinStack struct {
	stack    []int
	minStack []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		stack:    make([]int, 0),
		minStack: make([]int, 0),
	}
}

func (s *MinStack) Push(x int) {
	s.stack = append(s.stack, x)
	if len(s.minStack) == 0 || x < s.minStack[len(s.minStack)-1] {
		s.minStack = append(s.minStack, x)
	} else {
		s.minStack = append(s.minStack, s.minStack[len(s.minStack)-1])
	}
}

func (s *MinStack) Pop() {
	s.stack = s.stack[:len(s.stack)-1]
	s.minStack = s.minStack[:len(s.minStack)-1]
}

func (s *MinStack) Top() int {
	return s.stack[len(s.stack)-1]
}

func (s *MinStack) Min() int {
	return s.minStack[len(s.minStack)-1]
}
