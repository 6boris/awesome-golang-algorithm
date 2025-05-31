package Solution

type unionFind990 struct {
	father [26]int
}

func (u *unionFind990) findFather(x int) int {
	if x != u.father[x] {
		u.father[x] = u.findFather(u.father[x])
	}
	return u.father[x]
}
func (u *unionFind990) union(x, y int) {
	fx := u.findFather(x)
	fy := u.findFather(y)
	if fx < fy {
		u.father[fx] = fy
	} else {
		u.father[fy] = fx
	}
}

func Solution(equations []string) bool {
	uf := unionFind990{father: [26]int{}}
	for i := range 26 {
		uf.father[i] = i
	}

	for _, v := range equations {
		eq := v[1:3]
		a, b := int(v[0]-'a'), int(v[3]-'a')
		if eq != "==" {
			continue
		}
		uf.union(a, b)
	}
	for _, v := range equations {
		eq := v[1:3]
		a, b := int(v[0]-'a'), int(v[3]-'a')
		if eq == "==" {
			continue
		}
		fa := uf.findFather(a)
		fb := uf.findFather(b)
		if fa == fb {
			return false
		}
	}
	return true
}
