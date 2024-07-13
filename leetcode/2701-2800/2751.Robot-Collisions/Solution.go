package Solution

import "sort"

type tmp2751 struct {
	i    []int
	p, h []int
	d    []byte
}

func (t tmp2751) Len() int {
	return len(t.p)
}

func (t tmp2751) Swap(i, j int) {
	t.p[i], t.p[j] = t.p[j], t.p[i]
	t.h[i], t.h[j] = t.h[j], t.h[i]
	t.d[i], t.d[j] = t.d[j], t.d[i]
	t.i[i], t.i[j] = t.i[j], t.i[i]
}

func (t tmp2751) Less(i, j int) bool {
	return t.p[i] < t.p[j]
}

func Solution(positions []int, healths []int, directions string) []int {
	ans := make([]int, 0)
	l := len(positions)
	indies := make([]int, l)
	for i := 0; i < l; i++ {
		indies[i] = i
	}
	s := tmp2751{p: positions, h: healths, d: []byte(directions), i: indies}
	sort.Sort(s)

	stack := make([]int, len(positions))
	stackIndex := -1
	for i := 0; i < l; i++ {
		dir := s.d[i]
		if dir == 'L' {
			for stackIndex >= 0 && healths[stack[stackIndex]] < healths[i] {
				healths[i]--
				stackIndex--
			}

			if stackIndex == -1 {
				ans = append(ans, i)
				continue
			}
			if healths[stack[stackIndex]] == healths[i] {
				stackIndex--
				continue
			}

			healths[stack[stackIndex]]--

			continue
		}
		stackIndex++
		stack[stackIndex] = i
	}
	for i := 0; i <= stackIndex; i++ {
		ans = append(ans, stack[i])
	}
	sort.Slice(ans, func(i, j int) bool {
		return indies[ans[i]] < indies[ans[j]]
	})
	for i := 0; i < len(ans); i++ {
		ans[i] = healths[ans[i]]
	}

	return ans
}
