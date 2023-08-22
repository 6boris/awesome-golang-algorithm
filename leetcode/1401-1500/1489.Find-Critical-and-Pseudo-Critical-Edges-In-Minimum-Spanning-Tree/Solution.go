package Solution

import "sort"

func newUnion1489(n int) union1489 {
	u := union1489{father: make([]int, n), size: make([]int, n), maxsize: 0}
	for i := 0; i < n; i++ {
		u.father[i] = i
		u.size[i] = 1
	}
	return u
}

type union1489 struct {
	father []int
	// init to 1
	size    []int
	maxsize int
}

func (u *union1489) Find(x int) int {
	if x != u.father[x] {
		u.father[x] = u.Find(u.father[x])
	}
	return u.father[x]
}

func (u *union1489) union(x, y int) bool {
	fx := u.Find(x)
	fy := u.Find(y)
	if fx != fy {
		if u.size[fx] < u.size[fy] {
			fx, fy = fy, fx
		}
		u.father[fy] = fx
		u.size[fx] += u.size[fy]
		if u.size[fx] > u.maxsize {
			u.maxsize = u.size[fx]
		}
		return true
	}
	return false
}

func Solution(n int, edges [][]int) [][]int {
	u := newUnion1489(n)
	// kruskal 算法就是并查集+堆实现。每次将最小权重的边放到池子里。然后判断是否所有的点都已经加进去了
	for i := 0; i < len(edges); i++ {
		edges[i] = append(edges[i], i)
	}

	sort.Slice(edges, func(i, j int) bool {
		return edges[i][2] < edges[j][2]
	})

	minWeight := 0
	//计算树的最小权重
	for _, e := range edges {
		if u.union(e[0], e[1]) {
			minWeight += e[2]
		}
	}

	c := make([]int, 0)
	pc := make([]int, 0)
	for i := 0; i < len(edges); i++ {
		tmp := newUnion1489(n)
		w := 0
		for j := 0; j < len(edges); j++ {
			if i != j && tmp.union(edges[j][0], edges[j][1]) {
				w += edges[j][2]
			}
		}
		// 如果不用第i条边，组成一颗最小生成树
		if tmp.maxsize < n || w > minWeight {
			// 组成的权重比最小的要大，表明是临界边
			c = append(c, edges[i][3])
		} else {
			// 判断是否是伪临界边
			//  a pseudo-critical edge is that which can appear in some MSTs but not all.
			tmp1 := newUnion1489(n)
			w1 := edges[i][2]
			tmp1.union(edges[i][0], edges[i][1])
			for j := 0; j < len(edges); j++ {
				if j != i && tmp1.union(edges[j][0], edges[j][1]) {
					w1 += edges[j][2]
				}
			}

			if w1 == minWeight {
				pc = append(pc, edges[i][3])
			}
		}
	}

	return [][]int{c, pc}
}
