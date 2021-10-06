package Solution

type LRUCache struct {
	size, capacity int
	m              map[int]*Node
	head, tail     *Node
}

type Node struct {
	key, val   int
	prev, next *Node
}

func Constructor(capacity int) LRUCache {
	head, tail := &Node{key: 0, val: 0}, &Node{key: 0, val: 0}
	head.next = tail
	tail.prev = head
	return LRUCache{
		size:     0,
		capacity: capacity,
		head:     head,
		tail:     tail,
		m:        map[int]*Node{},
	}
}

func (this *LRUCache) Get(key int) int {
	if v, ok := this.m[key]; ok {
		this.removeNode(v)
		this.moveHead(v)
		return v.val
	}

	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if v, ok := this.m[key]; ok {
		v.val = value
		this.removeNode(v)
		this.moveHead(v)
	} else {
		node := &Node{key: key, val: value}
		this.m[key] = node
		this.size++
		this.moveHead(node)
		if this.size > this.capacity {
			tmp := this.tail.prev
			this.removeNode(tmp)
			this.size--
			delete(this.m, tmp.key)
		}
	}
}
func (this *LRUCache) moveHead(node *Node) {
	this.head.next.prev = node
	node.next = this.head.next
	node.prev = this.head
	this.head.next = node
}
func (this *LRUCache) removeNode(node *Node) {
	node.next.prev = node.prev
	node.prev.next = node.next
}
