package Solution

import (
	"math/bits"
	"strconv"
	"strings"
)

type BIT struct {
	tree []int
	n    int
}

func NewBIT(n int) *BIT {
	return &BIT{tree: make([]int, n+1), n: n}
}

func lowbit(x int) int { return x & -x }

func (b *BIT) Update(x, v int) {
	for x <= b.n {
		b.tree[x] ^= v
		x += lowbit(x)
	}
}

func (b *BIT) Query(x int) int {
	ret := 0
	for x > 0 {
		ret ^= b.tree[x]
		x -= lowbit(x)
	}
	return ret
}

func (b *BIT) UpdateRange(l, r, v int) {
	b.Update(l, v)
	b.Update(r+1, v)
}

func Solution(n int, edges [][]int, s string, queries []string) []bool {
	g := make([][]int, n)
	for _, e := range edges {
		g[e[0]] = append(g[e[0]], e[1])
		g[e[1]] = append(g[e[1]], e[0])
	}

	LOG := bits.Len(uint(n))
	fa := make([][]int, n)
	dep := make([]int, n)
	l := make([]int, n)
	r := make([]int, n)
	for i := range fa {
		fa[i] = make([]int, LOG)
	}
	timer := 0

	var dfs func(u, p int)
	dfs = func(u, p int) {
		timer++
		l[u] = timer
		fa[u][0] = p
		for _, v := range g[u] {
			if v == p {
				continue
			}
			dep[v] = dep[u] + 1
			dfs(v, u)
		}
		r[u] = timer
	}
	dfs(0, 0)

	for j := 1; j < LOG; j++ {
		for i := 0; i < n; i++ {
			fa[i][j] = fa[fa[i][j-1]][j-1]
		}
	}

	lca := func(u, v int) int {
		if dep[u] < dep[v] {
			u, v = v, u
		}
		diff := dep[u] - dep[v]
		for i := 0; i < LOG; i++ {
			if (diff>>i)&1 == 1 {
				u = fa[u][i]
			}
		}
		if u == v {
			return u
		}
		for i := LOG - 1; i >= 0; i-- {
			if fa[u][i] != fa[v][i] {
				u = fa[u][i]
				v = fa[v][i]
			}
		}
		return fa[u][0]
	}

	bit := NewBIT(n + 1)
	chars := []byte(s)
	for i := 0; i < n; i++ {
		bit.UpdateRange(l[i], r[i], 1<<(chars[i]-'a'))
	}

	ans := []bool{}
	for _, q := range queries {
		parts := strings.Fields(q)
		if parts[0] == "update" {
			u, _ := strconv.Atoi(parts[1])
			c := parts[2][0]
			bit.UpdateRange(l[u], r[u], (1<<(chars[u]-'a'))^(1<<(c-'a')))
			chars[u] = c
		} else {
			u, _ := strconv.Atoi(parts[1])
			v, _ := strconv.Atoi(parts[2])
			w := lca(u, v)
			mask := bit.Query(l[u]) ^ bit.Query(l[v]) ^ (1 << (chars[w] - 'a'))
			ans = append(ans, bits.OnesCount(uint(mask)) <= 1)
		}
	}
	return ans
}
