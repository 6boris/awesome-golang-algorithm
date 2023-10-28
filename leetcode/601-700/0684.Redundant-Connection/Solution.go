package Solution

type unionFind684 struct {
	father []int
}

func (u *unionFind684) findFather(x int) int {
	if u.father[x] != x {
		u.father[x] = u.findFather(u.father[x])
	}
	return u.father[x]
}

func (u *unionFind684) union(fx, fy int) {
	if fx < fy {
		u.father[fy] = fx
	} else {
		u.father[fx] = fy
	}
}

func Solution(edges [][]int) []int {
	n := len(edges)
	u := unionFind684{father: make([]int, n+1)}
	ans := make([]int, 2)
	for i := 1; i <= n; i++ {
		u.father[i] = i
	}
	for _, edge := range edges {
		from, to := edge[0], edge[1]
		ff, ft := u.findFather(from), u.findFather(to)
		if ff == ft {
			ans[0] = from
			ans[1] = to
			continue
		}
		u.union(ff, ft)
	}

	return ans
}
