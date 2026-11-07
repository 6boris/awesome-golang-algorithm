package Solution

import "container/heap"

type Item3970 struct {
	node     int
	lastChar byte
	count    int
	dist     int
	index    int // heap 内部索引
}

// PriorityQueue3970 实现 heap.Interface 接口
type PriorityQueue3970 []*Item3970

func (pq PriorityQueue3970) Len() int           { return len(pq) }
func (pq PriorityQueue3970) Less(i, j int) bool { return pq[i].dist < pq[j].dist }
func (pq PriorityQueue3970) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue3970) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item3970)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue3970) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.index = -1
	*pq = old[0 : n-1]
	return item
}

type Edge3970 struct {
	to, weight int
}

func Solution(n int, edges [][]int, labels string, k int) int {
	graph := make([][]Edge3970, n)
	for _, e := range edges {
		u, v, w := e[0], e[1], e[2]
		graph[u] = append(graph[u], Edge3970{v, w})
	}
	const INF = 1 << 60
	dist := make([][26]map[int]int, n)
	for i := range n {
		for c := range 26 {
			dist[i][c] = make(map[int]int)
		}
	}
	pq := &PriorityQueue3970{}
	heap.Init(pq)

	// 初始化起点 0
	startChar := labels[0] - 'a'
	startItem := &Item3970{
		node:     0,
		lastChar: startChar,
		count:    1,
		dist:     0,
	}
	heap.Push(pq, startItem)
	dist[0][startChar][1] = 0

	for pq.Len() > 0 {
		curr := heap.Pop(pq).(*Item3970)

		// 如果到达终点 n-1，因为是 Dijkstra，第一次弹出即为最短路
		if curr.node == n-1 {
			return curr.dist
		}

		// 剪枝：如果当前弹出的距离大于已记录的最短距离，跳过
		if d, ok := dist[curr.node][curr.lastChar][curr.count]; ok && curr.dist > d {
			continue
		}

		// 遍历邻居
		for _, edge := range graph[curr.node] {
			next := edge.to
			nextChar := labels[next] - 'a'
			nextWeight := curr.dist + edge.weight

			var nextCount int
			if nextChar == curr.lastChar {
				nextCount = curr.count + 1
				// 超过连续限制，跳过
				if nextCount > k {
					continue
				}
			} else {
				nextCount = 1
			}

			// 如果找到了更优的路径，更新并入队
			if d, ok := dist[next][nextChar][nextCount]; !ok || nextWeight < d {
				dist[next][nextChar][nextCount] = nextWeight
				heap.Push(pq, &Item3970{
					node:     next,
					lastChar: nextChar,
					count:    nextCount,
					dist:     nextWeight,
				})
			}
		}
	}
	return -1
}
