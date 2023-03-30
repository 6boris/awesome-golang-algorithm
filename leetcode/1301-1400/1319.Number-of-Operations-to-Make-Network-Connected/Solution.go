package Solution

type UnionFind1319 struct {
	Father []int
}

func (uf *UnionFind1319) FindFather(x int) int {
	if uf.Father[x] != x {
		uf.Father[x] = uf.FindFather(uf.Father[x])
	}
	return uf.Father[x]
}

func (uf *UnionFind1319) Union(x, y int) {
	x = uf.Father[x]
	y = uf.Father[y]
	if x < y {
		uf.Father[y] = x
	} else {
		uf.Father[x] = y
	}
}
func Solution(n int, connections [][]int) int {
	if len(connections) < n-1 {
		return -1
	}
	u := UnionFind1319{Father: make([]int, n)}
	for i := 0; i < n; i++ {
		u.Father[i] = -1
	}

	// 计算小团体
	// 计算没有任何连接的电脑个数

	for _, conn := range connections {
		if u.Father[conn[0]] == -1 {
			u.Father[conn[0]] = conn[0]
		}
		if u.Father[conn[1]] == -1 {
			u.Father[conn[1]] = conn[1]
		}
		af := u.FindFather(conn[0])
		bf := u.FindFather(conn[1])
		u.Union(af, bf)
	}
	ans := 0
	groups := make(map[int]int)
	for i := 0; i < n; i++ {
		if u.Father[i] == -1 {
			groups[-1]++
			continue
		}
		groups[u.FindFather(i)]++
	}
	singleGroups := len(groups)
	v, ok := groups[-1]
	if !ok {
		return singleGroups - 1
	}
	ans += v - 1
	singleGroups -= 1
	return ans + singleGroups
}
