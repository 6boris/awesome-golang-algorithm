package Solution

import "sort"

type MyCalendarTwo struct {
	store map[int]int
	keys  []int
}

func Constructor731() MyCalendarTwo {
	return MyCalendarTwo{
		store: make(map[int]int),
		keys:  make([]int, 0),
	}
}

func (this *MyCalendarTwo) Book(start int, end int) bool {
	_, ok1 := this.store[start]
	_, ok2 := this.store[end]
	tmp := make([]int, len(this.keys))
	copy(tmp, this.keys)
	this.store[start]++
	this.store[end]--

	needSort := false
	if !ok1 {
		needSort = true
		tmp = append(tmp, start)
	}
	if !ok2 {
		needSort = true
		tmp = append(tmp, end)
	}
	if needSort {
		sort.Ints(tmp)
	}
	ans := 0
	m := 0
	for _, key := range tmp {
		ans += this.store[key]
		if ans > m {
			m = ans
		}
	}
	if m >= 3 {
		if !ok1 {
			delete(this.store, start)
		} else {
			this.store[start]--
		}
		if !ok2 {
			delete(this.store, end)
		} else {
			this.store[end]++
		}
		return false
	}
	this.keys = tmp
	return true
}

func Solution(inputs [][]int) []bool {
	c := Constructor731()
	ans := make([]bool, len(inputs))
	for idx, in := range inputs {
		ans[idx] = c.Book(in[0], in[1])
	}
	return ans
}
