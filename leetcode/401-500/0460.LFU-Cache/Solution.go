package Solution

import (
	"container/heap"
	"fmt"
)

// 不是最优解法
type lfuItem struct {
	k int
	v int
	c int
	p int // 优先级
	i int // 堆中位置
}

type lfuItems []*lfuItem

func (li *lfuItems) String() {
	fmt.Println()
	for i := 0; i < len(*li); i++ {
		fmt.Printf("%+v\n", (*li)[i])
	}
}
func (li *lfuItems) Len() int {
	return len(*li)
}

func (li *lfuItems) Swap(i, j int) {
	(*li)[i], (*li)[j] = (*li)[j], (*li)[i]
	(*li)[i].i, (*li)[j].i = i, j
}

func (li *lfuItems) Less(i, j int) bool {
	if (*li)[i].c == (*li)[j].c {
		return (*li)[i].p < (*li)[j].p
	}
	return (*li)[i].c < (*li)[j].c
}

func (li *lfuItems) Push(x interface{}) {
	*li = append(*li, x.(*lfuItem))
}

func (li *lfuItems) Pop() interface{} {
	old := *li
	l := len(old)
	x := old[l-1]
	*li = old[:l-1]
	return x
}

type LFUCache struct {
	capSize int
	store   map[int]*lfuItem
	h       lfuItems
	next    int
}

func Constructor460(capacity int) LFUCache {
	return LFUCache{
		capSize: capacity,
		store:   make(map[int]*lfuItem),
		h:       lfuItems{},
		next:    0,
	}
}

func (this *LFUCache) Get(key int) int {
	v, ok := this.store[key]
	if !ok {
		return -1
	}

	v.c++
	v.p = this.next
	this.next++
	heap.Fix(&this.h, v.i)
	return v.v
}

func (this *LFUCache) Put(key int, value int) {
	if this.capSize == 0 {
		return
	}
	v, ok := this.store[key]
	if ok {
		v.v = value
		v.c++
		v.p = this.next
		this.next++
		heap.Fix(&this.h, v.i)
		return
	}

	var top *lfuItem
	if len(this.store) == this.capSize {
		top = heap.Pop(&this.h).(*lfuItem)
		delete(this.store, top.k)
	}
	targetIndex := len(this.store)
	if top != nil {
		targetIndex = top.i
	}
	n := &lfuItem{k: key, v: value, c: 1, i: targetIndex, p: this.next}
	this.next++
	heap.Push(&this.h, n)
	this.store[key] = n
}

func Solution(size int, ops []string, values [][]int) []int {
	ans := make([]int, 0)
	cr := Constructor460(size)
	for i, op := range ops {
		if op == "put" {
			cr.Put(values[i][0], values[i][1])
			continue
		}
		ans = append(ans, cr.Get(values[i][0]))
	}
	return ans
}
