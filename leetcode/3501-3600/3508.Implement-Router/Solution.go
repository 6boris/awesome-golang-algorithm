package Solution

import "sort"

type Router struct {
	memoryLimit int
	queue       [][3]int
	// source, dest, timestamp
	in map[[3]int]struct{}

	// dest: [source, timestamp]
	dest map[int][]int
}

func Constructor(memoryLimit int) Router {
	return Router{
		memoryLimit: memoryLimit,
		queue:       make([][3]int, 0),
		in:          map[[3]int]struct{}{},
		dest:        make(map[int][]int),
	}
}

func (this *Router) AddPacket(source int, destination int, timestamp int) bool {
	key := [3]int{source, destination, timestamp}
	if _, ok := this.in[key]; ok {
		return false
	}
	this.in[key] = struct{}{}
	cur := len(this.queue)
	if cur == this.memoryLimit {
		_ = this.ForwardPacket()
	}
	this.queue = append(this.queue, key)

	this.dest[destination] = append(this.dest[destination], timestamp)
	return true
}

func (this *Router) ForwardPacket() []int {
	if len(this.queue) == 0 {
		return []int{}
	}

	first := this.queue[0]
	this.queue = this.queue[1:]
	delete(this.in, first)
	// 同时移除队列的数据
	sources := this.dest[first[1]]
	index := sort.Search(len(sources), func(i int) bool {
		return sources[i] >= first[2]
	})
	sources = append(sources[:index], sources[index+1:]...)
	this.dest[first[1]] = sources
	return first[:]
}

func (this *Router) GetCount(destination int, startTime int, endTime int) int {
	sources, ok := this.dest[destination]
	if !ok {
		return 0
	}
	start := sort.Search(len(sources), func(i int) bool {
		return sources[i] >= startTime
	})
	end := sort.Search(len(sources), func(i int) bool {
		return sources[i] > endTime
	})
	return end - start
}

type op struct {
	name                           string
	source, destination, timestamp int
}

func Solution(memoryLimit int, ops []op) []any {
	ret := make([]any, 0)
	c := Constructor(memoryLimit)
	for _, o := range ops {
		if o.name == "add" {
			ret = append(ret, c.AddPacket(o.source, o.destination, o.timestamp))
			continue
		}

		if o.name == "for" {
			ret = append(ret, c.ForwardPacket())
			continue
		}

		if o.name == "get" {
			ret = append(ret, c.GetCount(o.source, o.destination, o.timestamp))
		}
	}
	return ret
}
