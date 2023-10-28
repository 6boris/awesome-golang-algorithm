package Solution

type unionFind1202 struct {
	fatehr []int
}

func (u *unionFind1202) findFather(x int) int {
	if u.fatehr[x] != x {
		u.fatehr[x] = u.findFather(u.fatehr[x])
	}
	return u.fatehr[x]
}
func (u *unionFind1202) union(x, y int) {
	fx := u.findFather(x)
	fy := u.findFather(y)
	if fx < fy {
		u.fatehr[fy] = fx
	} else {
		u.fatehr[fx] = fy
	}
}

type stroe1202 struct {
	bytes [26]int
	index int
}

func (s *stroe1202) selectByte() byte {
	for s.bytes[s.index] == 0 {
		s.index++
	}
	x := uint8(s.index) + 'a'
	s.bytes[s.index]--
	return x
}

func Solution(s string, pairs [][]int) string {
	bs := []byte(s)
	u := &unionFind1202{fatehr: make([]int, len(s))}
	for i := 0; i < len(s); i++ {
		u.fatehr[i] = i
	}
	for _, pair := range pairs {
		u.union(pair[0], pair[1])
	}

	children := make(map[int]*stroe1202)
	for idx := 0; idx < len(s); idx++ {
		fidx := u.findFather(idx)
		if _, ok := children[fidx]; !ok {
			children[fidx] = &stroe1202{bytes: [26]int{}, index: 0}
		}
		children[fidx].bytes[bs[idx]-'a']++
	}

	for idx := 0; idx < len(bs); idx++ {
		fidx := u.findFather(idx)
		bs[idx] = children[fidx].selectByte()
	}

	return string(bs)
}
