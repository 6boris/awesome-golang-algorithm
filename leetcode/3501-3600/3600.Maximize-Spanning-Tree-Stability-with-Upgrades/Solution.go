package Solution

import "sort"

const MAX_STABILITY = 200000

type DSU struct {
	parent []int
}

func NewDSU(parent []int) *DSU {
	p := make([]int, len(parent))
	copy(p, parent)
	return &DSU{parent: p}
}

func (d *DSU) find(x int) int {
	if d.parent[x] != x {
		d.parent[x] = d.find(d.parent[x])
	}
	return d.parent[x]
}

func (d *DSU) join(x, y int) {
	px := d.find(x)
	py := d.find(y)
	d.parent[px] = py
}

func Solution(n int, edges [][]int, k int) int {
	ans := -1

	if len(edges) < n-1 {
		return -1
	}
	mustEdges := [][]int{}
	optionalEdges := [][]int{}

	for _, e := range edges {
		if e[3] == 1 {
			mustEdges = append(mustEdges, e)
		} else {
			optionalEdges = append(optionalEdges, e)
		}
	}

	if len(mustEdges) > n-1 {
		return -1
	}

	sort.Slice(optionalEdges, func(i, j int) bool {
		return optionalEdges[i][2] > optionalEdges[j][2]
	})

	selectedInit := 0
	mustMinStability := MAX_STABILITY

	initParent := make([]int, n)
	for i := 0; i < n; i++ {
		initParent[i] = i
	}
	dsuInit := NewDSU(initParent)

	for _, e := range mustEdges {
		u, v, s := e[0], e[1], e[2]
		if dsuInit.find(u) == dsuInit.find(v) || selectedInit == n-1 {
			return -1
		}
		dsuInit.join(u, v)
		selectedInit++
		if s < mustMinStability {
			mustMinStability = s
		}
	}

	l, r := 0, mustMinStability
	for l < r {
		mid := l + (r-l+1)/2

		dsu := NewDSU(dsuInit.parent)
		selected := selectedInit
		doubledCount := 0

		for _, e := range optionalEdges {
			u, v, s := e[0], e[1], e[2]
			if dsu.find(u) == dsu.find(v) {
				continue
			}

			if s >= mid {
				dsu.join(u, v)
				selected++
			} else if doubledCount < k && s*2 >= mid {
				doubledCount++
				dsu.join(u, v)
				selected++
			} else {
				break
			}

			if selected == n-1 {
				break
			}
		}

		if selected != n-1 {
			r = mid - 1
		} else {
			ans = mid
			l = mid
		}
	}

	return ans
}
