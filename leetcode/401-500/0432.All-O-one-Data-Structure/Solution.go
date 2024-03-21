package Solution

import (
	"container/list"
)

type AllOne struct {
	list  *list.List
	data  map[string]int
	index map[int]*list.Element
}

func Constructor432() AllOne {
	a := AllOne{
		data:  make(map[string]int),
		index: make(map[int]*list.Element),
	}

	l := list.New()
	a.index[0] = l.PushBack(map[string]struct{}{})
	a.list = l
	return a
}

func (this *AllOne) Inc(key string) {
	source := this.data[key]

	node := this.index[source]
	next, ok := this.index[source+1]
	if !ok {
		next = this.list.InsertAfter(map[string]struct{}{}, node)
		this.index[source+1] = next
	}
	v := next.Value.(map[string]struct{})
	v[key] = struct{}{}
	next.Value = v
	if source > 0 {
		v1 := node.Value.(map[string]struct{})
		delete(v1, key)
		if len(v1) == 0 {
			this.list.Remove(node)
			delete(this.index, source)
		} else {
			node.Value = v1
		}
	}
	this.data[key] = source + 1
}

func (this *AllOne) Dec(key string) {
	source := this.data[key]
	if source == 0 {
		return
	}
	this.data[key]--
	if this.data[key] == 0 {
		delete(this.data, key)
	}
	node := this.index[source]
	if source > 1 {
		pre, ok := this.index[source-1]
		if !ok {
			pre = this.list.InsertBefore(map[string]struct{}{}, node)
			this.index[source-1] = pre
		}
		v1 := pre.Value.(map[string]struct{})
		v1[key] = struct{}{}
		pre.Value = v1
	}
	v := node.Value.(map[string]struct{})
	delete(v, key)
	if len(v) == 0 {
		this.list.Remove(node)
		delete(this.index, source)
	} else {
		node.Value = v
	}

}

func (this *AllOne) GetMinKey() string {
	if this.list.Len() == 1 {
		return ""
	}

	v := this.list.Front().Next().Value.(map[string]struct{})
	for k := range v {
		return k
	}
	return ""
}

func (this *AllOne) GetMaxKey() string {
	if this.list.Len() == 1 {
		return ""
	}
	v := this.list.Back().Value.(map[string]struct{})
	for k := range v {
		return k
	}
	return ""
}

type opt struct {
	name string
	v    string
}

func Solution(opts []opt) []string {
	c := Constructor432()
	ans := make([]string, 0)
	for _, o := range opts {
		if o.name == "i" {
			c.Inc(o.v)
			continue
		}
		if o.name == "d" {
			c.Dec(o.v)
			continue
		}
		if o.name == "max" {
			ans = append(ans, c.GetMaxKey())
			continue
		}
		ans = append(ans, c.GetMinKey())
	}
	return ans
}
