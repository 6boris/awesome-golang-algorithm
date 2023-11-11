package Solution

type Graph struct {
	adj [][]int
	n   int
}

func Constructor2642(n int, edges [][]int) Graph {
	adj := make([][]int, n)
	for i := 0; i < n; i++ {
		adj[i] = make([]int, n)
	}
	for _, edge := range edges {
		f, t, w := edge[0], edge[1], edge[2]
		adj[f][t] = w
	}
	return Graph{adj: adj, n: n}
}

func (this *Graph) AddEdge(edge []int) {
	f, t, w := edge[0], edge[1], edge[2]
	this.adj[f][t] = w
}

func (this *Graph) ShortestPath(node1 int, node2 int) int {
	distance := make([]int, this.n)
	for i := 0; i < this.n; i++ {
		distance[i] = -1
	}
	v := make([]bool, this.n)
	distance[node1] = 0
	for i := 0; i < this.n-1; i++ {
		seletNode := -1
		for from, dis := range distance {
			if dis == -1 {
				continue
			}
			if !v[from] && (seletNode == -1 || dis < distance[seletNode]) {
				seletNode = from
			}
		}
		if seletNode == -1 {
			break
		}
		v[seletNode] = true
		for next, dis := range this.adj[seletNode] {
			if dis == 0 {
				continue
			}
			if distance[next] == -1 || dis+distance[seletNode] < distance[next] {
				distance[next] = dis + distance[seletNode]
			}
		}
	}
	return distance[node2]
}

type opt struct {
	name string
	v    []int
}

func Solution(n int, edges [][]int, opts []opt) []int {
	c := Constructor2642(n, edges)
	ans := make([]int, 0)
	for _, op := range opts {
		if op.name == "s" {
			ans = append(ans, c.ShortestPath(op.v[0], op.v[1]))
			continue
		}
		c.AddEdge(op.v)
	}
	return ans
}
