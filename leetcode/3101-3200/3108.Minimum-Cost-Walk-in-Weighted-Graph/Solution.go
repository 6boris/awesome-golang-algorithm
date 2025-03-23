package Solution

type unionFind3108 struct {
	father []int
	weight []int
}

func (u *unionFind3108) find(x int) int {
	if u.father[x] != x {
		u.father[x] = u.find(u.father[x])
	}
	return u.father[x]
}

func (u *unionFind3108) union(x, y, w int) {
	fx := u.find(x)
	fy := u.find(y)
	if fx < fy {
		u.father[fy] = fx
		u.weight[fx] = u.weight[fx] & u.weight[fy] & w
	} else {
		u.father[fx] = fy
		u.weight[fy] = u.weight[fx] & u.weight[fy] & w
	}
}

func Solution(n int, edges [][]int, query [][]int) []int {
	u := &unionFind3108{father: make([]int, n), weight: make([]int, n)}
	for i := range n {
		u.father[i] = i
		u.weight[i] = 0xffffffff
	}

	for _, e := range edges {
		f, t, w := e[0], e[1], e[2]
		u.union(f, t, w)
	}
	ans := make([]int, len(query))
	for i, q := range query {
		f, t := q[0], q[1]
		ff := u.find(f)
		tf := u.find(t)
		if ff != tf {
			ans[i] = -1
			continue
		}
		ans[i] = u.weight[ff]
	}
	return ans
}
