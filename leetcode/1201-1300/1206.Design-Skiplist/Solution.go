package Solution

import "fmt"

type SkipListNode struct {
	Value   int
	Count   int
	Forward []*SkipListNode
}

const (
	level = 8
)

type Skiplist struct {
	DummyHead *SkipListNode
	loop      int
}

func newSkipListNodeWithValue(val int) *SkipListNode {
	return &SkipListNode{Value: val, Forward: make([]*SkipListNode, level), Count: 1}
}

func Constructor() Skiplist {
	return Skiplist{
		DummyHead: newSkipListNodeWithValue(-1),
		loop:      -1,
	}
}

func (this *Skiplist) search(target int) *SkipListNode {
	walker := this.DummyHead
	for l := level - 1; l >= 0; l-- {
		for walker.Forward[l] != nil {
			if walker.Forward[l].Value == target {
				return walker.Forward[l]
			}
			if walker.Forward[l].Value < target {
				walker = walker.Forward[l]
				continue
			}
			break
		}
	}
	return nil
}
func (this *Skiplist) Search(target int) bool {
	return this.search(target) != nil
}

func (this *Skiplist) Add(num int) {
	addLevel := (this.loop + 1) % level
	update := make([]*SkipListNode, level)
	walker := this.DummyHead
	for l := addLevel; l >= 0; l-- {
		for walker.Forward[l] != nil {
			if walker.Forward[l].Value == num {
				walker.Forward[l].Count++
				return
			}
			if walker.Forward[l].Value < num {
				walker = walker.Forward[l]
				continue
			}
			break
		}
		// update end idx
		update[l] = walker
	}
	addNode := newSkipListNodeWithValue(num)

	for i := 0; i <= addLevel; i++ {
		if update[i] == nil {
			this.DummyHead.Forward[i] = addNode
			continue
		}
		addNode.Forward[i] = update[i].Forward[i]
		update[i].Forward[i] = addNode
	}

	this.loop = addLevel
}

func (this *Skiplist) Erase(num int) bool {
	walker := this.DummyHead
	del := make([]*SkipListNode, level)
	found := -1
	for l := level - 1; l >= 0; l-- {
		for walker.Forward[l] != nil {
			if walker.Forward[l].Value < num {
				walker = walker.Forward[l]
				continue
			}
			if walker.Forward[l].Value == num {
				if found == -1 {
					found = l
				}

				if walker.Forward[l].Count > 1 {
					// same value
					walker.Forward[l].Count--
					return true
				}
			}
			break
		}
		del[l] = walker
	}
	for idx := 0; idx <= found; idx++ {
		if del[idx] == nil {
			continue
		}
		next := del[idx].Forward[idx]
		if next == nil {
			del[idx].Forward[idx] = nil
			continue
		}
		del[idx].Forward[idx] = next.Forward[idx]
	}

	return found != -1
}

func Solution(operations, ans []string, nums []int) bool {
	o := Constructor()
	for idx, operation := range operations {
		if operation == "add" {
			o.Add(nums[idx])
			continue
		}
		get := false
		if operation == "erase" {
			get = o.Erase(nums[idx])
		} else {
			get = o.Search(nums[idx])
		}
		getStr := fmt.Sprintf("%v", get)
		if getStr != ans[idx] {
			return false
		}
	}
	return true
}
