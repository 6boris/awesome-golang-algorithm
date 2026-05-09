package Solution

type unionFind1722 struct {
	father []int
}

func (u *unionFind1722) find(x int) int {
	if u.father[x] != x {
		u.father[x] = u.find(u.father[x])
	}
	return u.father[x]
}

func (u *unionFind1722) union(x, y int) {
	fx := u.find(x)
	fy := u.find(y)
	if fx < fy {
		u.father[fy] = fx
	} else {
		u.father[fx] = fy
	}
}

func Solution(source []int, target []int, allowedSwaps [][]int) int {
	size := len(source)
	u := unionFind1722{
		father: make([]int, size),
	}
	for i := range size {
		u.father[i] = i
	}
	for i := range allowedSwaps {
		u.union(allowedSwaps[i][0], allowedSwaps[i][1])
	}

	set := make(map[int]map[int]int, size)
	var ret int
	for i := range source {
		f := u.find(i)
		if _, ok := set[f]; !ok {
			set[f] = map[int]int{}
		}
		set[f][source[i]]++
	}
	for i := range target {
		f := u.find(i)
		if set[f][target[i]] == 0 {
			ret++
			continue
		}
		set[f][target[i]]--
	}

	return ret
}
