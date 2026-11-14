package Solution

import (
	"container/heap"

	"github.com/emirpasic/gods/trees/redblacktree"
)

type BusyServer struct {
	releaseTime int
	serverId    int
}

type MinHeap []BusyServer

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i].releaseTime < h[j].releaseTime }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(BusyServer))
}
func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func Solution(k int, arrival []int, load []int) []int {
	available := redblacktree.NewWithIntComparator()
	for i := 0; i < k; i++ {
		available.Put(i, true)
	}

	busyHeap := &MinHeap{}
	heap.Init(busyHeap)

	requestCount := make([]int, k)
	maxRequests := 0

	for i := 0; i < len(arrival); i++ {
		currTime := arrival[i]
		duration := load[i]

		for busyHeap.Len() > 0 && (*busyHeap)[0].releaseTime <= currTime {
			item := heap.Pop(busyHeap).(BusyServer)
			available.Put(item.serverId, true)
		}

		if available.Size() == 0 {
			continue
		}

		target := i % k
		var serverId int

		node, found := available.Ceiling(target)
		if found {
			serverId = node.Key.(int)
		} else {
			serverId = available.Left().Key.(int)
		}

		available.Remove(serverId)
		heap.Push(busyHeap, BusyServer{
			releaseTime: currTime + duration,
			serverId:    serverId,
		})

		requestCount[serverId]++
		if requestCount[serverId] > maxRequests {
			maxRequests = requestCount[serverId]
		}
	}

	var result []int
	for i := 0; i < k; i++ {
		if requestCount[i] == maxRequests {
			result = append(result, i)
		}
	}

	return result
}
