package Solution

type unionFind3493 struct {
	father [101]int
}

func (u *unionFind3493) find(x int) int {
	if u.father[x] != x {
		u.father[x] = u.find(u.father[x])
	}
	return u.father[x]
}

func (u *unionFind3493) union(x, y int) {
	fx := u.find(x)
	fy := u.find(y)
	if fx < fy {
		u.father[fy] = fx
	} else {
		u.father[fx] = fy
	}
}

func Solution(properties [][]int, k int) int {
	n := len(properties)
	intersect := func(a, b [101]bool) int {
		cnt := 0
		for i := range 101 {
			if a[i] && b[i] {
				cnt++
			}
		}
		return cnt
	}

	distinctNumArr := make([][101]bool, n)
	for i := range properties {
		for _, elem := range properties[i] {
			distinctNumArr[i][elem] = true
		}
	}

	u := &unionFind3493{}
	for i := range 101 {
		u.father[i] = i
	}
	for i := range n {
		for j := i + 1; j < n; j++ {
			if intersect(distinctNumArr[i], distinctNumArr[j]) >= k {
				u.union(i, j)
			}
		}
	}

	var f int
	components := make(map[int]struct{}, n)
	for i := range n {
		f = u.find(i)
		components[f] = struct{}{}

	}
	return len(components)
}
