package Solution

import (
	"container/list"
	"fmt"
)

type LRUCache struct {
	Capacity int
	Data     *list.List
	//*list.List
}

func Constructor(capacity int) LRUCache {
	lru := LRUCache{
		Capacity: capacity,
		Data:     list.New(),
	}
	return lru
}

func (this *LRUCache) Get(key int) int {
	for e := this.Data.Front(); e != nil; e = e.Next() {

		//fmt.Println(key, e.Value.(list.Element).Value.(map[int]int)[key])

		if v, ok := e.Value.(list.Element).Value.(map[int]int)[key]; ok {
			this.Data.MoveBefore(e, this.Data.Front())
			return v
		}
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	this.Data.PushFront(
		list.Element{
			Value: map[int]int{key: value},
		},
	)
	if this.Data.Len() > this.Capacity {
		this.Data.Remove(this.Data.Back())
	}
}

func (this *LRUCache) Print() {
	for e := this.Data.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value.(list.Element).Value.(map[int]int))
	}
}
