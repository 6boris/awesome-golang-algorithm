package Solution

import (
	"container/heap"
	"sort"
)

type roomHeap []int

func (m *roomHeap) Len() int {
	return len(*m)
}

func (m *roomHeap) Less(i, j int) bool {
	return (*m)[i] < (*m)[j]
}
func (m *roomHeap) Swap(i, j int) {
	(*m)[i], (*m)[j] = (*m)[j], (*m)[i]
}

func (m *roomHeap) Push(x any) {
	*m = append(*m, x.(int))
}

func (m *roomHeap) Pop() any {
	old := *m
	l := len(old)
	x := old[l-1]
	*m = old[:l-1]
	return x
}

type usedRoom struct {
	id  int
	end int64
}
type usedRoomHeap []usedRoom

func (u *usedRoomHeap) Len() int {
	return len(*u)
}

func (u *usedRoomHeap) Less(i, j int) bool {
	// note: Because there is an update to the end time of the schedule,
	// when end is equal, the smaller id is returned.
	if (*u)[i].end == (*u)[j].end {
		return (*u)[i].id < (*u)[j].id
	}
	return (*u)[i].end < (*u)[j].end
}

func (u *usedRoomHeap) Swap(i, j int) {
	(*u)[i], (*u)[j] = (*u)[j], (*u)[i]
}

func (u *usedRoomHeap) Push(x any) {
	*u = append(*u, x.(usedRoom))
}

func (u *usedRoomHeap) Pop() any {
	old := *u
	l := len(old)
	x := old[l-1]
	*u = old[:l-1]
	return x

}

func Solution(n int, meetings [][]int) int {
	count := make([]int, n)
	sort.Slice(meetings, func(i, j int) bool {
		return meetings[i][0] < meetings[j][0]
	})

	rooms := roomHeap{}
	for i := 0; i < n; i++ {
		heap.Push(&rooms, i)
	}
	usedRooms := usedRoomHeap{}
	// 1,20    2,10,    3,5   4,9,   6,8
	// 0       1        2
	for i := 0; i < len(meetings); i++ {
		start := int64(meetings[i][0])
		end := int64(meetings[i][1])
		for usedRooms.Len() > 0 && usedRooms[0].end <= start {
			top := heap.Pop(&usedRooms).(usedRoom)
			heap.Push(&rooms, top.id)
		}
		// [1, 100], [2,99], [10,20]
		// 0, 1 2
		if rooms.Len() > 0 {
			top := heap.Pop(&rooms).(int)
			count[top]++
			heap.Push(&usedRooms, usedRoom{id: top, end: end})
		} else {
			top := heap.Pop(&usedRooms).(usedRoom)
			// 0, 5
			// 5, 10
			count[top.id]++
			top.end += end - start
			heap.Push(&usedRooms, top)
		}
	}

	ans, cc := 0, 0
	for idx, c := range count {
		if c > cc {
			ans = idx
			cc = c
		}
	}
	return ans
}
