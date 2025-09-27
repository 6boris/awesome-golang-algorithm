package Solution

import (
	"container/heap"
	"sort"
)

type movie struct {
	shop, movie, price, index int
}

type movieList []*movie

func (m *movieList) Len() int {
	return len(*m)
}
func (m *movieList) Swap(i, j int) {
	(*m)[i], (*m)[j] = (*m)[j], (*m)[i]
	(*m)[i].index = i
	(*m)[j].index = j
}

func (m *movieList) Less(i, j int) bool {
	a, b := (*m)[i], (*m)[j]
	if a.price != b.price {
		return a.price < b.price
	}
	if a.shop != b.shop {
		return a.shop < b.shop
	}
	return a.movie < b.movie
}

func (m *movieList) Push(x any) {
	mv := x.(*movie)
	l := len(*m)
	mv.index = l
	*m = append(*m, mv)
}

func (m *movieList) Pop() any {
	old := *m
	l := len(old)
	x := old[l-1]
	*m = old[:l-1]
	return x
}

type MovieRentingSystem struct {
	// 每个电影自己的
	movieShops map[int]*movieList
	// 已经租界的电影列表
	rentQueue *movieList
	index     map[[2]int]*movie

	n int
}

func Constructor(n int, entries [][]int) MovieRentingSystem {
	instance := MovieRentingSystem{
		movieShops: make(map[int]*movieList),
		rentQueue:  &movieList{},
		index:      make(map[[2]int]*movie),
		n:          n,
	}
	// shop, movie, price
	// movie, shop
	for _, ent := range entries {
		// shop, movie
		key := [2]int{ent[0], ent[1]}
		if _, ok := instance.movieShops[ent[1]]; !ok {
			instance.movieShops[ent[1]] = &movieList{}
		}
		m := &movie{shop: ent[0], movie: ent[1], price: ent[2], index: 0}
		instance.index[key] = m
		heap.Push(instance.movieShops[ent[1]], m)
	}
	return instance
}

func (this *MovieRentingSystem) Search(m int) []int {
	top5 := [][]int{}
	list := this.movieShops[m]
	if list == nil {
		return nil
	}
	store := make([]*movie, 5)
	index := 0
	for list.Len() > 0 && index < 5 {
		top := heap.Pop(list).(*movie)
		store[index] = top
		index++
		top5 = append(top5, []int{top.shop, top.price})
	}
	for i := 0; i < index; i++ {
		heap.Push(list, store[i])
	}
	sort.Slice(top5, func(i, j int) bool {
		a, b := top5[i], top5[j]
		if a[1] != b[1] {
			return a[1] < b[1]
		}
		return a[0] < b[0]
	})
	ret := make([]int, len(top5))
	for i := range top5 {
		ret[i] = top5[i][0]
	}

	return ret
}

func (this *MovieRentingSystem) Rent(shop int, m int) {
	// 需要记录价格
	key := [2]int{shop, m}
	item := this.index[key]
	_ = heap.Remove(this.movieShops[m], item.index)
	heap.Push(this.rentQueue, item)
}

func (this *MovieRentingSystem) Drop(shop int, m int) {
	key := [2]int{shop, m}
	item := this.index[key]
	heap.Remove(this.rentQueue, item.index)
	heap.Push(this.movieShops[m], item)
}

func (this *MovieRentingSystem) Report() [][]int {
	top5 := [][]int{}
	store := make([]*movie, 5)
	index := 0
	for this.rentQueue.Len() > 0 && index < 5 {
		top := heap.Pop(this.rentQueue).(*movie)
		store[index] = top
		index++
		top5 = append(top5, []int{top.shop, top.movie, top.price})
	}
	for i := 0; i < index; i++ {
		heap.Push(this.rentQueue, store[i])
	}

	sort.Slice(top5, func(i, j int) bool {
		a, b := top5[i], top5[j]
		if a[2] != b[2] {
			return a[2] < b[2]
		}
		if a[0] != b[0] {
			return a[0] < b[0]
		}
		return a[1] < b[1]
	})
	ret := make([][]int, len(top5))
	for i := range top5 {
		ret[i] = top5[i][:2]
	}

	return ret
}

type op struct {
	name        string
	shop, movie int
}

func Solution(m int, entries [][]int, ops []op) []any {
	c := Constructor(m, entries)
	ret := make([]any, 0)
	for _, o := range ops {
		if o.name == "rent" {
			c.Rent(o.shop, o.movie)
			continue
		}
		if o.name == "drop" {
			c.Drop(o.shop, o.movie)
			continue
		}
		if o.name == "search" {
			ret = append(ret, c.Search(o.movie))
			continue
		}
		ret = append(ret, c.Report())
	}
	return ret
}
