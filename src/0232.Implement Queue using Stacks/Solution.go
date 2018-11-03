package Solution

type MyQueue struct {
	array []int
}

func Constructor() MyQueue {
	return MyQueue{}
}

func (stack *MyQueue) delet() int {
	temp := stack.array[len(stack.array)-1]
	stack.array = stack.array[:len(stack.array)-1]
	return temp
}

func (stack *MyQueue) Peek() int {
	if len(stack.array) == 0 {
		return 0
	}
	copied := MyQueue{}
	for len(stack.array) != 0 {
		copied.array = append(copied.array, stack.delet())
	}
	temp := copied.array[len(copied.array)-1]
	for len(copied.array) != 0 {
		stack.array = append(stack.array, copied.delet())
	}
	return temp
}

func (stack *MyQueue) Push(val int) {
	stack.array = append(stack.array, val)
}

func (stack *MyQueue) Pop() int {
	if len(stack.array) == 0 {
		return 0
	}
	copied := MyQueue{}
	for len(stack.array) != 1 {
		copied.array = append(copied.array, stack.delet())
	}
	temp := stack.delet()
	for len(copied.array) != 0 {
		stack.array = append(stack.array, copied.delet())
	}
	return temp
}

func (stack *MyQueue) Empty() bool {
	return len(stack.array) == 0
}
