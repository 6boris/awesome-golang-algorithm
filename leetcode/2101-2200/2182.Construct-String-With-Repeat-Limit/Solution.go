package Solution

import (
	"bytes"
	"container/heap"
)

type tmp2182 struct {
	b byte
	c int
}
type heap2182 []tmp2182

func (h *heap2182) Len() int {
	return len(*h)
}

func (h *heap2182) Less(i, j int) bool {
	return (*h)[i].b > (*h)[j].b
}

func (h *heap2182) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *heap2182) Push(x any) {
	*h = append(*h, x.(tmp2182))
}

func (h *heap2182) Pop() any {
	old := *h
	l := len(old)
	x := old[l-1]
	*h = old[:l-1]
	return x
}

func Solution(s string, repeatLimit int) string {
	c := make(map[byte]int)
	for _, b := range []byte(s) {
		c[b]++
	}
	list := &heap2182{}
	for k, v := range c {
		heap.Push(list, tmp2182{k, v})
	}
	buf := bytes.NewBuffer([]byte{})

	for list.Len() > 0 {
		top := heap.Pop(list).(tmp2182)
		for range min(top.c, repeatLimit) {
			buf.WriteByte(top.b)
		}
		top.c -= repeatLimit
		if top.c > 0 && list.Len() > 0 {
			nextPickOne := heap.Pop(list).(tmp2182)
			buf.WriteByte(nextPickOne.b)
			nextPickOne.c--
			if nextPickOne.c > 0 {
				heap.Push(list, nextPickOne)
			}
			heap.Push(list, top)
		}
	}
	return buf.String()
}
