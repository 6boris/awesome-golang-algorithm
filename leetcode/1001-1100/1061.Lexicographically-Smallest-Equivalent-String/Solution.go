package Solution

import "strings"

type UnionFind1061 struct {
	Father []int
}

func (uf *UnionFind1061) FindFather(x int) int {
	if uf.Father[x] != x {
		uf.Father[x] = uf.FindFather(uf.Father[x])
	}
	return uf.Father[x]
}

func (uf *UnionFind1061) Union(x, y int) {
	x = uf.Father[x]
	y = uf.Father[y]
	if x < y {
		uf.Father[y] = x
	} else {
		uf.Father[x] = y
	}
}

func Solution(s1, s2, baseStr string) string {
	uf := &UnionFind1061{
		Father: make([]int, 26),
	}
	for i := 0; i < 26; i++ {
		uf.Father[i] = i
	}
	for i := 0; i < len(s1); i++ {
		x := s1[i] - 'a'
		y := s2[i] - 'a'
		x1 := uf.FindFather(int(x))
		y1 := uf.FindFather(int(y))
		uf.Union(x1, y1)
	}
	sb := strings.Builder{}
	for _, b := range baseStr {
		f := uf.FindFather(int(b - 'a'))
		sb.WriteByte(byte(f) + 'a')
	}
	return sb.String()
}
