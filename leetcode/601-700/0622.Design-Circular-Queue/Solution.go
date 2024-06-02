package Solution

type node622 struct {
	val       int
	next, pre *node622
}

type node622List struct {
	root  *node622
	count int
}

func (n *node622List) len() int {
	return n.count
}

func (n *node622List) pushEnd(v int) {
	node := &node622{val: v}
	n.count++
	if n.count == 1 {
		node.next = node
		node.pre = node
		n.root = node
		return
	}
	n.root.pre.next = node
	node.pre = n.root.pre
	node.next = n.root
	n.root.pre = node
}

func (n *node622List) removeFront() {
	n.count--
	if n.count == 0 {
		n.root = nil
		return
	}
	// a,b,c
	newRoot := n.root.next
	newRoot.pre = n.root.pre
	n.root.pre.next = newRoot
	n.root = newRoot
}

func (n *node622List) Front() int {
	if n.count == 0 {
		return -1
	}
	return n.root.val
}
func (n *node622List) Back() int {
	if n.count == 0 {
		return -1
	}
	return n.root.pre.val
}

type MyCircularQueue struct {
	root  *node622List
	count int
}

func Constructor(k int) MyCircularQueue {
	return MyCircularQueue{root: &node622List{}, count: k}
}

func (this *MyCircularQueue) EnQueue(value int) bool {
	if this.root.count == this.count {
		return false
	}
	this.root.pushEnd(value)
	return true
}

func (this *MyCircularQueue) DeQueue() bool {
	if this.root.count == 0 {
		return false
	}
	this.root.removeFront()
	return true
}

func (this *MyCircularQueue) Front() int {
	return this.root.Front()
}

func (this *MyCircularQueue) Rear() int {
	return this.root.Back()
}

func (this *MyCircularQueue) IsEmpty() bool {
	return this.root.len() == 0
}

func (this *MyCircularQueue) IsFull() bool {
	return this.root.len() == this.count
}

type opt struct {
	name string
	val  int
}

func Solution(n int, opts []opt) []any {
	c := Constructor(n)
	ans := make([]any, len(opts))
	for i, op := range opts {
		if op.name == "en" {
			ans[i] = c.EnQueue(op.val)
			continue
		}
		if op.name == "de" {
			ans[i] = c.DeQueue()
			continue
		}
		if op.name == "fr" {
			ans[i] = c.Front()
			continue
		}
		if op.name == "re" {
			ans[i] = c.Rear()
			continue
		}
		if op.name == "em" {
			ans[i] = c.IsEmpty()
			continue
		}
		ans[i] = c.IsFull()
	}
	return ans
}
