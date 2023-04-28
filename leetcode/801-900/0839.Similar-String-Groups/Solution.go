package Solution

func isSimilar(a, b string) bool {
	if a == b {
		return true
	}
	diff := 0
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			diff++
			if diff > 2 {
				return false
			}
		}
	}
	return diff == 2
}

type unionFind839 struct {
	father []int
}

func (u *unionFind839) findFather(x int) int {
	f := u.father[x]
	if f != x {
		u.father[x] = u.findFather(f)
	}
	return u.father[x]
}
func (u *unionFind839) union(x, y int) {
	fx := u.findFather(x)
	fy := u.findFather(y)
	if fx < fy {
		u.father[y] = fx
	} else {
		u.father[x] = fy
	}
}

func Solution(strs []string) int {
	u := unionFind839{father: make([]int, len(strs))}
	for i := 0; i < len(strs); i++ {
		u.father[i] = i
	}

	for idx := 1; idx < len(strs); idx++ {
		for pre := 0; pre < idx; pre++ {
			if !isSimilar(strs[idx], strs[pre]) {
				continue
			}
			f1 := u.findFather(idx)
			f2 := u.findFather(pre)
			u.union(f1, f2)
		}
	}
	x := make(map[int]struct{})
	for i := 0; i < len(strs); i++ {
		x[u.findFather(i)] = struct{}{}
	}

	return len(x)
}
