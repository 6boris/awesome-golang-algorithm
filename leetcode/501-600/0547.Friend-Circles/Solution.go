package Solution

var (
	p   []int
	r   []int
	ans int
)

func findCircleNum(M [][]int) int {
	ans = len(M)
	initialize(ans)

	for i := 0; i < len(M); i++ {
		for j := 0; j < i; j++ {
			if 1 == M[i][j] {
				union(i, j)
			}
		}
	}
	return ans
}

func initialize(l int) {
	p = make([]int, l)
	r = make([]int, l)
	for i := range p {
		p[i] = i
		r[i] = 1
	}
}

func find(x int) int {
	if x != p[x] {
		p[x] = find(p[x])
	}
	return p[x]
}

func union(x, y int) {
	x, y = find(x), find(y)
	if x != y {
		if r[x] <= r[y] {
			p[x] = y
			r[y] += r[x]
		} else {
			p[y] = x
			r[x] += r[y]
		}
		ans--
	}
}
