package Solution

type UnionFind2316 struct {
	Father []int
}

func (uf *UnionFind2316) FindFather(x int) int {
	if uf.Father[x] != x {
		uf.Father[x] = uf.FindFather(uf.Father[x])
	}
	return uf.Father[x]
}

func (uf *UnionFind2316) Union(x, y int) {
	x = uf.Father[x]
	y = uf.Father[y]
	if x < y {
		uf.Father[y] = x
	} else {
		uf.Father[x] = y
	}
}

func Solution(n int, edges [][]int) int64 {
	u := UnionFind2316{Father: make([]int, n)}
	for i := 0; i < n; i++ {
		u.Father[i] = i
	}
	for _, e := range edges {
		af := u.FindFather(e[0])
		bf := u.FindFather(e[1])
		u.Union(af, bf)
	}
	groups := make(map[int]int64)
	for i := 0; i < n; i++ {
		groups[u.FindFather(i)]++
	}
	num := make([]int64, len(groups))
	index := 0
	for _, v := range groups {
		num[index] = v
		index++
	}
	ans := int64(0)
	for i := 0; i < len(num)-1; i++ {
		for j := i + 1; j < len(num); j++ {
			ans += num[i] * num[j]
		}
	}
	return ans
}
