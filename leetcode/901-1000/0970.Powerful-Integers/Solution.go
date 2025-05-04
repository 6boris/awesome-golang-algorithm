package Solution

func Solution(x int, y int, bound int) []int {
	xm := map[int]struct{}{
		1: struct{}{},
	}
	ym := map[int]struct{}{
		1: struct{}{},
	}
	if x != 1 {
		base := x
		for base <= bound {
			xm[base] = struct{}{}
			base *= x
		}
	}
	if y != 1 {
		base := 1
		for base <= bound {
			ym[base] = struct{}{}
			base *= y
		}
	}

	ans := make([]int, 0)
	exist := make(map[int]struct{})
	for xi := range xm {
		for yi := range ym {
			if r := xi + yi; r <= bound {
				if _, ok := exist[r]; !ok {
					exist[r] = struct{}{}
					ans = append(ans, r)
				}
			}
		}
	}
	return ans
}
