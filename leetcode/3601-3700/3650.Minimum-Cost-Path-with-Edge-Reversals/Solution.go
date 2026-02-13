package Solution

import (
	"container/heap"
	"math"
)

type Edge struct {
	to     int
	weight int
}

type Item struct {
	node int
	dist int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int            { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool  { return pq[i].dist < pq[j].dist }
func (pq PriorityQueue) Swap(i, j int)       { pq[i], pq[j] = pq[j], pq[i] }
func (pq *PriorityQueue) Push(x interface{}) { *pq = append(*pq, x.(*Item)) }
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

func Solution(n int, edges [][]int) int {
	adj := make([][]Edge, n)
	revAdj := make([][]Edge, n)
	for _, e := range edges {
		u, v, w := e[0], e[1], e[2]
		adj[u] = append(adj[u], Edge{v, w})
		revAdj[v] = append(revAdj[v], Edge{u, w})
	}

	dist := make([]int, n)
	for i := range dist {
		dist[i] = math.MaxInt32
	}
	dist[0] = 0

	pq := &PriorityQueue{}
	heap.Init(pq)
	heap.Push(pq, &Item{node: 0, dist: 0})

	for pq.Len() > 0 {
		current := heap.Pop(pq).(*Item)
		u := current.node
		d := current.dist

		if d > dist[u] {
			continue
		}
		if u == n-1 {
			return d
		}

		for _, e := range adj[u] {
			if dist[u]+e.weight < dist[e.to] {
				dist[e.to] = dist[u] + e.weight
				heap.Push(pq, &Item{node: e.to, dist: dist[e.to]})
			}
		}

		for _, e := range revAdj[u] {
			cost := e.weight * 2
			if dist[u]+cost < dist[e.to] {
				dist[e.to] = dist[u] + cost
				heap.Push(pq, &Item{node: e.to, dist: dist[e.to]})
			}
		}
	}

	if dist[n-1] == math.MaxInt32 {
		return -1
	}
	return dist[n-1]
}
