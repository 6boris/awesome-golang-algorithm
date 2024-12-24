package Solution

type pair2337 struct {
	c byte
	i int
}

func Solution(start string, target string) bool {
	a, b := make([]pair2337, 0), make([]pair2337, 0)
	for i, c := range []byte(start) {
		if c == '_' {
			continue
		}
		a = append(a, pair2337{c, i})
	}
	for i, c := range []byte(target) {
		if c == '_' {
			continue
		}
		b = append(b, pair2337{c, i})
	}
	if len(a) != len(b) {
		// _ 不相等，无法转换
		return false
	}

	for i := range len(a) {
		ac := a[i]
		bc := b[i]
		if ac.c != bc.c {
			return false
		}
		// 如果我此时是L
		if ac.c == 'L' && ac.i < bc.i {
			return false
		}
		if ac.c == 'R' && ac.i > bc.i {
			return false
		}
	}
	return true
}
