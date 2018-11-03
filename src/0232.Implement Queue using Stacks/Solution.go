package Solution

type MyQueue struct {
	array []int
}

func Constructor() MyQueue {
	return MyQueue{}
}

func (queue *MyQueue) delet() int {
	temp := queue.array[len(queue.array)-1]
	queue.array = queue.array[:len(queue.array)-1]
	return temp
}

func (queue *MyQueue) Peek() int {
	if len(queue.array) == 0 {
		return 0
	}
	copied := MyQueue{}
	for len(queue.array) != 0 {
		copied.array = append(copied.array, queue.delet())
	}
	temp := copied.array[len(copied.array)-1]
	for len(copied.array) != 0 {
		queue.array = append(queue.array, copied.delet())
	}
	return temp
}

func (queue *MyQueue) Push(val int) {
	queue.array = append(queue.array, val)
}

func (queue *MyQueue) Pop() int {
	if len(queue.array) == 0 {
		return 0
	}
	copied := MyQueue{}
	for len(queue.array) != 1 {
		copied.array = append(copied.array, queue.delet())
	}
	temp := queue.delet()
	for len(copied.array) != 0 {
		queue.array = append(queue.array, copied.delet())
	}
	return temp
}

func (queue *MyQueue) Empty() bool {
	return len(queue.array) == 0
}
