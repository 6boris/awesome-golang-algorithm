package Solution

type unionFind2685 struct {
	father []int
}

func (u *unionFind2685) find(x int) int {
	if u.father[x] != x {
		u.father[x] = u.find(u.father[x])
	}
	return u.father[x]
}
func (u *unionFind2685) union(x, y int) {
	fx := u.find(x)
	fy := u.find(y)
	if fx < fy {
		u.father[fy] = fx
	} else {
		u.father[fx] = fy
	}
}

func Solution(n int, edges [][]int) int {
	in := make([]int, n)
	for _, e := range edges {
		in[e[0]]++
		in[e[1]]++
	}
	u := &unionFind2685{father: make([]int, n)}
	for i := 0; i < n; i++ {
		u.father[i] = i
	}
	for _, e := range edges {
		u.union(e[0], e[1])
	}

	ans := 0
	groupMembers := make([]int, n)
	for i := 0; i < n; i++ {
		f := u.find(i)
		groupMembers[f]++
	}
	for i := 0; i < n; i++ {
		f := u.find(i)
		if groupMembers[f] == -1 {
			continue
		}
		if in[i] != groupMembers[f]-1 {
			groupMembers[f] = -1
		}
	}
	for i := 0; i < n; i++ {
		if groupMembers[i] > 0 {
			ans++
		}
	}
	return ans
}
