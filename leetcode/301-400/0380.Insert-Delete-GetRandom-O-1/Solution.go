package Solution

import "math/rand"

// 这道题先想的是map+双向链表，最后发现getRandom是过不去的
// 再就是考虑直接用两个数组，大小都是MAXUINT32+1, 直接将int转换成uint32, 正负数都是直接对应下标, 这个内存开的有点过于不合理了。后来改成了map
type RandomizedSet struct {
	bucket map[int]int
	stack  []int
	index  int
}

func Constructor380() RandomizedSet {
	return RandomizedSet{
		bucket: make(map[int]int),
		stack:  make([]int, 0),
		index:  -1,
	}
}

func (this *RandomizedSet) Insert(val int) bool {
	_, ok := this.bucket[val]
	if !ok {
		if this.index == len(this.stack)-1 {
			this.stack = append(this.stack, val)
			this.index++
		} else {
			this.index++
			this.stack[this.index] = val
		}
		this.bucket[val] = this.index
	}
	return !ok
}

func (this *RandomizedSet) Remove(val int) bool {
	removedIndex, ok := this.bucket[val]
	if ok {
		this.bucket[this.stack[this.index]] = removedIndex
		this.stack[removedIndex] = this.stack[this.index]

		this.index--
		delete(this.bucket, val)
	}
	return ok
}

func (this *RandomizedSet) GetRandom() int {
	i := rand.Intn(int(this.index) + 1)
	return this.stack[i]
}

type action struct {
	name     string
	v        int
	opResult bool
	exp      map[int]struct{}
}

func Solution(actions []action) bool {
	c := Constructor380()
	for _, act := range actions {
		if act.name == "insert" {
			if c.Insert(act.v) != act.opResult {
				return false
			}
			continue
		}
		if act.name == "remove" {
			if c.Remove(act.v) != act.opResult {
				return false
			}
			continue
		}
		e := c.GetRandom()
		if _, ok := act.exp[e]; !ok {
			return false
		}
	}
	return true
}
