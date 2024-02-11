package Solution

import "container/heap"

type seats []int

func (s *seats) Len() int {
	return len(*s)
}

func (s *seats) Less(i, j int) bool {
	return (*s)[i] < (*s)[j]
}

func (s *seats) Swap(i, j int) {
	(*s)[i], (*s)[j] = (*s)[j], (*s)[i]
}

func (s *seats) Push(x any) {
	*s = append(*s, x.(int))
}
func (s *seats) Pop() any {
	old := *s
	l := len(old)
	x := old[l-1]
	*s = old[:l-1]
	return x
}

type SeatManager struct {
	unused seats
}

func Constructor1845(n int) SeatManager {
	unused := seats{}
	for i := 1; i <= n; i++ {
		heap.Push(&unused, i)
	}
	return SeatManager{
		unused: unused,
	}
}

func (this *SeatManager) Reserve() int {
	x := heap.Pop(&this.unused).(int)
	return x
}

func (this *SeatManager) Unreserve(seatNumber int) {
	heap.Push(&this.unused, seatNumber)
}

type opt struct {
	name string
	val  int
}

func Solution(n int, opts []opt) []int {
	ans := make([]int, 0)
	c := Constructor1845(n)
	for _, op := range opts {
		if op.name == "r" {
			ans = append(ans, c.Reserve())
			continue
		}
		c.Unreserve(op.val)
	}
	return ans
}
